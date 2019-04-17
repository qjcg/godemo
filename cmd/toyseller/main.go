package main

import (
	"flag"
	"fmt"

	"github.com/qjcg/godemo/pkg/toy"
)

func main() {
	flagName := flag.String("n", "", "name of toy")
	flagWeight := flag.Int("w", 0, "weight of toy")
	flag.Parse()

	t := toy.New(*flagName, *flagWeight)
	fmt.Printf("Name: %s, Weight: %d, OnHand: %d, Sold %d\n",
		t.Name,
		t.Weight,
		t.UpdateOnHand(55),
		t.UpdateSold(12),
	)
}
