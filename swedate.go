// charticus - an astrology chart tool
// Copyright (c) 2023 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package charticus

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
)

// swedate.cpp

// flag for reading leap seconds from file
var init_leapseconds_done = false

/* Leap seconds were inserted at the end of the following days: */
var leap_seconds = []int{
	19720630,
	19721231,
	19731231,
	19741231,
	19751231,
	19761231,
	19771231,
	19781231,
	19791231,
	19810630,
	19820630,
	19830630,
	19850630,
	19871231,
	19891231,
	19901231,
	19920630,
	19930630,
	19940630,
	19951231,
	19970630,
	19981231,
	20051231,
	20081231,
	20120630,
	20150630,
	20161231,
}

const (
	J1972               = 2441317.5
	NLEAP_INIT          = 10
	NLEAP_SECONDS_SPACE = 100

	// Julian date for 1 Jan 1972 UTC, including leap seconds
	tjd_et_1972 = J1972 + (32.184+NLEAP_INIT)/86400.0
)

var NLEAP_SECONDS = len(leap_seconds)

// init_leapsec reads additional leap second dates from an external file.
//
// Returns the number of leap seconds in the table.
//
// TODO: add test
func init_leapsec() int {
	if !init_leapseconds_done {
		// ignore error message if file is missing
		data, err := os.ReadFile(filepath.Join(swed.ephepath, "seleapsec.json"))
		if err == nil {
			var entries []int
			// log the error if the json is invalid
			if err := json.Unmarshal(data, &entries); err != nil {
				log.Printf("%s: %v\n", filepath.Join(swed.ephepath, "seleapsec.json"), err)
			} else {
				leap_seconds = append(leap_seconds, entries...)
			}
		}
		// sort the list to ensure that its, uhm, sorted
		sort.Ints(leap_seconds)
		// set the flag so we don't repeat the load
		init_leapseconds_done = true
	}
	return len(leap_seconds)
}

// swe_date_conversion converts year, month, day to Julian date.
//
// uttime is hours UT
//
// Performs a sanity check after the conversion.
//
// Returns OK if date is valid, ERR if the date is invalid.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
func swe_date_conversion(year, month, day int, uttime float64, c byte, tjd *float64) RTSTATUS {
	gregflag := SE_JUL_CAL
	if c == 'g' {
		gregflag = SE_GREG_CAL
	}
	*tjd = swe_julday(year, month, day, uttime, gregflag)

	var rday, rmon, ryear int
	var rut float64
	swe_revjul(*tjd, gregflag, &ryear, &rmon, &rday, &rut)
	if ryear != year || rmon != month || rday != day {
		return ERR
	}
	return OK
}

