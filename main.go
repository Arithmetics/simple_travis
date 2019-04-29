package main

import (
	"fmt"
	"os"

	"github.com/EndFirstCorp/onedb/mgo"
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

func main() {
	dbConnection := os.Getenv("MONGO_CONNECTION")
	if dbConnection == "" {
		fmt.Println("MONGO_CONNECTION not set. Defaulting to localhost")
		dbConnection = "localhost"
	}
	m, err := mgo.Dial(dbConnection)

	if err != nil {
		fmt.Printf("Error connecting to MongoDB - %v \n", err)
	}

	insertIntoMongo(m, 6, 9)

	// fmt.Printf("%+v\n", len(readMongo(m)))
}
