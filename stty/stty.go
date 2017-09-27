package stty

import (
	"syscall"
)

const (
	tty  = "/usr/bin/tty"
	stty = "/bin/stty"
)

func IsTerm() bool {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(tty, []string{"tty", "-s"}, &attr)
	panicIf(err)

	_, err = syscall.Wait4(pid, &ws, 0, nil)
	panicIf(err)

	return ws.ExitStatus() == 0
}

func StartInvisible() (err error) {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(stty, []string{"stty", "-echo"}, &attr)
	if err != nil {
		return
	}
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	return
}

func StopInvisible() (err error) {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(stty, []string{"stty", "echo"}, &attr)
	if err != nil {
		return
	}
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	return
}

var attr = syscall.ProcAttr{
	Files: []uintptr{0, 1, 2},
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
