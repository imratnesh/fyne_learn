package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type calculator struct {
	window    fyne.Window
	display   *canvas.Text
	equation  string
	operation string
	firstNum  float64
}

func newCalculator() *calculator {
	calc := &calculator{}
	calc.window = app.New().NewWindow("Calculator")
	calc.display = canvas.NewText("", color.NRGBA{R: 0, G: 255, B: 255, A: 255})
	calc.display.TextStyle = fyne.TextStyle{Monospace: true}
	calc.display.Alignment = fyne.TextAlignTrailing
	return calc
}

func (c *calculator) createButton(text string, action func()) *widget.Button {
	return widget.NewButton(text, action)
}

func (c *calculator) numberPressed(num string) {
	if c.operation == "=" {
		c.equation = ""
		c.operation = ""
	}
	c.equation += num
	c.display.Text = c.equation
	c.display.Refresh()
}

func (c *calculator) operationPressed(op string) {
	if c.equation == "" {
		return
	}

	if c.operation == "" {
		c.firstNum, _ = strconv.ParseFloat(c.equation, 64)
		c.operation = op
		c.equation = ""
	} else {
		c.calculate()
		c.operation = op
	}
}

func (c *calculator) calculate() {
	if c.equation == "" || c.operation == "" {
		return
	}

	secondNum, _ := strconv.ParseFloat(c.equation, 64)
	var result float64

	switch c.operation {
	case "+":
		result = c.firstNum + secondNum
	case "-":
		result = c.firstNum - secondNum
	case "×":
		result = c.firstNum * secondNum
	case "÷":
		if secondNum == 0 {
			c.display.Text = "Error: Division by zero"
			c.display.Refresh()
			return
		}
		result = c.firstNum / secondNum
	}

	c.equation = fmt.Sprintf("%g", result)
	c.display.Text = c.equation
	c.display.Refresh()
	c.operation = "="
}

func (c *calculator) clear() {
	c.equation = ""
	c.operation = ""
	c.firstNum = 0
	c.display.Text = ""
	c.display.Refresh()
}

func (c *calculator) loadUI() {
	// Create number buttons
	buttons := make([][]fyne.CanvasObject, 5)
	numbers := []string{"7", "8", "9", "÷", "4", "5", "6", "×", "1", "2", "3", "-", "0", ".", "=", "+"}

	row := 0
	col := 0
	for _, num := range numbers {
		if col == 4 {
			col = 0
			row++
		}
		if buttons[row] == nil {
			buttons[row] = make([]fyne.CanvasObject, 4)
		}

		button := c.createButton(num, func() {
			switch num {
			case "+", "-", "×", "÷":
				c.operationPressed(num)
			case "=":
				c.calculate()
			default:
				c.numberPressed(num)
			}
		})
		buttons[row][col] = button
		col++
	}

	// Create clear button
	clearButton := c.createButton("C", c.clear)

	// Create grid layout for buttons
	grid := container.NewGridWithColumns(4)
	for _, row := range buttons {
		for _, button := range row {
			grid.Add(button)
		}
	}

	// Create main container
	content := container.NewVBox(
		container.NewPadded(c.display),
		clearButton,
		grid,
	)

	c.window.SetContent(content)
	c.window.Resize(fyne.NewSize(300, 400))
}

func main() {
	calc := newCalculator()
	calc.loadUI()
	calc.window.ShowAndRun()
}
