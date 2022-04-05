#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

cd protoactor-sleepsort
go get . && go run .
cd ..

cd ergo-sleepsort
go get . && go run .
cd ..
