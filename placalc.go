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

// placalc.cpp

func ast_file_posit(jd double, astfpp **FILE) int { panic("!implemented") }

func calc(planet int, jd_ad REAL8, int flag, missing ...any) int { panic("!implemented") }

func chi_file_posit(jd double, lrzfpp **FILE) int { panic("!implemented") }

func disturb(k *kor, al *REAL8, ar *REAL8, lk REAL8, rk REAL8, missing ...any) { panic("!implemented") }

func hel(planet int, t REAL8, al *REAL8, ar *REAL8, az *REAL8, alp *REAL8, missing ...any) int {
	panic("!implemented")
}

func helup(jd_ad REAL8) { panic("!implemented") }

func inpolq(n int, o int, p double, x []double, axu *double, adxu *double) int { panic("!implemented") }

func lrz_file_posit(jd double, lrzfpp **FILE) int { panic("!implemented") }

func moon(al *REAL8, ar *REAL8, az *REAL8) int { panic("!implemented") }

func outer_hel(planet int, jd_ad REAL8, al *REAL8, ar *REAL8, az *REAL8, missing ...any) int {
	panic("!implemented")
}

func togeo(lngearth REAL8, radearth REAL8, lng REAL8, rad REAL8, zet REAL8, missing ...any) {
	panic("!implemented")
}
