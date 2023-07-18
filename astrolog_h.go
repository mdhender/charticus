// charticus - an astrology chart tool
// Copyright (c) 2023 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation either version 3 of the License or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not see <https://www.gnu.org/licenses/>.

package charticus

import "math"

// Array index limits
const (
	objMax    = cObj + 1
	objMaxG   = objMax + 11
	cCnstl    = 88
	cZone     = 73
	cSector   = 36
	cPart     = 177
	cWeek     = 7
	cColor    = 16
	cRainbow  = 7
	cRay      = 7
	cRing     = 6
	cHasMoons = 11
)

// Math and conversion constants
const (
	rSqr2      = 1.41421356237309504880
	rSqr3      = 1.73205080756887729353
	rPhi       = 1.61803398874989484820
	rLog10     = 2.30258509299404568402
	rLog101    = 4.61512051684125945088
	rPi        = 3.14159265358979323846
	rPi2       = (rPi * 2.0)
	rPiHalf    = (rPi / 2.0)
	rDegMax    = 360.0
	rDegHalf   = 180.0
	rDegQuad   = 90.0
	rDegRad    = (rDegHalf / rPi)
	rMiToKm    = 1.609344
	rFtToM     = 0.3048
	rInToCm    = 2.54
	rAUToKm    = 149597870.7
	rLYToAU    = 63241.07708427
	rPCToAU    = 206264.8062471
	rDayInYear = 365.24219
	rEarthDist = 149.59787
	rEpoch2000 = -24.736467
	rJD2000    = 2451545.0
	rAxis      = 23.44578889
	rSmall     = (1.7453e-09)
	rLarge     = 10000.0
	rInvalid   = (1.23456789e-09)
	rRound     = 0.5
)

// miscellaneous constants
const (
	// sentinel for not a star in the rStarBrightMatrix
	rStarNot = 999.99

	WHEELCOLS   = 15  // Affects width of each house in wheel display.
	WHEELROWS   = 11  // Max no. of objects that can be in a wheel house.
	SCREENWIDTH = 80  // Number of columns to print interpretations in.
	MONTHSPACE  = 3   // Number of spaces between each calendar column.
	MAXINDAY    = 300 // Max number of aspects or transits displayable.
	MAXCROSS    = 750 // Max number of latitude crossings displayable.
	BIODAYS     = 14  // Days to include in graphic biorhythms.
	CREDITWIDTH = 74  // Number of text columns in the -Hc credit screen.
	MAXSWITCHES = 100 // Max number of switch parameters per input line.
	PSGUTTER    = 9   // Points of white space on PostScript page edge.
)

// Angle restrictions

type _angles int

const (
	arAsc _angles = 0 // Ascendant
	arMC  _angles = 1 // Midheaven
	arDes _angles = 2 // Descendant
	arIC  _angles = 3 // Nadir
)
const arMax = int(arIC) + 1

// Aspects

type _aspects int

const (
	aDis _aspects = -8 // Distances equal
	aNod _aspects = -7 // Latitude zero/node crossing
	aLen _aspects = -6 // Direction change (distance)
	aAlt _aspects = -5 // Direction change (latitude)
	aHou _aspects = -4 // 3D House change
	aDeg _aspects = -3 // Degree change
	aDir _aspects = -2 // Direction change (longitude)
	aSig _aspects = -1 // Sign change
	aCon _aspects = 1
	aOpp _aspects = 2
	aSqu _aspects = 3
	aTri _aspects = 4
	aSex _aspects = 5
	aInc _aspects = 6
	aSSx _aspects = 7
	aSSq _aspects = 8
	aSes _aspects = 9
	aQui _aspects = 10
	aBQn _aspects = 11
	aSQn _aspects = 12
	aSep _aspects = 13
	aNov _aspects = 14
	aBNv _aspects = 15
	aBSp _aspects = 16
	aTSp _aspects = 17
	aQNv _aspects = 18
	aDc3 _aspects = 19
	aUd1 _aspects = 20
	aUd2 _aspects = 21
	aUd3 _aspects = 22
	aUd4 _aspects = 23
	aUd5 _aspects = 24
)
const cAspect = int(aUd5)
const cAspect2 = cAspect + int(aOpp)
const cAspect3 = cAspect2 + int(aOpp)

