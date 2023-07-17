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

// swecl.cpp

func calc_dip(geoalt double, atpress double, attemp double, lapse_rate double) double {
	panic("!implemented")
}

func calc_mer_trans(...any) int32 { panic("!implemented") }

func calc_planet_star(tjd_et double, ipl int32, starname []byte, iflag int32, x *double, serr []byte) int32 {
	panic("!implemented")
}

func eclipse_how(tjd_ut double, ipl int32, starname []byte, ifl int32, missing ...any) int32 {
	panic("!implemented")
}

func eclipse_when_loc(tjd_start double, ifl int32, geopos *double, tret *double, attr *double, backward int32, serr []byte) int32 {
	panic("!implemented")
}

func eclipse_where(tjd_ut double, ipl int32, starname []byte, ifl int32, geopos *double, dcore *double, missing ...any) int32 {
	panic("!implemented")
}

func find_maximum(y00 double, y11 double, y2 double, dx double, missing ...any) int {
	panic("!implemented")
}

func find_zero(y00 double, y11 double, y2 double, dx double, missing ...any) int {
	panic("!implemented")
}

func get_dist_from_2_vectors(x1 *double, x2 *double) double { panic("!implemented") }

func get_gmsm(tjd_et double, ipl int32, iflag int32, r double, gmsm *double, serr []byte) int32 {
	panic("!implemented")
}

func get_sun_rad_plus_refr(ipl int32, dd double, rsmi int32, refr double) double {
	panic("!implemented")
}

func lun_eclipse_how(tjd_ut double, ifl int32, attr *double, missing ...any) int32 {
	panic("!implemented")
}

func occult_when_loc(tjd_start double, ipl int32, starname []byte, ifl int32, missing ...any) int32 {
	panic("!implemented")
}

func orbit_max_min_true_distance_helio(tjd_et double, ipl int, iflag int32, dmax *double, dmin *double, dtrue *double, missing ...any) int32 {
	panic("!implemented")
}

func osc_get_ecl_pos(ean double, pqr *double, xp *double) { panic("!implemented") }

func osc_get_orbit_constants(dp *double, pqr *double) { panic("!implemented") }

func osc_iterate_max_dist(ean double, pqr *double, xa *double, xb *double, deanopt *double, drmax *double, missing ...any) {
	panic("!implemented")
}

func osc_iterate_min_dist(ean double, pqr *double, xa *double, xb *double, deanopt *double, drmin *double, missing ...any) {
	panic("!implemented")
}

func rdi_twilight(rsmi int32) double { panic("!implemented") }

func rise_set_fast(...any) int32 { panic("!implemented") }

func swe_azalt(...any) { panic("!implemented") }

func swe_azalt_rev(...any) { panic("!implemented") }

func swe_gauquelin_sector(...any) int32 { panic("!implemented") }

func swe_get_orbital_elements(...any) int32 { panic("!implemented") }

func swe_lun_eclipse_how(...any) int32 { panic("!implemented") }

func swe_lun_eclipse_when(tjd_start double, ifl int32, ifltype int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_lun_eclipse_when_loc(tjd_start double, ifl int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_lun_occult_when_glob(...any) int32 { panic("!implemented") }

func swe_lun_occult_when_loc(tjd_start double, ipl int32, starname []byte, ifl int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_lun_occult_where(...any) int32 { panic("!implemented") }

func swe_nod_aps(tjd_et double, ipl int32, iflag int32, missing ...any) int32 { panic("!implemented") }

func swe_nod_aps_ut(tjd_ut double, ipl int32, iflag int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_orbit_max_min_true_distance(tjd_et double, ipl int32, iflag int32, dmax *double, dmin *double, dtrue *double, missing ...any) int32 {
	panic("!implemented")
}

func swe_pheno(tjd double, ipl int32, iflag int32, attr *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_pheno_ut(tjd_ut double, ipl int32, iflag int32, attr *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_refrac(inalt double, atpress double, attemp double, calc_flag int32) double {
	panic("!implemented")
}

func swe_refrac_extended(inalt double, geoalt double, atpress double, attemp double, lapse_rate double, calc_flag int32, missing ...any) double {
	panic("!implemented")
}

func swe_rise_trans(...any) int32 { panic("!implemented") }

func swe_rise_trans_true_hor(...any) int32 { panic("!implemented") }

func swe_set_lapse_rate(lapse_rate double) { panic("!implemented") }

func swe_sol_eclipse_how(...any) int32 { panic("!implemented") }

func swe_sol_eclipse_when_glob(tjd_start double, ifl int32, ifltype int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_sol_eclipse_when_loc(tjd_start double, ifl int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_sol_eclipse_where(...any) int32 { panic("!implemented") }
