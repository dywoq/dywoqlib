//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_SOURCE_HXX
#define DYWOQLIB_SOURCE_HXX

#include "__config.hxx"

#if __cplusplus >= 202002LL
#  if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE

class source {
private:
  const char *__file_name_;
  const char *__function_;
  int __line_;

public:
  DYWOQLIB_EXPORTED_FROM_ABI explicit source() noexcept;

  DYWOQLIB_EXPORTED_FROM_ABI explicit source(const char *__file_name,
                                             const char *__function,
                                             int __line) noexcept;

  DYWOQLIB_EXPORTED_FROM_ABI source &operator=(const source &__source) noexcept;

  DYWOQLIB_EXPORTED_FROM_ABI const char *file_name() const noexcept;

  DYWOQLIB_EXPORTED_FROM_ABI const char *function() const noexcept;

  DYWOQLIB_EXPORTED_FROM_ABI int line() const noexcept;
};

#    define DYWOQLIB_SOURCE_CURRENT                                            \
      ::dywoqlib::source(__FILE__, __func__, __LINE__)

DYWOQLIB_END_NAMESPACE
#  endif
#endif

#endif
