// Code generated for linux/arm64 by 'ccgo preprocessed.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_diff

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 21
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 21
const __FLT64X_DENORM_MIN__ = 3.64519953188247460252840593361941982e-4951
const __FLT64X_DIG__ = 18
const __FLT64X_EPSILON__ = 1.08420217248550443400745280086994171e-19
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 64
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 21
const __LDBL_DENORM_MIN__ = 3.64519953188247460252840593361941982e-4951
const __LDBL_DIG__ = 18
const __LDBL_EPSILON__ = 1.08420217248550443400745280086994171e-19
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 64
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const __LDBL_NORM_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __VERSION__ = "13.3.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type __u_char = uint8

type __u_short = uint16

type __u_int = uint32

type __u_long = uint64

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __int_least8_t = int8

type __uint_least8_t = uint8

type __int_least16_t = int16

type __uint_least16_t = uint16

type __int_least32_t = int32

type __uint_least32_t = uint32

type __int_least64_t = int64

type __uint_least64_t = uint64

type __quad_t = int64

type __u_quad_t = uint64

type __intmax_t = int64

type __uintmax_t = uint64

type __dev_t = uint64

type __uid_t = uint32

type __gid_t = uint32

type __ino_t = uint64

type __ino64_t = uint64

type __mode_t = uint32

type __nlink_t = uint64

type __off_t = int64

type __off64_t = int64

type __pid_t = int32

type __fsid_t = struct {
	F__val [2]int32
}

type __clock_t = int64

type __rlim_t = uint64

type __rlim64_t = uint64

type __id_t = uint32

type __time_t = int64

type __useconds_t = uint32

type __suseconds_t = int64

type __suseconds64_t = int64

type __daddr_t = int32

type __key_t = int32

type __clockid_t = int32

type __timer_t = uintptr

type __blksize_t = int64

type __blkcnt_t = int64

type __blkcnt64_t = int64

type __fsblkcnt_t = uint64

type __fsblkcnt64_t = uint64

type __fsfilcnt_t = uint64

type __fsfilcnt64_t = uint64

type __fsword_t = int64

type __ssize_t = int64

type __syscall_slong_t = int64

type __syscall_ulong_t = uint64

type __loff_t = int64

type __caddr_t = uintptr

type __intptr_t = int64

type __socklen_t = uint32

type __sig_atomic_t = int32

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast8_t = int8

type int_fast16_t = int64

type int_fast32_t = int64

type int_fast64_t = int64

type uint_fast8_t = uint8

type uint_fast16_t = uint64

type uint_fast32_t = uint64

type uint_fast64_t = uint64

type intptr_t = int64

type uintptr_t = uint64

type intmax_t = int64

type uintmax_t = uint64

type size_t = uint64

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type u_char = uint8

type u_short = uint16

type u_int = uint32

type u_long = uint64

type quad_t = int64

type u_quad_t = uint64

type fsid_t = struct {
	F__val [2]int32
}

type loff_t = int64

type ino_t = uint64

type dev_t = uint64

type gid_t = uint32

type mode_t = uint32

type nlink_t = uint64

type uid_t = uint32

type off_t = int64

type pid_t = int32

type id_t = uint32

type ssize_t = int64

type daddr_t = int32

type caddr_t = uintptr

type key_t = int32

type clock_t = int64

type clockid_t = int32

type time_t = int64

type timer_t = uintptr

type ulong = uint64

type ushort = uint16

type uint1 = uint32

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int32

func __bswap_16(tls *libc.TLS, __bsx __uint16_t) (r __uint16_t) {
	return libc.X__builtin_bswap16(tls, __bsx)
}

func __bswap_32(tls *libc.TLS, __bsx __uint32_t) (r __uint32_t) {
	return libc.X__builtin_bswap32(tls, __bsx)
}

func __bswap_64(tls *libc.TLS, __bsx __uint64_t) (r __uint64_t) {
	return libc.X__builtin_bswap64(tls, __bsx)
}

func __uint16_identity(tls *libc.TLS, __x __uint16_t) (r __uint16_t) {
	return __x
}

func __uint32_identity(tls *libc.TLS, __x __uint32_t) (r __uint32_t) {
	return __x
}

func __uint64_identity(tls *libc.TLS, __x __uint64_t) (r __uint64_t) {
	return __x
}

type __sigset_t = struct {
	F__val [16]uint64
}

type sigset_t = struct {
	F__val [16]uint64
}

type timeval = struct {
	Ftv_sec  __time_t
	Ftv_usec __suseconds_t
}

type timespec = struct {
	Ftv_sec  __time_t
	Ftv_nsec __syscall_slong_t
}

type suseconds_t = int64

type __fd_mask = int64

type fd_set = struct {
	F__fds_bits [16]__fd_mask
}

type fd_mask = int64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type __atomic_wide_counter = struct {
	F__value32 [0]struct {
		F__low  uint32
		F__high uint32
	}
	F__value64 uint64
}

type __pthread_list_t = struct {
	F__prev uintptr
	F__next uintptr
}

type __pthread_internal_list = __pthread_list_t

type __pthread_slist_t = struct {
	F__next uintptr
}

type __pthread_internal_slist = __pthread_slist_t

type __pthread_mutex_s = struct {
	F__lock    int32
	F__count   uint32
	F__owner   int32
	F__nusers  uint32
	F__kind    int32
	F__spins   int16
	F__elision int16
	F__list    __pthread_list_t
}

type __pthread_rwlock_arch_t = struct {
	F__readers       uint32
	F__writers       uint32
	F__wrphase_futex uint32
	F__writers_futex uint32
	F__pad3          uint32
	F__pad4          uint32
	F__cur_writer    int32
	F__shared        int32
	F__rwelision     int8
	F__pad1          [7]uint8
	F__pad2          uint64
	F__flags         uint32
}

type __pthread_cond_s = struct {
	F__wseq         __atomic_wide_counter
	F__g1_start     __atomic_wide_counter
	F__g_refs       [2]uint32
	F__g_size       [2]uint32
	F__g1_orig_size uint32
	F__wrefs        uint32
	F__g_signals    [2]uint32
}

type __tss_t = uint32

type __thrd_t = uint64

type __once_flag = struct {
	F__data int32
}

type pthread_t = uint64

type pthread_mutexattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type pthread_condattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type pthread_key_t = uint32

type pthread_once_t = int32

type pthread_attr_t = struct {
	F__align [0]int64
	F__size  [56]uint8
}

type pthread_mutex_t = struct {
	F__size  [0][40]uint8
	F__align [0]int64
	F__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	F__size  [0][48]uint8
	F__align [0]int64
	F__data  __pthread_cond_s
}

type pthread_rwlock_t = struct {
	F__size  [0][56]uint8
	F__align [0]int64
	F__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	F__align [0]int64
	F__size  [8]uint8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	F__align [0]int64
	F__size  [32]uint8
}

type pthread_barrierattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type random_data = struct {
	Ffptr      uintptr
	Frptr      uintptr
	Fstate     uintptr
	Frand_type int32
	Frand_deg  int32
	Frand_sep  int32
	Fend_ptr   uintptr
}

type drand48_data = struct {
	F__x     [3]uint16
	F__old_x [3]uint16
	F__c     uint16
	F__init  uint16
	F__a     uint64
}

type __compar_fn_t = uintptr

type TSStateId = uint16

type TSSymbol = uint16

type TSFieldId = uint16

type TSLanguage = struct {
	Fabi_version               uint32_t
	Fsymbol_count              uint32_t
	Falias_count               uint32_t
	Ftoken_count               uint32_t
	Fexternal_token_count      uint32_t
	Fstate_count               uint32_t
	Flarge_state_count         uint32_t
	Fproduction_id_count       uint32_t
	Ffield_count               uint32_t
	Fmax_alias_sequence_length uint16_t
	Fparse_table               uintptr
	Fsmall_parse_table         uintptr
	Fsmall_parse_table_map     uintptr
	Fparse_actions             uintptr
	Fsymbol_names              uintptr
	Ffield_names               uintptr
	Ffield_map_slices          uintptr
	Ffield_map_entries         uintptr
	Fsymbol_metadata           uintptr
	Fpublic_symbol_map         uintptr
	Falias_map                 uintptr
	Falias_sequences           uintptr
	Flex_modes                 uintptr
	Flex_fn                    uintptr
	Fkeyword_lex_fn            uintptr
	Fkeyword_capture_token     TSSymbol
	Fexternal_scanner          struct {
		Fstates      uintptr
		Fsymbol_map  uintptr
		Fcreate      uintptr
		Fdestroy     uintptr
		Fscan        uintptr
		Fserialize   uintptr
		Fdeserialize uintptr
	}
	Fprimary_state_ids          uintptr
	Fname                       uintptr
	Freserved_words             uintptr
	Fmax_reserved_word_set_size uint16_t
	Fsupertype_count            uint32_t
	Fsupertype_symbols          uintptr
	Fsupertype_map_slices       uintptr
	Fsupertype_map_entries      uintptr
	Fmetadata                   TSLanguageMetadata
}

type TSLanguageMetadata = struct {
	Fmajor_version uint8_t
	Fminor_version uint8_t
	Fpatch_version uint8_t
}

type TSFieldMapEntry = struct {
	Ffield_id    TSFieldId
	Fchild_index uint8_t
	Finherited   uint8
}

type TSMapSlice = struct {
	Findex  uint16_t
	Flength uint16_t
}

type TSSymbolMetadata = struct {
	Fvisible   uint8
	Fnamed     uint8
	Fsupertype uint8
}

type TSLexer = struct {
	Flookahead                  int32_t
	Fresult_symbol              TSSymbol
	Fadvance                    uintptr
	Fmark_end                   uintptr
	Fget_column                 uintptr
	Fis_at_included_range_start uintptr
	Feof                        uintptr
	Flog                        uintptr
}

