//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_VOID_HXX
#define DYWOQLIB_CONSTRAINTS_VOID_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct void_constraint : public is_constraint_valid<false> {};

template <> struct void_constraint<void> : public is_constraint_valid<true> {};

template <typename _Tp>
concept void_c = void_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif

#endif
