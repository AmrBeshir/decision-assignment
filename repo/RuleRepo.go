package repo

import "database/sql"

type RuleRepo struct {
	db *sql.DB
}

func CreateRuleRepo(db *sql.DB) *RuleRepo {
	return &RuleRepo{db}
}

type ruleSQL struct {
	id        int    `db:"id"`
	value     string `db:"value"`
	rulesetId int    `db:"ruleset_id"`
}
