//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_BASE_OF_HXX
#define DYWOQLIB_CONSTRAINTS_BASE_OF_HXX

#include "../__config.hxx"
#include "implementation.hxx"
#include "is_constraint_valid.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp, typename _Ut>
struct base_of_constraint
    : public is_constraint_valid<__implementation::is_base_of<_Tp, _Ut>()> {};

template <typename _Tp, typename _Ut>
concept base_of_c = base_of_constraint<_Tp, _Ut>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif

#endif

