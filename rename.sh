#!/usr/bin/env bash

VENDOR=${1}
PACKAGE=${2}

find ./ -type f -iname "*.go" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find ./ -type f -iname "*.yml" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find ./ -type f -iname "*.graphqls" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find ./ -type f -iname "*.mod" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find ./ -type f -iname "*Dockerfile*" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find ./ -type f -iname "*Makefile*" -exec sed -i '.bak' 's/arquil/${VENDOR}/g' {} \;
find . -name '*.bak' -exec rm -rf {} +

find ./ -type f -iname "*.go" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find ./ -type f -iname "*.yml" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find ./ -type f -iname "*.graphqls" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find ./ -type f -iname "*.mod" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find ./ -type f -iname "*Dockerfile*" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find ./ -type f -iname "*Makefile*" -exec sed -i '.bak' 's/accounts/${PACKAGE}/g' {} \;
find . -name '*.bak' -exec rm -rf {} +
