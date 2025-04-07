package main

import (
	"todo/cmd/cli/cmd"
	"todo/model"
)

const (
	PendingIcon    = "⏳"
	InProgressIcon = "🔄"
	CompletedIcon  = "✅"
	CancelledIcon  = "❌"
)

var iconMap = map[model.Status]string{
	model.StatusPending:    PendingIcon,
	model.StatusInProgress: InProgressIcon,
	model.StatusCompleted:  CompletedIcon,
	model.StatusCancelled:  CancelledIcon,
}

func main() {
	cmd.Execute()
}