// File codepage

type _charactercodepage int

const (
	ccNone  _charactercodepage = 0 // Unknown (don't ever convert)
	ccIBM   _charactercodepage = 1 // Codepage 437 (IBM/DOS/OEM)
	ccLatin _charactercodepage = 2 // Codepage ISO-8859-1 (Latin-1 or Windows-1252)
	ccUTF8  _charactercodepage = 3 // Codepage UTF8 (up to 3 bytes per character)
)

// Decan (or other) division types

type _decandivisiontype int

const (
	ddNone    _decandivisiontype = 0 // No division
	ddDecanR  _decandivisiontype = 1 // Decan/Face (ruler)
	ddDecanS  _decandivisiontype = 2 // Decan/Face (sign)
	ddChaldea _decandivisiontype = 3 // Chaldean Decan
	ddEgypt   _decandivisiontype = 4 // Egyptian Terms/Bounds
	ddPtolemy _decandivisiontype = 5 // Ptolemaic Terms/Bounds
	ddNavamsa _decandivisiontype = 6 // Navamsa division
	dd12      _decandivisiontype = 7 // 12th harmonic
	ddDwad    _decandivisiontype = 8 // Dwad
	dd27      _decandivisiontype = 9 // 27 Nakshatras
	ddMax     _decandivisiontype = 10
)

// House system models

type _housemodel int

const (
	hmEcliptic _housemodel = 0 // Standard 2D: Equator of houses is ecliptic
	hmPrime    _housemodel = 1 // 3D: Equator of houses is prime vertical
	hmHorizon  _housemodel = 2 // 3D: Equator of houses is local horizon
	hmEquator  _housemodel = 3 // 3D: Equator of houses is celestial equator
	cMethod    _housemodel = 3
)

// Object array index values
const (
	cPlanet = oVes
	cThing  = oLil
	oMain   = 10
	oCore   = 21
	cUran   = 9
	cDwarf  = 9
	cMoons  = 27
	cCOB    = 5
	cMoons2 = cMoons + cCOB
	cCust   = cUran + cDwarf + cMoons2
	cStar   = 50
	cuspLo  = oCore + 1
	cuspHi  = cuspLo + cSign - 1
	uranLo  = cuspHi + 1
	uranHi  = uranLo + cUran - 1
	dwarfLo = uranHi + 1
	dwarfHi = dwarfLo + cDwarf - 1
	moonsLo = dwarfHi + 1
	moonsHi = moonsLo + cMoons - 1
	cobLo   = moonsHi + 1
	cobHi   = cobLo + cCOB - 1
	starLo  = cobHi + 1
	starHi  = starLo + cStar - 1
	custLo  = uranLo
	custHi  = cobHi
	oNorm   = cobHi
	oNorm1  = starLo
)

// Objects
type _object = int

