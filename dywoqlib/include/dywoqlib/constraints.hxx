//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// Copyright 2025 dywoq
//
#ifndef DYWOQLIB_CONSTRAINTS_HXX
#define DYWOQLIB_CONSTRAINTS_HXX

// clang-format off
/** 

<dywoqlib/constraints.hxx> synopsis:

namespace dywoqlib {
	// must be derived by all constraints
	template<bool _Bv> struct is_constraint_valid {
		using status_type = _Bv;
		static constexpr bool status;
		constexpr operator bool() noexcept;
	};
	
	inline namespace constraints_library {
		// constraints
		template<typename _Tp>
		struct floating_point_constraint;

		template<typename _Tp>
		struct integer_constraint;

		template<typename _Tp>
		struct signed_integer_constraint;

		template<typename _Tp>
		struct unsigned_integer_constraint;

		template<typename _Tp, typename _Ut>
		struct same_as_constraint;

		template<typename _Tp>
		struct void_constraint;

		template<typename _Tp>
		struct array_constraint;

		template<typename _Tp>
		struct pointer_constraint;

		template<typename _Tp>
		struct function_constraint;

		template<typename _Tp>
		struct reference_constraint;

		template<typename _Tp>
		struct const_constraint;

		template<typename _Tp>
		struct volatile_constraint;

		template<typename _Tp>
		struct arithmetic_constraint;

		template<typename _Tp>
		struct enum_constraint;

		template<typename _Tp>
		struct class_or_union_constraint;

		template<typename _Tp>
		struct lvalue_reference_constraint;

		template<typename _Tp>
		struct rvalue_reference_constraint;

		template<typename _Tp>
		struct class_constraint;

		template<typename _Tp>
		struct union_constraint;

		// concepts
		template<typename _Tp>
		concept floating_point_c;

		template<typename _Tp>
		concept integer_c;

		template<typename _Tp>
		concept signed_integer_c;

 		template<typename _Tp>
		concept unsigned_integer_c;

		template<typename _Tp, typename _Ut>
		concept same_as_c;

		template<typename _Tp>
		concept void_c;

		template<typename _Tp>
		concept array_c;

		template<typename _Tp>
		concept pointer_c;

		template<typename _Tp>
		concept function_c;

		template<typename _Tp>
		concept reference_c;

		template<typename _Tp>
		concept const_c;

		template<typename _Tp>
		concept volatile_c;

		template<typename _Tp>
		concept arithmetic_c;

		template<typename _Tp>
		concept enum_c;

		template<typename _Tp>
		concept class_or_union_c;

		template<typename _Tp>
		concept lvalue_reference_constraint;

		template<typename _Tp>
		concept rvalue_reference_constraint;

		template<typename _Tp>
		concept class_c;

		template<typename _Tp>
		concept union_c;
	}

	template <typename _Tp> struct constraints {
  		template <typename _Ut> inline constexpr static bool same_as() noexcept;
		inline constexpr static bool voidc() noexcept;
		inline constexpr static bool integer() noexcept;
		inline constexpr static bool signed_integer() noexcept;
		inline constexpr static bool unsigned_integer() noexcept;
		inline constexpr static bool floating_point() noexcept;
		inline constexpr static bool array() noexcept;
		inline constexpr static bool pointer() noexcept;
		inline constexpr static bool function() noexcept;
		inline constexpr static bool reference() noexcept;
		inline constexpr static bool constc() noexcept;
		inline constexpr static bool volatilec() noexcept;
		inline constexpr static bool arithmetic() noexcept;
		inline constexpr static bool enumc() noexcept;
		inline constexpr static bool lvalue_reference() noexcept;
		inline constexpr static bool rvalue_reference() noexcept;
		inline constexpr static bool classc() noexcept;
		inline constexpr static bool unionc() noexcept;
 	};
} 

*/
// clang-format on

#include "__config.hxx"

#if DYWOQLIB_VERSION >= 202505LL
#  include "__constraints/arithmetic.hxx"          // IWYU pragma: keep
#  include "__constraints/array.hxx"               // IWYU pragma: keep
#  include "__constraints/class.hxx"               // IWYU pragma: keep
#  include "__constraints/const.hxx"               // IWYU pragma: keep
#  include "__constraints/constraints.hxx"         // IWYU pragma: keep
#  include "__constraints/enum.hxx"                // IWYU pragma: keep
#  include "__constraints/floating_point.hxx"      // IWYU pragma: keep
#  include "__constraints/function.hxx"            // IWYU pragma: keep
#  include "__constraints/integer.hxx"             // IWYU pragma: keep
#  include "__constraints/is_constraint_valid.hxx" // IWYU pragma: keep
#  include "__constraints/pointer.hxx"             // IWYU pragma: keep
#  include "__constraints/reference.hxx"           // IWYU pragma: keep
#  include "__constraints/same_as.hxx"             // IWYU pragma: keep
#  include "__constraints/void.hxx"                // IWYU pragma: keep
#  include "__constraints/volatile.hxx"            // IWYU pragma: keep
#endif

#endif
