*** Settings ***
Library     OperatingSystem
Library     supporting.py


*** Test Cases ***
Goal: Check that hash command works correctly
    Run And Check Expected Output
    ...    build/rcc holotree hash --json --show-blueprint --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_json_show.txt

    Run And Check Expected Output
    ...    build/rcc holotree hash --json --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_json.txt

    Run And Check Expected Output
    ...    build/rcc holotree hash --show-blueprint --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_txt_show.txt
    ...    use_stream=stderr

    Run And Check Expected Output
    ...    build/rcc holotree hash --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_txt.txt
    ...    use_stream=stderr

Goal: Check that hash command works correctly with devdeps
    Run And Check Expected Output
    ...    build/rcc holotree hash --json --show-blueprint --devdeps --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_devdeps_json_show.txt

    Run And Check Expected Output
    ...    build/rcc holotree hash --json --devdeps --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_devdeps_json.txt

    Run And Check Expected Output
    ...    build/rcc holotree hash --show-blueprint --devdeps --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_devdeps_txt_show.txt
    ...    use_stream=stderr

    Run And Check Expected Output
    ...    build/rcc holotree hash --devdeps --controller citests robot_tests/bare_action/package.yaml
    ...
    ...    robot_tests/expected/ht_hash/ht_hash_devdeps_txt.txt
    ...    use_stream=stderr
