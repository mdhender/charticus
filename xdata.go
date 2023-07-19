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

// Graphics Global Variables.

var (
	gs = GS{
		ft:          ftBmp,
		fPSComplete: true,
		fColor:      true,
		fText:       true,
		fBorder:     true,
		fLabel:      true,
		fKeepSquare: true,
		xWin:        DEFAULTX,
		yWin:        DEFAULTY,
		nScale:      200,
		nScaleText:  100,
		nAstLabel:   3,
		nLabelCity:  1,
		objTrack:    oMoo,
		chBmpMode:   BITMAPMODE,
		rBackPct:    25.0,
		nBackOrient: 1,
		xInch:       8.5,
		yInch:       11.0,
		nDecaSize:   25,
		nDecaLine:   11,
		nDecaFill:   1,
		nGridCell:   oCore,
		cspace:      1000,
		nRayWidth:   600,
		nGlyphCap:   1,
		nGlyphUra:   1,
		nGlyphPlu:   1,
		nGlyphLil:   2,
		nGlyphVer:   2,
		nGlyphEri:   1,
		fAltPalette: true,
		nDashMax:    7,
	}
)
