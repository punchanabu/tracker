package vo

import "strings"

type Address string

func NewAddress(address string) Address {
	return Address(strings.ToLower(address))
}

func (a Address) String() string {
	return string(a)
}

func (a Address) IsValid() bool {
	// TODO: add blockchain validation logic
	return true
}

