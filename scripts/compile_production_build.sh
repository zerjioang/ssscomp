#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#



cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Building production version..."
removeLogsRegex="log\.(Info|Debug|Error)\(\"[\w,\s]*\"\)"
go build
