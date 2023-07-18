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

import "math"

// general.cpp

func AddDay(month int, day int, year int, delta int) int { panic("!implemented") }

func AddTime(pci *CI, mode int, toadd int) { panic("!implemented") }

func AnsiColor(k int) { panic("!implemented") }

func Assert(f bool) { panic("!implemented") }

func CchSz(sz []byte) int { panic("!implemented") }

func ChDeg() byte { panic("!implemented") }

func ChIBMFromWch(wch wchar) uchar { panic("!implemented") }

func ChLatinFromWch(wch wchar) uchar { panic("!implemented") }

func ClearB(pb pbyte, cb int) { panic("!implemented") }

func ConvertSzFromUTF8(sz []byte) { panic("!implemented") }

func ConvertSzToLatin(sz []byte, szBuf []byte, cchBuf int) []byte { panic("!implemented") }

func CopyRgb(pbSrc []byte, pbDst []byte, cb int) { panic("!implemented") }

func CopyRgchToSz(pch []byte, cch int, sz []byte, cchMax int) { panic("!implemented") }

func CwchSz(sz []byte) int { panic("!implemented") }

func DayInMonth(month int, year int) int { panic("!implemented") }

func DayOfWeek(month int, day int, year int) int { panic("!implemented") }

func DaysInMonth(month int, year int) int { panic("!implemented") }

func DeallocateP(pv any) { panic("!implemented") }

func DecToDeg(d float64) float64 { panic("!implemented") }

func DegToDec(d float64) float64 { panic("!implemented") }

func Dignify(obj int, sign int) []byte { panic("!implemented") }

// Dvd returns integer division, like the "/" operator but always rounds result down.
//
// TODO: add test
func Dvd(x, y int) int {
	if y == 0 {
		return x
	} else if z := x / y; ((x >= 0) == (y >= 0)) || x-z*y == 0 {
		return z
	} else {
		return z - 1
	}
}

func EnsureRay() { panic("!implemented") }

func EnsureStarBright() { panic("!implemented") }

func ErrorArgv(szOpt []byte) { panic("!implemented") }

func ErrorEphem(sz []byte, l int) { panic("!implemented") }

func ErrorSwitch(szOpt []byte) { panic("!implemented") }

func FAppendCIList(pci *CI) bool { panic("!implemented") }

func FEnumerateCIList(nListAll int) bool { panic("!implemented") }

func FEqRgch(rgch1 []byte, rgch2 []byte, cch int, fInsensitive bool) bool { panic("!implemented") }

func FEqSzSubI(sz1 []byte, sz2 []byte) bool { panic("!implemented") }

func FErrorArgc(szOpt []byte, carg int, cargMax int) bool { panic("!implemented") }

func FErrorSubswitch(szOpt []byte, chSub byte, f bool) bool { panic("!implemented") }

func FErrorValN(szOpt []byte, f bool, nVal int, nPar int) bool { panic("!implemented") }

func FErrorValR(szOpt []byte, f bool, rVal float64, nPar int) bool { panic("!implemented") }

func FMatchSz(sz1 []byte, sz2 []byte) bool { panic("!implemented") }

func FSortCIList(nMethod int) bool { panic("!implemented") }

func FilterCIList(szName []byte, szLocation []byte) { panic("!implemented") }

func FormatR(sz []byte, r float64, n int) { panic("!implemented") }

func GetOrb(obj1 int, obj2 int, asp int) float64 { panic("!implemented") }

func GetTimeNow(mon *int, day *int, yea *int, tim *real, dst float64, zon float64) {
	panic("!implemented")
}

func KvBlend(kv1 KV, kv2 KV, rRatio float64) KV { panic("!implemented") }

func KvHue(deg float64) KV { panic("!implemented") }

func KvHue2(deg float64) KV { panic("!implemented") }

func Midpoint(deg1 float64, deg2 float64) float64 { panic("!implemented") }

func MinDifference(deg1 float64, deg2 float64) float64 { panic("!implemented") }

func MinDistance(deg1 float64, deg2 float64) float64 { panic("!implemented") }

// Mod is the modulus function for floating point values, in which we bring the given parameter to within the range of 0 to 360.
//
// TODO: add test
func Mod(d float64) float64 {
	// In most cases, value is only slightly out of range, so can test for it and avoid the more complicated arithmetic.
	if d < 0.0 {
		d += rDegMax
	} else if rDegMax < d {
		d -= rDegMax
	}
	if 0 <= d && d < rDegMax {
		return d
	}
	return d - math.Floor(d/rDegMax)*rDegMax
}

func Mod12(i int) int { panic("!implemented") }

func NCompareSz(sz1 []byte, sz2 []byte) int { panic("!implemented") }

