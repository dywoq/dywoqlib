//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//

#ifndef DYWOQLIB_SMART_PTR_HXX
#define DYWOQLIB_SMART_PTR_HXX

/** <dywoqlib/smart_ptr.hxx> synopsis

namespace dywoqlib {

enum class ptr_kind { shared, unique };

template <typename _Tp, ptr_kind _Kind> class smart_ptr {
public:
  using kind = ptr_kind;
  using type = _Tp;
  using pointer = _Tp *;
  using reference = _Tp &;

  explicit smart_ptr() noexcept;
  explicit smart_ptr(pointer __p_param_) noexcept;
  ~smart_ptr();
  smart_ptr(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::shared);

  smart_ptr(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::unique)
  = delete;

  [[nodiscard]] smart_ptr &operator=(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::shared);

  smart_ptr &operator=(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::unique)
  = delete;

  smart_ptr(smart_ptr &&__other_param_) noexcept;

  [[nodiscard]] smart_ptr &operator=(smart_ptr &&__other_param_) noexcept;

  [[nodiscard]] pointer get() const noexcept;

  void reset() noexcept;

  void reset(pointer __p_param_) noexcept;

  [[nodiscard]] size count() const noexcept
    requires(_Kind == ptr_kind::shared);

  [[nodiscard]] reference operator*() const noexcept;

  [[nodiscard]] pointer operator->() const noexcept;

  [[nodiscard]] explicit operator bool() const noexcept;

  static void reset(pointer &__p_param_) noexcept;
};

template <typename _Tp>
[[nodiscard]] smart_ptr<_Tp, ptr_kind::unique> make_unique_ptr(_Tp
__value_param) noexcept;

template <typename _Tp>
[[nodiscard]] smart_ptr<_Tp, ptr_kind::shared> make_shared_ptr(_Tp
__value_param) noexcept;

} // namespace dywoqlib

*/

#include "__config.hxx"
#include "types.hxx"

#if __cplusplus >= 202002LL
#  if DYWOQLIB_VERSION >= 202505LL
DYWOQLIB_BEGIN_NAMESPACE

enum class ptr_kind { shared, unique };

template <typename _Tp> struct __shared_control_block_ {
  DYWOQLIB_HIDDEN_FROM_ABI _Tp *__ptr_data_;
  DYWOQLIB_HIDDEN_FROM_ABI size __ref_count_data_;

  DYWOQLIB_HIDDEN_FROM_ABI explicit __shared_control_block_(_Tp *__p_param_)
      : __ptr_data_(__p_param_), __ref_count_data_(1) {
  }
  DYWOQLIB_HIDDEN_FROM_ABI ~__shared_control_block_() {
    delete __ptr_data_;
    __ptr_data_ = nullptr;
  }
};