// swe_jdet_to_utc does something
//
// Input:  tjd_et   Julian day number, terrestrial time (ephemeris time).
//
//	gregfalg Calendar flag
//
// Output: UTC year, month, day, hour, minute, second (decimal).
//   - Before 1 jan 1972 UTC, output UT1.
//     Note: UTC was introduced in 1961. From 1961 - 1971, the length of the
//     UTC second was regularly changed, so that UTC remained very close to UT1.
//   - From 1972 on, output is UTC.
//   - If delta_t - nleap - 32.184 > 1, the output is UT1.
//     Note: Like this we avoid errors greater than 1 second in case that the leap seconds table (or the Swiss Ephemeris version) has not been updated for a long time.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
// TODO: document what this does
func swe_jdet_to_utc(tjd_et float64, gregflag CALENDARFLAG, iyear, imonth, iday, ihour, imin *int, dsec *float64) {
	// if tjd_et is before 1 jan 1972 UTC, return UT1
	d := swe_deltat_ex(tjd_et, -1, nil)
	tjd_ut := tjd_et - swe_deltat_ex(tjd_et-d, -1, nil)
	tjd_ut = tjd_et - swe_deltat_ex(tjd_ut, -1, nil)
	if tjd_et < tjd_et_1972 {
		swe_revjul(tjd_ut, gregflag, iyear, imonth, iday, &d)
		*ihour = int(d)
		d -= float64(*ihour)
		d *= 60
		*imin = int(d)
		*dsec = (d - float64(*imin)) * 60.0
		return
	}

	// minimum number of leap seconds since 1972; we may be missing one leap second
	tabsiz_nleap := init_leapsec()
	var iyear2, imonth2, iday2 int
	swe_revjul(tjd_ut-1, SE_GREG_CAL, &iyear2, &imonth2, &iday2, &d)
	ndat := iyear2*10000 + imonth2*100 + iday2
	nleap := 0
	for i := 0; i < tabsiz_nleap; i++ {
		if ndat <= leap_seconds[i] {
			break
		}
		nleap++
	}

	// date of potentially missing leapsecond
	second_60 := 0
	if nleap < tabsiz_nleap {
		i := leap_seconds[nleap]
		iyear2 = i / 10000
		imonth2 = (i % 10000) / 100
		iday2 = i % 100
		tjd := swe_julday(iyear2, imonth2, iday2, 0, SE_GREG_CAL)
		swe_revjul(tjd+1, SE_GREG_CAL, &iyear2, &imonth2, &iday2, &d)
		var dret [10]float64
		swe_utc_to_jd(iyear2, imonth2, iday2, 0, 0, 0, SE_GREG_CAL, dret, nil)
		d = tjd_et - dret[0]
		if d >= 0 {
			nleap++
		} else if d < 0 && d > -1.0/86400.0 {
			second_60 = 1
		}
	}

	// UTC, still unsure about one leap second
	tjd := J1972 + (tjd_et - tjd_et_1972) - float64(nleap+second_60)/86400.0
	swe_revjul(tjd, SE_GREG_CAL, iyear, imonth, iday, &d)
	*ihour = int(d)
	d -= float64(*ihour)
	d *= 60
	*imin = int(d)
	*dsec = (d-float64(*imin))*60.0 + float64(second_60)

	// For input dates > today:
	//   If leap seconds table is not up-to-date, we'd better interpret the input time as UT1, not as UTC.
	//   How do we find out?
	//   Check, if delta_t - nleap - 32.184 > 0.9
	d = swe_deltat_ex(tjd_et, -1, nil)
	d = swe_deltat_ex(tjd_et-d, -1, nil)
	if d*86400.0-(double)(nleap+NLEAP_INIT)-32.184 >= 1.0 {
		swe_revjul(tjd_et-d, SE_GREG_CAL, iyear, imonth, iday, &d)
		*ihour = int(d)
		d -= float64(*ihour)
		d *= 60
		*imin = int(d)
		*dsec = (d - float64(*imin)) * 60.0
	}

	// convert Julian calendar to Gregorian if needed
	if gregflag == SE_JUL_CAL {
		tjd = swe_julday(*iyear, *imonth, *iday, 0, SE_GREG_CAL)
		swe_revjul(tjd, gregflag, iyear, imonth, iday, &d)
	}
}

// swe_jdut1_to_utc does something
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
// TODO: document what this does
func swe_jdut1_to_utc(tjd_ut float64, gregflag CALENDARFLAG, iyear, imonth, iday, ihour, imin *int, dsec *float64) {
	tjd_et := tjd_ut + swe_deltat_ex(tjd_ut, -1, nil)
	swe_jdet_to_utc(tjd_et, gregflag, iyear, imonth, iday, ihour, imin, dsec)
}

// swe_julday returns the absolute Julian day number (JD) for a given calendar date.
//
// The arguments are a calendar date: day, month, year as integers, hour as float64 with decimal fraction.
// If gregflag = SE_GREG_CAL, Gregorian calendar is used,
// if gregflag = SE_JUL_CAL, Julian calendar is used.
//
// The Julian day number is a system of numbering all days continuously within the time range of known human history.
// It should be familiar to every astrological or astronomical programmer.
// The time variable in astronomical theories is usually expressed in Julian days or
// Julian centuries (36525 days per century) relative to some start day;
// the start day is called 'the epoch'.
// The Julian day number is a float64 representing the number of
// days since JD = 0.0 on 1 Jan -4712, 12:00 noon (in the Julian calendar).
//
// Midnight has always a JD with fraction .5, because traditionally the astronomical day started at noon.
// This was practical because then there was no change of date during a night at the telescope.
// From this comes also the fact the noon ephemerides were printed
// before midnight ephemerides were introduced early in the 20th century.
//
// NOTE: The Julian day number must not be confused with the Julian calendar system.
//
// Be aware that we always use astronomical year numbering for the years before Christ, not the historical year numbering.
// Astronomical years are done with negative numbers, historical years with indicators BC or BCE (before common era).
//
//	Year  0 (astronomical)  = 1 BC
//	Year -1 (astronomical)  = 2 BC
//
// etc.
//
// Original author: Marc Pottenger, Los Angeles.
// With bug fix for year < -4711   15-aug-88 by Alois Treindl
//
// References: Oliver Montenbruck, Grundlagen der Ephemeridenrechnung, Verlag Sterne und Weltraum (1987), p.49 ff
//
// Related functions:
//
//	swe_revjul() reverse Julian day number: compute the calendar date from a given JD
//	swe_date_conversion() includes test for legal date values and notifies errors like 32 January.
//
// TODO: add test
func swe_julday(year, month, day int, hour float64, gregflag CALENDARFLAG) float64 {
	fyear, fmonth, fday := float64(year), float64(month), float64(day)
	u := fyear
	if month < 3 {
		u -= 1
	}
	u0 := u + 4712.0
	u1 := fmonth + 1.0
	if u1 < 4 {
		u1 += 12.0
	}
	jd := math.Floor(u0*365.25) + math.Floor(30.6*u1+0.000001) + fday + hour/24.0 - 63.5
	if gregflag == SE_GREG_CAL {
		u2 := math.Floor(math.Abs(u)/100) - math.Floor(math.Abs(u)/400)
		if u < 0.0 {
			u2 = -u2
		}
		jd = jd - u2 + 2
		if (u < 0.0) && (u/100 == math.Floor(u/100)) && (u/400 != math.Floor(u/400)) {
			jd -= 1
		}
	}
	return jd
}

