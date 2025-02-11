package main

import (
	"fmt"

	"github.com/Ayikoandrew/go-system-monitor/functions"
)

func main() {
	p := functions.GetProcessInfo()
	p = functions.SortByMemory(p)

	fmt.Println(p)
}
