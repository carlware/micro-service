package main

import (
	"condomilux/condo-admin/cli/cmd"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("Unable to complete command execution", "err", err)
	}
}
