// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./mkalldocs.sh

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gwaycc/goget/cmd/go/gointernal/get"
)

// Compare to $GOROOT/src/cmd/main.go
func main() {
	// show build version of go compiler
	for _, arg := range os.Args[1:] {
		if arg == "-ver" {
			fmt.Println(runtime.Version())
			return
		}
	}

	args := append([]string{"get"}, os.Args[1:]...)
	cmd := get.CmdGet
	cmd.Flag.Usage = func() { cmd.Usage() }
	if cmd.CustomFlags {
		args = args[1:]
	} else {
		cmd.Flag.Parse(args[1:])
		args = cmd.Flag.Args()
	}
	cmd.Run(cmd, args)
	return
}
