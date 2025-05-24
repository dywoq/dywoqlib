//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_POINTER_HXX
#define DYWOQLIB_CONSTRAINTS_POINTER_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if __cplusplus >= 202002LL
#  if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct pointer_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct pointer_constraint<_Tp *> : public is_constraint_valid<true> {};

template <typename _Tp>
concept pointer_c = pointer_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#  endif
#endif

#endif
