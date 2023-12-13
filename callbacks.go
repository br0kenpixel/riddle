package main

import "fmt"

func RunCode() {
	py_code := code.Text
	fmt.Printf("Executing: %s \n", py_code)
	result := Exec(uint8(runtimeSelector.SelectedIndex()), py_code)

	output.Text = result
	output.Refresh()
}

func ClearConsole() {
	output.Text = ""
	output.Refresh()
}
