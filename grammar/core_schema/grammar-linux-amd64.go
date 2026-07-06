// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions preprocessed.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_core_schema

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const _GNU_SOURCE = 1
const _LP64 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BITINT_MAXWIDTH__ = 8388608
const __BOOL_WIDTH__ = 8
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const __CLANG_ATOMIC_INT_LOCK_FREE = 2
const __CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const __CLANG_ATOMIC_LONG_LOCK_FREE = 2
const __CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const __CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const __CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __CONSTANT_CFSTRINGS__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DENORM_MIN__ = 4.9406564584124654e-324
const __DBL_DIG__ = 15
const __DBL_EPSILON__ = 2.2204460492503131e-16
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MAX__ = 1.7976931348623157e+308
const __DBL_MIN__ = 2.2250738585072014e-308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT128__ = 1
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.765625e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.5504e+4
const __FLT16_MIN__ = 6.103515625e-5
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209290e-7
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282347e+38
const __FLT_MIN__ = 1.17549435e-38
const __FLT_RADIX__ = 2
const __FPCLASS_NEGINF = 0x0004
const __FPCLASS_NEGNORMAL = 0x0008
const __FPCLASS_NEGSUBNORMAL = 0x0010
const __FPCLASS_NEGZERO = 0x0020
const __FPCLASS_POSINF = 0x0200
const __FPCLASS_POSNORMAL = 0x0100
const __FPCLASS_POSSUBNORMAL = 0x0080
const __FPCLASS_POSZERO = 0x0040
const __FPCLASS_QNAN = 0x0002
const __FPCLASS_SNAN = 0x0001
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
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 4
const __GXX_ABI_VERSION = 1002
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "L"
const __INT64_FMTd__ = "ld"
const __INT64_FMTi__ = "li"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "L"
const __INTMAX_FMTd__ = "ld"
const __INTMAX_FMTi__ = "li"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "ld"
const __INTPTR_FMTi__ = "li"
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_FMTd__ = "hd"
const __INT_FAST16_FMTi__ = "hi"
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_TYPE__ = "short"
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_FMTd__ = "d"
const __INT_FAST32_FMTi__ = "i"
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_FMTd__ = "ld"
const __INT_FAST64_FMTi__ = "li"
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_FMTd__ = "hhd"
const __INT_FAST8_FMTi__ = "hhi"
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_FMTd__ = "hd"
const __INT_LEAST16_FMTi__ = "hi"
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_TYPE__ = "short"
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_FMTd__ = "d"
const __INT_LEAST32_FMTi__ = "i"
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_FMTd__ = "ld"
const __INT_LEAST64_FMTi__ = "li"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_FMTd__ = "hhd"
const __INT_LEAST8_FMTi__ = "hhi"
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.2204460492503131e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.7976931348623157e+308
const __LDBL_MIN__ = 2.2250738585072014e-308
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX__ = 9223372036854775807
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MEMORY_SCOPE_DEVICE = 1
const __MEMORY_SCOPE_SINGLE = 4
const __MEMORY_SCOPE_SYSTEM = 0
const __MEMORY_SCOPE_WRKGRP = 2
const __MEMORY_SCOPE_WVFRNT = 3
const __MMX__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_INLINES = 1
const __OBJC_BOOL_IS_BOOL = 0
const __OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const __OPENCL_MEMORY_SCOPE_DEVICE = 2
const __OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const __OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const __OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __POINTER_WIDTH__ = 64
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_FMTd__ = "ld"
const __PTRDIFF_FMTi__ = "li"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
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
const __SIZE_FMTX__ = "lX"
const __SIZE_FMTo__ = "lo"
const __SIZE_FMTu__ = "lu"
const __SIZE_FMTx__ = "lx"
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __STDC_HOSTED__ = 1
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_FMTX__ = "hX"
const __UINT16_FMTo__ = "ho"
const __UINT16_FMTu__ = "hu"
const __UINT16_FMTx__ = "hx"
const __UINT16_MAX__ = 65535
const __UINT32_C_SUFFIX__ = "U"
const __UINT32_FMTX__ = "X"
const __UINT32_FMTo__ = "o"
const __UINT32_FMTu__ = "u"
const __UINT32_FMTx__ = "x"
const __UINT32_MAX__ = 4294967295
const __UINT64_C_SUFFIX__ = "UL"
const __UINT64_FMTX__ = "lX"
const __UINT64_FMTo__ = "lo"
const __UINT64_FMTu__ = "lu"
const __UINT64_FMTx__ = "lx"
const __UINT64_MAX__ = 18446744073709551615
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "UL"
const __UINTMAX_FMTX__ = "lX"
const __UINTMAX_FMTo__ = "lo"
const __UINTMAX_FMTu__ = "lu"
const __UINTMAX_FMTx__ = "lx"
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "lX"
const __UINTPTR_FMTo__ = "lo"
const __UINTPTR_FMTu__ = "lu"
const __UINTPTR_FMTx__ = "lx"
const __UINTPTR_MAX__ = 18446744073709551615
const __UINTPTR_WIDTH__ = 64
const __UINT_FAST16_FMTX__ = "hX"
const __UINT_FAST16_FMTo__ = "ho"
const __UINT_FAST16_FMTu__ = "hu"
const __UINT_FAST16_FMTx__ = "hx"
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_FMTX__ = "X"
const __UINT_FAST32_FMTo__ = "o"
const __UINT_FAST32_FMTu__ = "u"
const __UINT_FAST32_FMTx__ = "x"
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_FMTX__ = "lX"
const __UINT_FAST64_FMTo__ = "lo"
const __UINT_FAST64_FMTu__ = "lu"
const __UINT_FAST64_FMTx__ = "lx"
const __UINT_FAST64_MAX__ = 18446744073709551615
const __UINT_FAST8_FMTX__ = "hhX"
const __UINT_FAST8_FMTo__ = "hho"
const __UINT_FAST8_FMTu__ = "hhu"
const __UINT_FAST8_FMTx__ = "hhx"
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_FMTX__ = "hX"
const __UINT_LEAST16_FMTo__ = "ho"
const __UINT_LEAST16_FMTu__ = "hu"
const __UINT_LEAST16_FMTx__ = "hx"
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_FMTX__ = "X"
const __UINT_LEAST32_FMTo__ = "o"
const __UINT_LEAST32_FMTu__ = "u"
const __UINT_LEAST32_FMTx__ = "x"
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_FMTX__ = "lX"
const __UINT_LEAST64_FMTo__ = "lo"
const __UINT_LEAST64_FMTu__ = "lu"
const __UINT_LEAST64_FMTx__ = "lx"
const __UINT_LEAST64_MAX__ = 18446744073709551615
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __VERSION__ = "Ubuntu Clang 18.1.3 (1ubuntu1)"
const __WCHAR_MAX__ = 2147483647
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 4294967295
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 18
const __clang_minor__ = 1
const __clang_patchlevel__ = 3
const __clang_version__ = "18.1.3 (1ubuntu1)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __llvm__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __tune_k8__ = 1
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
	return libc.Uint16FromInt32(libc.Int32FromUint16(__bsx)>>libc.Int32FromInt32(8)&libc.Int32FromInt32(0xff) | libc.Int32FromUint16(__bsx)&libc.Int32FromInt32(0xff)<<libc.Int32FromInt32(8))
}

