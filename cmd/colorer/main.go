package main

import (
	"fmt"

	"github.com/claytonseraphin/internal/color"
)

func main() {
	redText := color.Text("This is a red text", color.Red)
	blueText := color.Text("This is a blue text", color.Blue)
	magentaText := color.Text("This is a magenta text", color.Magenta)
	CyanText := color.Text("This is a cyan text", color.Cyan)
	boldRedText := color.Text("This is a bold red text", color.Red, color.Bold)
	boldBlueUnderlineText := color.Text("This is a bold blue underline text with color",
		color.Blue, color.Bold, color.Underline)

	fmt.Println(redText)
	fmt.Println(blueText)
	fmt.Println(magentaText)
	fmt.Println(CyanText)
	fmt.Println(boldRedText)
	fmt.Println(boldBlueUnderlineText)
}
