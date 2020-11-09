package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/CrowdSurge/banner"
	"github.com/xyproto/rainbow"
)

func main() {
	text := filepath.Base(os.Args[0])
	if len(os.Args) > 1 {
		text = strings.Join(os.Args[1:], " ")
	}
	rainbow.NewWriter(2, 0.1, 0).Write([]byte(banner.PrintS(text) + "\n"))
}
