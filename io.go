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

// io.cpp

func BRead(file *FILE) byte { panic("!implemented") }

func FInputData(szFile []byte) flag { panic("!implemented") }

func FOutputAAFFile() flag { panic("!implemented") }

func FOutputChartList() flag { panic("!implemented") }

func FOutputDaedalusStar() flag { panic("!implemented") }

func FOutputData() flag { panic("!implemented") }

func FOutputQuickFile() flag { panic("!implemented") }

func FOutputSettings() flag { panic("!implemented") }

func FProcessAAFFile(szFile []byte, file *FILE) flag { panic("!implemented") }

func FProcessADBFile(szFile []byte, file *FILE) flag { panic("!implemented") }

func FProcessQuickFile(szFile []byte, file *FILE) flag { panic("!implemented") }

func FProcessSFTextFile(szFile []byte, file *FILE) flag { panic("!implemented") }

func FProcessSwitchFile(szFile []byte, file *FILE) flag { panic("!implemented") }

func FileOpen(szFile []byte, nFileMode int, szPath []byte) *FILE { panic("!implemented") }

func GetJPLHorizons(id int, obj *real, objalt *real, dir *real, dist *real, missing ...any) flag {
	panic("!implemented")
}

func GetURL(szUrl, szFile []byte) flag { panic("!implemented") }

func InputString(szPrompt []byte, sz []byte) { panic("!implemented") }

func LRead(file *FILE) dword { panic("!implemented") }

func NFromSz(sz []byte) int { panic("!implemented") }

func NInputRange(szPrompt []byte, low int, high int, pm int) int { panic("!implemented") }

func NParseSz(szEntry []byte, pm int) int { panic("!implemented") }

func OpenDir(szDir []byte) { panic("!implemented") }

func RFromSz(sz []byte) real { panic("!implemented") }

func RInputRange(szPrompt []byte, low real, high real, pm int) real { panic("!implemented") }

func RParseSz(szEntry []byte, pm int) real { panic("!implemented") }

func WRead(file *FILE) word { panic("!implemented") }
