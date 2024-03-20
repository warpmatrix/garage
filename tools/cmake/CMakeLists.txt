cmake_minimum_required(VERSION 3.10)

project(MyProject)

add_subdirectory(src)
include_directories(${CMAKE_SOURCE_DIR}/include)

add_executable(demo main.cpp)
target_link_libraries(demo my_lib)