const (
	oEar _object = 0
	oSun _object = 1
	oMoo _object = 2
	oMer _object = 3
	oVen _object = 4
	oMar _object = 5
	oJup _object = 6
	oSat _object = 7
	oUra _object = 8
	oNep _object = 9
	oPlu _object = 10
	oChi _object = 11
	oCer _object = 12
	oPal _object = 13
	oJun _object = 14
	oVes _object = 15
	oNod _object = 16
	oSou _object = 17
	oLil _object = 18
	oFor _object = 19
	oVtx _object = 20
	oEP  _object = 21
	oAsc _object = _object(int(cuspLo) - 1 + 1)
	o2nd _object = _object(int(cuspLo) - 1 + 2)
	o3rd _object = _object(int(cuspLo) - 1 + 3)
	oNad _object = _object(int(cuspLo) - 1 + 4)
	o5th _object = _object(int(cuspLo) - 1 + 5)
	o6th _object = _object(int(cuspLo) - 1 + 6)
	oDes _object = _object(int(cuspLo) - 1 + 7)
	o8th _object = _object(int(cuspLo) - 1 + 8)
	o9th _object = _object(int(cuspLo) - 1 + 9)
	oMC  _object = _object(int(cuspLo) - 1 + 10)
	o11h _object = _object(int(cuspLo) - 1 + 11)
	o12h _object = _object(int(cuspLo) - 1 + 12)
	oVul _object = _object(int(uranLo) + 0)      // Vulcan
	oVlk _object = _object(int(uranLo) + 7)      // Vulkanus
	oHyg _object = _object(int(dwarfLo) + 0)     // Hygiea
	oPho _object = _object(int(dwarfLo) + 1)     // Pholus
	oEri _object = _object(int(dwarfLo) + 2)     // Eris
	oHau _object = _object(int(dwarfLo) + 3)     // Haumea
	oMak _object = _object(int(dwarfLo) + 4)     // Makemake
	oGon _object = _object(int(dwarfLo) + 5)     // Gonggong
	oQua _object = _object(int(dwarfLo) + 6)     // Quaoar
	oSed _object = _object(int(dwarfLo) + 7)     // Sedna
	oOrc _object = _object(int(dwarfLo) + 8)     // Orcus
	oUrT _object = _object(int(moonsLo) + 14)    // Uranus' Titania
	oJuC _object = _object(int(cobLo) + 0)       // Jupiter COB
	oSaC _object = _object(int(cobLo) + 1)       // Saturn COB
	oUrC _object = _object(int(cobLo) + 2)       // Uranus COB
	oNeC _object = _object(int(cobLo) + 3)       // Neptune COB
	oPlC _object = _object(int(cobLo) + 4)       // Pluto COB
	oOri _object = _object(int(starLo) - 1 + 28) // Alnilam (Orion)
	oAlr _object = _object(int(starLo) - 1 + 30) // Alnair
	oPle _object = _object(int(starLo) - 1 + 45) // Alcyone (Pleiades)
	oAnd _object = _object(int(starLo) - 1 + 46) // Andromeda (M31)
	oGal _object = _object(int(starLo) - 1 + 49) // Galactic Center (Milky Way)
	cObj _object = 133
)

// Progressed chart types

type _progressiontype int

const (
	ptCast     _progressiontype = 0
	ptSolarArc _progressiontype = 1
	ptMixed    _progressiontype = 2
)

// Relationship chart modes

type _relationshipchart int

const (
	rcNone       _relationshipchart = 0
	rcSynastry   _relationshipchart = 1
	rcComposite  _relationshipchart = 2
	rcMidpoint   _relationshipchart = 3
	rcDifference _relationshipchart = 4
	rcBiorhythm  _relationshipchart = 5
	rcDual       _relationshipchart = -1
	rcTriWheel   _relationshipchart = -2
	rcQuadWheel  _relationshipchart = -3
	rcQuinWheel  _relationshipchart = -4
	rcHexaWheel  _relationshipchart = -5
	rcTransit    _relationshipchart = -6
	rcProgress   _relationshipchart = -7
)

// Rulership restrictions

type _rulerships int

const (
	rrStd _rulerships = 0 // Standard exoteric
	rrEso _rulerships = 1 // Esoteric
	rrHie _rulerships = 2 // Hierarchical
	rrExa _rulerships = 3 // Exaltation
	rrRay _rulerships = 4 // Ray rulership
)
const rrMax = int(rrRay) + 1

// Zodiac signs
type _signs int

const (
	sAri _signs = 1
	sTau _signs = 2
	sGem _signs = 3
	sCan _signs = 4
	sLeo _signs = 5
	sVir _signs = 6
	sLib _signs = 7
	sSco _signs = 8
	sSag _signs = 9
	sCap _signs = 10
	sAqu _signs = 11
	sPis _signs = 12
)
const cSign = int(sPis)

// Atlas values
const (
	cchSzAtl = 54
	cchSzZon = 13
	icnewMax = 252
	icnusMax = 51
	icncaMax = 13
	icnUS    = 234 // United States
	icnCA    = 38  // Canada
	icnFR    = 76  // France
	iznMax   = 425
	ilistMax = 200
)

