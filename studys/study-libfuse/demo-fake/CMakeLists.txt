cmake_minimum_required (VERSION 2.8)

# projectname is the same as the main-executable
project(demo-fake)

add_definitions('-g')
add_definitions('-Wall')
add_definitions('-D_FILE_OFFSET_BITS=64')
add_definitions('-DFUSE_USE_VERSION=26')

aux_source_directory(. DIR_SRCS)
add_executable(${PROJECT_NAME} ${DIR_SRCS})

target_link_libraries(${PROJECT_NAME} fuse)