type TSParseActionType = int32

const TSParseActionTypeShift = 0
const TSParseActionTypeReduce = 1
const TSParseActionTypeAccept = 2
const TSParseActionTypeRecover = 3

type TSParseAction = struct {
	Freduce [0]struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}
	Ftype_token [0]uint8_t
	Fshift      struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}
	F__ccgo_pad3 [2]byte
}

type TSLexMode = struct {
	Flex_state          uint16_t
	Fexternal_lex_state uint16_t
}

type TSLexerMode = struct {
	Flex_state            uint16_t
	Fexternal_lex_state   uint16_t
	Freserved_word_set_id uint16_t
}

type TSParseActionEntry = struct {
	Fentry [0]struct {
		Fcount    uint8_t
		Freusable uint8
	}
	Faction TSParseAction
}

type TSCharacterRange = struct {
	Fstart int32_t
	Fend   int32_t
}

func set_contains(tls *libc.TLS, ranges uintptr, len1 uint32_t, lookahead int32_t) (r uint8) {
	var half_size, index, mid_index, size uint32_t
	var range_token, range_token1 uintptr
	_, _, _, _, _, _ = half_size, index, mid_index, range_token, range_token1, size
	index = uint32(0)
	size = len1 - index
	for size > uint32(1) {
		half_size = size / uint32(2)
		mid_index = index + half_size
		range_token = ranges + uintptr(mid_index)*8
		if lookahead >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
			return libc.BoolUint8(1 != 0)
		} else {
			if lookahead > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				index = mid_index
			}
		}
		size = size - half_size
	}
	range_token1 = ranges + uintptr(index)*8
	return libc.BoolUint8(lookahead >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
}

type ts_symbol_identifiers = int32

const aux_sym_source_token1 = 1
const anon_sym_diff = 2
const aux_sym_command_token1 = 3
const anon_sym_new = 4
const anon_sym_deleted = 5
const anon_sym_file = 6
const anon_sym_mode = 7
const anon_sym_old = 8
const anon_sym_rename = 9
const anon_sym_from = 10
const anon_sym_to = 11
const anon_sym_Binary = 12
const anon_sym_files = 13
const anon_sym_and = 14
const anon_sym_differ = 15
const anon_sym_index = 16
const anon_sym_DOT_DOT = 17
const anon_sym_similarity = 18
const anon_sym_index2 = 19
const aux_sym_similarity_token1 = 20
const anon_sym_PERCENT = 21
const anon_sym_DASH_DASH_DASH = 22
const anon_sym_PLUS_PLUS_PLUS = 23
const anon_sym_AT_AT = 24
const anon_sym_AT_AT2 = 25
const aux_sym_location_token1 = 26
const anon_sym_PLUS = 27
const anon_sym_PLUS_PLUS = 28
const anon_sym_PLUS_PLUS_PLUS_PLUS = 29
const anon_sym_DASH = 30
const anon_sym_DASH_DASH = 31
const anon_sym_DASH_DASH_DASH_DASH = 32
const sym_context = 33
const anon_sym_POUND = 34
const sym_linerange = 35
const aux_sym_filename_token1 = 36
const sym_commit = 37
const sym_source = 38
const sym__line = 39
const sym_block = 40
const sym_hunks = 41
const sym_hunk = 42
const sym_changes = 43
const sym_command = 44
const sym_file_change = 45
const sym_binary_change = 46
const sym_index = 47
const sym_similarity = 48
const sym_old_file = 49
const sym_new_file = 50
const sym_location = 51
const sym_addition = 52
const sym_deletion = 53
const sym_comment = 54
const sym_filename = 55
const sym_mode = 56
const aux_sym_source_repeat1 = 57
const aux_sym_block_repeat1 = 58
const aux_sym_hunks_repeat1 = 59
const aux_sym_changes_repeat1 = 60
const aux_sym_changes_repeat2 = 61
const aux_sym_filename_repeat1 = 62
const alias_sym_score = 63

var ts_symbol_names = [64]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 18,
	3:  __ccgo_ts + 23,
	4:  __ccgo_ts + 32,
	5:  __ccgo_ts + 36,
	6:  __ccgo_ts + 44,
	7:  __ccgo_ts + 49,
	8:  __ccgo_ts + 54,
	9:  __ccgo_ts + 58,
	10: __ccgo_ts + 65,
	11: __ccgo_ts + 70,
	12: __ccgo_ts + 73,
	13: __ccgo_ts + 80,
	14: __ccgo_ts + 86,
	15: __ccgo_ts + 90,
	16: __ccgo_ts + 97,
	17: __ccgo_ts + 103,
	18: __ccgo_ts + 106,
	19: __ccgo_ts + 97,
	20: __ccgo_ts + 117,
	21: __ccgo_ts + 135,
	22: __ccgo_ts + 137,
	23: __ccgo_ts + 141,
	24: __ccgo_ts + 145,
	25: __ccgo_ts + 145,
	26: __ccgo_ts + 148,
	27: __ccgo_ts + 164,
	28: __ccgo_ts + 166,
	29: __ccgo_ts + 169,
	30: __ccgo_ts + 174,
	31: __ccgo_ts + 176,
	32: __ccgo_ts + 179,
	33: __ccgo_ts + 184,
	34: __ccgo_ts + 192,
	35: __ccgo_ts + 194,
	36: __ccgo_ts + 204,
	37: __ccgo_ts + 220,
	38: __ccgo_ts + 227,
	39: __ccgo_ts + 234,
	40: __ccgo_ts + 240,
	41: __ccgo_ts + 246,
	42: __ccgo_ts + 252,
	43: __ccgo_ts + 257,
	44: __ccgo_ts + 265,
	45: __ccgo_ts + 273,
	46: __ccgo_ts + 285,
	47: __ccgo_ts + 97,
	48: __ccgo_ts + 106,
	49: __ccgo_ts + 299,
	50: __ccgo_ts + 308,
	51: __ccgo_ts + 317,
	52: __ccgo_ts + 326,
	53: __ccgo_ts + 335,
	54: __ccgo_ts + 344,
	55: __ccgo_ts + 352,
	56: __ccgo_ts + 49,
	57: __ccgo_ts + 361,
	58: __ccgo_ts + 376,
	59: __ccgo_ts + 390,
	60: __ccgo_ts + 404,
	61: __ccgo_ts + 420,
	62: __ccgo_ts + 436,
	63: __ccgo_ts + 453,
}

var ts_symbol_map = [64]TSSymbol{
	1:  uint16(aux_sym_source_token1),
	2:  uint16(anon_sym_diff),
	3:  uint16(aux_sym_command_token1),
	4:  uint16(anon_sym_new),
	5:  uint16(anon_sym_deleted),
	6:  uint16(anon_sym_file),
	7:  uint16(anon_sym_mode),
	8:  uint16(anon_sym_old),
	9:  uint16(anon_sym_rename),
	10: uint16(anon_sym_from),
	11: uint16(anon_sym_to),
	12: uint16(anon_sym_Binary),
	13: uint16(anon_sym_files),
	14: uint16(anon_sym_and),
	15: uint16(anon_sym_differ),
	16: uint16(anon_sym_index),
	17: uint16(anon_sym_DOT_DOT),
	18: uint16(anon_sym_similarity),
	19: uint16(anon_sym_index),
	20: uint16(aux_sym_similarity_token1),
	21: uint16(anon_sym_PERCENT),
	22: uint16(anon_sym_DASH_DASH_DASH),
	23: uint16(anon_sym_PLUS_PLUS_PLUS),
	24: uint16(anon_sym_AT_AT),
	25: uint16(anon_sym_AT_AT),
	26: uint16(aux_sym_location_token1),
	27: uint16(anon_sym_PLUS),
	28: uint16(anon_sym_PLUS_PLUS),
	29: uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	30: uint16(anon_sym_DASH),
	31: uint16(anon_sym_DASH_DASH),
	32: uint16(anon_sym_DASH_DASH_DASH_DASH),
	33: uint16(sym_context),
	34: uint16(anon_sym_POUND),
	35: uint16(sym_linerange),
	36: uint16(aux_sym_filename_token1),
	37: uint16(sym_commit),
	38: uint16(sym_source),
	39: uint16(sym__line),
	40: uint16(sym_block),
	41: uint16(sym_hunks),
	42: uint16(sym_hunk),
	43: uint16(sym_changes),
	44: uint16(sym_command),
	45: uint16(sym_file_change),
	46: uint16(sym_binary_change),
	47: uint16(sym_index),
	48: uint16(sym_similarity),
	49: uint16(sym_old_file),
	50: uint16(sym_new_file),
	51: uint16(sym_location),
	52: uint16(sym_addition),
	53: uint16(sym_deletion),
	54: uint16(sym_comment),
	55: uint16(sym_filename),
	56: uint16(sym_mode),
	57: uint16(aux_sym_source_repeat1),
	58: uint16(aux_sym_block_repeat1),
	59: uint16(aux_sym_hunks_repeat1),
	60: uint16(aux_sym_changes_repeat1),
	61: uint16(aux_sym_changes_repeat2),
	62: uint16(aux_sym_filename_repeat1),
	63: uint16(alias_sym_score),
}

var ts_symbol_metadata = [64]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {},
	2: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	15: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	20: {},
	21: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	26: {},
	27: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	36: {},
	37: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	39: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {},
	63: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
}

type ts_field_identifiers = int32

const field_changes = 1
const field_location = 2

var ts_field_names = [3]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 257,
	2: __ccgo_ts + 317,
}

var ts_field_map_slices = [4]TSMapSlice{
	2: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [3]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_location),
	},
	1: {
		Ffield_id:    uint16(field_changes),
		Fchild_index: uint8(2),
	},
	2: {
		Ffield_id: uint16(field_location),
	},
}

