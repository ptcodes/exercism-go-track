// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "time"

// Hey should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9)
	return t
}
