package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

// Console struct
type Console struct {
	vm *otto.Otto
}

// NewConsole creates a console
func NewConsole() *Console {
	c := &Console{}
	c.vm = otto.New()

	return c
}

// RegisterFunction binds Go function to JS virtual machine
func (c Console) RegisterFunction(name string, function interface{}) {
	c.vm.Set(name, function)
}

// Run executes console in a blocking endless loop
func (c Console) Run() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		if text == ".exit\n" {
			return
		}

		value, err := c.vm.Run(text)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if value.IsDefined() {
				fmt.Println(value)
			}
		}

		// TODO: add up arrow command repeation
	}
}
