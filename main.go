package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("%v, commit %v, built at %v", version, commit, date)

	SlackToken := os.Getenv("SLACK_TOKEN")

	api := slack.New(SlackToken)

	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}
