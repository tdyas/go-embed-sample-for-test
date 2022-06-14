package main

import (
	_ "embed"
	"fmt"

	"github.com/tdyas/go-embed-sample-for-test/pkg"
)

func main() {
	fmt.Printf("%s\n", pkg.Message)
}
