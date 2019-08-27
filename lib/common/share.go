package common

import "fmt"

type Share struct {
	v int
}

func NewShare(v int) Share{
	return Share{v:v}
}

func NewSharePtr(v int) *Share{
	s :=NewShare(v)
	return &s
}

func (s Share) String() string {
	return fmt.Sprintf("share: %d", s.v)
}