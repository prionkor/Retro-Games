#!/usr/bin/env bash

fswatch -o engine game platform cmd | while read -r; do
	echo "Go files changed, rebuilding..."
	GOOS=js GOARCH=wasm go build \
		-o web/public/game.wasm \
		./cmd/wasm
done
