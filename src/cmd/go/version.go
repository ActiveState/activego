// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print Go version",
	Long:      `Version prints the Go version, as reported by runtime.Version.`,
}

func runVersion(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("ActiveGo 1803 [%s] provided by ActiveState Software, Inc.\n", runtime.ActiveGoBuild())
	fmt.Printf("Built %s\n\n", runtime.ActiveGoBuildDate())

	fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
