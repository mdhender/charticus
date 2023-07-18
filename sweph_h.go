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

// TODO: AS_BOOL is sometimes bool and sometimes it is int.

const (
	SEI_CURR_FPOS     = -1
	SEI_FILE_NMAXPLAN = 50
	SEI_NEPHFILES     = 7
	SEI_NMODELS       = 8
	SEI_NNODE_ETC     = 6

	SEI_EPSILON  = -2
	SEI_NUTATION = -1
	SEI_EMB      = 0
	SEI_EARTH    = 0
	SEI_SUN      = 0
	SEI_MOON     = 1
	SEI_MERCURY  = 2
	SEI_VENUS    = 3
	SEI_MARS     = 4
	SEI_JUPITER  = 5
	SEI_SATURN   = 6
	SEI_URANUS   = 7
	SEI_NEPTUNE  = 8
	SEI_PLUTO    = 9
	SEI_SUNBARY  = 10 // barycentric sun
	SEI_ANYBODY  = 11 // any asteroid
	SEI_CHIRON   = 12
	SEI_PHOLUS   = 13
	SEI_CERES    = 14
	SEI_PALLAS   = 15
	SEI_JUNO     = 16
	SEI_VESTA    = 17
	SEI_NPLANETS = 18
)

const (
	SWI_STAR_LENGTH = 40
)

// obliquity of ecliptic
type epsilon struct {
	teps float64 // jd
	eps  float64 // eps
	seps float64 // sin(eps)
	ceps float64 // cos(eps)
}

type file_data struct {
	fnam        [AS_MAXCH]byte         /* ephemeris file name */
	fversion    int                    /* version number of file */
	astnam      [50]byte               /* asteroid name, if asteroid file */
	sweph_denum int32                  /* DE number of JPL ephemeris, which this file is derived from. */
	fptr        *FILE                  /* ephemeris file pointer */
	tfstart     float64                /* file may be used from this date */
	tfend       float64                /*      through this date          */
	iflg        int32                  /* byte reorder flag and little/bigendian flag */
	npl         int32                  /* how many planets in file */
	ipl         [SEI_FILE_NMAXPLAN]int /* planet numbers */
}

type fixed_star struct {
	skey      [SWI_STAR_LENGTH + 2]byte // may be prefixed with comma, one char more
	starname  [SWI_STAR_LENGTH + 1]byte
	starbayer [SWI_STAR_LENGTH + 1]byte
	starno    [10]byte
	epoch     float64
	ra        float64
	de        float64
	ramot     float64
	demot     float64
	radvel    float64
	parall    float64
	mag       float64
}

type gen_const struct {
	clight       float64
	aunit        float64
	helgravconst float64
	ratme        float64
	sunradius    float64
}

type interpol struct {
	tjd_nut0  float64
	tjd_nut2  float64
	nut_dpsi0 float64
	nut_dpsi1 float64
	nut_dpsi2 float64
	nut_deps0 float64
	nut_deps1 float64
	nut_deps2 float64
}

type node_data struct {
	// result of most recent data evaluation for this body:
	teval float64    // time for which last computation was made
	iephe int32      // which ephemeris was used
	x     [6]float64 // position and speed vectors equatorial J2000
	xflgs int32      // hel., light-time, aberr., prec. flags etc.
	// return positions:
	//   xreturn+0	ecliptic polar coordinates
	//   xreturn+6	ecliptic cartesian coordinates
	//   xreturn+12	equatorial polar coordinates
	//   xreturn+18	equatorial cartesian coordinates
	xreturn [24]float64
}

// nutation
type nut struct {
	tnut   float64
	nutlo  [2]float64 // nutation in longitude and obliquity
	snut   float64    // sine of nutation in obliquity
	cnut   float64    // cosine of nutation in obliquity
	matrix [3][3]float64
}

