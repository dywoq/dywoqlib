package main

import (
	"fmt"
	"time"

	"github.com/dywoq/dywoqlib/console"
	"github.com/dywoq/dywoqlib/console/ansi"
)

// `console` package provides tools to work with console such as running commands and clearing the whole console screen.
// it also provides subpackage `ansi` - to work with ANSI escape codes.

func main() {
	// to run command, you would use console.Run(cmd string, args ...string).
	// 1. first argument is your command to run
	// 2. second arguments are parameters for your command.
	// this function returns a output ([]bytes slice) of the started command.
	// in our case, it will return "Hello, world!".
	output, err := console.Run("echo", "Hello, world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("output: %s\n", string(output))

	// you can also clear the entire console screen by using console.Clear()
	time.Sleep(2 * time.Second)
	console.Clear()

	// to create ANSI-colored message, you would use ansi.New(string).
	// the first parameter is a string value you want to be wrapped around ANSI escape codes.
	// ansi.New automatically sets a foreground and a background to ansi.None,
	// so you need it to set manually. you can use ansi.Base.SetFgColor(), as shown below:
	redText := ansi.New("hi!!").SetFgColor(ansi.Red)
	fmt.Printf("redText: %v\n", redText)

	// if you want to set the backgroud color, use ansi.Base.SetBgColor():
	greenBackground := ansi.New("bye!!").SetBgColor(ansi.Green)
	fmt.Printf("greenBackground: %v\n", greenBackground)

	// buf if you want to mix the background and foreground color together?
	// fortunately, SetFgColor and SetBgColor return ansi.Base - meaning you can use method chaining pattern.
	both := ansi.New("hi again!!").SetFgColor(ansi.Red).SetBgColor(ansi.Green)
	fmt.Printf("both: %v\n", both)

	// if you want to make ANSI-colored message without ansi.Base, it's possible.
	// you would use ansi.ApplyFg(string, ansi.Color):
	fg := ansi.ApplyFg("thank you!!", ansi.Red)
	fmt.Printf("fg: %s\n", fg)

	// if you want to make message with background without ansi.Base,
	// you would use ansi.ApplyBg(string, ansi.Color):
	bg := ansi.ApplyBg("no problem!!", ansi.Red)
	fmt.Printf("bg: %v\n", bg)

	// it's still possible to mix background and foreground, like in ansi.Base.
	// you just need to use ApplyBoth(string, Color, Color);
	// the first color is foreground, and the second one is background.
	result := ansi.ApplyBoth("i'm dywoq", ansi.Magenta, ansi.Black)
	fmt.Printf("result: %v\n", result)
}