func NCompareSzI(sz1 []byte, sz2 []byte) int { panic("!implemented") }

func ObjMoons(i int) int { panic("!implemented") }

func ObjOrbit(obj int) int { panic("!implemented") }

func PAllocate(cb int, szType []byte) pbyte { panic("!implemented") }

func PrintCh(ch byte) { panic("!implemented") }

func PrintCh2(ch byte) { panic("!implemented") }

func PrintError(sz []byte) { panic("!implemented") }

func PrintNotice(sz []byte) { panic("!implemented") }

func PrintProgress(sz []byte) { panic("!implemented") }

func PrintSz(sz []byte) { panic("!implemented") }

func PrintSzFormat(sz []byte, fPopup bool) { panic("!implemented") }

func PrintSzScreen(sz []byte) { panic("!implemented") }

func PrintTab(ch byte, cch int) { panic("!implemented") }

func PrintTab2(ch byte, cch int) { panic("!implemented") }

func PrintWarning(sz []byte) { panic("!implemented") }

func PrintZodiac(deg float64) { panic("!implemented") }

// RAngle returns the angle formed by a line from the origin to the x, y coordinate.
// This is just converting from rectangular to polar coordinates, however this doesn't involve the radius here.
//
// TODO: add test
func RAngle(x, y float64) float64 {
	var a float64
	if x != 0.0 {
		if y != 0.0 {
			a = RAtn(y / x)
		} else if x < 0.0 {
			a = rPi
		} else {
			a = 0.0
		}
	} else if y < 0.0 {
		a = -rPiHalf
	} else {
		a = rPiHalf
	}
	if a < 0.0 {
		a += rPi
	}
	if y < 0.0 {
		a += rPi
	}
	return a
}

func RObjDiam(obj int) float64 { panic("!implemented") }

// RSgn returns the sign of a number:
//
//	-1 if value negative
//	+1 if value positive,
//	 0 if it's zero.
func RSgn(r float64) float64 {
	if r < 0 {
		return -1
	} else if r == 0 {
		return 0
	}
	return 1
}

func RedoRestrictions() { panic("!implemented") }

func RgReallocate(rgElem any, cElem int, cbElem int, cElemNew int, missing ...any) pbyte {
	panic("!implemented")
}

func SetCentric(obj int) { panic("!implemented") }

func SphDistance(lon1, lat1, lon2, lat2 float64) float64 { panic("!implemented") }

func SphRatio(lon1, lat1, lon2, lat2, rRatio float64, missing ...any) { panic("!implemented") }

func SwapR(d1, d2 *real) { panic("!implemented") }

func SzAltitude(deg float64) []byte { panic("!implemented") }

func SzAspect(asp int) []byte { panic("!implemented") }

func SzAspectAbbrev(asp int) []byte { panic("!implemented") }

func SzCopy(szSrc []byte) []byte { panic("!implemented") }

func SzDate(mon int, day int, yea int, nFormat int) []byte { panic("!implemented") }

func SzDegree(deg float64) []byte { panic("!implemented") }

func SzDegree2(deg float64) []byte { panic("!implemented") }

func SzElevation(elv float64) []byte { panic("!implemented") }

func SzHMS(sec int) []byte { panic("!implemented") }

func SzInList(sz1 []byte, sz2 []byte, pisz *int) []byte { panic("!implemented") }

func SzLength(len float64) []byte { panic("!implemented") }

func SzLocation(lon float64, lat float64) []byte { panic("!implemented") }

func SzLookup(rgStrLook *StrLook, sz []byte) int { panic("!implemented") }

func SzPersist(szSrc []byte) []byte { panic("!implemented") }

func SzProcessProgname(szPath []byte) []byte { panic("!implemented") }

func SzTemperature(tmp float64) []byte { panic("!implemented") }

func SzTim(tim float64) []byte { panic("!implemented") }

func SzTime(hr int, min int, sec int) []byte { panic("!implemented") }

func SzZodiac(deg float64) []byte { panic("!implemented") }

func SzZone(zon float64) []byte { panic("!implemented") }

func Terminate(tc int) { panic("!implemented") }

func UTF8ToIBMSz(sz []byte) { panic("!implemented") }

func UTF8ToLatinSz(sz []byte) { panic("!implemented") }

func UTF8ToWch(pch *uchar, pwch *wchar) int { panic("!implemented") }

func VAngle(v1 *PT3R, v2 *PT3R) float64 { panic("!implemented") }

func WchFromChIBM(ch uchar) wchar { panic("!implemented") }

func WchFromChLatin(ch uchar) wchar { panic("!implemented") }

func WchToUTF8(wch wchar, sz []byte) int { panic("!implemented") }
