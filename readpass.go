package readpass

import (
	"bufio"
	"os"
	"syscall"
)

const tty = "/usr/bin/tty"
const stty = "/bin/stty"

type Reader struct {
	*bufio.Reader
}

func IsTerm() bool {
	var ws syscall.WaitStatus

	pid, err := syscall.ForkExec(tty, []string{"tty", "-s"}, &attr)
	panicIf(err)

	_, err = syscall.Wait4(pid, &ws, 0, nil)
	panicIf(err)

	return ws.ExitStatus() == 0
}

func NewStdinReader() Reader {
	return Reader{bufio.NewReader(os.Stdin)}
}

type ReadFunc func(Reader) (string, error)

func ReadLine(r Reader) (string, error) {
	return r.ReadString('\n')
}

func ReadPass(r Reader) (line string, err error) {
	var ws syscall.WaitStatus
	var pid int

	pid, err = syscall.ForkExec(stty, []string{"stty", "-echo"}, &attr)
	if err != nil {
		return
	}

	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		return
	}

	line, err = r.ReadString('\n')
	if err != nil {
		return
	}

	pid, err = syscall.ForkExec(stty, []string{"stty", "echo"}, &attr)
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