var ts_alias_sequences = [4][8]TSSymbol{
	0: {},
	1: {
		2: uint16(alias_sym_score),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [89]TSStateId{
	1:  uint16(1),
	2:  uint16(2),
	3:  uint16(3),
	4:  uint16(4),
	5:  uint16(5),
	6:  uint16(6),
	7:  uint16(7),
	8:  uint16(8),
	9:  uint16(9),
	10: uint16(10),
	11: uint16(11),
	12: uint16(12),
	13: uint16(13),
	14: uint16(14),
	15: uint16(15),
	16: uint16(16),
	17: uint16(17),
	18: uint16(18),
	19: uint16(19),
	20: uint16(20),
	21: uint16(21),
	22: uint16(22),
	23: uint16(23),
	24: uint16(24),
	25: uint16(25),
	26: uint16(26),
	27: uint16(27),
	28: uint16(28),
	29: uint16(29),
	30: uint16(30),
	31: uint16(31),
	32: uint16(32),
	33: uint16(33),
	34: uint16(34),
	35: uint16(35),
	36: uint16(24),
	37: uint16(25),
	38: uint16(24),
	39: uint16(25),
	40: uint16(40),
	41: uint16(41),
	42: uint16(42),
	43: uint16(43),
	44: uint16(44),
	45: uint16(45),
	46: uint16(46),
	47: uint16(47),
	48: uint16(48),
	49: uint16(49),
	50: uint16(50),
	51: uint16(51),
	52: uint16(52),
	53: uint16(53),
	54: uint16(54),
	55: uint16(55),
	56: uint16(56),
	57: uint16(57),
	58: uint16(58),
	59: uint16(59),
	60: uint16(60),
	61: uint16(61),
	62: uint16(62),
	63: uint16(63),
	64: uint16(64),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(68),
	69: uint16(69),
	70: uint16(70),
	71: uint16(71),
	72: uint16(72),
	73: uint16(73),
	74: uint16(74),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(80),
	81: uint16(81),
	82: uint16(82),
	83: uint16(83),
	84: uint16(84),
	85: uint16(85),
	86: uint16(86),
	87: uint16(87),
	88: uint16(88),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4, i5 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, i5, lookahead, result, skip
	result = libc.BoolUint8(0 != 0)
	skip = libc.BoolUint8(0 != 0)
	eof = libc.BoolUint8(0 != 0)
	goto start
	goto next_state
next_state:
	;
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, skip)
	goto start
start:
	;
	skip = libc.BoolUint8(0 != 0)
	lookahead = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(88)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('+') {
			state = uint16(3)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('+') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('.') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('@') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('@') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('a') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('a') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(8)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('a') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('d') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('d') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('d') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('d') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('d') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('d') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('d') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('e') {
			state = uint16(45)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('e') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('e') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('e') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('e') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('e') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('e') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('e') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('e') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('f') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('f') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(33)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('e') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('f') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('f') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('i') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('i') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('i') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(60)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('i') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('i') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('i') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('i') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('i') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('l') {
			state = uint16(21)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('l') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('l') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('l') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('l') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('l') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('m') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('m') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('m') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('n') {
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('n') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('n') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('n') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('n') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('o') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('o') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('o') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('r') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('r') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('r') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('s') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('t') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('t') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('w') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('x') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('x') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('y') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('y') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(74):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(75):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(76):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(77):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(78):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(79):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(80):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(81):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(82):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(83):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(84):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _4
		_4:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(85)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(85):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _5
		_5:
			;
			i4 = i4 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(85)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(86):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(86)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(87):
		if eof != 0 {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			state = uint16(157)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_source_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_diff)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_command_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_new)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_deleted)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_file)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_file)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_old)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rename)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_from)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_to)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Binary)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_files)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_and)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_and)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_differ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_differ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_index)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(109):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_similarity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_index2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(248)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(215)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(76)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(216)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(217)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(218)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(79)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(219)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(220)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(221)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(222)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(223)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(224)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(225)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(226)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(227)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(228)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(229)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(230)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(231)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(232)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(233)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(234)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(235)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(236)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(237)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(239)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(240)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(241)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(242)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(243)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(244)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(245)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(246)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(247)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_similarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(152):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(154):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT_AT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(157):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_location_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			state = uint16(157)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_location_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i5 = i5 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('@') {
			state = uint16(155)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(195)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(168):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(191)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(196)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(97)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(93)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(176)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(98)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(174):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(198)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(200)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(177):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(188)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(182)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(188)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(194)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(180):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(171)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(90)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(182):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(181)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(190)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(193)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(197)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(189)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(170)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(175)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(169)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(186)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(173)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(172)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(167)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(201)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(185)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(202)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(180)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('x') {
			state = uint16(107)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(101)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(109)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_context)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(205):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_linerange)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(',') {
			state = uint16(75)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_linerange)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(104)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(213)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(208)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(209)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(210)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(212):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(207)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(106)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_filename_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(216):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(232):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(240):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_commit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(247)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token = [44]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(204),
	6:  uint16('%'),
	7:  uint16(151),
	8:  uint16('+'),
	9:  uint16(159),
	10: uint16('-'),
	11: uint16(162),
	12: uint16('.'),
	13: uint16(4),
	14: uint16('@'),
	15: uint16(5),
	16: uint16('B'),
	17: uint16(38),
	18: uint16('a'),
	19: uint16(54),
	20: uint16('d'),
	21: uint16(18),
	22: uint16('f'),
	23: uint16(39),
	24: uint16('i'),
	25: uint16(56),
	26: uint16('m'),
	27: uint16(61),
	28: uint16('n'),
	29: uint16(20),
	30: uint16('o'),
	31: uint16(46),
	32: uint16('r'),
	33: uint16(29),
	34: uint16('s'),
	35: uint16(37),
	36: uint16('t'),
	37: uint16(59),
	38: uint16('b'),
	39: uint16(81),
	40: uint16('c'),
	41: uint16(81),
	42: uint16('e'),
	43: uint16(81),
}

var map_token1 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(204),
	6:  uint16('+'),
	7:  uint16(159),
	8:  uint16('-'),
	9:  uint16(162),
	10: uint16('@'),
	11: uint16(166),
	12: uint16('B'),
	13: uint16(184),
	14: uint16('d'),
	15: uint16(177),
	16: uint16('i'),
	17: uint16(192),
	18: uint16('n'),
	19: uint16(174),
	20: uint16('o'),
	21: uint16(187),
	22: uint16('r'),
	23: uint16(179),
	24: uint16('s'),
	25: uint16(183),
	26: uint16('\t'),
	27: uint16(165),
	28: uint16(0x0b),
	29: uint16(165),
	30: uint16('\f'),
	31: uint16(165),
	32: uint16(' '),
	33: uint16(165),
}

var map_token2 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('%'),
	5:  uint16(151),
	6:  uint16('.'),
	7:  uint16(4),
	8:  uint16('@'),
	9:  uint16(6),
	10: uint16('a'),
	11: uint16(54),
	12: uint16('d'),
	13: uint16(19),
	14: uint16('f'),
	15: uint16(39),
	16: uint16('i'),
	17: uint16(58),
	18: uint16('m'),
	19: uint16(61),
	20: uint16('n'),
	21: uint16(20),
	22: uint16('o'),
	23: uint16(46),
	24: uint16('r'),
	25: uint16(29),
	26: uint16('t'),
	27: uint16(59),
	28: uint16('b'),
	29: uint16(81),
	30: uint16('c'),
	31: uint16(81),
	32: uint16('e'),
	33: uint16(81),
}

var map_token3 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('+'),
	5:  uint16(2),
	6:  uint16('-'),
	7:  uint16(74),
	8:  uint16('@'),
	9:  uint16(6),
	10: uint16('d'),
	11: uint16(42),
	12: uint16('f'),
	13: uint16(43),
	14: uint16('i'),
	15: uint16(58),
	16: uint16('m'),
	17: uint16(61),
}

var map_token4 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('@'),
	5:  uint16(6),
	6:  uint16('d'),
	7:  uint16(42),
	8:  uint16('f'),
	9:  uint16(43),
	10: uint16('i'),
	11: uint16(58),
	12: uint16('m'),
	13: uint16(61),
	14: uint16('+'),
	15: uint16(74),
	16: uint16('-'),
	17: uint16(74),
}

var map_token5 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(89),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('d'),
	5:  uint16(178),
	6:  uint16('n'),
	7:  uint16(174),
	8:  uint16('o'),
	9:  uint16(187),
	10: uint16('r'),
	11: uint16(179),
	12: uint16('\t'),
	13: uint16(165),
	14: uint16(0x0b),
	15: uint16(165),
	16: uint16('\f'),
	17: uint16(165),
	18: uint16(' '),
	19: uint16(165),
}

