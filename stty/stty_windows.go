// +build windows

package stty

import "errors"

func IsInputTTY() (bool, error) { return false, errNA }

func StartInvisible() (err error) {
	return errNA
}

func StopInvisible() (err error) {
	return errNA
}

var errNA = errors.New("stty not available on windows")
