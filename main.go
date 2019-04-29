package main

import (
	"fmt"

	"github.com/sajari/fuzzy"
)

func goodFunc() string {
	return "looks good"
}

func badFunc() string {
	return "this one bad"
}

func main() {
	fmt.Printf("%+v\n", goodFunc())
	fmt.Printf("%+v\n", badFunc())

	model := fuzzy.NewModel()
	model.SetThreshold(1)
	model.SetDepth(5)
	words := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}
	model.Train(words)

	fmt.Println("\nSPELL CHECKS")
	fmt.Println("	Deletion test (yor) : ", model.SpellCheck("yor"))
	fmt.Println("	Swap test (uncel) : ", model.SpellCheck("uncel"))
	fmt.Println("	Replace test (dynemite) : ", model.SpellCheck("dynemite"))
	fmt.Println("	Insert test (dellicate) : ", model.SpellCheck("dellicate"))
	fmt.Println("	Two char test (dellicade) : ", model.SpellCheck("dellicade"))
}
