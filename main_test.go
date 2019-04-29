package main

import (
	"testing"

	"github.com/sajari/fuzzy"
)

func TestGoodFunc(t *testing.T) {

	if goodFunc() != "looks good" {
		t.Error("ERROR")
	}
}

func TestWords(t *testing.T) {
	model := fuzzy.NewModel()
	model.SetThreshold(1)
	model.SetDepth(5)
	words := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}
	model.Train(words)

	if model.SpellCheck("yor") != "your" {
		t.Error("yoooo bad")
	}
}

// func TestBadFunc(t *testing.T) {

// 	if badFunc() != "looks good" {
// 		t.Error("ERROR")
// 	}
// }
