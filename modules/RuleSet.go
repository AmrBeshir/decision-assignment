package modules

import (
	errors_enum "decision-assignment/errors"
	"errors"

	bitmap "github.com/kelindar/bitmap"
)

type RuleSet struct {
	id        uint32
	attribute string
	ruling    bitmap.Bitmap
	rules     []*Rule
	size      uint32
}

// Getters
func (rs *RuleSet) GetRuling() bitmap.Bitmap {
	return rs.ruling
}

func (rs *RuleSet) GetRules() []*Rule {
	return rs.rules
}

func (rs *RuleSet) GetAttribute() string {
	return rs.attribute
}

// Setters
func (rs *RuleSet) SetRuling(arg bitmap.Bitmap) {
	rs.ruling = arg
}

func (rs *RuleSet) SetRules(args []*Rule) {
	rs.rules = args
}

func (rs *RuleSet) SetAttribute(arg string) {
	rs.attribute = arg
}

func NewRuleSet(id uint32, attribute string) RuleSet {
	//intialize empty bitmap
	ruling := bitmap.FromBytes(make([]byte, 0))

	//intialize empty array of Rules
	rules := make([]*Rule, 0)
	return RuleSet{
		id,
		attribute,
		ruling,
		rules,
		0,
	}
}

func (rs *RuleSet) AddNewRuleAttr(id uint32, value string) {
	rule := NewRule(id, value)
	rs.rules = append(rs.rules, &rule)
	rs.size++
	rs.ruling.Grow(rs.size)
}

func (rs *RuleSet) AddNewRule(arg *Rule) {
	rule := NewRule(arg.GetId(), arg.GetValue())
	rs.rules = append(rs.rules, &rule)
	rs.size++
	rs.ruling.Grow(rs.size)
}

func (rs *RuleSet) FindIndexOfRule(value string) (uint32, error) {
	var i uint32
	for i = 0; i < rs.size; i++ {
		if rs.rules[i].GetValue() == value {
			return i, nil
		}
	}

	return 0, errors.New(errors_enum.VALUE_NOT_FOUND_IN_ARRAY)
}

func (rs *RuleSet) SetRuleToTrue(value string) error {
	i, err := rs.FindIndexOfRule(value)

	if err != nil {
		return err
	}

	rs.ruling.Set(i)

	return nil
}
