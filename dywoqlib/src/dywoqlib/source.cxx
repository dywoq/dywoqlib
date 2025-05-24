#include "../../include/dywoqlib/source.hxx"

DYWOQLIB_BEGIN_NAMESPACE

DYWOQLIB_EXPORTED_FROM_ABI const char *source_error_value =
    "__PRETTY_FUNCTION__ is not supported in your compiler";

DYWOQLIB_EXPORTED_FROM_ABI source::source() noexcept {
}

DYWOQLIB_EXPORTED_FROM_ABI source::source(const char *__file_name,
                                          const char *__function,
                                          int __line) noexcept
    : __file_name_(__file_name), __function_(__function), __line_(__line) {
}

DYWOQLIB_EXPORTED_FROM_ABI const char *source::file_name() const noexcept {
  return __file_name_;
}

DYWOQLIB_EXPORTED_FROM_ABI const char *source::function() const noexcept {
  return __function_;
}

DYWOQLIB_EXPORTED_FROM_ABI int source::line() const noexcept {
  return __line_;
}

DYWOQLIB_END_NAMESPACE