func __bswap_32(tls *libc.TLS, __bsx __uint32_t) (r __uint32_t) {
	return __bsx&libc.Uint32FromUint32(0xff000000)>>libc.Int32FromInt32(24) | __bsx&libc.Uint32FromUint32(0x00ff0000)>>libc.Int32FromInt32(8) | __bsx&libc.Uint32FromUint32(0x0000ff00)<<libc.Int32FromInt32(8) | __bsx&libc.Uint32FromUint32(0x000000ff)<<libc.Int32FromInt32(24)
}

func __bswap_64(tls *libc.TLS, __bsx __uint64_t) (r __uint64_t) {
	return uint64(__bsx&libc.Uint64FromUint64(0xff00000000000000)>>libc.Int32FromInt32(56) | __bsx&libc.Uint64FromUint64(0x00ff000000000000)>>libc.Int32FromInt32(40) | __bsx&libc.Uint64FromUint64(0x0000ff0000000000)>>libc.Int32FromInt32(24) | __bsx&libc.Uint64FromUint64(0x000000ff00000000)>>libc.Int32FromInt32(8) | __bsx&libc.Uint64FromUint64(0x00000000ff000000)<<libc.Int32FromInt32(8) | __bsx&libc.Uint64FromUint64(0x0000000000ff0000)<<libc.Int32FromInt32(24) | __bsx&libc.Uint64FromUint64(0x000000000000ff00)<<libc.Int32FromInt32(40) | __bsx&libc.Uint64FromUint64(0x00000000000000ff)<<libc.Int32FromInt32(56))
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
	F__size  [4]int8
}

type pthread_condattr_t = struct {
	F__align [0]int32
	F__size  [4]int8
}

type pthread_key_t = uint32

type pthread_once_t = int32

type pthread_attr_t = struct {
	F__align [0]int64
	F__size  [56]int8
}

