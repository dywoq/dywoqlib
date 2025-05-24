//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_INTEGER_HXX
#define DYWOQLIB_CONSTRAINTS_INTEGER_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if __cplusplus >= 202002LL 
#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct unsigned_integer_constraint : public is_constraint_valid<false> {};

template <typename _Tp>
struct signed_integer_constraint : public is_constraint_valid<false> {};

// clang-format off

// unsigned integer constraint
template <>
struct unsigned_integer_constraint<unsigned char> : public is_constraint_valid<true> {};
template <>
struct unsigned_integer_constraint<unsigned short> : public is_constraint_valid<true> {};
template <>
struct unsigned_integer_constraint<unsigned int> : public is_constraint_valid<true> {};
template <>
struct unsigned_integer_constraint<unsigned long long> : public is_constraint_valid<true> {};

// signed integer constraint
template <>
struct signed_integer_constraint<char> : public is_constraint_valid<true> {};
template <>
struct signed_integer_constraint<short> : public is_constraint_valid<true> {};
template <>
struct signed_integer_constraint<int> : public is_constraint_valid<true> {};
template <>
struct signed_integer_constraint<long long> : public is_constraint_valid<true> {};

// clang-format on

template <typename _Tp>
struct integer_constraint
    : public is_constraint_valid<signed_integer_constraint<_Tp>::status &&
                                 unsigned_integer_constraint<_Tp>::status> {};

template <typename _Tp>
concept integer_c = integer_constraint<_Tp>::status;

template <typename _Tp>
concept unsigned_integer_c = unsigned_integer_constraint<_Tp>::status;

template <typename _Tp>
concept signed_integer_c = signed_integer_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif
#endif

#endif
