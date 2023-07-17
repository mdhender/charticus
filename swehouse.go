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

// swehouse.cpp

func Asc1(x1 double, f double, sine double, cose double) double { panic("!implemented") }

func Asc2(x double, f double, sine double, cose double) double { panic("!implemented") }

func AscDash(x double, f double, sine double, cose double) double { panic("!implemented") }

func CalcH(th double, fi double, ekl double, hsy byte, hsp *houses) int { panic("!implemented") }

func apc_sector(n int, ph double, e double, az double) double { panic("!implemented") }

func armc_to_mc(armc double, eps double) double { panic("!implemented") }

func fix_asc_polar(asc double, armc double, eps double, geolat double) double { panic("!implemented") }

func sidereal_houses_ecl_t0(tjde double, missing ...any) int { panic("!implemented") }

func sidereal_houses_ssypl(tjde double, missing ...any) int { panic("!implemented") }

func sidereal_houses_trad(tjde double, missing ...any) int { panic("!implemented") }

func sunshine_init(lat double, dec double, xh []double) int { panic("!implemented") }

func sunshine_solution_makransky(ramc double, lat double, ecl double, hsp *houses) int {
	panic("!implemented")
}

func sunshine_solution_treindl(ramc double, lat double, ecl double, hsp *houses) int {
	panic("!implemented")
}

func swe_house_name(hsys int) []byte { panic("!implemented") }

func swe_house_pos(...any) double { panic("!implemented") }

func swe_houses(tjd_ut double, missing ...any) int { panic("!implemented") }

func swe_houses_armc(...any) int { panic("!implemented") }

func swe_houses_armc_ex2(...any) int { panic("!implemented") }

func swe_houses_ex(tjd_ut double, missing ...any) int { panic("!implemented") }

func swe_houses_ex2(tjd_ut double, missing ...any) int { panic("!implemented") }

func swi_armc_to_mc(armc double, eps double) double { panic("!implemented") }

func test_Asc1() { panic("!implemented") }