type _AtlasEntry struct {
	lon    float64        // Longitude of city
	lat    float64        // Latitude of city
	icn    int            // Country or region of city
	istate int            // State or province of city, if US or CA
	szNam  [cchSzAtl]byte // Name of city
	izn    int            // Time zone area of city
}
type AtlasEntry = _AtlasEntry

type _ChartInfo struct {
	mon int     // Month
	day int     // Day
	yea int     // Year
	tim float64 // Time in hours
	dst float64 // Daylight offset
	zon float64 // Time zone
	lon float64 // Longitude
	lat float64 // Latitude
	nam []byte  // Name for chart
	loc []byte  // Name of location
}
type CI = _ChartInfo

type _ChartPositions struct {
	obj    [objMax]float64    // The zodiac positions
	alt    [objMax]float64    // Ecliptic declination
	dir    [objMax]float64    // Retrogradation velocity
	diralt [objMax]float64    // Latitude velocity
	dirlen [objMax]float64    // Distance velocity
	pt     [objMax]PT3R       // X,Y,Z coordintes in space
	dist   [objMax]float64    // Distance to X,Y,Z coordinates
	cusp   [cSign + 1]float64 // House cusp positions
	cusp3  [cSign + 1]float64 // 3D house cusp positions
	house  [objMax]int        // House each object is in
	lonMC  float64            // 0 longitude converted to equatorial coordinates
}
type CP = _ChartPositions

// TODO: should nObj be _object instead of int?
type _InternalSettings struct {
	fHaveInfo     flag        // Do we need to prompt user for chart info?
	fDst          flag        // Has Daylight Saving Time been autodetected?
	fProgress     flag        // Are we doing a chart involving progression?
	fReturn       flag        // Are we doing a transit chart for returns?
	fMult         flag        // Have we already printed at least one text chart?
	fSeconds      flag        // Do we print locations to nearest second?
	fSzPersist    flag        // Are parameter strings persistent when processing?
	fSzInteract   flag        // Are we in middle of chart so some setting fixed?
	fNoEphFile    flag        // Have we already had a ephem file not found error?
	fSwissPathSet flag        // Has the Swiss Ephemeris path been set yet?
	szProgName    []byte      // The name and path of the executable running.
	rgszLine      [9][]byte   // The command lines to run before each -Yq chart.
	szFileScreen  []byte      // The file to send text output to as passed to -os.
	szFileOut     []byte      // The output chart filename string as passed to -o.
	rgszComment   *[]byte     // Points to any comment strings after -o filename.
	nContext      int         // Context of current or most recent chart cast.
	nObj          int         // Index of highest unrestricted object.
	cszComment    int         // The number of strings after -o that are comments.
	cchRow        int         // The current row text charts have scrolled to.
	cchCol        int         // The current column text charts are printing at.
	cchColMax     int         // Max column current text chart has printed at.
	nHTML         int         // HTML text output context mode for -kh.
	nHouseSystem  int         // Actual house system used to compute cusps for -c.
	nWheelRows    int         // Actual number of rows per house to use for -w.
	cae           int         // Number of atlas entries of city locations loaded.
	czcn          int         // Number of time zone change areas loaded.
	czce          int         // Total number of change entries in all zone areas.
	crun          int         // Number of time zone Daylight rule categories.
	crue          int         // Total number of rule entries in all categories.
	cci           int         // Number of user visible charts in chart list.
	cciAlloc      int         // Number of allocated positions in chart list.
	iciCur        int         // Index of most recently used chart in chart list.
	iciIndex1     int         // Index into chart list used for chart #1 with -5e.
	iciIndex2     int         // Index into chart list used for chart #2 with -5e3.
	cAlloc        int         // Number of memory allocations currently allocated.
	cAllocTotal   int         // Total memory allocations allocated this session.
	cbAllocSize   int         // Total bytes in all memory allocations allocated.
	rOff          float64     // Offset between sidereal and tropical zodiacs.
	rSid          float64     // Sidereal offset degrees to be added to locations.
	JD            float64     // Fractional Julian day for current chart.
	JDp           float64     // Julian day that a progressed chart indicates.
	Tp            float64     // Julian time used for progressed chart cusps.
	rgae          *AtlasEntry // List of atlas entries for city coordinates.
	rgzc          *ZoneChange // List of time zone change entries for zone areas.
	rgrun         *RuleName   // List of Daylight Saving change rule names.
	rgrue         *RuleEntry  // List of all Daylight Saving change rule entries.
	rgzonCol      *float64    // Cache of time zone offsets for each zone area.
	rgci          *CI         // List of chart information records for chart list.
	fileIn        *FILE       // The switch file currently being read from.
	S             *FILE       // File to write text to.
	T             float64     // Julian time for chart.
	MC            float64     // Midheaven at chart time.
	Asc           float64     // Ascendant at chart time.
	EP            float64     // East Point at chart time.
	Vtx           float64     // Vertex at chart time.
	RA            float64     // Right ascension at time.
	OB            float64     // Obliquity of ecliptic.
	rDeltaT       float64     // Delta-T at chart time, in days.
	jdDeltaT      float64     // JD for cached Delta-T offset above.
	rNut          float64     // Nutation offset.
}
type IS = _InternalSettings

