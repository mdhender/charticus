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

const (
	FALSE = false
	TRUE  = true
)

type CALENDARFLAG int

const (
	SE_JUL_CAL CALENDARFLAG = iota
	SE_GREG_CAL
)

type RTSTATUS int

const (
	OK                RTSTATUS = 0
	ERR                        = -1
	NOT_AVAILABLE              = -2
	BEYOND_EPH_LIMITS          = -3
)
