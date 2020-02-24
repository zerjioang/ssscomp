#!/bin/bash

#
# Copyright ssscomp Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

if [[ ! -f $GOPATH/bin/goreleaser ]]; then
	echo "Downloading gorelaser..."
	curl -sL https://git.io/goreleaser | bash
fi

echo "Generating new local release..."
git add . && \
git commit -s -m 'configuration of goreleaser updated' && \
git tag "$(date +'%Y%m%d%H%M%S')-$(git log --format=%h -1)" && \
goreleaser --skip-publish
echo "Done"