var ts_lex_modes = [89]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(82),
	},
	2: {
		Flex_state: uint16(82),
	},
	3: {
		Flex_state: uint16(82),
	},
	4: {
		Flex_state: uint16(82),
	},
	5: {
		Flex_state: uint16(82),
	},
	6: {
		Flex_state: uint16(82),
	},
	7: {
		Flex_state: uint16(82),
	},
	8: {
		Flex_state: uint16(82),
	},
	9: {
		Flex_state: uint16(82),
	},
	10: {
		Flex_state: uint16(82),
	},
	11: {
		Flex_state: uint16(82),
	},
	12: {
		Flex_state: uint16(82),
	},
	13: {
		Flex_state: uint16(82),
	},
	14: {
		Flex_state: uint16(82),
	},
	15: {
		Flex_state: uint16(82),
	},
	16: {
		Flex_state: uint16(82),
	},
	17: {
		Flex_state: uint16(82),
	},
	18: {
		Flex_state: uint16(82),
	},
	19: {
		Flex_state: uint16(86),
	},
	20: {
		Flex_state: uint16(86),
	},
	21: {
		Flex_state: uint16(82),
	},
	22: {
		Flex_state: uint16(82),
	},
	23: {
		Flex_state: uint16(84),
	},
	24: {
		Flex_state: uint16(86),
	},
	25: {
		Flex_state: uint16(86),
	},
	26: {
		Flex_state: uint16(86),
	},
	27: {
		Flex_state: uint16(86),
	},
	28: {
		Flex_state: uint16(86),
	},
	29: {
		Flex_state: uint16(87),
	},
	30: {
		Flex_state: uint16(87),
	},
	31: {
		Flex_state: uint16(86),
	},
	32: {
		Flex_state: uint16(86),
	},
	33: {
		Flex_state: uint16(87),
	},
	34: {
		Flex_state: uint16(86),
	},
	35: {
		Flex_state: uint16(87),
	},
	36: {
		Flex_state: uint16(8),
	},
	37: {
		Flex_state: uint16(8),
	},
	38: {
		Flex_state: uint16(14),
	},
	39: {
		Flex_state: uint16(14),
	},
	40: {},
	41: {},
	42: {
		Flex_state: uint16(84),
	},
	43: {},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
	51: {
		Flex_state: uint16(84),
	},
	52: {
		Flex_state: uint16(84),
	},
	53: {},
	54: {},
	55: {},
	56: {
		Flex_state: uint16(84),
	},
	57: {
		Flex_state: uint16(84),
	},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {
		Flex_state: uint16(84),
	},
	63: {
		Flex_state: uint16(33),
	},
	64: {},
	65: {},
	66: {
		Flex_state: uint16(33),
	},
	67: {},
	68: {},
	69: {},
	70: {
		Flex_state: uint16(84),
	},
	71: {},
	72: {
		Flex_state: uint16(84),
	},
	73: {
		Flex_state: uint16(84),
	},
	74: {},
	75: {},
	76: {},
	77: {},
	78: {},
	79: {
		Flex_state: uint16(84),
	},
	80: {},
	81: {},
	82: {
		Flex_state: uint16(84),
	},
	83: {
		Flex_state: uint16(73),
	},
	84: {
		Flex_state: uint16(33),
	},
	85: {},
	86: {},
	87: {},
	88: {
		Flex_state: uint16(84),
	},
}

var ts_parse_table = [4][63]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		27: uint16(1),
		28: uint16(1),
		29: uint16(1),
		30: uint16(1),
		31: uint16(1),
		32: uint16(1),
		34: uint16(1),
		37: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		2:  uint16(7),
		4:  uint16(9),
		5:  uint16(11),
		8:  uint16(13),
		9:  uint16(15),
		12: uint16(17),
		16: uint16(19),
		18: uint16(21),
		22: uint16(23),
		23: uint16(25),
		24: uint16(27),
		27: uint16(29),
		28: uint16(29),
		29: uint16(31),
		30: uint16(33),
		31: uint16(33),
		32: uint16(35),
		33: uint16(37),
		34: uint16(39),
		38: uint16(68),
		39: uint16(48),
		40: uint16(2),
		44: uint16(75),
		45: uint16(48),
		46: uint16(48),
		47: uint16(48),
		48: uint16(48),
		49: uint16(48),
		50: uint16(48),
		51: uint16(48),
		52: uint16(48),
		53: uint16(48),
		54: uint16(48),
		57: uint16(2),
	},
	2: {
		0:  uint16(41),
		1:  uint16(43),
		2:  uint16(7),
		4:  uint16(9),
		5:  uint16(11),
		8:  uint16(13),
		9:  uint16(15),
		12: uint16(17),
		16: uint16(19),
		18: uint16(21),
		22: uint16(23),
		23: uint16(25),
		24: uint16(27),
		27: uint16(29),
		28: uint16(29),
		29: uint16(31),
		30: uint16(33),
		31: uint16(33),
		32: uint16(35),
		33: uint16(45),
		34: uint16(39),
		39: uint16(53),
		40: uint16(3),
		44: uint16(75),
		45: uint16(53),
		46: uint16(53),
		47: uint16(53),
		48: uint16(53),
		49: uint16(53),
		50: uint16(53),
		51: uint16(53),
		52: uint16(53),
		53: uint16(53),
		54: uint16(53),
		57: uint16(3),
	},
	3: {
		0:  uint16(47),
		1:  uint16(49),
		2:  uint16(52),
		4:  uint16(55),
		5:  uint16(58),
		8:  uint16(61),
		9:  uint16(64),
		12: uint16(67),
		16: uint16(70),
		18: uint16(73),
		22: uint16(76),
		23: uint16(79),
		24: uint16(82),
		27: uint16(85),
		28: uint16(85),
		29: uint16(88),
		30: uint16(91),
		31: uint16(91),
		32: uint16(94),
		33: uint16(97),
		34: uint16(100),
		39: uint16(76),
		40: uint16(3),
		44: uint16(75),
		45: uint16(76),
		46: uint16(76),
		47: uint16(76),
		48: uint16(76),
		49: uint16(76),
		50: uint16(76),
		51: uint16(76),
		52: uint16(76),
		53: uint16(76),
		54: uint16(76),
		57: uint16(3),
	},
}

