/*
Copyright © 2023 x65534
*/
package main

import (
	"github.com/x65534/swfx/cmd/swfx/cmd"
	_ "github.com/x65534/swfx/cmd/swfx/cmd/extract"
	_ "github.com/x65534/swfx/cmd/swfx/cmd/ls"
)

func main() {
	cmd.Execute()
}
