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

// wdialog.cpp

func DlgAbort(hdlg HWND, uint message, wParam WPARAM, lParam LPARAM) BOOL { panic("!implemented") }

func DlgAbortProc(hdc HDC, nCode int) flag { panic("!implemented") }

func DlgAbout(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgAspect(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgCalc(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgChart(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgColor(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgCommand(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgCustom(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgCustomS(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgDefault(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgDisplay(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgGraphics(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgInfo(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgInfoAll(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgList(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgMoons(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgObject(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgObject2(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgObjectM(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgOpenChart() flag { panic("!implemented") }

func DlgOpenDir() flag { panic("!implemented") }

func DlgPrint() flag { panic("!implemented") }

func DlgRestrict(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgSaveChart() flag { panic("!implemented") }

func DlgStar(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func DlgTransit(hdlg HWND, uint message, wParam WORD, lParam LONG) flag { panic("!implemented") }

func ErrorEnsure(n int, sz []byte) { panic("!implemented") }

func FDlgInfoAtlas(hdlg HWND, wParam WORD, fDefault flag) flag { panic("!implemented") }

func GetEditN(hdlg HWND, id int) int { panic("!implemented") }

func GetEditR(hdlg HWND, id int) real { panic("!implemented") }

func GetRadio(hdlg HWND, id1, id2 int) int { panic("!implemented") }

func KvDialog() KV { panic("!implemented") }

func SetEditColor(hdlg HWND, id int, ki KI, nExtra int) { panic("!implemented") }

func SetEditMDYT(hdlg HWND, idMon int, idDay int, idYea int, idTim int, missing ...any) {
	panic("!implemented")
}

func SetEditR(hdlg HWND, id int, r real, n int) { panic("!implemented") }

func SetEditSZOA(hdlg HWND, idDst int, idZon int, idLon int, idLat int, missing ...any) {
	panic("!implemented")
}

func SetEditSz(hdlg HWND, id int, sz []byte) { panic("!implemented") }

func SetListSz(hdlg HWND, id int, sz []byte) int { panic("!implemented") }
