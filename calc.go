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

func CastChart(nContext int) float64 { panic("!implemented") }

func CastSectors() { panic("!implemented") }

func ComputeEphem(t float64) { panic("!implemented") }

func ComputeHouses(housesystem int) { panic("!implemented") }

func ComputeInHouses() { panic("!implemented") }

func ComputeStars(t float64, Off float64) { panic("!implemented") }

func CoorXform(azi *float64, alt *float64, tilt float64) { panic("!implemented") }

func CoorXformFast(azi *float64, alt *float64, sinazi float64, cosazi float64, missing ...any) {
	panic("!implemented")
}

func CreateElemTable(pet *ET) { panic("!implemented") }

func Decan(deg float64) float64 { panic("!implemented") }

func Dwad(deg float64) float64 { panic("!implemented") }

func FAcceptAspect(obj1 int, asp int, obj2 int) flag { panic("!implemented") }

func FCreateGrid(fFlip flag) flag { panic("!implemented") }

func FCreateGridRelation(fMidpoint flag) flag { panic("!implemented") }

func FEnsureGrid() flag { panic("!implemented") }

func FSwissPlanet(ind int, jd float64, indCent int, missing ...any) flag { panic("!implemented") }

func FSwissPlanetData(jd float64, ind int, rPhase *float64, rDiam *float64, rMag *float64) flag {
	panic("!implemented")
}

func GetAspect(planet1 *float64, planet2 *float64, missing ...any) int { panic("!implemented") }

func GetDistance(space1 *PT3R, space2 *PT3R, missing ...any) int { panic("!implemented") }

func GetParallel(planet1 *float64, planet2 *float64, missing ...any) int { panic("!implemented") }

func HouseAlcabitius() { panic("!implemented") }

func HouseEqualGeneric(rOffset float64) { panic("!implemented") }

func HousePorphyry(Asc float64) { panic("!implemented") }

func HousePullenSinusoidalDelta(Asc float64) { panic("!implemented") }

func HousePullenSinusoidalRatio(Asc float64) { panic("!implemented") }

func HouseSripati() { panic("!implemented") }

func JulianToMdy(float64 JD, mon *int, day *int, yea *int) { panic("!implemented") }

func MdyToJulian(mon int, day int, yea int) long { panic("!implemented") }

func MdytszToJulian(mon int, day int, yea int, tim float64, dst float64, zon float64) float64 {
	panic("!implemented")
}

func NCheckEclipse(obj1 int, obj2 int, prPct *float64) int { panic("!implemented") }

func NCheckEclipseAny(obj1 int, asp int, obj2 int, prEclipse *float64) int { panic("!implemented") }

func NCheckEclipseLunar(iEar int, iMoo int, iSun int, prPct *float64) int { panic("!implemented") }

func NCheckEclipseSolar(iEar int, iMoo int, iSun int, prPct *float64) int { panic("!implemented") }

func NHousePlaceIn(rLon float64, rLat float64) int { panic("!implemented") }

func Navamsa(deg float64) float64 { panic("!implemented") }

func ObjTerm(pos float64, nType int) int { panic("!implemented") }

func ProcessPlanet(ind int, aber float64) { panic("!implemented") }

func RHousePlaceIn3D(rLon float64, rLat float64) float64 { panic("!implemented") }

func RHousePlaceIn3DCore(rLon float64, rLat float64) float64 { panic("!implemented") }

// RecToPol transform rectangular coordinates in x, y to polar coordinates.
//
// TODO: add test
// TODO: convert to return values rather than receive pointers to update
func RecToPol(x, y float64, a, r *float64) {
	*r = RLength2(x, y)
	*a = RAngle(x, y)
}

func RecToSph3(rx float64, ry float64, rz float64, azi *float64, alt *float64) { panic("!implemented") }

func SortPlanets() { panic("!implemented") }

func SphToRec(r float64, azi float64, alt float64, rx *float64, ry *float64, rz *float64) {
	panic("!implemented")
}

func SwissComputeAsteroid(jd float64, pes *ES, fBack flag) flag { panic("!implemented") }

func SwissComputeStar(jd float64, pes *ES) flag { panic("!implemented") }

func SwissComputeStars(jd float64, fInitBright flag) { panic("!implemented") }

func SwissEnsurePath() { panic("!implemented") }

func SwissGetFileData(jt1 *float64, jt2 *float64) { panic("!implemented") }

func SwissGetObjName(sz []byte, iobj int) { panic("!implemented") }

func SwissHouse(jd float64, lon float64, lat float64, housesystem int, asc *float64, missing ...any) {
	panic("!implemented")
}

func SwissJulDay(month int, day int, year int, hour float64, gregflag int) float64 {
	panic("!implemented")
}

func SwissLatLmt(jd float64) float64 { panic("!implemented") }

func SwissRefract(rAlt float64) float64 { panic("!implemented") }

func SwissRevJul(jd float64, gregflag int, missing ...any) { panic("!implemented") }

func SwissTestStar(sz []byte) flag { panic("!implemented") }
