package core

import (
	"errors"
	"slices"
)

type User struct {
	Id        int
	Interests []int
	Companion int
	Age       int
	Sex       int
}

func (u *User) AddOrRemoveInterest(i int) {
	err := u.AddInterest(i)

	if err != nil {
		u.RemoveInterest(i)
	}
}

func (u *User) AddInterest(i int) error {
	if u.Interests == nil {
		u.Interests = []int{}
	}

	if !slices.Contains(u.Interests, i) {
		u.Interests = append(u.Interests, i)
	} else {
		return errors.New("User alredy has this interest")
	}

	return nil
}

func (u *User) RemoveInterest(i int) {
	newInterests := []int{}

	for j := 0; j < len(u.Interests); j++ {
		if u.Interests[j] != i {
			newInterests = append(newInterests, u.Interests[j])
		}
	}

	u.Interests = newInterests
}

func (u *User) SetSex(s int) {
	u.Sex = s
}

func (u *User) SetAge(a int) {
	u.Sex = a
}

func (u *User) SetCompanion(c int) {
	u.Companion = c
}

func (u *User) RemoveCompanion() {
	u.Companion = 0
}
