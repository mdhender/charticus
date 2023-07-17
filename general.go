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

// general.cpp

func AddDay(month int, day int, year int, delta int) int { panic("!implemented") }

func AddTime(pci *CI, mode int, toadd int) { panic("!implemented") }

func AnsiColor(k int) { panic("!implemented") }

func Assert(f flag) { panic("!implemented") }

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

func DecToDeg(d real) real { panic("!implemented") }

func DegToDec(d real) real { panic("!implemented") }

func Dignify(obj int, sign int) []byte { panic("!implemented") }

func Dvd(x long, y long) long { panic("!implemented") }

func EnsureRay() { panic("!implemented") }

func EnsureStarBright() { panic("!implemented") }

func ErrorArgv(szOpt []byte) { panic("!implemented") }

func ErrorEphem(sz []byte, l long) { panic("!implemented") }

func ErrorSwitch(szOpt []byte) { panic("!implemented") }

func FAppendCIList(pci *CI) flag { panic("!implemented") }

func FEnumerateCIList(nListAll int) flag { panic("!implemented") }

func FEqRgch(rgch1 []byte, rgch2 []byte, cch int, fInsensitive flag) flag { panic("!implemented") }

func FEqSzSubI(sz1 []byte, sz2 []byte) flag { panic("!implemented") }

func FErrorArgc(szOpt []byte, carg int, cargMax int) flag { panic("!implemented") }

func FErrorSubswitch(szOpt []byte, chSub byte, f flag) flag { panic("!implemented") }

func FErrorValN(szOpt []byte, f flag, nVal int, nPar int) flag { panic("!implemented") }

func FErrorValR(szOpt []byte, f flag, rVal real, nPar int) flag { panic("!implemented") }

func FMatchSz(sz1 []byte, sz2 []byte) flag { panic("!implemented") }

func FSortCIList(nMethod int) flag { panic("!implemented") }

func FilterCIList(szName []byte, szLocation []byte) { panic("!implemented") }

func FormatR(sz []byte, r real, n int) { panic("!implemented") }

func GetOrb(obj1 int, obj2 int, asp int) real { panic("!implemented") }

func GetTimeNow(mon *int, day *int, yea *int, tim *real, dst real, zon real) { panic("!implemented") }

func KvBlend(kv1 KV, kv2 KV, rRatio real) KV { panic("!implemented") }

func KvHue(deg real) KV { panic("!implemented") }

func KvHue2(deg real) KV { panic("!implemented") }

func Midpoint(deg1 real, deg2 real) real { panic("!implemented") }

func MinDifference(deg1 real, deg2 real) real { panic("!implemented") }

func MinDistance(deg1 real, deg2 real) real { panic("!implemented") }

func Mod(d real) real { panic("!implemented") }

func Mod12(i int) int { panic("!implemented") }

func NCompareSz(sz1 []byte, sz2 []byte) int { panic("!implemented") }

func NCompareSzI(sz1 []byte, sz2 []byte) int { panic("!implemented") }

func ObjMoons(i int) int { panic("!implemented") }

func ObjOrbit(obj int) int { panic("!implemented") }

func PAllocate(cb long, szType []byte) pbyte { panic("!implemented") }

func PrintCh(ch byte) { panic("!implemented") }

func PrintCh2(ch byte) { panic("!implemented") }

func PrintError(sz []byte) { panic("!implemented") }

func PrintNotice(sz []byte) { panic("!implemented") }

func PrintProgress(sz []byte) { panic("!implemented") }

func PrintSz(sz []byte) { panic("!implemented") }

func PrintSzFormat(sz []byte, fPopup flag) { panic("!implemented") }

func PrintSzScreen(sz []byte) { panic("!implemented") }

func PrintTab(ch byte, cch int) { panic("!implemented") }

func PrintTab2(ch byte, cch int) { panic("!implemented") }

func PrintWarning(sz []byte) { panic("!implemented") }

func PrintZodiac(deg real) { panic("!implemented") }

func RAngle(x real, y real) real { panic("!implemented") }

func RObjDiam(obj int) real { panic("!implemented") }

func RSgn(r real) real { panic("!implemented") }

func RedoRestrictions() { panic("!implemented") }

func RgReallocate(rgElem any, cElem int, cbElem int, cElemNew int, missing ...any) pbyte {
	panic("!implemented")
}

func SetCentric(obj int) { panic("!implemented") }

func SphDistance(lon1, lat1, lon2, lat2 real) real { panic("!implemented") }

func SphRatio(lon1, lat1, lon2, lat2, rRatio real, missing ...any) { panic("!implemented") }

func SwapR(d1, d2 *real) { panic("!implemented") }

func SzAltitude(deg real) []byte { panic("!implemented") }

func SzAspect(asp int) []byte { panic("!implemented") }

func SzAspectAbbrev(asp int) []byte { panic("!implemented") }

func SzCopy(szSrc []byte) []byte { panic("!implemented") }

func SzDate(mon int, day int, yea int, nFormat int) []byte { panic("!implemented") }

func SzDegree(deg real) []byte { panic("!implemented") }

func SzDegree2(deg real) []byte { panic("!implemented") }

func SzElevation(elv real) []byte { panic("!implemented") }

func SzHMS(sec int) []byte { panic("!implemented") }

func SzInList(sz1 []byte, sz2 []byte, pisz *int) []byte { panic("!implemented") }

func SzLength(len real) []byte { panic("!implemented") }

func SzLocation(lon real, lat real) []byte { panic("!implemented") }

func SzLookup(rgStrLook *StrLook, sz []byte) int { panic("!implemented") }

func SzPersist(szSrc []byte) []byte { panic("!implemented") }

func SzProcessProgname(szPath []byte) []byte { panic("!implemented") }

func SzTemperature(tmp real) []byte { panic("!implemented") }

func SzTim(tim real) []byte { panic("!implemented") }

func SzTime(hr int, min int, sec int) []byte { panic("!implemented") }

func SzZodiac(deg real) []byte { panic("!implemented") }

func SzZone(zon real) []byte { panic("!implemented") }

func Terminate(tc int) { panic("!implemented") }

func UTF8ToIBMSz(sz []byte) { panic("!implemented") }

func UTF8ToLatinSz(sz []byte) { panic("!implemented") }

func UTF8ToWch(pch *uchar, pwch *wchar) int { panic("!implemented") }

func VAngle(v1 *PT3R, v2 *PT3R) real { panic("!implemented") }

func WchFromChIBM(ch uchar) wchar { panic("!implemented") }

func WchFromChLatin(ch uchar) wchar { panic("!implemented") }

func WchToUTF8(wch wchar, sz []byte) int { panic("!implemented") }
