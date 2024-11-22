package components

import (
	"crypto/md5"
	"fmt"
)

// ID generates an id based on the seed string.
// will create duplicates if the seed is the same.
func ID(seed string) string {
	h := md5.Sum([]byte("111" + seed))
	s := fmt.Sprintf("id-%x", h)
	if len(s) > 10 {
		s = s[:10]
	}
	return s
}
