#!/bin/bash

headerLines=2

run_test () {
  inputFile=$1
  expectedLines=$2
  testNumber=$3
  <$inputFile go run main.go -e | wc | awk '{print $1}' | \
    if [[ $(($(xargs) - $headerLines)) == $expectedLines ]]; then
      echo Test $testNumber [$inputFile] success
    else
      echo Test $testNumber [$inputFile] FAILED
    fi
}

# Test for multiple line breaks
run_test "input_files/input.txt" 3 1
# Tests temp file
run_test "input_files/input2.txt" 1 2
# Tests temp file without output
run_test "input_files/input3.txt" 0 3
# Tests temp file without output and line with line break
run_test "input_files/input4.txt" 1 4
# Tests temp file with/ without output and line with line break
run_test "input_files/input5.txt" 2 5
# Tests temp file with/ without output and line with multiple line breaks
run_test "input_files/input6.txt" 10 6
# Tests string continues without newline when previous temp file is completed
run_test "input_files/input7.txt" 1 7

# create test file - 10MB
# base64 /dev/urandom | head -c 10000000 > input8.txt
# create test file - 5GB
# base64 /dev/urandom | head -c 5000000000 > input_5gb.txt
# remove \n chars:
# cat input8.txt | tr -d '\n'

# Analyze profiler results
# go tool pprof -http=:8080 heap_profile.out
