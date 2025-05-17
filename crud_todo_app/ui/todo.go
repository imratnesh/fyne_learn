package ui

import (
	"crud_todo_app/database"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TodoPage represents the todo list screen
type TodoPage struct {
	window   fyne.Window
	db       *database.DB
	userID   int
	taskList *widget.List
}

// NewTodoPage creates a new todo page
func NewTodoPage(window fyne.Window, db *database.DB, userID int) *TodoPage {
	return &TodoPage{
		window: window,
		db:     db,
		userID: userID,
	}
}

func getPriorityColor(p database.Priority) color.Color {
	switch p {
	case database.PriorityHigh:
		return color.NRGBA{R: 255, G: 68, B: 68, A: 255} // Red
	case database.PriorityMedium:
		return color.NRGBA{R: 255, G: 187, B: 51, A: 255} // Orange
	default:
		return color.NRGBA{R: 51, G: 181, B: 229, A: 255} // Blue
	}
}

// Show displays the todo page
func (t *TodoPage) Show() {
	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("Enter new task")

	prioritySelect := widget.NewSelect([]string{
		string(database.PriorityLow),
		string(database.PriorityMedium),
		string(database.PriorityHigh),
	}, nil)
	prioritySelect.SetSelected(string(database.PriorityMedium))
	prioritySelect.PlaceHolder = "Priority"

	t.taskList = widget.NewList(
		func() int {
			tasks, _ := t.db.GetTasks(t.userID)
			return len(tasks)
		},
		func() fyne.CanvasObject {
			priorityLabel := canvas.NewText("", color.White)
			priorityLabel.TextSize = 14
			priorityLabel.TextStyle = fyne.TextStyle{Bold: true}
			priorityLabel.Resize(fyne.NewSize(20, 20))

			taskLabel := canvas.NewText("", color.White)
			taskLabel.TextSize = 14
			taskLabel.TextStyle = fyne.TextStyle{Bold: true}

			editBtn := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)

			leftContent := container.NewHBox(
				container.NewPadded(priorityLabel),
				container.NewPadded(taskLabel),
			)
			rightContent := container.NewHBox(editBtn, deleteBtn)

			return container.NewHBox(
				leftContent,
				layout.NewSpacer(),
				rightContent,
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			tasks, _ := t.db.GetTasks(t.userID)
			if i < len(tasks) {
				task := tasks[i]
				box := o.(*fyne.Container)
				leftContent := box.Objects[0].(*fyne.Container)
				priorityBox := leftContent.Objects[0].(*fyne.Container)
				taskBox := leftContent.Objects[1].(*fyne.Container)
				priorityLabel := priorityBox.Objects[0].(*canvas.Text)
				taskLabel := taskBox.Objects[0].(*canvas.Text)
				rightContent := box.Objects[2].(*fyne.Container)
				editBtn := rightContent.Objects[0].(*widget.Button)
				deleteBtn := rightContent.Objects[1].(*widget.Button)

				priorityLabel.Text = string(task.Priority)
				priorityLabel.Color = getPriorityColor(task.Priority)
				priorityLabel.Refresh()

				taskLabel.Text = task.Text
				taskLabel.Color = color.White
				taskLabel.Refresh()

				editBtn.OnTapped = func() {
					editEntry := widget.NewEntry()
					editEntry.SetText(task.Text)
					editPriority := widget.NewSelect([]string{
						string(database.PriorityLow),
						string(database.PriorityMedium),
						string(database.PriorityHigh),
					}, nil)
					editPriority.SetSelected(string(task.Priority))

					dialog.ShowForm("Edit Task", "Save", "Cancel",
						[]*widget.FormItem{
							{Text: "Task", Widget: editEntry},
							{Text: "Priority", Widget: editPriority},
						},
						func(b bool) {
							if b && editEntry.Text != "" {
								err := t.db.UpdateTask(task.ID, editEntry.Text,
									database.Priority(editPriority.Selected))
								if err == nil {
									t.taskList.Refresh()
								}
							}
						}, t.window)
				}

				deleteBtn.OnTapped = func() {
					dialog.ShowConfirm("Delete Task",
						fmt.Sprintf("Delete '%s'?", task.Text),
						func(b bool) {
							if b {
								err := t.db.DeleteTask(task.ID)
								if err == nil {
									t.taskList.Refresh()
								}
							}
						}, t.window)
				}
			}
		},
	)

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		task := taskEntry.Text
		if task != "" {
			err := t.db.AddTask(t.userID, task,
				database.Priority(prioritySelect.Selected))
			if err == nil {
				taskEntry.SetText("")
				t.taskList.Refresh()
			}
		}
	})

	header := widget.NewLabelWithStyle("Tasks", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	content := container.NewBorder(
		container.NewVBox(
			header,
			container.NewGridWithColumns(2,
				taskEntry,
				container.NewGridWithColumns(2,
					prioritySelect,
					addButton,
				),
				// prioritySelect,
				// addButton,
			),
		),
		nil, nil, nil,
		t.taskList,
	)
	// Set minimum sizes for components
	taskEntry.Resize(fyne.NewSize(500, taskEntry.MinSize().Height))
	prioritySelect.Resize(fyne.NewSize(150, prioritySelect.MinSize().Height))
	addButton.Resize(fyne.NewSize(100, addButton.MinSize().Height))

	t.window.SetContent(content)
	// Set a reasonable default window size
	t.window.Resize(fyne.NewSize(800, 600))
}
