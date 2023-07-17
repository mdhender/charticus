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

// atlas.cpp

func AdjustTime(mon *int, day *int, yea *int, tim *int) { panic("!implemented") }

func DisplayAtlasLookup(szIn []byte, lDialog size_t, piae *int) flag { panic("!implemented") }

func DisplayAtlasNearby(lon real, lat real, lDialog size_t, piae *int, missing ...any) flag {
	panic("!implemented")
}

func DisplayTimezoneChanges(iznIn int, lDialog size_t, ci *CI) flag { panic("!implemented") }

func FEnsureAtlas() flag { panic("!implemented") }

func FEnsureTimezoneChanges() flag { panic("!implemented") }

func FLoadAtlas(file *FILE, cae int) flag { panic("!implemented") }

func FLoadZoneChanges(file *FILE, izcnMax int, izceMax int) flag { panic("!implemented") }

func FLoadZoneLinks(file *FILE, czl int) flag { panic("!implemented") }

func FLoadZoneRules(file *FILE, irunMax int, irueMax int) flag { panic("!implemented") }

func FSetDstZon(ci *CI, izn int, missing ...any) flag { panic("!implemented") }

func ILookupCN(szAbb []byte, rgcn *CN, ccn int) int { panic("!implemented") }

func NParseHMS(sz []byte) int { panic("!implemented") }

func SzCity(iae int) []byte { panic("!implemented") }

func ZondefFromIzn(izn int) real { panic("!implemented") }
