#!/bin/bash

echo "Generating new mkdocs project"
docker exec cc_mkdocs_container mkdocs new ssscomp_docs
if [[ -d data/ssscomp_docs ]]; then
	echo "Chaincode documentation folder successfully created"
	ls -alh data/ssscomp_docs
	sudo chown $USER -R data
	echo ""
	echo "Edit the files and once you finish, execute generate_docs.sh"
	echo ""
fi