var ts_small_parse_table = [1025]uint16_t{
	0:    uint16(13),
	1:    uint16(9),
	2:    uint16(1),
	3:    uint16(anon_sym_new),
	4:    uint16(11),
	5:    uint16(1),
	6:    uint16(anon_sym_deleted),
	7:    uint16(13),
	8:    uint16(1),
	9:    uint16(anon_sym_old),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(anon_sym_rename),
	13:   uint16(17),
	14:   uint16(1),
	15:   uint16(anon_sym_Binary),
	16:   uint16(19),
	17:   uint16(1),
	18:   uint16(anon_sym_index),
	19:   uint16(21),
	20:   uint16(1),
	21:   uint16(anon_sym_similarity),
	22:   uint16(107),
	23:   uint16(1),
	24:   uint16(anon_sym_DASH_DASH_DASH),
	25:   uint16(5),
	26:   uint16(1),
	27:   uint16(aux_sym_block_repeat1),
	28:   uint16(74),
	29:   uint16(1),
	30:   uint16(sym_old_file),
	31:   uint16(61),
	32:   uint16(4),
	33:   uint16(sym_file_change),
	34:   uint16(sym_binary_change),
	35:   uint16(sym_index),
	36:   uint16(sym_similarity),
	37:   uint16(103),
	38:   uint16(6),
	40:   uint16(anon_sym_diff),
	41:   uint16(anon_sym_AT_AT),
	42:   uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	43:   uint16(anon_sym_DASH_DASH_DASH_DASH),
	44:   uint16(anon_sym_POUND),
	45:   uint16(105),
	46:   uint16(7),
	47:   uint16(aux_sym_source_token1),
	48:   uint16(anon_sym_PLUS_PLUS_PLUS),
	49:   uint16(anon_sym_PLUS),
	50:   uint16(anon_sym_PLUS_PLUS),
	51:   uint16(anon_sym_DASH),
	52:   uint16(anon_sym_DASH_DASH),
	53:   uint16(sym_context),
	54:   uint16(13),
	55:   uint16(9),
	56:   uint16(1),
	57:   uint16(anon_sym_new),
	58:   uint16(11),
	59:   uint16(1),
	60:   uint16(anon_sym_deleted),
	61:   uint16(13),
	62:   uint16(1),
	63:   uint16(anon_sym_old),
	64:   uint16(15),
	65:   uint16(1),
	66:   uint16(anon_sym_rename),
	67:   uint16(17),
	68:   uint16(1),
	69:   uint16(anon_sym_Binary),
	70:   uint16(19),
	71:   uint16(1),
	72:   uint16(anon_sym_index),
	73:   uint16(21),
	74:   uint16(1),
	75:   uint16(anon_sym_similarity),
	76:   uint16(107),
	77:   uint16(1),
	78:   uint16(anon_sym_DASH_DASH_DASH),
	79:   uint16(6),
	80:   uint16(1),
	81:   uint16(aux_sym_block_repeat1),
	82:   uint16(86),
	83:   uint16(1),
	84:   uint16(sym_old_file),
	85:   uint16(61),
	86:   uint16(4),
	87:   uint16(sym_file_change),
	88:   uint16(sym_binary_change),
	89:   uint16(sym_index),
	90:   uint16(sym_similarity),
	91:   uint16(109),
	92:   uint16(6),
	94:   uint16(anon_sym_diff),
	95:   uint16(anon_sym_AT_AT),
	96:   uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	97:   uint16(anon_sym_DASH_DASH_DASH_DASH),
	98:   uint16(anon_sym_POUND),
	99:   uint16(111),
	100:  uint16(7),
	101:  uint16(aux_sym_source_token1),
	102:  uint16(anon_sym_PLUS_PLUS_PLUS),
	103:  uint16(anon_sym_PLUS),
	104:  uint16(anon_sym_PLUS_PLUS),
	105:  uint16(anon_sym_DASH),
	106:  uint16(anon_sym_DASH_DASH),
	107:  uint16(sym_context),
	108:  uint16(11),
	109:  uint16(117),
	110:  uint16(1),
	111:  uint16(anon_sym_new),
	112:  uint16(120),
	113:  uint16(1),
	114:  uint16(anon_sym_deleted),
	115:  uint16(123),
	116:  uint16(1),
	117:  uint16(anon_sym_old),
	118:  uint16(126),
	119:  uint16(1),
	120:  uint16(anon_sym_rename),
	121:  uint16(129),
	122:  uint16(1),
	123:  uint16(anon_sym_Binary),
	124:  uint16(132),
	125:  uint16(1),
	126:  uint16(anon_sym_index),
	127:  uint16(135),
	128:  uint16(1),
	129:  uint16(anon_sym_similarity),
	130:  uint16(6),
	131:  uint16(1),
	132:  uint16(aux_sym_block_repeat1),
	133:  uint16(61),
	134:  uint16(4),
	135:  uint16(sym_file_change),
	136:  uint16(sym_binary_change),
	137:  uint16(sym_index),
	138:  uint16(sym_similarity),
	139:  uint16(113),
	140:  uint16(6),
	142:  uint16(anon_sym_diff),
	143:  uint16(anon_sym_AT_AT),
	144:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	145:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	146:  uint16(anon_sym_POUND),
	147:  uint16(115),
	148:  uint16(8),
	149:  uint16(aux_sym_source_token1),
	150:  uint16(anon_sym_DASH_DASH_DASH),
	151:  uint16(anon_sym_PLUS_PLUS_PLUS),
	152:  uint16(anon_sym_PLUS),
	153:  uint16(anon_sym_PLUS_PLUS),
	154:  uint16(anon_sym_DASH),
	155:  uint16(anon_sym_DASH_DASH),
	156:  uint16(sym_context),
	157:  uint16(12),
	158:  uint16(31),
	159:  uint16(1),
	160:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	161:  uint16(35),
	162:  uint16(1),
	163:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	164:  uint16(142),
	165:  uint16(1),
	166:  uint16(anon_sym_DASH_DASH_DASH),
	167:  uint16(144),
	168:  uint16(1),
	169:  uint16(anon_sym_PLUS_PLUS_PLUS),
	170:  uint16(146),
	171:  uint16(1),
	172:  uint16(sym_context),
	173:  uint16(11),
	174:  uint16(1),
	175:  uint16(aux_sym_changes_repeat2),
	176:  uint16(18),
	177:  uint16(1),
	178:  uint16(sym_changes),
	179:  uint16(29),
	180:  uint16(2),
	181:  uint16(anon_sym_PLUS),
	182:  uint16(anon_sym_PLUS_PLUS),
	183:  uint16(33),
	184:  uint16(2),
	185:  uint16(anon_sym_DASH),
	186:  uint16(anon_sym_DASH_DASH),
	187:  uint16(40),
	188:  uint16(2),
	189:  uint16(sym_addition),
	190:  uint16(sym_deletion),
	191:  uint16(140),
	192:  uint16(5),
	193:  uint16(aux_sym_source_token1),
	194:  uint16(anon_sym_new),
	195:  uint16(anon_sym_deleted),
	196:  uint16(anon_sym_old),
	197:  uint16(anon_sym_rename),
	198:  uint16(138),
	199:  uint16(7),
	201:  uint16(anon_sym_diff),
	202:  uint16(anon_sym_Binary),
	203:  uint16(anon_sym_index),
	204:  uint16(anon_sym_similarity),
	205:  uint16(anon_sym_AT_AT),
	206:  uint16(anon_sym_POUND),
	207:  uint16(5),
	208:  uint16(152),
	209:  uint16(1),
	210:  uint16(anon_sym_AT_AT),
	211:  uint16(71),
	212:  uint16(1),
	213:  uint16(sym_location),
	214:  uint16(8),
	215:  uint16(2),
	216:  uint16(sym_hunk),
	217:  uint16(aux_sym_hunks_repeat1),
	218:  uint16(148),
	219:  uint16(8),
	221:  uint16(anon_sym_diff),
	222:  uint16(anon_sym_Binary),
	223:  uint16(anon_sym_index),
	224:  uint16(anon_sym_similarity),
	225:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	226:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	227:  uint16(anon_sym_POUND),
	228:  uint16(150),
	229:  uint16(12),
	230:  uint16(aux_sym_source_token1),
	231:  uint16(anon_sym_new),
	232:  uint16(anon_sym_deleted),
	233:  uint16(anon_sym_old),
	234:  uint16(anon_sym_rename),
	235:  uint16(anon_sym_DASH_DASH_DASH),
	236:  uint16(anon_sym_PLUS_PLUS_PLUS),
	237:  uint16(anon_sym_PLUS),
	238:  uint16(anon_sym_PLUS_PLUS),
	239:  uint16(anon_sym_DASH),
	240:  uint16(anon_sym_DASH_DASH),
	241:  uint16(sym_context),
	242:  uint16(5),
	243:  uint16(27),
	244:  uint16(1),
	245:  uint16(anon_sym_AT_AT),
	246:  uint16(71),
	247:  uint16(1),
	248:  uint16(sym_location),
	249:  uint16(8),
	250:  uint16(2),
	251:  uint16(sym_hunk),
	252:  uint16(aux_sym_hunks_repeat1),
	253:  uint16(155),
	254:  uint16(8),
	256:  uint16(anon_sym_diff),
	257:  uint16(anon_sym_Binary),
	258:  uint16(anon_sym_index),
	259:  uint16(anon_sym_similarity),
	260:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	261:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	262:  uint16(anon_sym_POUND),
	263:  uint16(157),
	264:  uint16(12),
	265:  uint16(aux_sym_source_token1),
	266:  uint16(anon_sym_new),
	267:  uint16(anon_sym_deleted),
	268:  uint16(anon_sym_old),
	269:  uint16(anon_sym_rename),
	270:  uint16(anon_sym_DASH_DASH_DASH),
	271:  uint16(anon_sym_PLUS_PLUS_PLUS),
	272:  uint16(anon_sym_PLUS),
	273:  uint16(anon_sym_PLUS_PLUS),
	274:  uint16(anon_sym_DASH),
	275:  uint16(anon_sym_DASH_DASH),
	276:  uint16(sym_context),
	277:  uint16(11),
	278:  uint16(163),
	279:  uint16(1),
	280:  uint16(anon_sym_DASH_DASH_DASH),
	281:  uint16(166),
	282:  uint16(1),
	283:  uint16(anon_sym_PLUS_PLUS_PLUS),
	284:  uint16(172),
	285:  uint16(1),
	286:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	287:  uint16(178),
	288:  uint16(1),
	289:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	290:  uint16(181),
	291:  uint16(1),
	292:  uint16(sym_context),
	293:  uint16(10),
	294:  uint16(1),
	295:  uint16(aux_sym_changes_repeat2),
	296:  uint16(169),
	297:  uint16(2),
	298:  uint16(anon_sym_PLUS),
	299:  uint16(anon_sym_PLUS_PLUS),
	300:  uint16(175),
	301:  uint16(2),
	302:  uint16(anon_sym_DASH),
	303:  uint16(anon_sym_DASH_DASH),
	304:  uint16(40),
	305:  uint16(2),
	306:  uint16(sym_addition),
	307:  uint16(sym_deletion),
	308:  uint16(161),
	309:  uint16(5),
	310:  uint16(aux_sym_source_token1),
	311:  uint16(anon_sym_new),
	312:  uint16(anon_sym_deleted),
	313:  uint16(anon_sym_old),
	314:  uint16(anon_sym_rename),
	315:  uint16(159),
	316:  uint16(7),
	318:  uint16(anon_sym_diff),
	319:  uint16(anon_sym_Binary),
	320:  uint16(anon_sym_index),
	321:  uint16(anon_sym_similarity),
	322:  uint16(anon_sym_AT_AT),
	323:  uint16(anon_sym_POUND),
	324:  uint16(11),
	325:  uint16(31),
	326:  uint16(1),
	327:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	328:  uint16(35),
	329:  uint16(1),
	330:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	331:  uint16(142),
	332:  uint16(1),
	333:  uint16(anon_sym_DASH_DASH_DASH),
	334:  uint16(144),
	335:  uint16(1),
	336:  uint16(anon_sym_PLUS_PLUS_PLUS),
	337:  uint16(146),
	338:  uint16(1),
	339:  uint16(sym_context),
	340:  uint16(10),
	341:  uint16(1),
	342:  uint16(aux_sym_changes_repeat2),
	343:  uint16(29),
	344:  uint16(2),
	345:  uint16(anon_sym_PLUS),
	346:  uint16(anon_sym_PLUS_PLUS),
	347:  uint16(33),
	348:  uint16(2),
	349:  uint16(anon_sym_DASH),
	350:  uint16(anon_sym_DASH_DASH),
	351:  uint16(40),
	352:  uint16(2),
	353:  uint16(sym_addition),
	354:  uint16(sym_deletion),
	355:  uint16(186),
	356:  uint16(5),
	357:  uint16(aux_sym_source_token1),
	358:  uint16(anon_sym_new),
	359:  uint16(anon_sym_deleted),
	360:  uint16(anon_sym_old),
	361:  uint16(anon_sym_rename),
	362:  uint16(184),
	363:  uint16(7),
	365:  uint16(anon_sym_diff),
	366:  uint16(anon_sym_Binary),
	367:  uint16(anon_sym_index),
	368:  uint16(anon_sym_similarity),
	369:  uint16(anon_sym_AT_AT),
	370:  uint16(anon_sym_POUND),
	371:  uint16(4),
	372:  uint16(188),
	373:  uint16(1),
	374:  uint16(aux_sym_source_token1),
	375:  uint16(13),
	376:  uint16(1),
	377:  uint16(aux_sym_changes_repeat1),
	378:  uint16(159),
	379:  uint16(9),
	381:  uint16(anon_sym_diff),
	382:  uint16(anon_sym_Binary),
	383:  uint16(anon_sym_index),
	384:  uint16(anon_sym_similarity),
	385:  uint16(anon_sym_AT_AT),
	386:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	387:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	388:  uint16(anon_sym_POUND),
	389:  uint16(161),
	390:  uint16(11),
	391:  uint16(anon_sym_new),
	392:  uint16(anon_sym_deleted),
	393:  uint16(anon_sym_old),
	394:  uint16(anon_sym_rename),
	395:  uint16(anon_sym_DASH_DASH_DASH),
	396:  uint16(anon_sym_PLUS_PLUS_PLUS),
	397:  uint16(anon_sym_PLUS),
	398:  uint16(anon_sym_PLUS_PLUS),
	399:  uint16(anon_sym_DASH),
	400:  uint16(anon_sym_DASH_DASH),
	401:  uint16(sym_context),
	402:  uint16(4),
	403:  uint16(192),
	404:  uint16(1),
	405:  uint16(aux_sym_source_token1),
	406:  uint16(13),
	407:  uint16(1),
	408:  uint16(aux_sym_changes_repeat1),
	409:  uint16(190),
	410:  uint16(9),
	412:  uint16(anon_sym_diff),
	413:  uint16(anon_sym_Binary),
	414:  uint16(anon_sym_index),
	415:  uint16(anon_sym_similarity),
	416:  uint16(anon_sym_AT_AT),
	417:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	418:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	419:  uint16(anon_sym_POUND),
	420:  uint16(195),
	421:  uint16(11),
	422:  uint16(anon_sym_new),
	423:  uint16(anon_sym_deleted),
	424:  uint16(anon_sym_old),
	425:  uint16(anon_sym_rename),
	426:  uint16(anon_sym_DASH_DASH_DASH),
	427:  uint16(anon_sym_PLUS_PLUS_PLUS),
	428:  uint16(anon_sym_PLUS),
	429:  uint16(anon_sym_PLUS_PLUS),
	430:  uint16(anon_sym_DASH),
	431:  uint16(anon_sym_DASH_DASH),
	432:  uint16(sym_context),
	433:  uint16(2),
	434:  uint16(47),
	435:  uint16(9),
	437:  uint16(anon_sym_diff),
	438:  uint16(anon_sym_Binary),
	439:  uint16(anon_sym_index),
	440:  uint16(anon_sym_similarity),
	441:  uint16(anon_sym_AT_AT),
	442:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	443:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	444:  uint16(anon_sym_POUND),
	445:  uint16(197),
	446:  uint16(12),
	447:  uint16(aux_sym_source_token1),
	448:  uint16(anon_sym_new),
	449:  uint16(anon_sym_deleted),
	450:  uint16(anon_sym_old),
	451:  uint16(anon_sym_rename),
	452:  uint16(anon_sym_DASH_DASH_DASH),
	453:  uint16(anon_sym_PLUS_PLUS_PLUS),
	454:  uint16(anon_sym_PLUS),
	455:  uint16(anon_sym_PLUS_PLUS),
	456:  uint16(anon_sym_DASH),
	457:  uint16(anon_sym_DASH_DASH),
	458:  uint16(sym_context),
	459:  uint16(2),
	460:  uint16(113),
	461:  uint16(9),
	463:  uint16(anon_sym_diff),
	464:  uint16(anon_sym_Binary),
	465:  uint16(anon_sym_index),
	466:  uint16(anon_sym_similarity),
	467:  uint16(anon_sym_AT_AT),
	468:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	469:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	470:  uint16(anon_sym_POUND),
	471:  uint16(115),
	472:  uint16(12),
	473:  uint16(aux_sym_source_token1),
	474:  uint16(anon_sym_new),
	475:  uint16(anon_sym_deleted),
	476:  uint16(anon_sym_old),
	477:  uint16(anon_sym_rename),
	478:  uint16(anon_sym_DASH_DASH_DASH),
	479:  uint16(anon_sym_PLUS_PLUS_PLUS),
	480:  uint16(anon_sym_PLUS),
	481:  uint16(anon_sym_PLUS_PLUS),
	482:  uint16(anon_sym_DASH),
	483:  uint16(anon_sym_DASH_DASH),
	484:  uint16(sym_context),
	485:  uint16(2),
	486:  uint16(199),
	487:  uint16(9),
	489:  uint16(anon_sym_diff),
	490:  uint16(anon_sym_Binary),
	491:  uint16(anon_sym_index),
	492:  uint16(anon_sym_similarity),
	493:  uint16(anon_sym_AT_AT),
	494:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	495:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	496:  uint16(anon_sym_POUND),
	497:  uint16(201),
	498:  uint16(12),
	499:  uint16(aux_sym_source_token1),
	500:  uint16(anon_sym_new),
	501:  uint16(anon_sym_deleted),
	502:  uint16(anon_sym_old),
	503:  uint16(anon_sym_rename),
	504:  uint16(anon_sym_DASH_DASH_DASH),
	505:  uint16(anon_sym_PLUS_PLUS_PLUS),
	506:  uint16(anon_sym_PLUS),
	507:  uint16(anon_sym_PLUS_PLUS),
	508:  uint16(anon_sym_DASH),
	509:  uint16(anon_sym_DASH_DASH),
	510:  uint16(sym_context),
	511:  uint16(2),
	512:  uint16(203),
	513:  uint16(9),
	515:  uint16(anon_sym_diff),
	516:  uint16(anon_sym_Binary),
	517:  uint16(anon_sym_index),
	518:  uint16(anon_sym_similarity),
	519:  uint16(anon_sym_AT_AT),
	520:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	521:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	522:  uint16(anon_sym_POUND),
	523:  uint16(205),
	524:  uint16(12),
	525:  uint16(aux_sym_source_token1),
	526:  uint16(anon_sym_new),
	527:  uint16(anon_sym_deleted),
	528:  uint16(anon_sym_old),
	529:  uint16(anon_sym_rename),
	530:  uint16(anon_sym_DASH_DASH_DASH),
	531:  uint16(anon_sym_PLUS_PLUS_PLUS),
	532:  uint16(anon_sym_PLUS),
	533:  uint16(anon_sym_PLUS_PLUS),
	534:  uint16(anon_sym_DASH),
	535:  uint16(anon_sym_DASH_DASH),
	536:  uint16(sym_context),
	537:  uint16(2),
	538:  uint16(207),
	539:  uint16(9),
	541:  uint16(anon_sym_diff),
	542:  uint16(anon_sym_Binary),
	543:  uint16(anon_sym_index),
	544:  uint16(anon_sym_similarity),
	545:  uint16(anon_sym_AT_AT),
	546:  uint16(anon_sym_PLUS_PLUS_PLUS_PLUS),
	547:  uint16(anon_sym_DASH_DASH_DASH_DASH),
	548:  uint16(anon_sym_POUND),
	549:  uint16(209),
	550:  uint16(12),
	551:  uint16(aux_sym_source_token1),
	552:  uint16(anon_sym_new),
	553:  uint16(anon_sym_deleted),
	554:  uint16(anon_sym_old),
	555:  uint16(anon_sym_rename),
	556:  uint16(anon_sym_DASH_DASH_DASH),
	557:  uint16(anon_sym_PLUS_PLUS_PLUS),
	558:  uint16(anon_sym_PLUS),
	559:  uint16(anon_sym_PLUS_PLUS),
	560:  uint16(anon_sym_DASH),
	561:  uint16(anon_sym_DASH_DASH),
	562:  uint16(sym_context),
	563:  uint16(4),
	564:  uint16(213),
	565:  uint16(1),
	566:  uint16(aux_sym_filename_token1),
	567:  uint16(24),
	568:  uint16(1),
	569:  uint16(aux_sym_filename_repeat1),
	570:  uint16(55),
	571:  uint16(1),
	572:  uint16(sym_filename),
	573:  uint16(211),
	574:  uint16(2),
	576:  uint16(aux_sym_source_token1),
	577:  uint16(4),
	578:  uint16(213),
	579:  uint16(1),
	580:  uint16(aux_sym_filename_token1),
	581:  uint16(24),
	582:  uint16(1),
	583:  uint16(aux_sym_filename_repeat1),
	584:  uint16(47),
	585:  uint16(1),
	586:  uint16(sym_filename),
	587:  uint16(215),
	588:  uint16(2),
	590:  uint16(aux_sym_source_token1),
	591:  uint16(4),
	592:  uint16(27),
	593:  uint16(1),
	594:  uint16(anon_sym_AT_AT),
	595:  uint16(17),
	596:  uint16(1),
	597:  uint16(sym_hunks),
	598:  uint16(71),
	599:  uint16(1),
	600:  uint16(sym_location),
	601:  uint16(9),
	602:  uint16(2),
	603:  uint16(sym_hunk),
	604:  uint16(aux_sym_hunks_repeat1),
	605:  uint16(4),
	606:  uint16(27),
	607:  uint16(1),
	608:  uint16(anon_sym_AT_AT),
	609:  uint16(16),
	610:  uint16(1),
	611:  uint16(sym_hunks),
	612:  uint16(71),
	613:  uint16(1),
	614:  uint16(sym_location),
	615:  uint16(9),
	616:  uint16(2),
	617:  uint16(sym_hunk),
	618:  uint16(aux_sym_hunks_repeat1),
	619:  uint16(3),
	620:  uint16(219),
	621:  uint16(1),
	622:  uint16(aux_sym_similarity_token1),
	623:  uint16(59),
	624:  uint16(1),
	625:  uint16(sym_mode),
	626:  uint16(217),
	627:  uint16(2),
	629:  uint16(aux_sym_source_token1),
	630:  uint16(3),
	631:  uint16(223),
	632:  uint16(1),
	633:  uint16(aux_sym_filename_token1),
	634:  uint16(25),
	635:  uint16(1),
	636:  uint16(aux_sym_filename_repeat1),
	637:  uint16(221),
	638:  uint16(2),
	640:  uint16(aux_sym_source_token1),
	641:  uint16(3),
	642:  uint16(227),
	643:  uint16(1),
	644:  uint16(aux_sym_filename_token1),
	645:  uint16(25),
	646:  uint16(1),
	647:  uint16(aux_sym_filename_repeat1),
	648:  uint16(225),
	649:  uint16(2),
	651:  uint16(aux_sym_source_token1),
	652:  uint16(3),
	653:  uint16(230),
	654:  uint16(1),
	655:  uint16(aux_sym_filename_token1),
	656:  uint16(36),
	657:  uint16(1),
	658:  uint16(aux_sym_filename_repeat1),
	659:  uint16(64),
	660:  uint16(1),
	661:  uint16(sym_filename),
	662:  uint16(3),
	663:  uint16(213),
	664:  uint16(1),
	665:  uint16(aux_sym_filename_token1),
	666:  uint16(24),
	667:  uint16(1),
	668:  uint16(aux_sym_filename_repeat1),
	669:  uint16(41),
	670:  uint16(1),
	671:  uint16(sym_filename),
	672:  uint16(3),
	673:  uint16(213),
	674:  uint16(1),
	675:  uint16(aux_sym_filename_token1),
	676:  uint16(24),
	677:  uint16(1),
	678:  uint16(aux_sym_filename_repeat1),
	679:  uint16(47),
	680:  uint16(1),
	681:  uint16(sym_filename),
	682:  uint16(3),
	683:  uint16(211),
	684:  uint16(1),
	686:  uint16(232),
	687:  uint16(1),
	688:  uint16(aux_sym_source_token1),
	689:  uint16(234),
	690:  uint16(1),
	691:  uint16(aux_sym_location_token1),
	692:  uint16(3),
	693:  uint16(215),
	694:  uint16(1),
	696:  uint16(236),
	697:  uint16(1),
	698:  uint16(aux_sym_source_token1),
	699:  uint16(238),
	700:  uint16(1),
	701:  uint16(aux_sym_location_token1),
	702:  uint16(3),
	703:  uint16(240),
	704:  uint16(1),
	705:  uint16(aux_sym_filename_token1),
	706:  uint16(38),
	707:  uint16(1),
	708:  uint16(aux_sym_filename_repeat1),
	709:  uint16(88),
	710:  uint16(1),
	711:  uint16(sym_filename),
	712:  uint16(3),
	713:  uint16(213),
	714:  uint16(1),
	715:  uint16(aux_sym_filename_token1),
	716:  uint16(24),
	717:  uint16(1),
	718:  uint16(aux_sym_filename_repeat1),
	719:  uint16(85),
	720:  uint16(1),
	721:  uint16(sym_filename),
	722:  uint16(3),
	723:  uint16(242),
	724:  uint16(1),
	726:  uint16(244),
	727:  uint16(1),
	728:  uint16(aux_sym_source_token1),
	729:  uint16(246),
	730:  uint16(1),
	731:  uint16(aux_sym_location_token1),
	732:  uint16(3),
	733:  uint16(213),
	734:  uint16(1),
	735:  uint16(aux_sym_filename_token1),
	736:  uint16(24),
	737:  uint16(1),
	738:  uint16(aux_sym_filename_repeat1),
	739:  uint16(55),
	740:  uint16(1),
	741:  uint16(sym_filename),
	742:  uint16(3),
	743:  uint16(248),
	744:  uint16(1),
	746:  uint16(250),
	747:  uint16(1),
	748:  uint16(aux_sym_source_token1),
	749:  uint16(252),
	750:  uint16(1),
	751:  uint16(aux_sym_location_token1),
	752:  uint16(3),
	753:  uint16(254),
	754:  uint16(1),
	755:  uint16(anon_sym_and),
	756:  uint16(256),
	757:  uint16(1),
	758:  uint16(aux_sym_filename_token1),
	759:  uint16(37),
	760:  uint16(1),
	761:  uint16(aux_sym_filename_repeat1),
	762:  uint16(3),
	763:  uint16(258),
	764:  uint16(1),
	765:  uint16(anon_sym_and),
	766:  uint16(260),
	767:  uint16(1),
	768:  uint16(aux_sym_filename_token1),
	769:  uint16(37),
	770:  uint16(1),
	771:  uint16(aux_sym_filename_repeat1),
	772:  uint16(3),
	773:  uint16(254),
	774:  uint16(1),
	775:  uint16(anon_sym_differ),
	776:  uint16(263),
	777:  uint16(1),
	778:  uint16(aux_sym_filename_token1),
	779:  uint16(39),
	780:  uint16(1),
	781:  uint16(aux_sym_filename_repeat1),
	782:  uint16(3),
	783:  uint16(258),
	784:  uint16(1),
	785:  uint16(anon_sym_differ),
	786:  uint16(265),
	787:  uint16(1),
	788:  uint16(aux_sym_filename_token1),
	789:  uint16(39),
	790:  uint16(1),
	791:  uint16(aux_sym_filename_repeat1),
	792:  uint16(2),
	793:  uint16(268),
	794:  uint16(1),
	795:  uint16(aux_sym_source_token1),
	796:  uint16(12),
	797:  uint16(1),
	798:  uint16(aux_sym_changes_repeat1),
	799:  uint16(1),
	800:  uint16(270),
	801:  uint16(2),
	803:  uint16(aux_sym_source_token1),
	804:  uint16(2),
	805:  uint16(272),
	806:  uint16(1),
	807:  uint16(anon_sym_PLUS_PLUS_PLUS),
	808:  uint16(69),
	809:  uint16(1),
	810:  uint16(sym_new_file),
	811:  uint16(1),
	812:  uint16(274),
	813:  uint16(2),
	815:  uint16(aux_sym_source_token1),
	816:  uint16(1),
	817:  uint16(276),
	818:  uint16(2),
	820:  uint16(aux_sym_source_token1),
	821:  uint16(1),
	822:  uint16(278),
	823:  uint16(2),
	825:  uint16(aux_sym_source_token1),
	826:  uint16(1),
	827:  uint16(280),
	828:  uint16(2),
	830:  uint16(aux_sym_source_token1),
	831:  uint16(1),
	832:  uint16(282),
	833:  uint16(2),
	835:  uint16(aux_sym_source_token1),
	836:  uint16(2),
	837:  uint16(41),
	838:  uint16(1),
	840:  uint16(284),
	841:  uint16(1),
	842:  uint16(aux_sym_source_token1),
	843:  uint16(1),
	844:  uint16(286),
	845:  uint16(2),
	846:  uint16(anon_sym_from),
	847:  uint16(anon_sym_to),
	848:  uint16(1),
	849:  uint16(288),
	850:  uint16(2),
	852:  uint16(aux_sym_source_token1),
	853:  uint16(2),
	854:  uint16(290),
	855:  uint16(1),
	856:  uint16(anon_sym_file),
	857:  uint16(292),
	858:  uint16(1),
	859:  uint16(anon_sym_mode),
	860:  uint16(2),
	861:  uint16(219),
	862:  uint16(1),
	863:  uint16(aux_sym_similarity_token1),
	864:  uint16(41),
	865:  uint16(1),
	866:  uint16(sym_mode),
	867:  uint16(2),
	868:  uint16(284),
	869:  uint16(1),
	870:  uint16(aux_sym_source_token1),
	871:  uint16(294),
	872:  uint16(1),
	874:  uint16(1),
	875:  uint16(296),
	876:  uint16(2),
	878:  uint16(aux_sym_source_token1),
	879:  uint16(1),
	880:  uint16(298),
	881:  uint16(2),
	883:  uint16(aux_sym_source_token1),
	884:  uint16(2),
	885:  uint16(219),
	886:  uint16(1),
	887:  uint16(aux_sym_similarity_token1),
	888:  uint16(50),
	889:  uint16(1),
	890:  uint16(sym_mode),
	891:  uint16(2),
	892:  uint16(272),
	893:  uint16(1),
	894:  uint16(anon_sym_PLUS_PLUS_PLUS),
	895:  uint16(65),
	896:  uint16(1),
	897:  uint16(sym_new_file),
	898:  uint16(1),
	899:  uint16(300),
	900:  uint16(2),
	902:  uint16(aux_sym_source_token1),
	903:  uint16(1),
	904:  uint16(302),
	905:  uint16(2),
	907:  uint16(aux_sym_source_token1),
	908:  uint16(1),
	909:  uint16(304),
	910:  uint16(2),
	912:  uint16(aux_sym_source_token1),
	913:  uint16(1),
	914:  uint16(306),
	915:  uint16(1),
	916:  uint16(aux_sym_source_token1),
	917:  uint16(1),
	918:  uint16(308),
	919:  uint16(1),
	920:  uint16(aux_sym_similarity_token1),
	921:  uint16(1),
	922:  uint16(310),
	923:  uint16(1),
	924:  uint16(sym_commit),
	925:  uint16(1),
	926:  uint16(312),
	927:  uint16(1),
	928:  uint16(anon_sym_and),
	929:  uint16(1),
	930:  uint16(314),
	931:  uint16(1),
	932:  uint16(aux_sym_source_token1),
	933:  uint16(1),
	934:  uint16(316),
	935:  uint16(1),
	936:  uint16(sym_commit),
	937:  uint16(1),
	938:  uint16(318),
	939:  uint16(1),
	940:  uint16(anon_sym_PERCENT),
	941:  uint16(1),
	942:  uint16(320),
	943:  uint16(1),
	945:  uint16(1),
	946:  uint16(322),
	947:  uint16(1),
	948:  uint16(aux_sym_source_token1),
	949:  uint16(1),
	950:  uint16(324),
	951:  uint16(1),
	952:  uint16(anon_sym_AT_AT2),
	953:  uint16(1),
	954:  uint16(326),
	955:  uint16(1),
	956:  uint16(aux_sym_source_token1),
	957:  uint16(1),
	958:  uint16(328),
	959:  uint16(1),
	960:  uint16(anon_sym_index2),
	961:  uint16(1),
	962:  uint16(330),
	963:  uint16(1),
	964:  uint16(sym_linerange),
	965:  uint16(1),
	966:  uint16(332),
	967:  uint16(1),
	968:  uint16(aux_sym_source_token1),
	969:  uint16(1),
	970:  uint16(334),
	971:  uint16(1),
	972:  uint16(aux_sym_source_token1),
	973:  uint16(1),
	974:  uint16(284),
	975:  uint16(1),
	976:  uint16(aux_sym_source_token1),
	977:  uint16(1),
	978:  uint16(215),
	979:  uint16(1),
	980:  uint16(aux_sym_source_token1),
	981:  uint16(1),
	982:  uint16(211),
	983:  uint16(1),
	984:  uint16(aux_sym_source_token1),
	985:  uint16(1),
	986:  uint16(290),
	987:  uint16(1),
	988:  uint16(anon_sym_file),
	989:  uint16(1),
	990:  uint16(292),
	991:  uint16(1),
	992:  uint16(anon_sym_mode),
	993:  uint16(1),
	994:  uint16(336),
	995:  uint16(1),
	996:  uint16(anon_sym_mode),
	997:  uint16(1),
	998:  uint16(338),
	999:  uint16(1),
	1000: uint16(sym_linerange),
	1001: uint16(1),
	1002: uint16(340),
	1003: uint16(1),
	1004: uint16(aux_sym_command_token1),
	1005: uint16(1),
	1006: uint16(342),
	1007: uint16(1),
	1008: uint16(anon_sym_files),
	1009: uint16(1),
	1010: uint16(344),
	1011: uint16(1),
	1012: uint16(aux_sym_source_token1),
	1013: uint16(1),
	1014: uint16(346),
	1015: uint16(1),
	1016: uint16(aux_sym_source_token1),
	1017: uint16(1),
	1018: uint16(348),
	1019: uint16(1),
	1020: uint16(anon_sym_DOT_DOT),
	1021: uint16(1),
	1022: uint16(350),
	1023: uint16(1),
	1024: uint16(anon_sym_differ),
}

