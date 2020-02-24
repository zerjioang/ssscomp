#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#



cd "$(dirname "$0")"
# move to project root dir from ./scripts to ./
cd ..

## SCRIPT GLOBAL VARIABLES
virtualenv_name="python36_venv"

#installing requirements
# sudo apt-get install python3.6
sudo apt-get install python3-pip
sudo pip3 install virtualenv

# check if virtualenv is created
if [[ ! -d $virtualenv_name ]]; then
	echo "Virtual ENV does not exists. Creating it"
	virtualenv -p python3 $virtualenv_name
	echo "Virtualenv created at $(pwd)/$virtualenv_name"
else
	echo "Virtual ENV found at $(pwd)/$virtualenv_name"
fi

. ./scripts/activate_venv.sh

#upgrade to latest version pip and setup tools
echo "Upgrading PIP..."
pip install --upgrade pip

echo "Installing mkdocs..."
pip install mkdocs

echo "Installing material theme..."
pip install mkdocs-material

echo "Updating APIDOCS..."
cd ..
mkdocs build

. ./scripts/deactivate_venv.sh