package main

import (
	"fmt"
	"os"
)

func main() {
	svc, err := NewService("mig-agent", "MIG Agent", "Mozilla InvestiGator Agent")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "platform is %v\n", svc.String())
}
