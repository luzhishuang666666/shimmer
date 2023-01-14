package main

import (
	"shimmer/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

func main() {
	cmd.Execute()
}
