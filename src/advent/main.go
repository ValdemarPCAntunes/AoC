package main

import (
	"os"
)

func GetLineEnding() string {
	if os.PathSeparator == '\\' {
		return "\r\n"
	}
	return "\n"
}

func main() {
	ps := MakePromptSolver()
	for {
		user_input, _ := ps.reader.ReadString('\n')
		user_input = user_input[:len(user_input)-len(GetLineEnding())]
		ps.Exec(user_input)
	}
}
