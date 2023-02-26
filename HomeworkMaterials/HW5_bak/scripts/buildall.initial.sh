#!/bin/bash
cd ../src
echo ""
echo "First time building images."
echo ""
echo "Configuring contract and lib go.mod files."
cd contract
go mod init dclass.ucscx.edu/contract
cd ..
cd lib
go mod init dclass.ucscx.edu/lib
cd ..
cd frontend
echo ""
echo "Creating and configuring frontend go.mod file, in order to use local modules."
go mod init dclass.ucscx.edu/frontend
go mod edit -replace dclass.ucscx.edu/gorilla/mux=../gorilla/mux
go mod edit -replace dclass.ucscx.edu/contract=../contract
go mod edit -replace dclass.ucscx.edu/lib=../lib
go mod tidy
echo ""
echo "Compiling frontend.go."
go build frontend.go
cd ../backend
echo ""
echo "Creating and configuring backend go.mod file, in order to use local modules."
go mod init dclass.ucscx.edu/backend
go mod edit -replace dclass.ucscx.edu/gorilla/mux=../gorilla/mux
go mod edit -replace dclass.ucscx.edu/contract=../contract
go mod edit -replace dclass.ucscx.edu/lib=../lib
go mod tidy
echo ""
echo "Compiling backend.go."
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
