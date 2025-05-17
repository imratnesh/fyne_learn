package main

import (
	"log"

	"crud_todo_app/database"
	"crud_todo_app/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	// Initialize the app
	myApp := app.New()
	myWindow := myApp.NewWindow("To-Do App")

	// Initialize database
	db, err := database.NewDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Create login page
	loginPage := ui.NewLoginPage(myWindow, db, func(userID int) {
		// On successful login, show todo page
		todoPage := ui.NewTodoPage(myWindow, db, userID)
		todoPage.Show()
	})

	// Show login page initially
	loginPage.Show()

	// Run the application
	myWindow.ShowAndRun()
}