type _OrbitalElements struct {
	ma0, ma1, ma2 float64 // Mean anomaly.
	ec0, ec1, ec2 float64 // Eccentricity.
	sma           float64 // Semi-major axis.
	ap0, ap1, ap2 float64 // Argument of perihelion.
	an0, an1, an2 float64 // Ascending node.
	in0, in1, in2 float64 // Inclination.
}
type OE = _OrbitalElements

type _PT2S struct {
	x int
	y int
}
type PT2S = _PT2S

type _PT3R struct {
	x float64
	y float64
	z float64
}
type PT3R = _PT3R

type _TimezoneRuleEntry struct {
	yea1    int  // Start year rule applies to
	yea2    int  // End year rule applies to
	mon     byte // Month in each year rule takes place
	daytype byte // Type of day (0=num, 1=lastDOW, 2=DOW>=num, 3=DOW<=num)
	daynum  byte // Day within month rule takes place
	dayweek byte // Day of week (DOW) rule takes place
	tim     int  // Time within day rule takes place
	timtype int  // Type of time (0=local, 1=standard, 2=UTC)
	dst     int  // Rule applies this Daylight offset (in seconds before UTC)
}
type RuleEntry = _TimezoneRuleEntry

type _TimezoneRuleName struct {
	szNam [cchSzZon]byte // Name of rule
	irue  int            // Start index of rule entries
}
type RuleName = _TimezoneRuleName

type _TimezoneChange struct {
	zon     int  // Time zone value (in seconds before UTC)
	irun    int  // Daylight Saving rule (if any)
	dst     int  // Daylight offset to always use (if no rule)
	yea     int  // Year time zone value ends
	mon     byte // Month time zone value ends
	day     byte // Day time zone value ends
	tim     int  // Time that time zone value ends (in seconds)
	timtype int  // Type of time (0=local, 1=standard, 2=UTC)
}
type ZoneChange = _TimezoneChange

