package main

import (
	"database/sql"
	"decision-assignment/modules"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "amr"
	password = "root"
	dbname   = "decision_assignment"
)

func main() {
	connectionString := "postgresql://amr:root@localhost:5432/decision_assignment"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Ping()

	fmt.Println(os.Getenv("DB_CONNECTION"))

	testing()

}

func testing() {

	e := modules.InitEngine()
	e.AddRuleSet("lead_type", []*modules.Rule{
		modules.NewRule(0, "Primary"),
		modules.NewRule(0, "Nawy Now"),
	})
	e.AddRuleSet("method_of_contact", []*modules.Rule{
		modules.NewRule(0, "Facebook"),
		modules.NewRule(0, "Google"),
	})
	err := e.SetTrue("method_of_contact", "Google")
	fmt.Println(err)
}
