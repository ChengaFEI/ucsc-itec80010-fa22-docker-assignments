#!/bin/bash
echo ""
echo "Removing images from Docker local cache."
docker image rm frontend backend
docker images
echo ""
echo "Deleting binaries and config.txt files Dockerfile build directories."
rm ../images/frontend/frontend
rm ../images/frontend/config.txt
rm ../images/backend/backend
rm ../images/backend/config.txt
ls -R ../images
