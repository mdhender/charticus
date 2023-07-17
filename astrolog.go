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

// astrolog.cpp

func Action() { panic("!implemented") }

func FProcessCommandLine(szLine []byte) flag { panic("!implemented") }

func FProcessSwitches(argc int, argv [][]byte) flag { panic("!implemented") }

func FinalizeProgram(fSkip flag) { panic("!implemented") }

func InitColors() { panic("!implemented") }

func InitProgram() { panic("!implemented") }

func InitVariables() { panic("!implemented") }

func NParseCommandLine(szLine []byte, argv [][]byte) int { panic("!implemented") }

func NProcessSwitchesRare(argc int, argv [][]byte, pos int, missing ...any) int {
	panic("!implemented")
}

func NPromptSwitches(line []byte, argv [MAXSWITCHES][]byte) int { panic("!implemented") }

func main(argc int, argv [][]byte) int { panic("!implemented") }
