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

// swedate.cpp

func init_leapsec() int { panic("!implemented") }

func swe_date_conversion(y int, missing ...any) int { panic("!implemented") }

func swe_jdet_to_utc(tjd_et double, gregflag int32, iyear *int32, imonth *int32, iday *int32, ihour *int32, imin *int32, missing ...any) {
	panic("!implemented")
}

func swe_jdut1_to_utc(tjd_ut double, gregflag int32, iyear *int32, imonth *int32, iday *int32, ihour *int32, imin *int32, missing ...any) {
	panic("!implemented")
}

func swe_julday(year int, month int, day int, hour double, gregflag int) double {
	panic("!implemented")
}

func swe_revjul(jd double, gregflag int, missing ...any) { panic("!implemented") }

func swe_utc_time_zone(...any) { panic("!implemented") }

func swe_utc_to_jd(iyear int32, imonth int32, iday int32, ihour int32, imin int32, dsec double, gregflag int32, dret *double, missing ...any) int32 {
	panic("!implemented")
}
