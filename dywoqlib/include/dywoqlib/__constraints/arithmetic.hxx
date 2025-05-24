//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_ARITHMETIC_HXX
#define DYWOQLIB_CONSTRAINTS_ARITHMETIC_HXX

#include "../__config.hxx"
#include "floating_point.hxx"
#include "integer.hxx"
#include "is_constraint_valid.hxx"

#if __cplusplus >= 202002LL
#  if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct arithmetic_constraint
    : public is_constraint_valid<integer_constraint<_Tp>::status &&
                                 floating_point_constraint<_Tp>::status> {};

template <typename _Tp>
concept arithmetic_c = arithmetic_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#  endif
#endif

#endif
