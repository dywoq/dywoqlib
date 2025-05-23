//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_IMPLEMENTATION_HXX
#define DYWOQLIB_CONSTRAINTS_IMPLEMENTATION_HXX

#include "../__config.hxx"

#if __cplusplus >= 202002LL 
#if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE
inline namespace constraints_library {

class __implementation {
public:
  template <typename _Tp>
  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI inline static constexpr bool
  is_enum() noexcept {
#  if __has_builtin(__is_enum) || defined(_MSC_VER)
    return __is_enum(_Tp);
#  else
#    if defined(__clang__) || defined(__GNUC__)
#      warning [dywoqlib/__constraints/implementation.hxx] Built-in __is_enum can't be found in your compiler. The returning value will be always true
#    elif defined(__MSC_VER)
#      pragma message(                                                         \
          "[dywoqlib/__constraints/implementation.hxx] Built-in __is_enum can't be found in your compiler.
    The returning value will be always true ")
#    endif
    return true;
#  endif
  }

  template <typename _Tp>
  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI inline static constexpr bool
  is_class() noexcept {
#  if __has_builtin(__is_class) || defined(_MSC_VER)
    return __is_class(_Tp);
#  else
#    if defined(__clang__) || defined(__GNUC__)
#      warning [dywoqlib/__constraints/implementation.hxx] Built-in __is_class can't be found in your compiler. \
		The returning value will be always true
#    elif defined(__MSC_VER)
#      pragma message(                                                         \
          "[dywoqlib/__constraints/implementation.hxx] Built-in __is_class can't be found in your compiler.
    The returning value will be always true ")
#    endif
    return true;
#  endif
  }

  template <typename _Tp>
  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI inline static constexpr bool
  is_union() noexcept {
#  if __has_builtin(__is_union) || defined(_MSC_VER)
    return __is_union(_Tp);
#  else
#    if defined(__clang__) || defined(__GNUC__)
#      warning [dywoqlib/__constraints/implementation.hxx] Built-in __is_union can't be found in your compiler. \
		The returning value will be always true
#    elif defined(__MSC_VER)
#      pragma message(                                                         \
          "[dywoqlib/__constraints/implementation.hxx] Built-in __is_union can't be found in your compiler.
    The returning value will be always true ")
#    endif
    return true;
#  endif
  }

  template <typename _Tp, typename _Ut>
  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI inline static constexpr bool
  is_base_of() noexcept {
#  if __has_builtin(__is_base_of) || defined(_MSC_VER)
    return __is_base_of(_Tp, _Ut);
#  else
#    if defined(__clang__) || defined(__GNUC__)
#      warning [dywoqlib/__constraints/implementation.hxx] Built-in __is_base_of can't be found in your compiler. \
		The returning value will be always true
#    elif defined(__MSC_VER)
#      pragma message(                                                         \
          "[dywoqlib/__constraints/implementation.hxx] Built-in __is_base_of can't be found in your compiler.
    The returning value will be always true ")
#    endif
    return true;
#  endif
  }
};

} // namespace constraints_library
DYWOQLIB_END_NAMESPACE
#endif
#endif

#endif
