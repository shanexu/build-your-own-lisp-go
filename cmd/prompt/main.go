package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit\n")

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "lispy> ",
		HistoryFile:     ".lispy.his",
		InterruptPrompt: "^C",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if errors.Is(err, readline.ErrInterrupt) || err == io.EOF {
				break
			}
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("No you're a %s\n", line)
	}
}
