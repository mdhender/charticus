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

// DEFAULT_LOC
const (
	DEFAULT_LOC_LON = 122.0 + 19.0/60.0 + 55.0/3600.0
	DEFAULT_LOC_LAT = 47.0 + 36.0/60.0 + 22.0/3600.0
)

// TODO: add test
func IoeFromObj(obj _object) int {
	if obj < oMoo {
		return 0
	} else if obj <= cPlanet {
		return int(obj) - 2
	}
	return int(obj) - uranLo + int(cPlanet) - 2
}

// TODO: add test
func JulianDayFromTime(t float64) float64 {
	return t*36525.0 + 2415020.0
}

func PtAdd2(pt, pt2 PT3R) PT3R {
	return PT3R{x: pt.x + pt2.x, y: pt.y + pt2.y, z: pt.z + pt2.z}
}
func PtCross(pt, pt1, pt2 PT3R) PT3R {
	return PT3R{
		x: pt1.y*pt2.z - pt1.z*pt2.y,
		y: pt1.z*pt2.x - pt1.x*pt2.z,
		z: pt1.x*pt2.y - pt1.y*pt2.x,
	}
}
func PtDiv(pt PT3R, r float64) PT3R {
	return PT3R{
		x: pt.x / r,
		y: pt.y / r,
		z: pt.z / r,
	}
}
func PtDot(pt1, pt2 PT3R) float64 {
	return pt1.x*pt2.x + pt1.y*pt2.y + pt1.z*pt2.z
}
func PtLen(pt PT3R) float64 {
	return RLength3(pt.x, pt.y, pt.z)
}
func PtMul(pt PT3R, r float64) PT3R {
	return PT3R{
		x: pt.x * r,
		y: pt.y * r,
		z: pt.z * r,
	}
}
func PtNeg2(pt, pt2 PT3R) PT3R {
	return PT3R{
		x: -pt2.x,
		y: -pt2.y,
		z: -pt2.z,
	}
}
func PtSet(pt PT3R, x, y, z float64) PT3R {
	return PT3R{x: x, y: y, z: z}
}
func PtSub2(pt, pt2 PT3R) PT3R {
	return PT3R{
		x: pt.x - pt2.x,
		y: pt.y - pt2.y,
		z: pt.z - pt2.z,
	}
}
func PtVec(pt, pt1, pt2 PT3R) PT3R {
	return PtSub2(pt2, pt1)
}
func PtZero(pt PT3R) PT3R {
	return PT3R{}
}

func MM() int     { return ciCore.mon }
func DD() int     { return ciCore.day }
func YY() int     { return ciCore.yea }
func TT() float64 { return ciCore.tim }
func ZZ() float64 { return ciCore.zon }
func SS() float64 { return ciCore.dst }
func OO() float64 { return ciCore.lon }
func AA() float64 { return ciCore.lat }
