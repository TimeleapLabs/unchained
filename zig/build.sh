#!/usr/bin/env bash

if [ ! -d "vendor/blst" ]; then
  git clone https://github.com/supranational/blst vendor/blst
fi

zig build -Doptimize=ReleaseSmall
