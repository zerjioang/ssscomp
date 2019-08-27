package common

import "fmt"

// Secret Sharing individual item, also known as Share
type Share struct {
	v int
}

// Creates new Share
func NewShare(v int) Share {
	return Share{v: v}
}

// Creates new Share as a ptr
func NewSharePtr(v int) *Share {
	s := NewShare(v)
	return &s
}

// String representation of a Share using `fmt` package
func (s Share) String() string {
	return fmt.Sprintf("share: %d", s.v)
}
