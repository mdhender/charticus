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

// xdevice.cpp

func BeginFileX() flag { panic("!implemented") }

func BmGetXY(x int, y int) KI { panic("!implemented") }

func BmpCopyBlock(bs *Bitmap, x1 int, y1 int, x2 int, y2 int, missing ...any) { panic("!implemented") }

func BmpCopyBlock2(bs *Bitmap, x1, y1, x2, y2 real, missing ...any) { panic("!implemented") }

func BmpCopyWin(b *Bitmap, hdc HDC, x int, y int) { panic("!implemented") }

func BmpGetXY(b *Bitmap, x int, y int) KV { panic("!implemented") }

func BmpSetAll(b *Bitmap, kv KV) { panic("!implemented") }

func BmpSetXY(b *Bitmap, x int, y int, kv KV) { panic("!implemented") }

func EndFileX() { panic("!implemented") }

func FAllocateBmp(b *Bitmap, x int, y int) flag { panic("!implemented") }

func FBmpDrawBack(bDest *Bitmap) flag { panic("!implemented") }

func FBmpDrawMap() flag { panic("!implemented") }

func FBmpDrawMap2(x1 int, y1 int, x2 int, y2 int, missing ...any) flag { panic("!implemented") }

func FLoadBmp(szFile []byte, bmp *Bitmap, fNoHeader flag) flag { panic("!implemented") }

func FReadBmp(b *Bitmap, file *FILE, fNoHeader flag) flag { panic("!implemented") }

func GetXY(x int, y int) KI { panic("!implemented") }

func MetaInit() { panic("!implemented") }

func MetaLong(l long) { panic("!implemented") }

func MetaSelect() { panic("!implemented") }

func MetaSz(sz []byte) { panic("!implemented") }

func MetaWord(w word) { panic("!implemented") }

func PsBegin() { panic("!implemented") }

func PsDash(dashoff int) { panic("!implemented") }

func PsEnd() { panic("!implemented") }

func PsFont(nFont int) { panic("!implemented") }

func PsLineCap(fLineCap flag) { panic("!implemented") }

func PsLineWidth(linewidth int) { panic("!implemented") }

func PsStroke(n int) { panic("!implemented") }

func PsStrokeForce() { panic("!implemented") }

func SetXY(x int, y int, ki KI) { panic("!implemented") }

func WireChartOrbit() { panic("!implemented") }

func WireChartSphere() { panic("!implemented") }

func WireCircle(x int, y int, z int, r real, tilt real, rot real) { panic("!implemented") }

func WireDrawGlobe(fSky flag, deg real) { panic("!implemented") }

func WireGlobeCalc(x1 real, y1 real, u *int, v *int, w *int, rz int, deg real) { panic("!implemented") }

func WireLine(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int) { panic("!implemented") }

func WireMapCalc(x1 real, y1 real, xp *int, yp *int, zp *int, fSky flag, missing ...any) {
	panic("!implemented")
}

func WireNum(n int) { panic("!implemented") }

func WireOctahedron(x int, y int, z int, r int) { panic("!implemented") }

func WireSphere(x int, y int, z int, r int) { panic("!implemented") }

func WireSphereEarth(azi real, alt real, zr int, xp *int, yp *int, zp *int) { panic("!implemented") }

func WireSphereLocal(azi real, alt real, zr int, xp *int, yp *int, zp *int) { panic("!implemented") }

func WireSphereMeridian(azi real, alt real, zr int, xp *int, yp *int, zp *int) { panic("!implemented") }

func WireSpherePrime(azi real, alt real, zr int, xp *int, yp *int, zp *int) { panic("!implemented") }

func WireSphereZodiac(lon real, lat real, zr int, xp *int, yp *int, zp *int) { panic("!implemented") }

func WireSpot(x int, y int, z int) { panic("!implemented") }

func WireStar(x int, y int, z int, pes *ES) { panic("!implemented") }

func WriteAscii(file *FILE) { panic("!implemented") }

func WriteBmp(file *FILE) { panic("!implemented") }

func WriteBmp2(b *Bitmap, file *FILE) { panic("!implemented") }

func WriteMeta(file *FILE) { panic("!implemented") }

func WriteWire(file *FILE) { panic("!implemented") }

func WriteXBitmap(file *FILE, name []byte, mode byte) { panic("!implemented") }

func _GetP(pb []byte) KV { panic("!implemented") }

func _GetRGB(pb []byte, r *int, g *int, b *int) { panic("!implemented") }

func _GetXY(b *Bitmap, x int, y int) KV { panic("!implemented") }

func _IbXY(b *Bitmap, x int, y int) long { panic("!implemented") }

func _PbXY(b *Bitmap, x int, y int) []byte { panic("!implemented") }

func _SetRGB(pb []byte, r int, g int, b int) { panic("!implemented") }
