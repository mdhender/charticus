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

// express.cpp

func CsCreateTrie(rgszIn [][]uchar, cszIn int, rgsOut TRIE, csTrieMax int, missing ...any) int {
	panic("!implemented")
}

func ExpFinalize() { panic("!implemented") }

func ExpGetString(i int) []byte { panic("!implemented") }

func ExpSetMacro(i int, sz []byte) flag { panic("!implemented") }

func ExpSetN(i int, n int) flag { panic("!implemented") }

func ExpSetR(i int, r real) flag { panic("!implemented") }

func ExpSetString(i int, sz []byte, fList flag) flag { panic("!implemented") }

func FCreateTries() flag { panic("!implemented") }

func FEnsureExpMacro(cszNew int) flag { panic("!implemented") }

func FEnsureExpStr(cszNew int) flag { panic("!implemented") }

func FEnsureParVar(cparNew int) flag { panic("!implemented") }

func FEvalFunction(ifun int, rgpar *PAR, rgpchEval [2][]byte) flag { panic("!implemented") }

func GetParameter(sz []byte, ppar *PAR) { panic("!implemented") }

func ILookupTrie(rgsIn TRIE, rgch []byte, cch int, missing ...any) int { panic("!implemented") }

func LFromRgch(rgch []byte, cch int) long { panic("!implemented") }

func NExpGet(i int) int { panic("!implemented") }

func NParseExpression(sz []byte) long { panic("!implemented") }

func PchFormatExpression(sz []byte, i int) []byte { panic("!implemented") }

func PchFormatString(sz []byte, i int) []byte { panic("!implemented") }

func PchGetParameter(pchCur []byte, rgpar *PAR, ifun int, missing ...any) []byte {
	panic("!implemented")
}

func RExpGet(i int) real { panic("!implemented") }

func RParseExpression(sz []byte) real { panic("!implemented") }

func ShowParseExpression(sz []byte) flag { panic("!implemented") }

func UpperRgch(rgch []byte, cch int) { panic("!implemented") }
