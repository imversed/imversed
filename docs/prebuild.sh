#!/usr/bin/env bash

rm -rf src/modules && mkdir -p src/modules

cp ../x/README.md ./src/modules/README.md

for D in ../x/*; do
  if [ -d "${D}" ] && [ -d "${D}/spec" ]; then
    rm -rf "src/modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "src/modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done
