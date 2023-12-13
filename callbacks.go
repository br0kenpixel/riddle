package main

func RunCode() {
	py_code := code.Text
	result := Exec(uint8(runtimeSelector.SelectedIndex()), py_code)

	output.Text = result
	output.Refresh()
}

func ClearConsole() {
	output.Text = ""
	output.Refresh()
}
