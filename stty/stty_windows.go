// +build windows

package stty

import "errors"

func IsTerminal() bool { return false }

func StartInvisible() (err error) {
	return errNA
}

func StopInvisible() (err error) {
	return errNA
}

var errNA = errors.New("stty not available on windows")
