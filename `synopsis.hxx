namespace dywoqlib {
	inline namespace constraints_library {
		template<typename _Tp>
		struct class_constraint;

		template<typename _Tp>
		struct union_constraint;

		template<typename _Tp>
		concept class_c = /** implementation */;

		template<typename _Tp>
		concept union_c = /** implementation */;
	}

	template<typename _Tp>
	struct constraints {
		inline constexpr static bool classc() noexcept;
		inline constexpr static bool unionc() noexcept;
	};
}