template <typename _Tp, ptr_kind _Kind> class smart_ptr {
public:
  using kind = ptr_kind;
  using type = _Tp;
  using pointer = _Tp *;
  using reference = _Tp &;

private:
  pointer __ptr_data_;
  __shared_control_block_<_Tp> *__control_block_;

  DYWOQLIB_HIDDEN_FROM_ABI void __add_ref_() noexcept
    requires(_Kind == ptr_kind::shared)
  {
    if (__control_block_) {
      __control_block_->__ref_count_data_++;
    }
  }

  DYWOQLIB_HIDDEN_FROM_ABI void __release_shared_() noexcept
    requires(_Kind == ptr_kind::shared)
  {
    if (__control_block_) {
      if (--__control_block_->__ref_count_data_ == 0) {
        delete __control_block_;
        __control_block_ = nullptr;
      }
    }
  }

  DYWOQLIB_HIDDEN_FROM_ABI void __release_unique_() noexcept {
    delete __ptr_data_;
    __ptr_data_ = nullptr;
  }

public:
  DYWOQLIB_HIDDEN_FROM_ABI explicit smart_ptr() noexcept
      : __ptr_data_(nullptr), __control_block_(nullptr) {
  }

  DYWOQLIB_HIDDEN_FROM_ABI explicit smart_ptr(pointer __p_param_) noexcept
      : __ptr_data_(__p_param_), __control_block_(nullptr) {
    if constexpr (_Kind == ptr_kind::shared) {
      if (__p_param_) {
        __control_block_ = new __shared_control_block_<_Tp>(__p_param_);
      }
    }
  }

  DYWOQLIB_HIDDEN_FROM_ABI ~smart_ptr() noexcept {
    if constexpr (_Kind == ptr_kind::shared) {
      __release_shared_();
    } else {
      __release_unique_();
    }
  }

  DYWOQLIB_HIDDEN_FROM_ABI smart_ptr(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::shared)
      : __ptr_data_(__other_param_.__ptr_data_),
        __control_block_(__other_param_.__control_block_) {
    __add_ref_();
  }

  DYWOQLIB_HIDDEN_FROM_ABI smart_ptr(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::unique)
  = delete;

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI smart_ptr &
  operator=(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::shared)
  {
    if (this != &__other_param_) {
      __release_shared_();
      __ptr_data_ = __other_param_.__ptr_data_;
      __control_block_ = __other_param_.__control_block_;
      __add_ref_();
    }
    return *this;
  }

  DYWOQLIB_HIDDEN_FROM_ABI smart_ptr &
  operator=(const smart_ptr &__other_param_) noexcept
    requires(_Kind == ptr_kind::unique)
  = delete;

  DYWOQLIB_HIDDEN_FROM_ABI smart_ptr(smart_ptr &&__other_param_) noexcept
      : __ptr_data_(__other_param_.__ptr_data_),
        __control_block_(_Kind == ptr_kind::shared
                             ? __other_param_.__control_block_
                             : nullptr) {
    if constexpr (_Kind == ptr_kind::shared) {
      __other_param_.__control_block_ = nullptr;
    }
    __other_param_.__ptr_data_ = nullptr;
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI smart_ptr &
  operator=(smart_ptr &&__other_param_) noexcept {
    if (this != &__other_param_) {
      if constexpr (_Kind == ptr_kind::shared) {
        __release_shared_();
        __control_block_ = __other_param_.__control_block_;
        __other_param_.__control_block_ = nullptr;
      } else {
        __release_unique_();
      }
      __ptr_data_ = __other_param_.__ptr_data_;
      __other_param_.__ptr_data_ = nullptr;
    }
    return *this;
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI pointer get() const noexcept {
    if constexpr (_Kind == ptr_kind::shared) {
      return __control_block_ ? __control_block_->__ptr_data_ : nullptr;
    }
    return __ptr_data_;
  }

  DYWOQLIB_HIDDEN_FROM_ABI void reset() noexcept {
    if constexpr (_Kind == ptr_kind::shared) {
      __release_shared_();
      __ptr_data_ = nullptr;
      __control_block_ = nullptr;
    } else {
      __release_unique_();
    }
  }

  DYWOQLIB_HIDDEN_FROM_ABI void reset(pointer __p_param_) noexcept {
    if constexpr (_Kind == ptr_kind::shared) {
      if (__control_block_ && __control_block_->__ptr_data_ == __p_param_) {
        return;
      }
      __release_shared_();
      if (__p_param_) {
        __control_block_ = new __shared_control_block_<_Tp>(__p_param_);
        __ptr_data_ = __p_param_;
      } else {
        __control_block_ = nullptr;
        __ptr_data_ = nullptr;
      }
    } else {
      if (__ptr_data_ == __p_param_) {
        return;
      }
      __release_unique_();
      __ptr_data_ = __p_param_;
    }
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI size count() const noexcept
    requires(_Kind == ptr_kind::shared)
  {
    return __control_block_ ? __control_block_->__ref_count_data_ : 0;
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI reference operator*() const noexcept {
    return *get();
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI pointer operator->() const noexcept {
    return get();
  }

  [[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI explicit
  operator bool() const noexcept {
    return get() != nullptr;
  }

  DYWOQLIB_HIDDEN_FROM_ABI static void reset(pointer &__p_param_) noexcept {
    delete __p_param_;
    __p_param_ = nullptr;
  }
};

template <typename _Tp>
[[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI smart_ptr<_Tp, ptr_kind::unique>
make_unique_ptr(_Tp __value_param) noexcept {
  return smart_ptr<_Tp, ptr_kind::unique>(new _Tp(__value_param));
}

template <typename _Tp>
[[nodiscard]] DYWOQLIB_HIDDEN_FROM_ABI smart_ptr<_Tp, ptr_kind::shared>
make_shared_ptr(_Tp __value_param) noexcept {
  return smart_ptr<_Tp, ptr_kind::shared>(new _Tp(__value_param));
}

DYWOQLIB_END_NAMESPACE
#  endif
#endif

#endif