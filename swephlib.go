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

// swephlib.cpp

func adjust_for_tidacc(ans double, Y double, tid_acc double, tid_acc0 double, adjust_after_1955 AS_BOOL) double {
	panic("!implemented")
}

func bessel(v *double, n int, t double) double { panic("!implemented") }

func calc_deltat(tjd double, iflag int32, deltat *double, serr []byte) int32 { panic("!implemented") }

func calc_nutation(J double, iflag int32, nutlo *double) int { panic("!implemented") }

func calc_nutation_iau1980(J double, nutlo *double) int { panic("!implemented") }

func calc_nutation_iau2000ab(J double, nutlo *double) int { panic("!implemented") }

func calc_nutation_woolard(J double, nutlo *double) int { panic("!implemented") }

func deltat_aa(tjd double, tid_acc double) double { panic("!implemented") }

func deltat_espenak_meeus_1620(tjd double, tid_acc double) double { panic("!implemented") }

func deltat_longterm_morrison_stephenson(tjd double) double { panic("!implemented") }

func deltat_stephenson_etc_2016(tjd double, tid_acc double) double { panic("!implemented") }

func deltat_stephenson_morrison_1997_1600(tjd double, tid_acc double) double { panic("!implemented") }

func deltat_stephenson_morrison_2004_1600(tjd double, tid_acc double) double { panic("!implemented") }

func epsiln_owen_1986(tjd double, eps *double) { panic("!implemented") }

func get_deltat_model(dtmod int, s []byte) { panic("!implemented") }

func get_frame_bias_model(biasmod int, s []byte) { panic("!implemented") }

func get_nutation_model(nutmod int, iflag int32, s []byte) { panic("!implemented") }

func get_owen_t0_icof(tjd double, t0 *double, icof *int) { panic("!implemented") }

func get_precession_model(precmod int, iflag int32, s []byte) { panic("!implemented") }

func get_sidt_model(sidtmod int, s []byte) { panic("!implemented") }

func init_crc32() { panic("!implemented") }

func init_dt() int { panic("!implemented") }

func owen_pre_matrix(tjd double, rp *double, iflag int) { panic("!implemented") }

func pre_pecl(tjd double, vec *double) { panic("!implemented") }

func pre_pequ(tjd double, veq *double) { panic("!implemented") }

func pre_pmat(tjd double, rp *double) { panic("!implemented") }

func precess_1(R *double, J double, direction int, prec_method int) int { panic("!implemented") }

func precess_2(R *double, J double, iflag int32, direction int, prec_method int) int {
	panic("!implemented")
}

func precess_3(R *double, J double, direction int, iflag int, prec_meth int) int {
	panic("!implemented")
}

func quadratic_intp(ym double, y0 double, yp double, x double) double { panic("!implemented") }

func set_astro_models(samod []byte) { panic("!implemented") }

func sidtime_long_term(tjd_ut double, eps double, nut double) double { panic("!implemented") }

func sidtime_non_polynomial_part(tt double) double { panic("!implemented") }

func split_deg_nakshatra(ddeg double, roundflag int32, ideg *int32, imin *int32, isec *int32, dsecfr *double, inak *int32) {
	panic("!implemented")
}

func swe_cotrans(xpo *double, xpn *double, eps double) { panic("!implemented") }

func swe_cotrans_sp(xpo *double, xpn *double, eps double) { panic("!implemented") }

func swe_cs2degstr(t CSEC, a []byte) []byte { panic("!implemented") }

func swe_cs2lonlatstr(t CSEC, pchar byte, mchar byte, sp []byte) []byte { panic("!implemented") }

func swe_cs2timestr(t CSEC, sep int, suppressZero AS_BOOL, a []byte) []byte { panic("!implemented") }

func swe_csnorm(p centisec) centisec { panic("!implemented") }

func swe_csroundsec(x centisec) centisec { panic("!implemented") }

func swe_d2l(x double) int32 { panic("!implemented") }

func swe_day_of_week(jd double) int { panic("!implemented") }

func swe_deg_midp(x1 double, x0 double) double { panic("!implemented") }

func swe_degnorm(x double) double { panic("!implemented") }

func swe_deltat(tjd double) double { panic("!implemented") }

