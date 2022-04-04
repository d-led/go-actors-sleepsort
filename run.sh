#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

cd protoactor-sleepsort
echo "--=== protoactor-sleepsort ===--"
go get . && go run .
cd ..
