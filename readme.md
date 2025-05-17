/*
# Installation Guide for Fyne and Go

This guide will walk you through the steps to install Go and the Fyne toolkit on your system.

## Step 1: Install Go

1. **Download Go:**
   - Visit the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
   - Download the installer for your operating system.

2. **Install Go:**
   - Run the downloaded installer and follow the on-screen instructions.
   - Verify the installation by opening a terminal or command prompt and typing:
     ```
     go version
     ```
   - You should see the installed version of Go.

## Step 2: Set Up Go Environment

1. **Configure Go Environment Variables:**
   - Set the `GOPATH` environment variable to your workspace directory.
   - Add the Go binary path to your `PATH` environment variable.

2. **Verify Go Environment:**
   - Open a terminal or command prompt and type:
     ```
     go env
     ```
   - Ensure that the environment variables are set correctly.

## Step 3: Install Fyne

1. **Install Fyne CLI:**
   - Open a terminal or command prompt and run:
     ```
     go install fyne.io/fyne/v2/cmd/fyne@latest
     ```

2. **Verify Fyne Installation:**
   - Check that the Fyne CLI is installed by typing:
     ```
     fyne version
     ```
   - You should see the installed version of Fyne.

## Step 4: Create a Simple Fyne Application

1. **Create a New Directory for Your Project:**
   - Open a terminal or command prompt and create a new directory:
     ```
     mkdir my-fyne-app
     cd my-fyne-app
     ```

2. **Initialize a New Go Module:**
   - Run the following command to initialize a new Go module:
     ```
     go mod init my-fyne-app
     ```

3. **Create a Simple Fyne Application:**
   - Create a new file named `main.go` and add the following code:
     ```go
     package main

     import (
         "fyne.io/fyne/v2/app"
         "fyne.io/fyne/v2/container"
         "fyne.io/fyne/v2/widget"
     )

     func main() {
         myApp := app.New()
         myWindow := myApp.NewWindow("Hello Fyne")

         myWindow.SetContent(container.NewVBox(
             widget.NewLabel("Hello Fyne!"),
         ))

         myWindow.ShowAndRun()
     }
     ```

4. **Run Your Fyne Application:**
   - Execute the following command to run your application:
     ```
     go run main.go
     ```

Congratulations! You have successfully installed Go and Fyne, and created a simple Fyne application.
*/