type _UserSettings struct {
	// Chart types
	fListing       bool // -v
	fWheel         bool // -w
	fGrid          bool // -g
	fAspList       bool // -a
	fMidpoint      bool // -m
	fHorizon       bool // -Z
	fOrbit         bool // -S
	fSector        bool // -l
	fInfluence     bool // -j
	fEsoteric      bool // -7
	fAstroGraph    bool // -L
	fCalendar      bool // -K
	fInDay         bool // -d
	fInDayInf      bool // -D
	fEphemeris     bool // -E
	fArabic        bool // -P
	fHorizonSearch bool // -Zd
	fTransit       bool // -t
	fTransitInf    bool // -T
	fInDayGra      bool // -B
	fTransitGra    bool // -V
	fAtlasLook     bool // -N
	fAtlasNear     bool // -Nl
	fZoneChange    bool // -Nz
	fMoonChart     bool // -8
	// Chart suboptions
	fVelocity      bool // -v0
	fListDecan     bool // -v3
	fWheelReverse  bool // -w0
	fGridConfig    bool // -g0 (on by default)
	fGridMidpoint  bool // -gm
	nAppSep        bool // -ga
	fParallel      bool // -gp
	fDistance      bool // -gd
	fAspSummary    bool // -a0 (on by default)
	fMidSummary    bool // -m0 (on by default)
	fMidAspect     bool // -ma
	fPrimeVert     bool // -Z0
	fSectorApprox  bool // -l0
	fInfluenceSign bool // -j0 (on by default)
	fLatitudeCross bool // -L0 (on by default)
	fCalendarYear  bool // -Ky
	fInDayMonth    bool // -dm
	fInDayYear     bool // -dy
	fGraphAll      bool // -B0
	fArabicFlip    bool // -P0
	fMoonChartSep  bool // -80
	// Table chart types
	fCredit     bool // -Hc
	fSwitch     bool // -H
	fSwitchRare bool // -Y
	fKeyGraph   bool // -HX
	fSign       bool // -HC
	fObject     bool // -HO
	fAspect     bool // -HA
	fConstel    bool // -HF
	fOrbitData  bool // -HS
	fRay        bool // -H7
	fMeaning    bool // -HI
	// Main flags
	fLoop       bool // -Q
	fSidereal   bool // -s
	fCusp       bool // -C
	fUranian    bool // -u
	fDwarf      bool // -u0
	fMoons      bool // -u8
	fCOB        bool // -ub
	fStar       bool // -U
	fProgress   bool // Are we doing a -p progressed chart?
	fInterpret  bool // Is -I interpretation switch in effect?
	fHouse3D    bool // -c3
	fAspect3D   bool // -A3
	fAspectLat  bool // -Ap
	fParallel2  bool // -AP
	fDecan      bool // -3
	fFlip       bool // -f
	fGeodetic   bool // -G
	fIndian     bool // -J
	fNavamsa    bool // -9
	fEphemFiles bool // -b
	fWriteFile  bool // -o
	fAnsiColor  bool // -k
	fGraphics   bool // -X
	// Main subflags
	fNoSwitches bool
	fLoopInit   bool // -Q0
	fSeconds    bool // -b0
	fPlacalcAst bool // -ba
	fPlacalcPla bool // -bp
	fMatrixPla  bool // -bm
	fMatrixStar bool // -bU
	fEquator    bool // -sr
	fEquator2   bool // -sr0
	fAnsiChar   bool // -k0
	fTextHTML   bool // -kh
	fSolarWhole bool // -10
	fListAuto   bool // -5
	// Obscure flags
	fTruePos      bool // -YT
	fTopoPos      bool // -YV
	fRefract      bool // -Yf
	fBarycenter   bool // -Yh
	fMoonMove     bool // -Ym
	fSidereal2    bool // -Ys
	fTrueNode     bool // -Yn
	fNoNutation   bool // -Yn0
	fEuroDate     bool // -Yd
	fEuroTime     bool // -Yt
	fEuroDist     bool // -Yv
	fRound        bool // -Yr
	fSmartCusp    bool // -YC
	fSmartSave    bool // -YO
	fClip80       bool // -Y8
	fWriteOld     bool // -Yo
	fHouseAngle   bool // -Yc
	fPolarAsc     bool // -Yp
	fEclipse      bool // -Yu
	fEclipseAny   bool // -Yu0
	fObjRotWhole  bool // -Y10
	fIgnoreSign   bool // -YR0
	fIgnoreDir    bool // -YR0
	fIgnoreDiralt bool // -YR1
	fIgnoreDirlen bool // -YR1
	fIgnoreAlt0   bool // -YR2
	fIgnoreDisequ bool // -YR2
	fIgnoreAuto   bool // -YRh
	fStarsList    bool // -YRU0
	fStarMagDist  bool // -YUb
	fStarMagAbs   bool // -YUb0
	fNoWrite      bool // -0o
	fNoRead       bool // -0i
	fNoQuit       bool // -0q
	fNoGraphics   bool // -0X
	fNoPlacalc    bool // -0b
	fNoNetwork    bool // -0n
	fNoExp        bool // -0~
	fExpOff       bool // -~0
	// Value settings
	nDecanType   _decandivisiontype // -v3
	nAspectSort  int                // -a
	nEphemYears  int                // -Ey
	nEphemRate   int                // -E0
	nEphemFactor int                // -E0
	nArabicSort  int                // -P
	nRel         _relationshipchart // What relationship chart is in effect, if any?
	nSwissEph    int                // -bs
	nHouseSystem int                // -c
	nHouse3D     _housemodel        // -c3
	nAsp         int                // -A
	objCenter    _object            // -h
	nDwad        int                // -4
	nStarSort    int                // -U
	rHarmonic    float64            // Harmonic chart value passed to -x switch.
	objOnAsc     int                // Planet value passed to -1 or -2 switch.
	objRequire   int                // Required object passed to -RO switch.
	nDegForm     int                // -s
	nProgress    _progressiontype   // -p0
	nDivision    int                // -d
	nScreenWidth int                // -I
	nWriteFormat int                // -o
	nListAll     int                // -5e
	dstDef       float64            // -z0
	zonDef       float64            // -z
	lonDef       float64            // -zl
	latDef       float64            // -zl
	elvDef       float64            // -zv
	tmpDef       float64            // -zf
	namDef       []byte             // -zj
	locDef       []byte             // -zj
	rgszPath     [10][]byte         // -Yi
	szADB        []byte             // -Y5i
	szAstColor   []byte             // -YkE
	szStarsColor []byte             // -YkU
	szStarsList  []byte             // -YRU
	// Value subsettings
	nWheelRows       int     // Number of rows per house to use for -w wheel.
	nAstroGraphStep  int     // Latitude step rate passed to -L switch.
	nAstroGraphDist  int     // Maximum crossing distance passed to -L0 switch.
	nArabicParts     int     // Arabic parts to include value passed to -P.
	nAtlasList       int     // Number of rows to display value passed to -N.
	rZodiacOffset    float64 // Position shifting value passed to -s switch.
	rZodiacOffsetAll float64 // Position shifting value passed to -Ys switch.
	rProgDay         float64 // Progression day value passed to -pd switch.
	rProgCusp        float64 // Progression cusp ratio value passed to -pC.
	nRatio1          int     // Chart ratio factors passed to -rc or -rm.
	nRatio2          int
	nCharset         _charactercodepage // -Ya
	nCharsetOut      _charactercodepage // -Yao
	nScrollRow       int                // -YQ
	cSequenceLine    int                // -Yq
	lTimeAddition    int                // -Yz
	rDeltaT          float64            // -Yz0
	rObjAddition     float64            // -YzO
	rCuspAddition    float64            // -YzC
	objRot1          _object            // -Y1
	objRot2          _object            // -Y1
	nHorizon         int                // -YZ
	nArabicNight     int                // -YP
	nBioday          int                // -Yb
	nSignDiv         int                // -YRd
	iExpADB          int                // -~5i
	cExpADB          int                // -~5i
	// AstroExpression hooks
	szExpConfig  []byte // -~g
	szExpAspList []byte // -~a
	szExpAspSumm []byte // -~a0
	szExpMid     []byte // -~m
	szExpMidAsp  []byte // -~ma
	szExpInf     []byte // -~j
	szExpInf0    []byte // -~j0
	szExpEso     []byte // -~7
	szExpCross   []byte // -~L
	szExpEph     []byte // -~E
	szExpRis     []byte // -~Zd
	szExpDay     []byte // -~d
	szExpVoid    []byte // -~dv
	szExpTra     []byte // -~t
	szExpPart    []byte // -~P
	szExpObj     []byte // -~O
	szExpHou     []byte // -~C
	szExpAsp     []byte // -~A
	szExpProg    []byte // -~p
	szExpProg0   []byte // -~p0
	szExpColObj  []byte // -~kO
	szExpColAsp  []byte // -~kA
	szExpColFill []byte // -~kv
	szExpFontSig []byte // -~F
	szExpFontHou []byte // -~FC
	szExpFontObj []byte // -~FO
	szExpFontAsp []byte // -~FA
	szExpSort    []byte // -~v
	szExpDecan   []byte // -~v3
	szExpStar    []byte // -~U
	szExpAst     []byte // -~U0
	szExpCity    []byte // -~XL
	szExpSidebar []byte // -~Xt
	szExpKey     []byte // -~XQ
	szExpMenu    []byte // -~WQ
	szExpCast1   []byte // -~q1
	szExpCast2   []byte // -~q2
	szExpDisp1   []byte // -~Q1
	szExpDisp2   []byte // -~Q2
	szExpDisp3   []byte // -~Q3
	szExpListS   []byte // -~5s
	szExpListF   []byte // -~5f
	szExpListY   []byte // -~5Y
	szExpADB     []byte // -~5i
}
type US = _UserSettings

