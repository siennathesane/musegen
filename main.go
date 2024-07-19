package main

import (
	"context"
	"fmt"
	"musgen/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCmd().ExecuteContext(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
