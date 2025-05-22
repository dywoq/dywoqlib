cmake_minimum_required(VERSION 4.0.0)
project(dywoqlib CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

file(GLOB SOURCES src/dywoqlib/*.cxx)

target_sources(${PROJECT_NAME} PUBLIC ${SOURCES})
target_include_directories(${PROJECT_NAME} PUBLIC include/dywoqlib)
