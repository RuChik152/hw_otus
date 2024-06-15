#!/usr/bin/env bash
set -xeuo pipefail

go test -cover

echo "PASS"