var ts_small_parse_table_map = [85]uint32_t{
	1:  uint32(54),
	2:  uint32(108),
	3:  uint32(157),
	4:  uint32(207),
	5:  uint32(242),
	6:  uint32(277),
	7:  uint32(324),
	8:  uint32(371),
	9:  uint32(402),
	10: uint32(433),
	11: uint32(459),
	12: uint32(485),
	13: uint32(511),
	14: uint32(537),
	15: uint32(563),
	16: uint32(577),
	17: uint32(591),
	18: uint32(605),
	19: uint32(619),
	20: uint32(630),
	21: uint32(641),
	22: uint32(652),
	23: uint32(662),
	24: uint32(672),
	25: uint32(682),
	26: uint32(692),
	27: uint32(702),
	28: uint32(712),
	29: uint32(722),
	30: uint32(732),
	31: uint32(742),
	32: uint32(752),
	33: uint32(762),
	34: uint32(772),
	35: uint32(782),
	36: uint32(792),
	37: uint32(799),
	38: uint32(804),
	39: uint32(811),
	40: uint32(816),
	41: uint32(821),
	42: uint32(826),
	43: uint32(831),
	44: uint32(836),
	45: uint32(843),
	46: uint32(848),
	47: uint32(853),
	48: uint32(860),
	49: uint32(867),
	50: uint32(874),
	51: uint32(879),
	52: uint32(884),
	53: uint32(891),
	54: uint32(898),
	55: uint32(903),
	56: uint32(908),
	57: uint32(913),
	58: uint32(917),
	59: uint32(921),
	60: uint32(925),
	61: uint32(929),
	62: uint32(933),
	63: uint32(937),
	64: uint32(941),
	65: uint32(945),
	66: uint32(949),
	67: uint32(953),
	68: uint32(957),
	69: uint32(961),
	70: uint32(965),
	71: uint32(969),
	72: uint32(973),
	73: uint32(977),
	74: uint32(981),
	75: uint32(985),
	76: uint32(989),
	77: uint32(993),
	78: uint32(997),
	79: uint32(1001),
	80: uint32(1005),
	81: uint32(1009),
	82: uint32(1013),
	83: uint32(1017),
	84: uint32(1021),
}

