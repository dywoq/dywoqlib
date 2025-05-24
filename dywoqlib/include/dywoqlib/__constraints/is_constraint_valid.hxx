//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_IS_CONSTRAINT_VALID_HXX
#define DYWOQLIB_CONSTRAINTS_IS_CONSTRAINT_VALID_HXX

#include "../__config.hxx"

#if __cplusplus >= 202002LL 
#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE

template <bool _Bv> struct is_constraint_valid {
  using status_type DYWOQLIB_NODEBUG = decltype(_Bv);
  DYWOQLIB_NODEBUG static constexpr bool status = _Bv;
  DYWOQLIB_NODEBUG [[nodiscard]] constexpr operator bool() noexcept {
    return status;
  }
};

DYWOQLIB_END_NAMESPACE
#endif
#endif

#endif
