//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_VOLATILE_HXX
#define DYWOQLIB_CONSTRAINTS_VOLATILE_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct volatile_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct volatile_constraint<volatile _Tp> : public is_constraint_valid<false> {};

template <typename _Tp>
concept volatile_c = volatile_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif

#endif