var ts_parse_actions = [352]TSParseActionEntry{
	0: {},
	1: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	2: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeRecover)})),
	3: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	4: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_source),
	})))),
	5: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	6: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
	}})))),
	7: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	8: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
	}})))),
	9: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	10: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
	}})))),
	11: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	12: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
	}})))),
	13: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	14: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
	}})))),
	15: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	16: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
	}})))),
	17: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	18: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
	19: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	20: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
	}})))),
	21: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	22: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
	}})))),
	23: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	24: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
	}})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
	}})))),
	27: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	28: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
	}})))),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
	}})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
	}})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
	}})))),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	42: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source),
	})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
	}})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
	}})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(83)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(80)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(20)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(19)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(82)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(29)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	89: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(29)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(35)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	103: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block),
	})))),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block),
	})))),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
	}})))),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
	})))),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
	})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(80)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_hunk),
		Fproduction_id: uint16(2),
	})))),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_hunk),
		Fproduction_id: uint16(2),
	})))),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
	}})))),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
	}})))),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
	}})))),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_hunks_repeat1),
	})))),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_hunks_repeat1),
	})))),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_hunks_repeat1),
	})))),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(82)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_hunks),
	})))),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_hunks),
	})))),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(78)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(29)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(29)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat2),
	})))),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(40)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	184: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_changes),
	})))),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_changes),
	})))),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
	}})))),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat1),
	})))),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat1),
	})))),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_changes_repeat1),
	})))),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_block),
	})))),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_block),
	})))),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_block),
	})))),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_block),
	})))),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_hunk),
		Fproduction_id: uint16(3),
	})))),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_hunk),
		Fproduction_id: uint16(3),
	})))),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_addition),
	})))),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
	}})))),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_deletion),
	})))),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_index),
	})))),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
	}})))),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_filename),
	})))),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
	}})))),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filename_repeat1),
	})))),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filename_repeat1),
	})))),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(25)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
	}})))),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_addition),
	})))),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
	}})))),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_deletion),
	})))),
	238: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
	}})))),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
	}})))),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_location),
	})))),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_location),
	})))),
	246: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
	}})))),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
	}})))),
	254: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_filename),
	})))),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
	}})))),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filename_repeat1),
	})))),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filename_repeat1),
	})))),
	262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
	}})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filename_repeat1),
	})))),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
	}})))),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_file_change),
	})))),
	272: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
	}})))),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_binary_change),
	})))),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_addition),
	})))),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_deletion),
	})))),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_comment),
	})))),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_old_file),
	})))),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
	}})))),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
	}})))),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_file_change),
	})))),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
	}})))),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
	}})))),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_source),
	})))),
	296: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_similarity),
		Fproduction_id: uint16(1),
	})))),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_new_file),
	})))),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_mode),
	})))),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_index),
	})))),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_location),
	})))),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
	}})))),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
	}})))),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
	}})))),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
	}})))),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
	}})))),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
	}})))),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
	}})))),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
	}})))),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
	}})))),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
	}})))),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
	}})))),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
	}})))),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
	}})))),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
	}})))),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
	}})))),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
	}})))),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
	}})))),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
	}})))),
	344: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_command),
	})))),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
	}})))),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
	}})))),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
	}})))),
}