// swe_revjul is the inverse function to swe_julday(), see the description there.
//
// Arguments are julian day number, calendar flag.
// Return values are the calendar day, month, year and the hour of the day with decimal fraction (0 .. 23.999999).
//
// Be aware that we use astronomical year numbering for the years before Christ, not the historical year numbering.
// Astronomical years are done with negative numbers, historical years with indicators BC or BCE (before common era).
//
//	Year    0 (astronomical) =   1 BC historical year
//	Year   -1 (astronomical) =   2 BC historical year
//	Year -234 (astronomical) = 235 BC historical year
//
// etc.
//
// Original author Mark Pottenger, Los Angeles.
// Bug fix for year < -4711 16-aug-88 Alois Treindl.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
func swe_revjul(jd float64, gregflag CALENDARFLAG, jyear, jmon, jday *int, jut *float64) {
	u0 := jd + 32082.5

	var u1 float64
	if gregflag == SE_GREG_CAL {
		u1 = u0 + math.Floor(u0/36525.0) - math.Floor(u0/146100.0) - 38.0
		if jd >= 1830691.5 {
			u1 += 1
		}
		u0 = u0 + math.Floor(u1/36525.0) - math.Floor(u1/146100.0) - 38.0
	}

	u2 := math.Floor(u0 + 123.0)
	u3 := math.Floor((u2 - 122.2) / 365.25)
	u4 := math.Floor((u2 - math.Floor(365.25*u3)) / 30.6001)

	*jmon = int(u4 - 1.0)
	if *jmon > 12 {
		*jmon -= 12
	}
	*jday = int(u2 - math.Floor(365.25*u3) - math.Floor(30.6001*u4))
	*jyear = int(u3 + math.Floor((u4-2.0)/12.0) - 4800)
	*jut = (jd - math.Floor(jd+0.5) + 0.5) * 24.0
}

// swe_utc_time_zone transforms local time to UTC or UTC to local time
//
// input
//
//	iyear, imonth, iday, ihour, dsec  date and time
//	d_timezone		timezone offset
//
// output
//
//	year, month, day, hour, sec
//
// For time zones east of Greenwich, d_timezone is positive.
// For time zones west of Greenwich, d_timezone is negative.
//
// For conversion from local time to utc, use +d_timezone.
// For conversion from utc to local time, use -d_timezone.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
func swe_utc_time_zone(iyear, imonth, iday, ihour, imin int, dsec float64, d_timezone float64, iyear_out, imonth_out, iday_out, ihour_out, imin_out *int, dsec_out *float64) {
	have_leapsec := dsec >= 60.0
	if have_leapsec {
		dsec -= 1.0
	}
	tjd := swe_julday(iyear, imonth, iday, 0, SE_GREG_CAL)
	dhour := float64(ihour) + float64(imin)/60.0 + dsec/3600.0 - d_timezone
	if dhour < 0.0 {
		tjd -= 1.0
		dhour += 24.0
	} else if dhour >= 24.0 {
		tjd += 1.0
		dhour -= 24.0
	}
	var d float64
	swe_revjul(tjd+0.001, SE_GREG_CAL, iyear_out, imonth_out, iday_out, &d)
	*ihour_out = int(dhour)
	d = (dhour - float64(*ihour_out)) * 60
	*imin_out = int(d)
	*dsec_out = (d - float64(*imin_out)) * 60
	if have_leapsec {
		*dsec_out += 1.0
	}
}

