// Code generated by "string2enum -samepkg -linecomment -type Hero"; DO NOT EDIT.

package d2enum

import "fmt"

// HeroFromString returns the Hero enum corresponding to s.
func HeroFromString(s string) Hero {
	if len(s) == 0 {
		return 0
	}
	for i := range _Hero_index[:len(_Hero_index)-1] {
		if s == _Hero_name[_Hero_index[i]:_Hero_index[i+1]] {
			return Hero(i)
		}
	}
	panic(fmt.Errorf("unable to locate Hero enum corresponding to %q", s))
}
