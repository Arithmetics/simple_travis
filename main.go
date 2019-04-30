package main

import (
	"fmt"
	"os"

	"github.com/EndFirstCorp/onedb"
	"github.com/EndFirstCorp/onedb/mgo"
	cool "github.com/arithmetics/simple_dependency"
	"gopkg.in/mgo.v2/bson"
)

// Basket full of fruit
type Basket struct {
	Apples  int `bson:"apples"      json:"apples"`
	Oranges int `bson:"oranges"               json:"oranges"`
}

func goodFunc() string {
	return "looks good"
}

func badFunc() string {
	return "this one bad"
}

func getCool() string {
	return cool.Cool()
}

func envFunc() string {
	return os.Getenv("DATADIR")
}

func insertIntoMongo(m mgo.Sessioner, x int, y int) error {

	e := Basket{
		Apples:  x,
		Oranges: y,
	}

	err := m.DB("travis").C("basket").Insert(e)

	return err
}

func readMongo(m mgo.Sessioner) ([]Basket, error) {
	baskets := []Basket{}

	err := m.DB("travis").C("basket").Find(bson.M{}).All(&baskets)

	return baskets, err
}

// Fun whatever
type Fun struct {
	Cool     int32
	VeryCool int32
}

func main() {

	psqlConnection := os.Getenv("PSQLCONN")
	if psqlConnection == "" {
		psqlConnection = "postgresql://brocktillotson:uiop7890&*()@localhost:5432/travis_ci_test"
	}
	conn, _ := onedb.NewPgxFromURI(psqlConnection)

	fun := []Fun{}

	query := onedb.NewSqlQuery(`SELECT cool, verycool FROM fun`)

	if err := conn.QueryStruct(query, &fun); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", fun)

	// fmt.Printf("%+v\n", len(readMongo(m)))
}
