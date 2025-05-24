//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_TYPES_HXX
#define DYWOQLIB_TYPES_HXX

#include "__config.hxx"

#if __cplusplus >= 202002LL
#  if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE

#    if DYWOQLIB_SYSTEM_ARCHITECTURE == 64
using size DYWOQLIB_NODEBUG = unsigned long long;
#    elif DYWOQLIB_SYSTEM_ARCHITECTURE == 32
using size DYWOQLIB_NODEBUG = unsigned int;
#    endif

DYWOQLIB_END_NAMESPACE
#  endif
#endif

#endif