type pthread_mutex_t = struct {
	F__size  [0][40]int8
	F__align [0]int64
	F__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	F__size  [0][48]int8
	F__align [0]int64
	F__data  __pthread_cond_s
}

type pthread_rwlock_t = struct {
	F__size  [0][56]int8
	F__align [0]int64
	F__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	F__align [0]int64
	F__size  [8]int8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	F__align [0]int64
	F__size  [32]int8
}

type pthread_barrierattr_t = struct {
	F__align [0]int32
	F__size  [4]int8
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

const sym_null = 1
const sym_bool = 2
const sym_int = 3
const sym_float = 4
const sym_scalar = 5

var ts_symbol_names = [6]uintptr{
	0: __ccgo_ts,
	1: __ccgo_ts + 4,
	2: __ccgo_ts + 9,
	3: __ccgo_ts + 14,
	4: __ccgo_ts + 18,
	5: __ccgo_ts + 24,
}

var ts_symbol_map = [6]TSSymbol{
	1: uint16(sym_null),
	2: uint16(sym_bool),
	3: uint16(sym_int),
	4: uint16(sym_float),
	5: uint16(sym_scalar),
}

var ts_symbol_metadata = [6]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
}

var ts_alias_sequences = [1][1]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [4]TSStateId{
	1: uint16(1),
	2: uint16(2),
	3: uint16(3),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i uint32_t
	var lookahead int32_t
	_, _, _, _, _ = eof, i, lookahead, result, skip
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
			state = uint16(34)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('.') {
			state = uint16(7)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('A') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('A') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('E') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('F') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('I') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('I') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('L') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('L') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('L') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('N') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('N') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('R') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('S') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('U') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('U') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('a') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('a') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('f') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('l') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('l') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('l') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('n') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('n') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('r') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('s') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('u') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(32)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(31):
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(32):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(33):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(43)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token = [22]uint16_t{
	0:  uint16('.'),
	1:  uint16(6),
	2:  uint16('0'),
	3:  uint16(37),
	4:  uint16('F'),
	5:  uint16(2),
	6:  uint16('N'),
	7:  uint16(16),
	8:  uint16('T'),
	9:  uint16(13),
	10: uint16('f'),
	11: uint16(17),
	12: uint16('n'),
	13: uint16(29),
	14: uint16('t'),
	15: uint16(26),
	16: uint16('~'),
	17: uint16(35),
	18: uint16('+'),
	19: uint16(1),
	20: uint16('-'),
	21: uint16(1),
}

var ts_lex_modes = [4]TSLexMode{}

var ts_parse_table = [2][6]uint16_t{
	0: {
		0: uint16(1),
		1: uint16(1),
		2: uint16(1),
		3: uint16(1),
		4: uint16(1),
	},
	1: {
		1: uint16(3),
		2: uint16(3),
		3: uint16(5),
		4: uint16(5),
		5: uint16(3),
	},
}

var ts_small_parse_table = [8]uint16_t{
	0: uint16(1),
	1: uint16(7),
	2: uint16(1),
	4: uint16(1),
	5: uint16(9),
	6: uint16(1),
}

var ts_small_parse_table_map = [2]uint32_t{
	1: uint32(4),
}

var ts_parse_actions = [11]TSParseActionEntry{
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_scalar),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	10: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
}

func tree_sitter_core_schema(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(14),
	Fsymbol_count:              uint32(6),
	Ftoken_count:               uint32(5),
	Fstate_count:               uint32(4),
	Flarge_state_count:         uint32(2),
	Fproduction_id_count:       uint32(1),
	Fmax_alias_sequence_length: uint16(1),
	Fparse_table:               uintptr(unsafe.Pointer(&ts_parse_table)),
	Fsmall_parse_table:         uintptr(unsafe.Pointer(&ts_small_parse_table)),
	Fsmall_parse_table_map:     uintptr(unsafe.Pointer(&ts_small_parse_table_map)),
	Fparse_actions:             uintptr(unsafe.Pointer(&ts_parse_actions)),
	Fsymbol_names:              uintptr(unsafe.Pointer(&ts_symbol_names)),
	Fsymbol_metadata:           uintptr(unsafe.Pointer(&ts_symbol_metadata)),
	Fpublic_symbol_map:         uintptr(unsafe.Pointer(&ts_symbol_map)),
	Falias_map:                 uintptr(unsafe.Pointer(&ts_non_terminal_alias_map)),
	Falias_sequences:           uintptr(unsafe.Pointer(&ts_alias_sequences)),
	Flex_modes:                 uintptr(unsafe.Pointer(&ts_lex_modes)),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
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

var __ccgo_ts1 = "end\x00null\x00bool\x00int\x00float\x00scalar\x00"