// DATA CONFIGURATION SECTION: These settings describe particulars of
// your own location and where the program looks for certain info. It is
// recommended that these values be changed appropriately, although the
// program will still run if they are left alone. They may also be
// specified within the default settings file or in the program itself.
const (
	// Change numbers to longitude and latitude of your current location. Use
	// negative values for eastern or southern degrees.
	DEFAULT_LONG = 122.0 + 20.0/60.0 // was DM(122, 20)
	DEFAULT_LAT  = 47.0 + 36.0/60.0  // was DM(47, 36)

	// Change this number to the time zone of your current location in hours
	// before (west of) UTC. Use negative values for eastern zones.
	DEFAULT_ZONE = 8.00
)

// OPTIONAL CONFIGURATION SECTION: Although not necessary, one may like
// to change some of the values below: These constants affect some of
// the default parameters and related options. They may also be
// specified within the default settings file or in the program itself.
const (
	// Normally, Placidus houses are used (unless the user specifies otherwise).
	// If you want a different default system, change this number to a value
	// from 0-21 (values same as in -c switch).
	DEFAULT_SYSTEM = 0

	// Default number of aspects to include in charts.
	DEFAULT_ASPECTS = 5

	// Greater numbers means more accuracy but slower calculation, of exact
	// aspect and transit times.
	DIVISIONS = 48
)

