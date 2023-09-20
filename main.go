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
	ruleset1 := modules.NewRuleSet(1, "lead_type")
	ruleset2 := modules.NewRuleSet(2, "method_of_contact")
	rule1 := modules.NewRule(1, "Priamry")
	rule2 := modules.NewRule(2, "Nawy Now")
	rule3 := modules.NewRule(3, "Whats App")
	rule4 := modules.NewRule(4, "Facebook")
	rule5 := modules.NewRule(5, "Google")
	user1 := modules.NewUser(1, "AMR")
	user2 := modules.NewUser(2, "WAEL")
	user3 := modules.NewUser(3, "WEW")

	ruleset1.AddNewRule(&rule1)
	ruleset1.AddNewRule(&rule2)
	ruleset2.AddNewRule(&rule3)
	ruleset2.AddNewRule(&rule4)
	ruleset2.AddNewRule(&rule5)

	user1.SetRuleForUser(1, 1)
	user1.SetRuleForUser(2, 3)
	user2.SetRuleForUser(1, 1)
	user2.SetRuleForUser(2, 4)
	user3.SetRuleForUser(1, 2)

	fmt.Println("RULESET 1")
	fmt.Println(ruleset1)
	fmt.Println("RULESET 2")
	fmt.Println(ruleset2)
	fmt.Println("AMR")
	fmt.Println(user1)
	fmt.Println("WAEL")
	fmt.Println(user2)
	fmt.Println("WEW")
	fmt.Println(user3)

}
