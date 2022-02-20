package main

import (
	"fmt"

	"github.com/skeptycal/types"
)

func main() {
	r := types.GenerateRoster(35)

	fmt.Println(r)
}
