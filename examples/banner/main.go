package main

import (
	"github.com/CrowdSurge/banner"
	"github.com/xyproto/rainbow"
)

func main() {
	w := rainbow.NewWriter(2, 0.1, 0)
	w.Write([]byte(banner.PrintS("rainbow")+"\n"))
}
