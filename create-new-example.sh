#!/usr/bin/env bash

source bootstrap.sh

rm -rf go-example
./build/go-james new --path go-example --package github.com/pieterclaerhout/go-example