// swe_utc_to_jd does something
//
// Input:  Clock time UTC, year, month, day, hour, minute, second (decimal).
//
//	gregflag  Calendar flag
//	serr      error string
//
// Output: An array of float64:
//
//	dret[0] = Julian day number TT (ET)
//	dret[1] = Julian day number UT1
//
// Function returns OK or Error.
//
//   - Before 1972, swe_utc_to_jd() treats its input time as UT1.
//     Note: UTC was introduced in 1961. From 1961 - 1971, the length of the
//     UTC second was regularly changed, so that UTC remained very close to UT1.
//   - From 1972 on, input time is treated as UTC.
//   - If delta_t - nleap - 32.184 > 1, the input time is treated as UT1.
//     Note: Like this we avoid errors greater than 1 second in case that
//     the leap seconds table (or the Swiss Ephemeris version) is not updated
//     for a long time.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
// TODO: document what this does
func swe_utc_to_jd(iyear, imonth, iday, ihour, imin int, dsec float64, gregflag CALENDARFLAG, dret [10]float64, serr *[]byte) RTSTATUS {
	// error handling: invalid iyear etc.
	tjd_ut1 := swe_julday(iyear, imonth, iday, 0, gregflag)
	var iyear2, imonth2, iday2 int
	var ut float64
	swe_revjul(tjd_ut1, gregflag, &iyear2, &imonth2, &iday2, &ut)
	if iyear != iyear2 || imonth != imonth2 || iday != iday2 {
		if serr != nil {
			*serr = []byte(fmt.Sprintf("invalid date: year = %d, month = %d, day = %d", iyear, imonth, iday))
		}
		return ERR
	} else if ihour < 0 || ihour > 23 || imin < 0 || imin > 59 || dsec < 0 || dsec >= 61 || (dsec >= 60 && (imin < 59 || ihour < 23 || tjd_ut1 < J1972)) {
		if serr != nil {
			*serr = []byte(fmt.Sprintf("invalid time: %d:%d:%.2f", ihour, imin, dsec))
		}
		return ERR
	}
	dhour := float64(ihour) + (float64(imin))/60.0 + dsec/3600.0

	// before 1972, we treat input date as UT1
	if tjd_ut1 < J1972 {
		dret[1] = swe_julday(iyear, imonth, iday, dhour, gregflag)
		dret[0] = dret[1] + swe_deltat_ex(dret[1], -1, nil)
		return OK
	}
	// if gregflag = Julian calendar, convert to gregorian calendar
	if gregflag == SE_JUL_CAL {
		gregflag = SE_GREG_CAL
		swe_revjul(tjd_ut1, gregflag, &iyear, &imonth, &iday, &ut)
	}
	// number of leap seconds since 1972:
	tabsiz_nleap := init_leapsec()
	nleap := NLEAP_INIT // initial difference between UTC and TAI in 1972
	ndat := iyear*10000 + imonth*100 + iday
	for i := 0; i < tabsiz_nleap && ndat <= leap_seconds[i]; i++ {
		nleap++
	}
	// For input dates > today:
	//   If leap seconds table is not up-to-date, we'd better interpret the input time as UT1, not as UTC.
	//   How do we find out?
	//   Check, if delta_t - nleap - 32.184 > 0.9
	d := swe_deltat_ex(tjd_ut1, -1, nil) * 86400.0
	if outOfDate := d-float64(nleap)-32.184 >= 1.0; outOfDate {
		dret[1] = tjd_ut1 + dhour/24.0
		dret[0] = dret[1] + swe_deltat_ex(dret[1], -1, nil)
		return OK
	}
	// if input second is 60: is it a valid leap second ?
	if dsec >= 60 {
		j := 0
		for i := 0; i < tabsiz_nleap; i++ {
			if ndat == leap_seconds[i] {
				j = 1
				break
			}
		}
		if j != 1 {
			if serr != nil {
				*serr = []byte(fmt.Sprintf("invalid time (no leap second!): %d:%d:%.2f", ihour, imin, dsec))
			}
			return ERR
		}
	}

	// convert UTC to ET and UT1

	// the number of days between input date and 1 jan 1972:
	d = tjd_ut1 - J1972
	// SI time since 1972, ignoring leap seconds:
	d += float64(ihour)/24.0 + float64(imin)/1440.0 + dsec/86400.0
	// ET (TT)
	tjd_et := tjd_et_1972 + d + (float64(nleap-NLEAP_INIT))/86400.0
	d = swe_deltat_ex(tjd_et, -1, nil)
	tjd_ut1 = tjd_et - swe_deltat_ex(tjd_et-d, -1, nil)
	tjd_ut1 = tjd_et - swe_deltat_ex(tjd_ut1, -1, nil)
	dret[0] = tjd_et
	dret[1] = tjd_ut1
	return OK
}