func tree_sitter_diff(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(15),
	Fsymbol_count:              uint32(63),
	Falias_count:               uint32(1),
	Ftoken_count:               uint32(38),
	Fstate_count:               uint32(89),
	Flarge_state_count:         uint32(4),
	Fproduction_id_count:       uint32(4),
	Ffield_count:               uint32(2),
	Fmax_alias_sequence_length: uint16(8),
	Fparse_table:               uintptr(unsafe.Pointer(&ts_parse_table)),
	Fsmall_parse_table:         uintptr(unsafe.Pointer(&ts_small_parse_table)),
	Fsmall_parse_table_map:     uintptr(unsafe.Pointer(&ts_small_parse_table_map)),
	Fparse_actions:             uintptr(unsafe.Pointer(&ts_parse_actions)),
	Fsymbol_names:              uintptr(unsafe.Pointer(&ts_symbol_names)),
	Ffield_names:               uintptr(unsafe.Pointer(&ts_field_names)),
	Ffield_map_slices:          uintptr(unsafe.Pointer(&ts_field_map_slices)),
	Ffield_map_entries:         uintptr(unsafe.Pointer(&ts_field_map_entries)),
	Fsymbol_metadata:           uintptr(unsafe.Pointer(&ts_symbol_metadata)),
	Fpublic_symbol_map:         uintptr(unsafe.Pointer(&ts_symbol_map)),
	Falias_map:                 uintptr(unsafe.Pointer(&ts_non_terminal_alias_map)),
	Falias_sequences:           uintptr(unsafe.Pointer(&ts_alias_sequences)),
	Flex_modes:                 uintptr(unsafe.Pointer(&ts_lex_modes)),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 18,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(1),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00source_token1\x00diff\x00argument\x00new\x00deleted\x00file\x00mode\x00old\x00rename\x00from\x00to\x00Binary\x00files\x00and\x00differ\x00index\x00..\x00similarity\x00similarity_token1\x00%\x00---\x00+++\x00@@\x00location_token1\x00+\x00++\x00++++\x00-\x00--\x00----\x00context\x00#\x00linerange\x00filename_token1\x00commit\x00source\x00_line\x00block\x00hunks\x00hunk\x00changes\x00command\x00file_change\x00binary_change\x00old_file\x00new_file\x00location\x00addition\x00deletion\x00comment\x00filename\x00source_repeat1\x00block_repeat1\x00hunks_repeat1\x00changes_repeat1\x00changes_repeat2\x00filename_repeat1\x00score\x00"
