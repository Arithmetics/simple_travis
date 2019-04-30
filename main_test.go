package main

import (
	"os"
	"testing"

	"github.com/EndFirstCorp/onedb"
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

// func TestInsertIntoMongo(t *testing.T) {
// 	m, err := mgo.Dial("localhost")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := insertIntoMongo(m, 12, 32); err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestReadMongo(t *testing.T) {
	m, err := mgo.Dial("localhost")

	if err != nil {
		t.Fatal(err)
	}

	if err := insertIntoMongo(m, 12, 32); err != nil {
		t.Fatal(err)
	}

	if err := insertIntoMongo(m, 54, 99); err != nil {
		t.Fatal(err)
	}

	basket, err := readMongo(m)

	if len(basket) != 2 {
		t.Errorf("wanted two baskets, got %v \n", len(basket))
	}
}

func TestGetCool(t *testing.T) {
	if getCool() != "cool" {
		t.Error("dependency management not working")
	}
}

func TestPsql(t *testing.T) {
	psqlConnection := os.Getenv("PSQLCONN")
	if psqlConnection == "" {
		psqlConnection = "postgresql://brocktillotson:uiop7890&*()@localhost:5432/travis_ci_test"
	}
	conn, _ := onedb.NewPgxFromURI(psqlConnection)

	fun := []Fun{}

	in := onedb.NewSqlQuery("INSERT INTO fun (cool, verycool) VALUES (32, 99)")

	conn.Execute(in)

	in2 := onedb.NewSqlQuery("INSERT INTO fun (cool, verycool) VALUES (32, 99)")

	conn.Execute(in2)

	query := onedb.NewSqlQuery(`SELECT cool, verycool FROM fun`)

	conn.QueryStruct(query, &fun)

	if len(fun) != 2 {
		t.Error("not enough records")
	}

}
