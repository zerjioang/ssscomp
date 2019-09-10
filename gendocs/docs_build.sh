#!/bin/bash

echo "Generating production files..."
docker exec cc_mkdocs_container sh -c "cd ssscomp_docs && mkdocs build"
if [[ -d data/ssscomp_docs/site ]]; then
	echo "Chaincode documentation successfuly built"
	ls -alh data/ssscomp_docs/site
	echo "Setting ownership of the docs..."
	sudo chown $USER -R data/ssscomp_docs/site
	echo "Moving documentation to gitlab pages folder..."
	rm -rf ../site ../public
	mv ./data/ssscomp_docs/site ../
	mv ../site ../public
fi
