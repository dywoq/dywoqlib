//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// 2025 - dywoq
//
#ifndef DYWOQLIB_CONFIG_HXX
#define DYWOQLIB_CONFIG_HXX

// abi
#if defined(_WIN32) && defined(_MSC_VER)
#define DYWOQLIB_EXPORTED_FROM_ABI __declspec(dllexport)
#define DYWOQLIB_HIDDEN_FROM_ABI
#elif defined(__clang__) || defined(__GNUC__)
#define DYWOQLIB_EXPORTED_FROM_ABI __attribute__((visibility("default")))
#define DYWOQLIB_HIDDEN_FROM_ABI __attribute__((visibility("hidden")))
#else
#define DYWOQLIB_EXPORTED_FROM_ABI
#define DYWOQLIB_HIDDEN_FROM_ABI
#endif

#endif
