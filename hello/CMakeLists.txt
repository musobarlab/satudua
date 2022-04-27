cmake_minimum_required(VERSION 2.8)

project (hello)

# omit for now
# set(CMAKE_BUILD_TYPE Release)

set( SOURCE_FILES 
    src/Hello.c
)

include_directories( ${CMAKE_CURRENT_SOURCE_DIR}/include )
# SHARED (dynamic) lib
# add_library( ${PROJECT_NAME} SHARED ${SOURCE_FILES} )
# STATIC lib
add_library( ${PROJECT_NAME} STATIC ${SOURCE_FILES} )