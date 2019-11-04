#!/bin/bash

echo "Running container..."

if [[ ! -d data ]]; then
	echo "Creating data dir"
	mkdir data
fi

docker run \
	-d \
	-p 8888:8888 \
	--name cc_mkdocs_container \
	-v $(pwd)/data:/tmp/data \
	cc_mkdocs:latest
