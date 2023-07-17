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

// swejpl.cpp

func fsizer(serr []byte) int32 { panic("!implemented") }

func interp(buf *double, t double, intv double, ncfin int32, missing ...any) int {
	panic("!implemented")
}

func read_const_jpl(ss *double, serr []byte) int { panic("!implemented") }

func reorder(x []byte, size int, number int) { panic("!implemented") }

func state(et double, list *int32, do_bary int, missing ...any) int { panic("!implemented") }

func swi_close_jpl_file() { panic("!implemented") }

func swi_get_jpl_denum() int32 { panic("!implemented") }

func swi_open_jpl_file(ss *double, fname []byte, fpath []byte, serr []byte) int {
	panic("!implemented")
}

func swi_pleph(et double, ntarg int, ncent int, rrd *double, serr []byte) int { panic("!implemented") }
