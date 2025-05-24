cmake_minimum_required(VERSION 4.0.0)
project(dywoqlib)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

file(GLOB SOURCES src/dywoqlib/*.cxx)

add_library(${PROJECT_NAME} SHARED ${SOURCES})

target_include_directories(${PROJECT_NAME} PUBLIC include/dywoqlib)
target_compile_options(${PROJECT_NAME} PRIVATE 
	"-fno-exceptions"
	"-fno-rtti"
	"-D__dywoqlib=202505LL"
)