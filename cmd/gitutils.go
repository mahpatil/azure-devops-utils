package main

import (
	"flag"
	"fmt"
	"inspiredbytech/azure-devops-utils/actions"
)

func main() {
	var command string
	flag.StringVar(&command, "cmd", "", "a string")
	flag.Parse()

	var t actions.Action
	fmt.Println(command)
	switch command {
	case "git-branches":
		t = actions.AzureDevOpsGitBranches{}
	}
	if t != nil {
		t.Invoke()
	}
}
