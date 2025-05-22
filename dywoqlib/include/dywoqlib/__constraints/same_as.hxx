//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_SAME_AS_HXX
#define DYWOQLIB_CONSTRAINTS_SAME_AS_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp, typename _Ut>
struct same_as_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct same_as_constraint<_Tp, _Tp> : public is_constraint_valid<true> {};

template <typename _Tp, typename _Ut>
concept same_as_c = same_as_constraint<_Tp, _Ut>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif

#endif
