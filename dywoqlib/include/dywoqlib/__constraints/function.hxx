//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_FUNCTION_HXX
#define DYWOQLIB_CONSTRAINTS_FUNCTION_HXX

#include "../__config.hxx"
#include "is_constraint_valid.hxx"

#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

template <typename _Tp>
struct function_constraint : public is_constraint_valid<false> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...)> : public is_constraint_valid<true> {
};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) const>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) volatile>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) const volatile>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) noexcept>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) const noexcept>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) volatile noexcept>
    : public is_constraint_valid<true> {};

template <typename _Ret, typename... _Args>
struct function_constraint<_Ret(_Args...) const volatile noexcept>
    : public is_constraint_valid<true> {};

template <typename _Tp>
concept function_c = function_constraint<_Tp>::status;

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif

#endif
