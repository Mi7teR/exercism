// Package bob is solution for task https://exercism.org/tracks/go/exercises/bob
package bob

import "strings"

// Hey returns Bob's response to the given input.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}

		return "Whoa, chill out!"
	}

	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}

	return "Whatever."
}
