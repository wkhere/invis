// +build linux darwin

package stty

import (
	"syscall"
)

const (
	tty  = "/usr/bin/tty"
	stty = "/bin/stty"
)

func IsTerminal() (bool, error) {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(tty, []string{"tty", "-s"}, &attr)
	if err != nil {
		return false, err
	}

	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		return false, err
	}

	return ws.ExitStatus() == 0, nil
}

func StartInvisible() error {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(stty, []string{"stty", "-echo"}, &attr)
	if err != nil {
		return err
	}
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	return err
}

func StopInvisible() error {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(stty, []string{"stty", "echo"}, &attr)
	if err != nil {
		return err
	}
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	return err
}

var attr = syscall.ProcAttr{
	Files: []uintptr{0, 1, 2},
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
