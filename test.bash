#!/bin/bash

set -e

wget "$DB_URL" -O "$TMP_PATH"
sed -i '1,4d' "$TMP_PATH"
grep "iPhone OS 12" "$TMP_PATH" > "$DB_PATH"
grep "Chrome/56.0" "$TMP_PATH" >> "$DB_PATH"
go test -v -cpu="$CPU" -benchmem -bench=.
