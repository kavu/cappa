#!/bin/bash

set -e

wget "$DB_URL" -O "$TMP_PATH"
sed -i '1,4d' "$TMP_PATH"
sed -n '145500,146000p' "$TMP_PATH" > "$DB_PATH"
go test -v -cpu="$CPU" -benchmem -bench=.
