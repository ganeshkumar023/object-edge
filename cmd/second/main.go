package main

import (
	"fmt"
	"strings"
)

type skipString struct {
	pos         int
	cur         int
	data        string
	transformed []rune
}

func NewSkipString(pos int, data string) skipString {
	return skipString{
		pos:  pos,
		cur:  1,
		data: strings.ToLower(data),
	}
}

func (s *skipString) TransformRune(pos int) {
	val := s.data[pos]
	if (val >= 'a' && val <= 'z') || (val >= 0 && val <= 9) {
		if s.cur == 3 {
			s.cur = 1
			if val >= 'a' && val <= 'z' {
				val = val - 32
			}
		} else {
			s.cur = s.cur + 1
		}
	}
	s.transformed = append(s.transformed, rune(val))
}

func (s *skipString) GetValueAsRuneSlice() []rune {
	return []rune(s.data)
}

func (s skipString) String() string {
	return string(s.transformed)
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// And change your code to look like this:

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
