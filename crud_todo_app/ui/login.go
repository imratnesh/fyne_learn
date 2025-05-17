package ui

import (
	"crud_todo_app/database"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// LoginPage represents the login screen
type LoginPage struct {
	window    fyne.Window
	onSuccess func(userID int)
	db        *database.DB
}

// NewLoginPage creates a new login page
func NewLoginPage(window fyne.Window, db *database.DB, onSuccess func(userID int)) *LoginPage {
	return &LoginPage{
		window:    window,
		onSuccess: onSuccess,
		db:        db,
	}
}

// Show displays the login page
func (l *LoginPage) Show() {
	title := canvas.NewText("Todo App", nil)
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	username := widget.NewEntry()
	username.SetPlaceHolder("Enter Username")
	username.Text = "ratnesh" // Set default username
	username.Resize(fyne.NewSize(400, 40))
	username.TextStyle = fyne.TextStyle{Bold: true}

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Enter Password")
	password.Text = "ratnesh" // Set default password
	password.Resize(fyne.NewSize(400, 40))
	password.TextStyle = fyne.TextStyle{Bold: true}

	submitButton := widget.NewButton("Submit", nil)
	submitButton.Resize(fyne.NewSize(200, 40))
	submitButton.Importance = widget.HighImportance

	// Create form without submit
	formContainer := container.NewVBox()
	formContainer.Add(container.NewPadded(username))
	formContainer.Add(layout.NewSpacer())
	formContainer.Add(container.NewPadded(password))

	// Handle submit action
	submitButton.OnTapped = func() {
		userID, success := l.db.VerifyUser(username.Text, password.Text)
		if success {
			l.onSuccess(userID)
		} else {
			dialog.ShowError(fmt.Errorf("Invalid username or password"), l.window)
		}
	}

	buttonContainer := container.NewHBox(layout.NewSpacer(), submitButton, layout.NewSpacer())

	hint := widget.NewRichTextFromMarkdown(`**Default Login:**
- Username: ratnesh
- Password: ratnesh`)

	content := container.NewVBox(
		container.NewPadded(title),
		container.NewPadded(formContainer),
		container.NewPadded(buttonContainer),
		container.NewPadded(hint),
	)

	centered := container.NewCenter(content)
	padded := container.NewPadded(centered)
	l.window.SetContent(padded)
	l.window.Resize(fyne.NewSize(1080, 1920))
}
