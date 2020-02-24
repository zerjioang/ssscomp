#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Testing GO *_test.go files"
files=$(go list ./... | grep -v /vendor)
go test -v -race $files