func swe_deltat_ex(tjd double, iflag int32, serr []byte) double { panic("!implemented") }

func swe_difcs2n(p1, p2 centisec) centisec { panic("!implemented") }

func swe_difcsn(p1, p2 centisec) centisec { panic("!implemented") }

func swe_difdeg2n(p1 double, p2 double) double { panic("!implemented") }

func swe_difdegn(p1 double, p2 double) double { panic("!implemented") }

func swe_difrad2n(p1 double, p2 double) double { panic("!implemented") }

func swe_get_astro_models(samod []byte, sdet []byte, iflag int32) { panic("!implemented") }

func swe_get_tid_acc() double { panic("!implemented") }

func swe_rad_midp(x1 double, x0 double) double { panic("!implemented") }

func swe_radnorm(x double) double { panic("!implemented") }

func swe_set_astro_models(samod []byte, iflag int32) { panic("!implemented") }

func swe_set_delta_t_userdef(dt double) { panic("!implemented") }

func swe_set_interpolate_nut(do_interpolate AS_BOOL) { panic("!implemented") }

func swe_set_tid_acc(t_acc double) { panic("!implemented") }

func swe_sidtime(tjd_ut double) double { panic("!implemented") }

func swe_sidtime0(tjd double, eps double, nut double) double { panic("!implemented") }

func swe_split_deg(ddeg double, roundflag int32, ideg *int32, imin *int32, isec *int32, dsecfr *double, isgn *int32) {
	panic("!implemented")
}

func swi_FK4_FK5(xp *double, tjd double) { panic("!implemented") }

func swi_FK5_FK4(xp *double, tjd double) { panic("!implemented") }

func swi_angnorm(x double) double { panic("!implemented") }

func swi_approx_jplhor(x *double, tjd double, iflag int32, backward AS_BOOL) { panic("!implemented") }

func swi_bias(x *double, tjd double, iflag int32, backward AS_BOOL) { panic("!implemented") }

func swi_cartpol(x *double, l *double) { panic("!implemented") }

func swi_cartpol_sp(x *double, l *double) { panic("!implemented") }

func swi_coortrf(xpo *double, xpn *double, eps double) { panic("!implemented") }

func swi_coortrf2(xpo *double, xpn *double, sineps double, coseps double) { panic("!implemented") }

func swi_crc32(buf []byte, len int) uint32 { panic("!implemented") }

func swi_cross_prod(a *double, b *double, x *double) { panic("!implemented") }

func swi_cutstr(s []byte, cutlist []byte, cpos [][]byte, nmax int) int { panic("!implemented") }

func swi_dot_prod_unit(x *double, y *double) double { panic("!implemented") }

func swi_echeb(x double, coef *double, ncf int) double { panic("!implemented") }

func swi_edcheb(x double, coef *double, ncf int) double { panic("!implemented") }

func swi_epsiln(J double, iflag int32) double { panic("!implemented") }

func swi_gen_filename(tjd double, ipli int, fname []byte) { panic("!implemented") }

func swi_get_tid_acc(tjd_ut double, iflag int32, denum int32, denumret *int32, tid_acc *double, serr []byte) int32 {
	panic("!implemented")
}

func swi_guess_ephe_flag() int32 { panic("!implemented") }

func swi_icrs2fk5(x *double, iflag int32, backward AS_BOOL) { panic("!implemented") }

func swi_kepler(E double, M double, ecce double) double { panic("!implemented") }

func swi_ldp_peps(tjd double, dpre *double, deps *double) { panic("!implemented") }

func swi_mod2PI(x double) double { panic("!implemented") }

func swi_nutation(tjd double, iflag int32, nutlo *double) int { panic("!implemented") }

func swi_polcart(l *double, x *double) { panic("!implemented") }

func swi_polcart_sp(l *double, x *double) { panic("!implemented") }

func swi_precess(R *double, J double, iflag int32, direction int) int { panic("!implemented") }

func swi_right_trim(s []byte) []byte { panic("!implemented") }

func swi_set_tid_acc(tjd_ut double, iflag int32, denum int32, serr []byte) int32 {
	panic("!implemented")
}

func swi_strcpy(to []byte, from []byte) []byte { panic("!implemented") }
