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

var (
	ciCore = CI{mon: 11, day: 19, yea: 1971, tim: HM(11, 1), zon: 8.0, lon: DEFAULT_LOC_LON, lat: DEFAULT_LOC_LAT}
	ciFive = CI{mon: -1}
	ciFour = CI{mon: -1}
	ciGreg = CI{mon: 10, day: 15, yea: 1582}
	ciHexa = CI{mon: -1}
	ciMain = CI{mon: -1}
	ciSave = CI{mon: 4, day: 9, yea: 2023, tim: HMS(1, 9, 11), zon: 8.0, lon: DEFAULT_LOC_LON, lat: DEFAULT_LOC_LAT}
	ciThre = CI{mon: -1}
	ciTran = CI{mon: 1, day: 1, yea: 2023}
	ciTwin = CI{mon: 9, day: 11, yea: 1991, tim: HMS(0, 0, 38), lon: DEFAULT_LOC_LON, lat: DEFAULT_LOC_LAT}
	cp0    CP
	cp1    CP
	cp2    CP
	cp3    CP
	cp4    CP
	cp5    CP
	cp6    CP

	// Restriction status of each object, as specified with -R switch.
	ignore = [objMax]bool{
		true,
		// Planets
		false, false, false, false, false, false, false, false, false, false,
		// Minors
		false, false, false, false, false, false, true, false, false, false, false,
		// Cusps
		true, true, true, true, true, true, true, true, true, true, true, true,
		// Uranians
		true, true, true, true, true, true, true, true, true,
		// Dwarfs
		true, true, true, true, true, true, true, true, true,
		// Moons
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		// Stars
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	}

	// Restriction of objects when transiting, as specified with -RT switch.
	ignore2 = [objMax]bool{
		true,
		// Planets
		false, true, false, false, false, false, false, false, false, false,
		// Minors
		false, false, false, false, false, false, true, false, true, true, true,
		// Cusps
		true, true, true, true, true, true, true, true, true, true, true, true,
		// Uranians
		true, true, true, true, true, true, true, true, true,
		// Dwarfs
		true, true, true, true, true, true, true, true, true,
		// Moons
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		// Stars
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	}

	// Restriction status of each aspect, as specified with -RA switch.
	ignorea = [cAspect + 1]bool{
		false,
		false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	}
	// Restrictions for -Zd chart events.
	ignorez = [arMax]bool{
		false, false, false, false,
	}
	// Restrictions for rulership types.
	ignore7 = [rrMax]bool{
		false, true, true, false, true,
	}

	ignoreMem  [objMax]bool
	ignore2Mem [objMax]bool
	ignoreaMem [cAspect + 1]bool
	ignorezMem [arMax]bool
	ignore7Mem [rrMax]bool
	ignorefMem [6]bool

	is = IS{
		nObj:      int(cObj),
		iciCur:    -1,
		iciIndex1: -1,
		iciIndex2: -1,
		OB:        rAxis,
		jdDeltaT:  rInvalid,
	}

	us = US{
		// Chart types
		// Chart suboptions
		fGridConfig:    true,
		fAspSummary:    true,
		fMidSummary:    true,
		fInfluenceSign: true,
		fLatitudeCross: true,
		// Table chart types
		// Main flags
		fEphemFiles: true,
		// Main subflags
		// Obscure flags
		fEclipse:      true,
		fIgnoreDiralt: true,
		fIgnoreDirlen: true,
		fIgnoreAlt0:   true,
		// Value settings
		nDecanType:   ddDecanR,
		nEphemFactor: 1,
		nRel:         rcNone,
		nHouseSystem: DEFAULT_SYSTEM,
		nHouse3D:     hmPrime,
		nAsp:         DEFAULT_ASPECTS,
		objCenter:    oEar,
		rHarmonic:    1.0,
		objRequire:   -1,
		nProgress:    ptCast,
		nDivision:    DIVISIONS,
		nScreenWidth: SCREENWIDTH,
		zonDef:       DEFAULT_ZONE,
		lonDef:       DEFAULT_LONG,
		latDef:       DEFAULT_LAT,
		tmpDef:       15.0,
		// Value subsettings
		nAstroGraphStep: 5,
		nAstroGraphDist: 200,
		nArabicParts:    cPart,
		rProgCusp:       1.0,
		nRatio1:         1,
		nRatio2:         1,
		nCharset:        ccNone,
		nCharsetOut:     ccNone,
		nScrollRow:      24,
		rDeltaT:         rInvalid,
		objRot1:         oEar,
		objRot2:         oEar,
		nBioday:         BIODAYS,
		// AstroExpressions
	}
)
