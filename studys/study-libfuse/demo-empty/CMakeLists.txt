cmake_minimum_required (VERSION 2.8)

# projectname is the same as the main-executable
project(demo-empty)

add_definitions('-g')
add_definitions('-Wall')
add_definitions('-D_FILE_OFFSET_BITS=64')
add_definitions('-DFUSE_USE_VERSION=26')

add_executable(${PROJECT_NAME} ${PROJECT_NAME}.c)

target_link_libraries(${PROJECT_NAME} fuse)
