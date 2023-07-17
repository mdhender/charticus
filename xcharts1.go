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

// xcharts1.cpp

func DrawArrow(x1 int, y1 int, x2 int, y2 int) { panic("!implemented") }

func DrawCalendar(mon int, yea int, x1 int, y1 int, x2 int, y2 int) { panic("!implemented") }

func DrawCalendarAspect(pid *InDayInfo, i int, iMax int, nVoid int, missing ...any) flag {
	panic("!implemented")
}

func EarToHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func EarToHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func EarToTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func EclToHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func EclToHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func EclToTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func EquToHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func EquToHorizon2(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func EquToHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func EquToHorizonSky2(lon real, lat real, pcr *CIRC, xp *int, yp *int, missing ...any) {
	panic("!implemented")
}

func EquToTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func EquToTelescope2(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func FSphereEarth(azi real, alt real, pcr *CIRC, xp *int, yp *int) flag { panic("!implemented") }

func FSphereLocal(azi real, alt real, pcr *CIRC, xp *int, yp *int) flag { panic("!implemented") }

func FSphereMeridian(azi real, alt real, pcr *CIRC, xp *int, yp *int) flag { panic("!implemented") }

func FSpherePrime(azi real, alt real, pcr *CIRC, xp *int, yp *int) flag { panic("!implemented") }

func FSphereZodiac(lon real, lat real, pcr *CIRC, xp *int, yp *int) flag { panic("!implemented") }

func FormatGridCell(sz []byte, x int, y int, type_ int, fWide flag) KI { panic("!implemented") }

func LocToHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func LocToHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func LocToTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func OrbitPlot(pxp *real, pyp *real, pzp *real, sz real, obj int, rgspc *PT3R) { panic("!implemented") }

func OrbitRecord() { panic("!implemented") }

func PlotHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func PlotHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func PlotTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func PriToHorizon(lon real, lat real, x1 int, y1 int, xs int, ys int, missing ...any) {
	panic("!implemented")
}

func PriToHorizonSky(lon real, lat real, pcr *CIRC, xp *int, yp *int) { panic("!implemented") }

func PriToTelescope(lon real, lat real, pte *TELE, missing ...any) { panic("!implemented") }

func XChartAstroGraph() { panic("!implemented") }

func XChartCalendar() { panic("!implemented") }

func XChartDispositor() { panic("!implemented") }

func XChartGrid(x0, y0 int) { panic("!implemented") }

func XChartHorizon() { panic("!implemented") }

func XChartHorizonSky() { panic("!implemented") }

func XChartIndian() { panic("!implemented") }

func XChartLocal() { panic("!implemented") }

func XChartMidpoint() { panic("!implemented") }

func XChartMoons() { panic("!implemented") }

func XChartOrbit() { panic("!implemented") }

func XChartSector() { panic("!implemented") }

func XChartSphere() { panic("!implemented") }

func XChartTelescope() { panic("!implemented") }

func XChartWheel() { panic("!implemented") }
