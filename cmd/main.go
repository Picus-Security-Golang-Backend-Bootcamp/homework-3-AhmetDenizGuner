package main

import (
	"fmt"
	"os"
	"path/filepath"

	//data "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/data"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/command_handler"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	model "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
)

var bookSlice []model.Book = []model.Book{}

func init() {
	database.Connect()

	err := database.InitiliazeDatabase()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {

	args := os.Args

	//if user didn't give any argument to the program, return this warning message that includes possible commands and stop program
	if len(args) == 1 {
		printDefaultErrorMessage(args[0])
		return
	}

	switch args[1] {
	case "list":
		command_handler.List(args)
	case "search":
		command_handler.Search(args)
	case "buy":
		command_handler.Buy(args)
	case "delete":
		command_handler.Delete(args)
	case "add":
		command_handler.Add(args)
	case "update":
		command_handler.Update(args)
	default:
		printDefaultErrorMessage(args[0])
	}

}

//printDefaultErrorMessage prints general message to console
func printDefaultErrorMessage(projectPath string) {
	//taking executable project name from path
	projectName := filepath.Base(projectPath)
	fmt.Printf("%s uygulamasinda kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n buy => satin alma islemi icin\n delete => silme islemi icin \n", projectName)
}
