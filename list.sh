#!/bin/bash

find tests -iname "*.proto" | while read line; do
    name=$(echo ${line} | cut -d "/" -f2)
    file=$(echo ${line}| cut -d "/" -f3)
    protoc --debug_out="tests/${name}:tests/." ./tests/${name}/${file}
done
