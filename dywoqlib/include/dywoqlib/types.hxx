//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// 2025 - dywoq
//
#ifndef DYWOQLIB_TYPES_HXX
#define DYWOQLIB_TYPES_HXX

#include "__config.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
namespace types {

#  if DYWOQLIB_SYSTEM_ARCHITECTURE == 64
using size = unsigned long long;
#  elif DYWOQLIB_SYSTEM_ARCHITECTURE == 32
using size = unsigned int int;
#  endif

} // namespace types
DYWOQLIB_END_NAMESPACE
#endif

#endif
