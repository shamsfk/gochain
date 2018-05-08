package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

// Console struct
type Console struct {
	commands []string
}

// NewConsole creates a console
func NewConsole() *Console {
	return &Console{}
}

// Run executes console in a blocking endless loop
func (c *Console) Run() {
	vm := otto.New()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		value, err := vm.Run(text)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(value)
		}

		if c.commands[len(c.commands)-1] != text {
			c.commands = append(c.commands, text)
		}

		// TODO: add up arrow command repeation
	}
}
