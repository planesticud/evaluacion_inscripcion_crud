package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsParametricas_20190805_081420 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsParametricas_20190805_081420{}
	m.Created = "20190805_081420"

	migration.Register("InsertsParametricas_20190805_081420", m)
}

// Run the migrations
func (m *InsertsParametricas_20190805_081420) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190805_081420_inserts_parametricas.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *InsertsParametricas_20190805_081420) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190805_081420_inserts_parametricas.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}