func DFromR(x float64) float64 {
	return x * rDegRad
}

func DM(d, m float64) float64     { return d + m/60.0 }
func DMS(d, m, s float64) float64 { return DM(d, m) + s/3600.0 }

func FBetween(v, v1, v2 int) bool {
	return v1 <= v && v <= v2
}

func HM(h, m float64) float64     { return h + m/60.0 }
func HMS(h, m, s float64) float64 { return HM(h, m) + s/3600.0 }

func RAbs(x float64) float64 {
	return math.Abs(x)
}

func RACos(x float64) float64 {
	return math.Acos(x)
}

func RAngleD(x, y float64) float64 {
	return DFromR(RAngle(x, y))
}

func RASin(x float64) float64 {
	return math.Asin(x)
}

func RAtn(x float64) float64 {
	return math.Atan(x)
}

func RCos(x float64) float64 {
	return math.Cos(x)
}

func RCosD(x float64) float64 {
	return math.Cos(RFromD(x))
}

// RFloor rounds to the largest integral value not greater than x
func RFloor(x float64) float64 {
	return math.Floor(x)
}

func RFromD(x float64) float64 {
	// TODO: performance - turn divide into multiply?
	return (x) / rDegRad
}

func RFract(x float64) float64 {
	return x - math.Floor(x)
}

func RLength2(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func RLength3(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}

func RSgn2(x float64) float64 {
	if x < 0 {
		return -1
	}
	return 1
}

func RSin(x float64) float64 {
	return math.Sin(x)
}

func RSinD(x float64) float64 {
	return math.Sin(RFromD(x))
}

func RSqr(x float64) float64 {
	return math.Sqrt(x)
}

func RTan(x float64) float64 {
	return math.Tan(x)
}

func RTanD(x float64) float64 {
	return math.Tan(RFromD(x))
}

func Sq(n float64) float64 {
	return n * n
}
