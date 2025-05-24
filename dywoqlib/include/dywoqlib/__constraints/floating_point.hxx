//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_FLOATING_POINT_HXX
#define DYWOQLIB_CONSTRAINTS_FLOATING_POINT_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if __cplusplus >= 202002LL 
#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct floating_point_constraint : public is_constraint_valid<false> {};

template <>
struct floating_point_constraint<float> : public is_constraint_valid<true> {};

template <>
struct floating_point_constraint<double> : public is_constraint_valid<true> {};

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif
#endif

#endif
