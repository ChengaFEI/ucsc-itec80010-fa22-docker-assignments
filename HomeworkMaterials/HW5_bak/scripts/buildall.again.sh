#!/bin/bash
cd ../src/frontend
echo ""
echo "2nd and later times building images."
pwd
echo ""
echo "Compiling frontend.go."
go build frontend.go
cd ../backend
echo ""
echo "Compiling backend.go."
pwd
go build backend.go
cd ../.. # back to the top of dir tree.
echo ""
echo "Copying frontend and backend binaries to both images subdirectories."
pwd
cp src/frontend/frontend images/frontend
cp src/backend/backend images/backend
echo ""
echo "Copying config.txt to both images subdirectories."
pwd
##stat src/etc/config.txt
cp src/etc/config.txt images/frontend
cp src/etc/config.txt images/backend
cd scripts
echo ""
echo "Confirming what is in both images subdirectories."
ls -R ../images
echo ""
echo "Building frontend and backend container images."
docker build -t frontend ../images/frontend/.
docker build -t backend ../images/backend/.
docker images
