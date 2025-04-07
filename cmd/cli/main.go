package main

import (
	"flag"
	"fmt"
	"os"
	"todo/model"
	"todo/persist"
	"todo/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Item{})
	itemService := service.NewItemService(db)
	switch os.Args[1] {
	case "add":
		addCmd(itemService)
	case "list":
		listCmd(itemService)
	case "update":
		updateCmd(itemService)
	case "delete":
		deleteCmd(itemService)
	default:
		fmt.Println("expected 'add' or 'list' subcommands")
		os.Exit(1)
	}
}
func deleteCmd(service *service.ItemService) error {
	delete := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := delete.Uint("id", 0, "ID of the item")
	deleteName := delete.String("name", "", "Name of the item")
	deleteStatus := delete.String("status", "", "Status of the item")
	deleteComments := delete.String("comments", "", "Comments for the item")
	delete.Parse(os.Args[2:])
	deleteModel := struct {
		Name     string
		Status   model.Status
		Comments string
	}{}
	if *deleteId == 0 && *deleteName == "" && *deleteStatus == "" && *deleteComments == "" {
		fmt.Println("ID or Name or Status or Comments is required")
		return nil
	}
	if *deleteId != 0 {
		err := service.Delete(persist.WithId(uint(*deleteId)))
		return err
	}
	if *deleteName != "" {
		if len(*deleteName) > 255 {
			fmt.Println("Name too long")
			return nil
		}
		deleteModel.Name = *deleteName
	}
	if *deleteStatus != "" {
		if *deleteStatus != "completed" && *deleteStatus != "pending" && *deleteStatus != "in_progress" && *deleteStatus != "cancelled" {
			fmt.Println("Invalid status")
			return nil
		}
		deleteModel.Status = model.Status(*deleteStatus)
	}
	if *deleteComments != "" {
		if len(*deleteComments) > 255 {
			fmt.Println("Comments too long")
			return nil
		}
		deleteModel.Comments = *deleteComments
	}
	var err error
	if deleteModel.Status != "" {
		err = service.Delete(
			persist.WithStatus(deleteModel.Status),
			persist.WithName(deleteModel.Name),
			persist.WithComments(deleteModel.Comments))
	} else {
		err = service.Delete(
			persist.WithName(deleteModel.Name),
			persist.WithComments(deleteModel.Comments))
	}
	return err
}

func addCmd(service *service.ItemService) error {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	addName := add.String("name", "", "Name of the item")
	addStatus := add.String("status", "", "Status of the item")
	addComments := add.String("comments", "", "Comments for the item")
	add.Parse(os.Args[2:])
	status := model.StatusPending
	if *addName == "" {
		fmt.Println("Name is required")
		return nil
	}
	if *addStatus != "" {
		if *addStatus != "completed" && *addStatus != "pending" && *addStatus != "in_progress" && *addStatus != "cancelled" {
			fmt.Println("Invalid status")
			return nil
		}
		status = model.Status(*addStatus)
	}
	if *addComments != "" {
		if len(*addComments) > 255 {
			fmt.Println("Comments too long")
			return nil
		}
	}
	item := model.Item{
		Name:     *addName,
		Status:   status,
		Comments: *addComments,
	}
	err := service.Create(item)
	return err
}
func updateCmd(service *service.ItemService) error {
	update := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := update.Uint("id", 0, "ID of the item")
	updateName := update.String("name", "", "Name of the item")
	updateStatus := update.String("status", "", "Status of the item")
	updateComments := update.String("comments", "", "Comments for the item")
	update.Parse(os.Args[2:])
	updateModel := struct {
		Name     string
		Status   string
		Comments string
	}{}
	if *updateId == 0 {
		fmt.Println("ID is required")
		return nil
	}
	if *updateName != "" {
		if len(*updateName) > 255 {
			fmt.Println("Name too long")
			return nil
		}
		updateModel.Name = *updateName
	}
	if *updateStatus != "" {
		if *updateStatus != "completed" && *updateStatus != "pending" && *updateStatus != "in_progress" && *updateStatus != "cancelled" {
			fmt.Println("Invalid status")
			return nil
		}
		updateModel.Status = *updateStatus
	}
	if *updateComments != "" {
		if len(*updateComments) > 255 {
			fmt.Println("Comments too long")
			return nil
		}
		updateModel.Comments = *updateComments
	}
	item, err := service.Get(persist.WithId(uint(*updateId)))
	if err != nil {
		return err
	}
	if updateModel.Name != "" {
		item.Name = *updateName
	}
	if updateModel.Status != "" {
		item.Status = model.Status(*updateStatus)
	}
	if updateModel.Comments != "" {
		item.Comments = *updateComments
	}
	err = service.Update(item, persist.WithId(uint(*updateId)))
	return err
}
func listCmd(service *service.ItemService) error {
	list := flag.NewFlagSet("list", flag.ExitOnError)
	listName := list.String("name", "", "Name of the item")
	listStatus := list.String("status", "", "Status of the item")
	listComments := list.String("comments", "", "Comments for the item")
	list.Parse(os.Args[2:])
	if *listName != "" {
		if len(*listName) > 255 {
			fmt.Println("Name too long")
			return nil
		}
	}
	if *listStatus != "" {
		if *listStatus != "completed" && *listStatus != "pending" && *listStatus != "in_progress" && *listStatus != "cancelled" {
			fmt.Println("Invalid status")
			return nil
		}
	}
	if *listComments != "" {
		if len(*listComments) > 255 {
			fmt.Println("Comments too long")
			return nil
		}
	}
	items := []model.Item{}
	var err error
	if *listStatus != "" {
		items, err = service.List(
			persist.WithName(*listName),
			persist.WithStatus(model.Status(*listStatus)),
			persist.WithComments(*listComments),
		)
	} else {
		items, err = service.List(
			persist.WithName(*listName),
			persist.WithComments(*listComments),
		)
	}
	if err != nil {
		return err
	}
	for _, item := range items {
		fmt.Printf("[%s] - %d - %s - %s \n", iconMap[item.Status], item.ID, item.Name, item.Status)
	}
	return nil
}
