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

// xcharts0.cpp

func DrawAspectLine(obj1 int, obj2 int, cx int, cy int, missing ...any) { panic("!implemented") }

func DrawChartX() { panic("!implemented") }

func DrawFillWheel(x int, y int, i int, typ int) flag { panic("!implemented") }

func DrawInfo(pci *CI, szHeader []byte, fAll flag) { panic("!implemented") }

func DrawLeyLine(l1, f1, l2, f2 real) { panic("!implemented") }

func DrawLeyLines(deg real) { panic("!implemented") }

func DrawMap(fSky flag, fGlobe flag, deg real) { panic("!implemented") }

func DrawMapLine(lon1, lat1, lon2, lat2 real, missing ...any) { panic("!implemented") }

func DrawMapSquare(lon1, lat1, lon2, lat2 real, missing ...any) { panic("!implemented") }

func DrawMapTriangle(lon1, lat1, lon2, lat2, lon3 real, missing ...any) { panic("!implemented") }

func DrawMapTriangles(fGlobe flag, nScl int, pcr *CIRC, deg real) { panic("!implemented") }

func DrawObjects(rgod *ObjDraw, cod int, zEdge int) { panic("!implemented") }

func DrawPrint(sz []byte, m int, n int) int { panic("!implemented") }

func DrawRing(iRing int, iRingMax int, missing ...any) { panic("!implemented") }

func DrawSidebar() { panic("!implemented") }

func DrawSymbolRing(real *symbol, xplanet *real, dir *real, cx int, cy int, missing ...any) {
	panic("!implemented")
}

func DrawWheel(real *xsign, xhouse *real, cx int, cy int, unitx real, missing ...any) {
	panic("!implemented")
}

func DrawZodiac(deg real, n int) int { panic("!implemented") }

func EnumConstelLines(x1 *int, y1 *int, x2 *int, y2 *int, iConst *int) flag { panic("!implemented") }

func EnumStarsLines(fInit flag, ppes1, ppes2 **ES) flag { panic("!implemented") }

func EnumWorldLines(x1 *int, y1 *int, x2 *int, y2 *int, kRainbow *int) flag { panic("!implemented") }

func FGlobeCalc(x1 real, y1 real, u *int, v *int, pcr *CIRC, deg real) flag { panic("!implemented") }

func FMapCalc(x1 real, y1 real, xp *int, yp *int, fGlobe flag, fSky flag, missing ...any) flag {
	panic("!implemented")
}

func FReadWorldData(nam *[]byte, loc *[]byte, lin *[]byte) flag { panic("!implemented") }
