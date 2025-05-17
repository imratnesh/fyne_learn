# Todo App with Fyne

A modern Todo application built with Go and Fyne framework, featuring user authentication and task management with priority levels.

## Features

- User authentication (login system)
- Create, read, update, and delete tasks
- Priority levels for tasks (High, Medium, Low)
- Modern and responsive UI
- Cross-platform support

## Prerequisites

- Go 1.16 or higher
- Fyne dependencies:
  - For macOS: Xcode Command Line Tools
  - For Linux: gcc and X11 development files
  - For Windows: gcc (MinGW)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/imratnesh/fyne_learn.git
cd fyne_learn
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

To run the application:
```bash
cd crud_todo_app
go run main.go
```

Default login credentials:
- Username: ratnesh
- Password: ratnesh

## Project Structure

```
crud_todo_app/
├── database/         # Database operations
│   └── db.go
├── ui/              # User interface components
│   ├── login.go     # Login page
│   └── todo.go      # Todo list page
├── main.go          # Application entry point
└── go.mod           # Go module file
```

## Git Commands Used

1. Initialize git repository:
```bash
git init
```

2. Add files to git:
```bash
git add .
```

3. Create initial commit:
```bash
git commit -m "Initial commit"
```

4. Create GitHub repository and push code:
```bash
# Install GitHub CLI if not installed
brew install gh

# Login to GitHub
gh auth login

# Create repository and push code
gh repo create fyne_learn --public --source=. --remote=origin --push
```

5. If you need to remove large files from git history:
```bash
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch [large-file-path]' --prune-empty --tag-name-filter cat -- --all
git push -f origin master
```

## Building for Different Platforms

To build the application for different platforms:

```bash
# For macOS
go build -o todo-app

# For Windows
GOOS=windows GOARCH=amd64 go build -o todo-app.exe

# For Linux
GOOS=linux GOARCH=amd64 go build -o todo-app
```

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.
