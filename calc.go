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

// calc.cpp

func CastChart(nContext int) real { panic("!implemented") }

func CastSectors() { panic("!implemented") }

func ComputeEphem(t real) { panic("!implemented") }

func ComputeHouses(housesystem int) { panic("!implemented") }

func ComputeInHouses() { panic("!implemented") }

func ComputeStars(t real, Off real) { panic("!implemented") }

func CoorXform(azi *real, alt *real, tilt real) { panic("!implemented") }

func CoorXformFast(azi *real, alt *real, sinazi real, cosazi real, missing ...any) {
	panic("!implemented")
}

func CreateElemTable(pet *ET) { panic("!implemented") }

func Decan(deg real) real { panic("!implemented") }

func Dwad(deg real) real { panic("!implemented") }

func FAcceptAspect(obj1 int, asp int, obj2 int) flag { panic("!implemented") }

func FCreateGrid(fFlip flag) flag { panic("!implemented") }

func FCreateGridRelation(fMidpoint flag) flag { panic("!implemented") }

func FEnsureGrid() flag { panic("!implemented") }

func FSwissPlanet(ind int, jd real, indCent int, missing ...any) flag { panic("!implemented") }

func FSwissPlanetData(jd real, ind int, rPhase *real, rDiam *real, rMag *real) flag {
	panic("!implemented")
}

func GetAspect(planet1 *real, planet2 *real, missing ...any) int { panic("!implemented") }

func GetDistance(space1 *PT3R, space2 *PT3R, missing ...any) int { panic("!implemented") }

func GetParallel(planet1 *real, planet2 *real, missing ...any) int { panic("!implemented") }

func HouseAlcabitius() { panic("!implemented") }

func HouseEqualGeneric(rOffset real) { panic("!implemented") }

func HousePorphyry(Asc real) { panic("!implemented") }

func HousePullenSinusoidalDelta(Asc real) { panic("!implemented") }

func HousePullenSinusoidalRatio(Asc real) { panic("!implemented") }

func HouseSripati() { panic("!implemented") }

func JulianToMdy(real JD, mon *int, day *int, yea *int) { panic("!implemented") }

func MdyToJulian(mon int, day int, yea int) long { panic("!implemented") }

func MdytszToJulian(mon int, day int, yea int, tim real, dst real, zon real) real {
	panic("!implemented")
}

func NCheckEclipse(obj1 int, obj2 int, prPct *real) int { panic("!implemented") }

func NCheckEclipseAny(obj1 int, asp int, obj2 int, prEclipse *real) int { panic("!implemented") }

func NCheckEclipseLunar(iEar int, iMoo int, iSun int, prPct *real) int { panic("!implemented") }

func NCheckEclipseSolar(iEar int, iMoo int, iSun int, prPct *real) int { panic("!implemented") }

func NHousePlaceIn(rLon real, rLat real) int { panic("!implemented") }

func Navamsa(deg real) real { panic("!implemented") }

func ObjTerm(pos real, nType int) int { panic("!implemented") }

func ProcessPlanet(ind int, aber real) { panic("!implemented") }

func RHousePlaceIn3D(rLon real, rLat real) real { panic("!implemented") }

func RHousePlaceIn3DCore(rLon real, rLat real) real { panic("!implemented") }

func RecToPol(x real, y real, a *real, r *real) { panic("!implemented") }

func RecToSph3(rx real, ry real, rz real, azi *real, alt *real) { panic("!implemented") }

func SortPlanets() { panic("!implemented") }

func SphToRec(r real, azi real, alt real, rx *real, ry *real, rz *real) { panic("!implemented") }

func SwissComputeAsteroid(jd real, pes *ES, fBack flag) flag { panic("!implemented") }

func SwissComputeStar(jd real, pes *ES) flag { panic("!implemented") }

func SwissComputeStars(jd real, fInitBright flag) { panic("!implemented") }

func SwissEnsurePath() { panic("!implemented") }

func SwissGetFileData(jt1 *real, jt2 *real) { panic("!implemented") }

func SwissGetObjName(sz []byte, iobj int) { panic("!implemented") }

func SwissHouse(jd real, lon real, lat real, housesystem int, asc *real, missing ...any) {
	panic("!implemented")
}

func SwissJulDay(month int, day int, year int, hour real, gregflag int) real { panic("!implemented") }

func SwissLatLmt(jd real) real { panic("!implemented") }

func SwissRefract(rAlt real) real { panic("!implemented") }

func SwissRevJul(jd real, gregflag int, missing ...any) { panic("!implemented") }

func SwissTestStar(sz []byte) flag { panic("!implemented") }
