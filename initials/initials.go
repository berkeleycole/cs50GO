package initials

import (
	"errors"
	"strings"
)

var (
	HasNumErr = errors.New("You cannot have numbers in your name.")
)

func GetInitials(name string) (string, error) {
	if strings.IndexAny(name, "1234567890") >= 0 {
		return "", HasNumErr
	}

	//name to all caps
	caps := strings.ToUpper(name)

	var initials []string

	initials = append(initials, string(caps[0]))

	//loop through all chars
	for i := range name {
		if caps[i] == 32 {
			initials = append(initials, string(caps[i+1]))
		}
	}

	ret := strings.Join(initials, "")

	return ret, nil
}
