package main

import (
	"os/exec"
	"strings"
)

type Runtime struct {
	path string
	name string
}

func GetRuntimes() []Runtime {
	test_commands := []string{"python", "python3", "python3.9", "python3.10"}
	results := make([]Runtime, 0)

	for i := range test_commands {
		cmd := exec.Command(test_commands[i], "--version")
		stdout, err := cmd.Output()

		if err != nil {
			continue
		}

		stdout_str := string(stdout)
		stdout_str = strings.TrimSpace(stdout_str)
		results = append(results, Runtime{path: test_commands[i], name: stdout_str})
	}

	return results
}

func Exec(index uint8, code string) string {
	py_exec := runtimes[index]
	cmd := exec.Command(py_exec.path, "-c", code)
	stdout, _ := cmd.CombinedOutput()

	return string(stdout)
}
