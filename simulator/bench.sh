#!/bin/bash

set -e

# go run ./cmd/main.go -config configs/alibabablock.toml # File not in repo
go run ./cmd/main.go -config configs/ds1.toml
go run ./cmd/main.go -config configs/loop.toml
go run ./cmd/main.go -config configs/oltp.toml
go run ./cmd/main.go -config configs/p3.toml
go run ./cmd/main.go -config configs/p8.toml
go run ./cmd/main.go -config configs/s3.toml
# go run ./cmd/main.go -config configs/twitter.toml # File not in repo
# go run ./cmd/main.go -config configs/twitter1.toml # File not in repo
# go run ./cmd/main.go -config configs/wikicdn.toml # File not in repo
go run ./cmd/main.go -config configs/zipf.toml
