#!/usr/bin/env bash

fswatch -o engine games platform cmd | while read -r; do
	echo "Go files changed, rebuilding..."
	GOOS=js GOARCH=wasm go build \
		-o web/public/game.wasm \
		./cmd/wasm
done
