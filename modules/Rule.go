package modules

type Rule struct {
	id    uint32
	value string
}

// Getters
func (r *Rule) GetId() uint32 {
	return r.id
}

func (r *Rule) GetValue() string {
	return r.value
}

// Setters
func (r *Rule) SetId(id uint32) {
	r.id = id
}

func (r *Rule) SetValue(value string) {
	r.value = value
}

func NewRule(id uint32, value string) *Rule {
	return &Rule{id, value}
}

func (r *Rule) isEqual(arg string) bool {
	return arg == r.value
}

func (r *Rule) isNotEqual(arg string) bool {
	return arg != r.value
}
