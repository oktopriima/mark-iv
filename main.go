package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-iv/jobs"
	"github.com/oktopriima/mark-iv/tasks"
	"github.com/oktopriima/mark-v/configurations"
	"github.com/oktopriima/mark-v/database"
	"time"
)

// private variable declaration
var cfg configurations.Config
var db *gorm.DB

func init() {
	var err error
	cfg = configurations.NewConfig("yaml")

	// create connection database
	db, err = connectToDatabase(cfg)
	if err != nil {
		panic(err)
	}

	// beautify section
	fmt.Println(" __  __           _     _")
	fmt.Println("|  \\/  | __ _ _ _| |_  (_)_   _")
	fmt.Println("| |\\/| |/ _' | '_|  ,/ | |\\\\,//")
	fmt.Println("|_|  |_|\\__,_|_| |_|_\\ |_| \\_/")
	fmt.Println("                                   ")
}

// call connection to database
func connectToDatabase(cfg configurations.Config) (*gorm.DB, error) {
	db, err := database.MysqlConnection(cfg)
	return db, err
}

// create http client request to third party
func main() {
	schedule := gocron.NewScheduler()
	schedule.Every(1).Minutes().Do(runningTask)
	<-schedule.Start()
}

func runningTask() {
	var err error
	taskList, err := tasks.ParseTask(cfg.GetString("task_list_file"))
	if err != nil {
		// handle error
		panic(err)
	}

	now := time.Now().Format("15:04")
	for _, task := range taskList.Task {
		fmt.Println(now, task.ExecutingTime)
		if now == task.ExecutingTime {
			// register task
			switch task.Key {
			case "HTTP_REQUEST_EXAMPLE":
				// write some information
				fmt.Printf("task name : %s.\ndescription : %s. \n"+
					"executing at %s \n", task.Name, task.Description, task.ExecutingTime)

				// executing jobs
				j := jobs.NewHttpRequestJobs(cfg, db)
				j.GetHttpRequest()

			case "HTTP_REQUEST_EXAMPLE_POST":
				// write some information
				fmt.Printf("task name : %s.\ndescription : %s. \n"+
					"executing at %s \n", task.Name, task.Description, task.ExecutingTime)

				// executing jobs
				j := jobs.NewHttpRequestJobs(cfg, db)
				j.PostHttpRequest()
			}
		}
	}
}
