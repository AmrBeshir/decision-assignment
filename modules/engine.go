package modules

import (
	errors_enum "decision-assignment/errors"
	"errors"
	"fmt"
	"reflect"
)

type Engine struct {
	ruleSets []*RuleSet
	users    []*User
}

func InitEngine() Engine {
	return Engine{
		ruleSets: make([]*RuleSet, 0),
		users:    make([]*User, 0),
	}
}

// params: the attribute of the rule set
// return: index of the rule set
func (e *Engine) AddRuleSet(attribute string, rules []*Rule) int {
	ruleSet := NewRuleSetWithRules(0, attribute, rules)
	e.ruleSets = append(e.ruleSets, &ruleSet)
	return len(e.ruleSets)
}

func (e *Engine) SetTrue(ruleSetAttribute string, ruleAttribute string) error {

	ruleSet, err := e.findRuleSet(ruleSetAttribute)
	if err != nil {
		return err
	}

	err = ruleSet.SetRuleToTrue(ruleAttribute)
	if err != nil {
		return err
	}
	for _, el := range e.ruleSets {
		fmt.Println(*el)
	}
	return nil
}

func (e *Engine) Run(comparable Comparable) {
	r := reflect.ValueOf(comparable).Elem()
	count := r.NumField()
	for i := 0; i < count; i++ {
		fmt.Println(r.Type().Field(i).Name)
	}
}

func (e *Engine) Select(s Selector) {
	fmt.Println(s)
}

func (e *Engine) findRuleSet(attribute string) (*RuleSet, error) {
	for _, ruleSet := range e.ruleSets {
		if ruleSet.GetAttribute() == attribute {
			return ruleSet, nil
		}
	}
	return nil, errors.New(errors_enum.VALUE_NOT_FOUND_IN_ARRAY)

}
