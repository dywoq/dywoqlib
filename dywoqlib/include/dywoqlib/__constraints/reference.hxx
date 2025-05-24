//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_REFERENCE_HXX
#define DYWOQLIB_CONSTRAINTS_REFERENCE_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if __cplusplus >= 202002LL 
#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct reference_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct lvalue_reference_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct rvalue_reference_constraint : public is_constraint_valid<false> {};

// reference constraint
template <typename _Tp>
struct reference_constraint<_Tp &> : public is_constraint_valid<true> {};
template <typename _Tp>
struct reference_constraint<_Tp &&> : public is_constraint_valid<true> {};

// lvalue reference constraint
template <typename _Tp>
struct lvalue_reference_constraint<_Tp &> : public is_constraint_valid<true> {};

// rvalue reference constraint
// clang-format off
template <typename _Tp>
struct rvalue_reference_constraint<_Tp &&> : public is_constraint_valid<true> {};
// clang-format on

template <typename _Tp>
concept reference_c = reference_constraint<_Tp>::status;

template <typename _Tp>
concept lvalue_reference_c = lvalue_reference_constraint<_Tp>::status;

template <typename _Tp>
concept rvalue_reference_c = rvalue_reference_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif
#endif

#endif
