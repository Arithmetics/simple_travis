package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/EndFirstCorp/onedb/mgo"
	"github.com/sajari/fuzzy"
)

func TestGoodFunc(t *testing.T) {

	if goodFunc() != "looks good" {
		t.Error("ERROR")
	}
}

func TestEnv(t *testing.T) {
	if envFunc() != "/Users/brocktillotson/6DDATA" {
		t.Error("not the right env")
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

func TestInsertIntoMongo(t *testing.T) {
	dbConnection := os.Getenv("MONGO_CONNECTION")
	if dbConnection == "" {
		fmt.Println("MONGO_CONNECTION not set. Defaulting to localhost")
		dbConnection = "localhost"
	}

	m, err := mgo.Dial(dbConnection)

	if err != nil {
		t.Fatal(err)
	}

	if err := insertIntoMongo(m, 12, 32); err != nil {
		t.Fatal(err)
	}

}
