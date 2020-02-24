#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

# on CI env variable CI=true is set
echo "Generating new release..."

# install required packages after successful built to release this version
apt install -y rpm snapd upx

# needed for the snap pipe
snap install snapcraft --classic

# docker login is required if you want to push docker images.
# DOCKER_PASSWORD should be a secret in your .travis.yml configuration
# - test -n "$TRAVIS_TAG" && docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"


echo "Downloading gorelaser..."
curl -sL https://git.io/goreleaser | bash

echo "Done"