type plan_data struct {
	// the following data are read from file only once, immediately after file has been opened
	ibdy int /* internal body number */
	// contains several bit flags describing the data:
	//   SEI_FLG_HELIO: true if helio, false if bary
	//   SEI_FLG_ROTATE: TRUE if coefficients are referred to coordinate system of orbital plane
	//   SEI_FLG_ELLIPSE: TRUE if reference ellipse
	iflg int32
	ncoe int // # of coefficients of ephemeris polynomial, is polynomial order + 1
	// where is the segment index on the file
	lndx0 int32 // file position of begin of planet's index
	nndx  int32 // number of index entries on file: computed
	// file contains ephemeris for tfstart thru tfend for this particular planet !!!
	tfstart float64
	tfend   float64
	dseg    float64 // segment size (days covered by a polynomial)  */
	// orbital elements:
	telem float64 // epoch of elements
	prot  float64
	qrot  float64
	dprot float64
	dqrot float64
	rmax  float64 // normalisation factor of cheby coefficients
	// in addition, if reference ellipse is used:
	peri  float64
	dperi float64
	// pointer to cheby coeffs of reference ellipse, size of data is 2 x ncoe
	refep *float64
	// unpacked segment information, only updated when a segment is read:
	tseg0 float64  // start jd of current segment
	tseg1 float64  // end jd of current segment
	segp  *float64 // pointer to unpacked cheby coeffs of segment; the size is 3 x ncoe
	neval int      // how many coefficients to evaluate. this may be less than ncoe
	// result of most recent data evaluation for this body:
	teval float64    // time for which previous computation was made
	iephe int32      // which ephemeris was used
	x     [6]float64 // position and speed vectors equatorial J2000
	xflgs int32      // hel., light-time, aberr., prec. flags etc.
	// return positions:
	//   xreturn+0	ecliptic polar coordinates
	//   xreturn+6	ecliptic cartesian coordinates
	//   xreturn+12	equatorial polar coordinates
	//   xreturn+18	equatorial cartesian coordinates
	xreturn [24]float64
}

type save_positions struct {
	ipl      int
	tsave    float64
	iflgsave int32
	// position at t = tsave,
	// in ecliptic polar (offset 0),
	//    ecliptic cartesian (offset 6),
	//    equatorial polar (offset 12),
	//    and equatorial cartesian coordinates (offset 18).
	// 6 float64s each for position and speed coordinates.
	//
	xsaves [24]float64
}

type sid_data struct {
	sid_mode int32
	ayan_t0  float64
	t0       float64
	t0_is_UT bool
}

// if this is changed, then also update initialisation in sweph.cpp
type swe_data struct {
	ephe_path_is_set       bool
	jpl_file_is_open       bool
	fixfp                  *FILE  // fixed stars file pointer
	ephepath               string // was [AS_MAXCH]byte
	jplfnam                [AS_MAXCH]byte
	jpldenum               int32
	last_epheflag          int32
	geopos_is_set          bool
	ayana_is_set           bool
	is_old_starfile        bool
	eop_tjd_beg            float64
	eop_tjd_beg_horizons   float64
	eop_tjd_end            float64
	eop_tjd_end_add        float64
	eop_dpsi_loaded        int
	tid_acc                float64
	is_tid_acc_manual      bool
	init_dt_done           bool
	swed_is_initialised    bool
	delta_t_userdef_is_set bool
	delta_t_userdef        float64
	ast_G                  float64
	ast_H                  float64
	ast_diam               float64
	astelem                [AS_MAXCH * 10]byte
	i_saved_planet_name    int
	saved_planet_name      [80]byte
	dpsi                   *float64
	deps                   *float64
	timeout                int32
	astro_models           [SEI_NMODELS]int32
	do_interpolate_nut     bool
	interpol               interpol
	fidat                  [SEI_NEPHFILES]file_data
	gcdat                  gen_const
	pldat                  [SEI_NPLANETS]plan_data
	nddat                  [SEI_NNODE_ETC]plan_data
	savedat                [SE_NPLANETS + 1]save_positions
	oec                    epsilon
	oec2000                epsilon
	nut                    nut
	nut2000                nut
	nutv                   nut
	topd                   topo_data
	sidd                   sid_data
	n_fixstars_real        int // real number of fixed stars in sefstars.txt
	n_fixstars_named       int // number of fixed stars with tradtional name
	n_fixstars_records     int // number of fixed stars records in fixed_stars
	fixed_stars            *fixed_star
}

type topo_data struct {
	geolon float64
	geolat float64
	geoalt float64
	teval  float64
	tjd_ut float64
	xobs   [6]float64
}
