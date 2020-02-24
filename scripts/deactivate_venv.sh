#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#



#get current script dir
p=$(cd "$(dirname "$0")" && pwd -P)
cd $p

echo "Deactivating virtualenv..."
deactivate
echo "virtualenv deactivated"