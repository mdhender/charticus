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

// xgeneral.cpp

func AdjustGlyph(ch *int, x *int, y *int, fi *int, nScale *int, missing ...any) {
	panic("!implemented")
}

func ClipGreater(x1 *int, y1 *int, x2 *int, y2 *int, s int) { panic("!implemented") }

func ClipLesser(x1 *int, y1 *int, x2 *int, y2 *int, s int) { panic("!implemented") }

func DrawArc(x1 int, y1 int, x2 int, y2 int, rRotate real, t1 real, t2 real) { panic("!implemented") }

func DrawAspect(asp int, x int, y int) { panic("!implemented") }

func DrawBlock(x1 int, y1 int, x2 int, y2 int) { panic("!implemented") }

func DrawBox(x1 int, y1 int, x2 int, y2 int, xsiz int, ysiz int) { panic("!implemented") }

func DrawClearScreen() { panic("!implemented") }

func DrawColor(col KI) { panic("!implemented") }

func DrawCrescent(x1 int, y1 int, x2 int, y2 int, rProp real, rRotate real, missing ...any) {
	panic("!implemented")
}

func DrawDash(x1 int, y1 int, x2 int, y2 int, skip int) { panic("!implemented") }

func DrawEllipse2(x1 int, y1 int, x2 int, y2 int) { panic("!implemented") }

func DrawFill(x int, y int, kv KV) { panic("!implemented") }

func DrawHouse(i int, x int, y int) { panic("!implemented") }

func DrawNakshatra(i int, x int, y int) { panic("!implemented") }

func DrawObject(obj int, x int, y int) { panic("!implemented") }

func DrawPoint(x int, y int) { panic("!implemented") }

func DrawSign(i int, x int, y int) { panic("!implemented") }

func DrawSpot(x int, y int) { panic("!implemented") }

func DrawStar(x int, y int, pes *ES) { panic("!implemented") }

func DrawSz(sz []byte, x int, y int, dt int) { panic("!implemented") }

func DrawThick(fThick flag) { panic("!implemented") }

func DrawTurtle(sz []byte, x0 int, y0 int) { panic("!implemented") }

func DrawWrap(x1 int, y1 int, x2 int, y2 int, xmin int, xmax int) { panic("!implemented") }

func FDrawClip(x1 int, y1 int, x2 int, y2 int, xl int, yl int, xh int, yh int, missing ...any) flag {
	panic("!implemented")
}

func KiCity(iae int) int { panic("!implemented") }

func NFromPch(str *[]byte) int { panic("!implemented") }

func WinClearScreen(ki KI) { panic("!implemented") }

func WinDrawGlyph(ch int, x int, y int, fi int, nScale int) flag { panic("!implemented") }
