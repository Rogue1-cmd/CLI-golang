package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

// Constants for environment variable settings
const (
	markName  = "CLI_REMINDER"
	markValue = "1"
)

func main() {
	//checking if there are enough arguments provided
	if len(os.Args) < 3 {
		fmt.Printf("Usage:%s <hh:mm> <text message \n>", os.Args[0])
		os.Exit(1)
	}
	//getting current time
	now := time.Now()

	//initializing a new when parser
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...) //adds common parsing rules

	//parsing the provided time argument
	t, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//check if parsing result is nil
	if t == nil {
		fmt.Println("Unable to parse time")
		os.Exit(2)
	}

	//check if the provided time is in the past
	if now.After(t.Time) {
		fmt.Println("set a future time!")
		os.Exit(3)
	}

	//calculating difference between now and the provided time
	diff := t.Time.Sub(now)

	//checks if environment variable is set for the reminder
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)

		//shows reminder notifications using beeep.Alert
		beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		//if environment variable is not set, fork a new process and execute the command again
		cmd := exec.Command(os.Args[0], os.Args[1:]...)

		//set environment variable in new process
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))

		//start the command
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		//inform user that the command will be displayed after the calculated time difference
		fmt.Println("Reminder will be displayed after", diff.Round(time.Second))
		os.Exit(0)
	}
}
