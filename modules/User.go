package modules

import (
	"fmt"

	"github.com/kelindar/bitmap"
)

type User struct {
	id     uint32
	name   string
	ruling []bitmap.Bitmap
	size   uint32
}

// Setters
func (u *User) SetName(name string) {
	u.name = name
}

// Getters
func (u *User) GetName() string {
	return u.name
}

func NewUser(id uint32, name string) User {
	ruling := make([]bitmap.Bitmap, 0)
	return User{
		id,
		name,
		ruling,
		0,
	}
}

func (u *User) SetRuleForUser(ruleSetIndex uint32, ruleIndex uint32) {
	if u.size <= ruleSetIndex {
		u.size = ruleSetIndex + 1
		newRuling := make([]bitmap.Bitmap, u.size)
		copy(newRuling, u.ruling)
		u.ruling = newRuling
		fmt.Println(ruleSetIndex)
		fmt.Println(u.ruling)
	}
	u.ruling[ruleSetIndex].Set(ruleIndex)
}
