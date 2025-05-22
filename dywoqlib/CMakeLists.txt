cmake_minimum_required(VERSION 4.0.0)
project(dywoqlib CXX)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED True)

file(GLOB SOURCES src/dywoqlib/*.cxx)

target_sources(${PROJECT_NAME} PUBLIC ${SOURCES})
target_include_directories(${PROJECT_NAME} PUBLIC include/dywoqlib)
target_compile_options(${PROJECT_NAME} PRIVATE 
	"-fno-exceptions"
	"-fno-rtti"
	"-D__dywoqlib=202505LL"
)