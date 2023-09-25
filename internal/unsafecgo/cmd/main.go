package main

import (
	"github.com/bisque-io/mdbx-go/internal/unsafecgo"
)

func main() {
	//cgo.CGO()
	unsafecgo.NonBlocking((*byte)(nil), 0, 0)
}
