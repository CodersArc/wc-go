// main.go
// This is the entry point for the wc-go application.

/*
Copyright © 2024 CodersArc <codersarc@gmail.com>
Licensed under the MIT License
See LICENSE file in the project root for full license information.
This file is part of the wc-go project, a learning project to implement a Unix-style
wc program as a Golang CLI.
For more information, visit https://github.com/codersarc/wc-go
*/
package main

import (
	"log"

	"github.com/codersarc/wc-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
