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

// sweph.cpp
func aberr_light(xx *double, xe *double) { panic("!implemented") }

func app_pos_etc_mean(ipl int, iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_etc_moon(iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_etc_plan(ipli int, iplmoon int, iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_etc_plan_osc(ipl int, ipli int, iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_etc_sbar(iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_etc_sun(iflag int32, serr []byte) int { panic("!implemented") }

func app_pos_rest(pdp *plan_data, iflag int32, missing ...any) int { panic("!implemented") }

func calc_center_body(ipli int32, iflag int32, xx *double, xcom *double, serr []byte) int {
	panic("!implemented")
}

func calc_epsilon(tjd double, iflag int32, e *epsilon) { panic("!implemented") }

func calc_speed(x0 *double, x1 *double, x2 *double, dt double) { panic("!implemented") }

func denormalize_positions(x0 *double, x1 *double, x2 *double) { panic("!implemented") }

func do_fread(trg any, size int, count int, corrsize int, fp *FILE, fpos int32, freord int, fendian int, ifno int, missing ...any) int {
	panic("!implemented")
}

func embofs(xemb *double, xmoon *double) { panic("!implemented") }

func fixedstar_name_compare(star1 any, star2 any) int { panic("!implemented") }

func fixstar_calc_from_struct(stardata *fixed_star, tjd double, iflag int32, star []byte, xx *double, serr []byte) int32 {
	panic("!implemented")
}

func fixstar_cut_string(srecord []byte, star []byte, stardata *fixed_star, serr []byte) int32 {
	panic("!implemented")
}

func fixstar_format_search_name(star []byte, sstar []byte, serr []byte) int32 { panic("!implemented") }

func free_planets() { panic("!implemented") }

func fstar_node_compare(node1 any, node2 any) int { panic("!implemented") }

func get_aya_correction(iflag int, corr *double, serr []byte) int { panic("!implemented") }

func get_builtin_star(star []byte, sstar []byte, srecord []byte) AS_BOOL { panic("!implemented") }

func get_new_segment(tjd double, ipli int, ifno int, serr []byte) int { panic("!implemented") }

func intp_apsides(tjd double, ipl int, iflag int32, serr []byte) int { panic("!implemented") }

func jplplan(tjd double, ipli int, iflag int32, do_save AS_BOOL, missing ...any) int {
	panic("!implemented")
}

func load_all_fixed_stars(serr []byte) int32 { panic("!implemented") }

func load_dpsi_deps() { panic("!implemented") }

func lunar_osc_elem(tjd double, ipl int, iflag int32, serr []byte) int { panic("!implemented") }

func main_planet(tjd double, ipli int, iplmoon int, epheflag int32, iflag int32, missing ...any) int {
	panic("!implemented")
}

func main_planet_bary(tjd double, ipli int, epheflag int32, iflag int32, do_save AS_BOOL, missing ...any) int {
	panic("!implemented")
}

func meff(r double) double { panic("!implemented") }

func nut_matrix(nu *nut, oec *epsilon) { panic("!implemented") }

func open_jpl_file(ss *double, fname []byte, fpath []byte, serr []byte) int { panic("!implemented") }

func plaus_iflag(iflag int32, ipl int32, tjd double, serr []byte) int32 { panic("!implemented") }

func read_const(ifno int, serr []byte) int { panic("!implemented") }

func rot_back(ipli int) { panic("!implemented") }

func save_star_in_struct(nrecs int, fstp *fixed_star, serr []byte) int32 { panic("!implemented") }

func search_star_in_list(sstar []byte, stardata *fixed_star, serr []byte) int32 {
	panic("!implemented")
}

func swe_calc(tjd double, ipl int, iflag int32, missing ...any) int32 { panic("!implemented") }

func swe_calc_pctr(tjd double, ipl int32, iplctr int32, iflag int32, xxret *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_calc_ut(tjd_ut double, ipl int32, iflag int32, missing ...any) int32 { panic("!implemented") }

func swe_close() { panic("!implemented") }

func swe_fixstar(star []byte, tjd double, iflag int32, missing ...any) int32 { panic("!implemented") }

func swe_fixstar2(star []byte, tjd double, iflag int32, missing ...any) int32 { panic("!implemented") }

func swe_fixstar2_mag(star []byte, mag *double, serr []byte) int32 { panic("!implemented") }

func swe_fixstar2_ut(star []byte, tjd_ut double, iflag int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_fixstar_mag(star []byte, mag *double, serr []byte) int32 { panic("!implemented") }

func swe_fixstar_ut(star []byte, tjd_ut double, iflag int32, missing ...any) int32 {
	panic("!implemented")
}

func swe_get_ayanamsa(tjd_et double) double { panic("!implemented") }

func swe_get_ayanamsa_ex(tjd_et double, iflag int32, daya *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_get_ayanamsa_ex_ut(tjd_ut double, iflag int32, daya *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_get_ayanamsa_name(isidmode int32) []byte { panic("!implemented") }

func swe_get_ayanamsa_ut(tjd_ut double) double { panic("!implemented") }

func swe_get_current_file_data(ifno int, tfstart *double, tfend *double, denum *int) []byte {
	panic("!implemented")
}

func swe_get_library_path(s []byte) []byte { panic("!implemented") }

func swe_get_planet_name(ipl int, s []byte) []byte { panic("!implemented") }

func swe_helio_cross(ipl int, x2cross double, jd_et double, iflag int, dir int, jd_cross *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_helio_cross_ut(ipl int, x2cross double, jd_ut double, iflag int, dir int, jd_cross *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_lat_to_lmt(tjd_lat double, geolon double, tjd_lmt *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_lmt_to_lat(tjd_lmt double, geolon double, tjd_lat *double, serr []byte) int32 {
	panic("!implemented")
}

func swe_mooncross(x2cross double, jd_et double, int flag, serr []byte) double { panic("!implemented") }

func swe_mooncross_node(jd_et double, int flag, xlon *double, xla *double, serr []byte) double {
	panic("!implemented")
}

func swe_mooncross_node_ut(jd_ut double, int flag, xlon *double, xla *double, serr []byte) double {
	panic("!implemented")
}

func swe_mooncross_ut(x2cross double, jd_ut double, int flag, serr []byte) double {
	panic("!implemented")
}

func swe_set_ephe_path(path []byte) { panic("!implemented") }

func swe_set_jpl_file(fname []byte) { panic("!implemented") }

func swe_set_sid_mode(sid_mode int32, t0 double, ayan_t0 double) { panic("!implemented") }

func swe_set_topo(geolon double, geolat double, geoalt double) { panic("!implemented") }

func swe_solcross(x2cross double, jd_et double, int flag, serr []byte) double { panic("!implemented") }

func swe_solcross_ut(x2cross double, jd_ut double, int flag, serr []byte) double {
	panic("!implemented")
}

func swe_time_equ(tjd_ut double, E *double, serr []byte) int32 { panic("!implemented") }

func swe_version(s []byte) []byte { panic("!implemented") }

func swecalc(tjd double, ipl int, iplmoon int32, iflag int32, x *double, serr []byte) int32 {
	panic("!implemented")
}

func swemoon(tjd double, iflag int32, do_save AS_BOOL, xpret *double, serr []byte) int {
	panic("!implemented")
}

func sweph(tjd double, ipli int, ifno int, iflag int32, xsunb *double, do_save AS_BOOL, xpret *double, serr []byte) int {
	panic("!implemented")
}

func sweplan(tjd double, ipli int, ifno int, iflag int32, do_save AS_BOOL, missing ...any) int {
	panic("!implemented")
}

func swi_aberr_light(xx *double, xe *double, iflag int32) { panic("!implemented") }

func swi_aberr_light_ex(xx *double, xe *double, xe_dt *double, dt double, iflag int32) {
	panic("!implemented")
}

func swi_check_ecliptic(tjd double, iflag int32) { panic("!implemented") }

func swi_check_nutation(tjd double, iflag int32) { panic("!implemented") }

func swi_close_keep_topo_etc() { panic("!implemented") }

func swi_deflect_light(xx *double, dt double, iflag int32) { panic("!implemented") }

func swi_fixstar_calc_from_record(srecord []byte, tjd double, iflag int32, star []byte, xx *double, serr []byte) int32 {
	panic("!implemented")
}

func swi_fixstar_load_record(star []byte, srecord []byte, sname []byte, sbayer []byte, dparams *double, serr []byte) int32 {
	panic("!implemented")
}

func swi_fopen(ifno int, fname []byte, ephepath []byte, serr []byte) *FILE { panic("!implemented") }

func swi_force_app_pos_etc() { panic("!implemented") }

func swi_get_ayanamsa_ex(tjd_et double, iflag int32, daya *double, serr []byte) int32 {
	panic("!implemented")
}

func swi_get_ayanamsa_with_speed(tjd_et double, iflag int32, daya *double, serr []byte) int32 {
	panic("!implemented")
}

func swi_get_denum(ipli int32, iflag int32) int32 { panic("!implemented") }

func swi_get_observer(tjd double, iflag int32, missing ...any) int { panic("!implemented") }

func swi_init_swed_if_start() int32 { panic("!implemented") }

func swi_nutate(xx *double, iflag int32, backward AS_BOOL) { panic("!implemented") }

func swi_plan_for_osc_elem(iflag int32, tjd double, xx *double) int { panic("!implemented") }

func swi_precess_speed(xx *double, t double, iflag int32, direction int) { panic("!implemented") }

func swi_trop_ra2sid_lon(xin *double, xout *double, xoutr *double, iflag int32) int {
	panic("!implemented")
}

func swi_trop_ra2sid_lon_sosy(xin *double, xout *double, iflag int32) int { panic("!implemented") }

func trace_swe_fixstar(swtch int, star []byte, tjd double, iflag int32, xx *double, serr []byte) {
	panic("!implemented")
}

func trace_swe_get_planet_name(swtch int, ipl int, s []byte) { panic("!implemented") }
