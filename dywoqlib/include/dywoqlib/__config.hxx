//
//						dywoqlib (C++)
//
// Part of repository: https://github.com/dywoq/dywoqlib
// Under Apache License 2.0
//
// 2025 - dywoq
//
#ifndef DYWOQLIB___CONFIG_HXX
#define DYWOQLIB___CONFIG_HXX

// abi
#if defined(_WIN32) && defined(_MSC_VER)
#  define DYWOQLIB_EXPORTED_FROM_ABI __declspec(dllexport)
#  define DYWOQLIB_HIDDEN_FROM_ABI
#elif defined(__clang__) || defined(__GNUC__)
#  define DYWOQLIB_EXPORTED_FROM_ABI __attribute__((visibility("default")))
#  define DYWOQLIB_HIDDEN_FROM_ABI __attribute__((visibility("hidden")))
#else
#  define DYWOQLIB_EXPORTED_FROM_ABI
#  define DYWOQLIB_HIDDEN_FROM_ABI
#endif

// version
#if defined(__dywoqlib)
#  define DYWOQLIB_VERSION __dywoqlib
#else
#  define DYWOQLIB_VERSION 202505LL
#endif

// namespace
#define DYWOQLIB_BEGIN_NAMESPACE namespace dywoqlib {
#define DYWOQLIB_END_NAMESPACE }

#define DYWOQLIB_BEGIN_IMPLEMENTATION_NAMESPACE namespace __implementation {
#define DYWOQLIB_END_IMPLEMENTATION_NAMESPACE }

// detecting 32 bit and 64 bit system architectures
#if defined(_M_IX86) || defined(__i386__) || defined(__arm__) ||               \
    defined(__PPC__) || defined(_LP32) || defined(__ILP32__)
#  define DYWOQLIB_SYSTEM_ARCHITECTURE 32
#elif defined(_M_X64) || defined(_M_AMD64) || defined(__x86_64__) ||           \
    defined(__aarch64__) || defined(__powerpc64__) || defined(__LP64__)
#  define DYWOQLIB_SYSTEM_ARCHITECTURE 64
#endif

#endif
