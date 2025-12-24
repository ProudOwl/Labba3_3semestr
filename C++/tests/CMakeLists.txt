message(STATUS "=== Конфигурация тестов LABS3 ===")

include(GoogleTest)

set(TEST_SOURCES
    ArrayTest.cpp
    CBTTest.cpp
    DoublyListTest.cpp
    HashTest.cpp
    SinglyListTest.cpp
    QueueTest.cpp
    StackTest.cpp
)

foreach(TFILE ${TEST_SOURCES})

    message(STATUS "Добавляем тест: ${TFILE}")

    get_filename_component(TNAME ${TFILE} NAME_WE)

    add_executable(${TNAME} ${TFILE})
    target_link_libraries(${TNAME}
    PRIVATE
        DataStructures
        gtest
        gtest_main
)


    gtest_discover_tests(${TNAME})

    message(STATUS "Собираем бинарник: ${TNAME}")

endforeach()
