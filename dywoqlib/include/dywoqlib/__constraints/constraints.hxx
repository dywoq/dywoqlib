//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_CONSTRAINTS_HXX
#define DYWOQLIB_CONSTRAINTS_CONSTRAINTS_HXX

#include "../__config.hxx"
#include "arithmetic.hxx"
#include "array.hxx"
#include "const.hxx"
#include "enum.hxx"
#include "function.hxx"
#include "integer.hxx"
#include "pointer.hxx"
#include "reference.hxx"
#include "same_as.hxx"
#include "void.hxx"
#include "volatile.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE

template <typename _Tp> struct constraints {
  template <typename _Ut> inline constexpr static bool same_as() noexcept {
    return same_as_constraint<_Tp, _Ut>::status;
  }
  inline constexpr static bool voidc() noexcept {
    return void_constraint<_Tp>::status;
  }
  inline constexpr static bool integer() noexcept {
    return integer_constraint<_Tp>::status;
  }
  inline constexpr static bool signed_integer() noexcept {
    return signed_integer_constraint<_Tp>::status;
  }
  inline constexpr static bool unsigned_integer() noexcept {
    return unsigned_integer_constraint<_Tp>::status;
  }
  inline constexpr static bool array() noexcept {
    return array_constraint<_Tp>::status;
  }
  inline constexpr static bool pointer() noexcept {
    return pointer_constraint<_Tp>::status;
  }
  inline constexpr static bool function() noexcept {
    return function_constraint<_Tp>::status;
  }
  inline constexpr static bool reference() noexcept {
    return reference_constraint<_Tp>::status;
  }
  inline constexpr static bool lvalue_reference() noexcept {
    return lvalue_reference_constraint<_Tp>::status;
  }
  inline constexpr static bool rvalue_reference() noexcept {
    return rvalue_reference_constraint<_Tp>::status;
  }
  inline constexpr static bool constc() noexcept {
    return const_constraint<_Tp>::status;
  }
  inline constexpr static bool volatilec() noexcept {
    return volatile_constraint<_Tp>::status;
  }
  inline constexpr static bool arithmetic() noexcept {
    return arithmetic_constraint<_Tp>::status;
  }
  inline constexpr static bool enumc() noexcept {
    return enum_constraint<_Tp>::status;
  }
};

DYWOQLIB_END_NAMESPACE
#endif

#endif
