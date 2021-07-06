package exec

import (
	"bufio"
	"io"
	"os"
	go_exec "os/exec"
)

// Executes the `cmd` with `args`, and returns the output as list of string.
func RunCmd(cmd string, args ...string) ([]string, error) {
	var outputs []string
	cmdExec := go_exec.Command(cmd, args...)
	cmdExec.Stdin = os.Stdin

	stdout, err := cmdExec.StdoutPipe()
	if err != nil {
		return outputs, err
	}

	cmdExec.Start()

	r := bufio.NewReader(stdout)
	for {
		line, _, err := r.ReadLine()
		if err == nil {
			outputs = append(outputs, string(line))
		} else if err == io.EOF {
			break // detect exec error or return code
		} else {
			return nil, err
		}
	}
	// This includes non-zero return code
	if err = cmdExec.Wait(); err != nil {
		return outputs, err
	}
	return outputs, nil
}
