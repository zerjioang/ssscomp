from ctypes import *


# Calling shared_object from python
#
# The standard C numeric types are available under the names:
#
# * C.char (byte)
# * C.schar (signed char)
# * C.uchar (unsigned char)
# * C.short
# * C.ushort (unsigned short)
# * C.int
# * C.uint (unsigned int)
# * C.long
# * C.ulong (unsigned long)
# * C.longlong (long long)
# * C.ulonglong (unsigned long long)
# * C.float
# * C.double
# * C.complexfloat (complex float)
# * C.complexdouble (complex double)
#
# The C type void* is represented by Go's unsafe.Pointer.
#
# The C types __int128_t and __uint128_t are represented by [16]byte.
#
# conversion table from Go to Python
#
# shared object function argument definition via
# function_name.argtypes = ...
#
# shared object function return argument definition via
# get_image_size.restype = ...
#
# Relation table:
# |---------------------------------------|
# | Golang         | Python               |
# |---------------------------------------|
# | *C.char        | c_char_p             |
# | *uint          | POINTER(c_ulonglong) |
# | uint           |          c_ulonglong |
# |---------------------------------------|


# A Go String primitive wrapper for Python
# Usage example:
#
# msg = GoString(b"Hello Python!", 13)
#
# Calling go function (i.e Log) via shared_object would be
# lib.Log(msg)
class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]


# A Go Slice primitive wrapper for Python
# Usage example:
#
# nums = GoSlice((c_void_p * 5)(74, 4, 122, 9, 12), 5, 5)
#
# Calling go function (i.e Sort) via shared_object would be
# nums = GoSlice((c_void_p * 5)(74, 4, 122, 9, 12), 5, 5)
# lib.Sort.argtypes = [GoSlice]
# lib.Sort.restype = None
# lib.Sort(nums)
class GoSlice(Structure):
    _fields_ = [
        ("data", POINTER(c_void_p)),
        ("len", c_longlong), ("cap", c_longlong)
    ]


class GoInterface(Structure):
    _fields_ = [
        ("t", POINTER(c_void_p)),
        ("v", POINTER(c_void_p))
    ]


class GoError(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]


class SMPCAdditive(Structure):
    _fields_ = [
        ("r0", POINTER(c_void_p)),
        ("r1", GoInterface)
    ]
