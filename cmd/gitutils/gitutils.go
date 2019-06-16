package main

import (
	"flag"
	"fmt"
	"inspiredbytech/azure-devops-utils/api"
	"log"
	"os"
)

const azureDevOpsGitURL = "https://dev.azure.com/%s/%s/_apis/git/repositories"
const azureDevOpsToken = "AZURE_DEVOPS_TOKEN"
const azureDevOpsOrg = "AZURE_DEVOPS_ORG"
const azureDevOpsProject = "AZURE_DEVOPS_PROJECT"

func main() {
	var command string
	flag.StringVar(&command, "cmd", "", "a string")
	flag.Parse()

	var t Action
	fmt.Println(command)
	switch command {
	case "git-branches":
		t = AzureDevOpsGitBranches{}
	}
	if t != nil {
		t.Invoke()
	}
}

type Action interface {
	Invoke()
}

type AzureDevOpsGitBranches struct {
}

func (AzureDevOpsGitBranches) Invoke() {
	var baseUrl = azureDevOpsGitURL
	var token = os.Getenv(azureDevOpsToken)
	var org = os.Getenv(azureDevOpsOrg)
	var project = os.Getenv(azureDevOpsProject)

	baseUrl = fmt.Sprintf(baseUrl, org, project)
	log.Print(org, "-", project)
	var client = api.NewApiClient(token)
	var retVal = AzureApiReturn{}
	err := client.Get(baseUrl, &retVal)
	if err != nil {
		log.Print(err)
	}

	log.Printf("Received %d repositories", retVal.Count)

	var url2 = baseUrl + "/%s/refs"
	var repoUrl = ""
	var branchInfo = RepoRef{}
	fmt.Printf("Repo#, Branch#, Repo Name, Branch Name, Creator\n")
	for i, repo := range retVal.Repos {
		repoUrl = fmt.Sprintf(url2, repo.Name)
		//log.Print(repoUrl)
		branchInfo = RepoRef{}
		err := client.Get(repoUrl, &branchInfo)
		if err != nil {
			log.Print(err)
		} else {
			for j, ref := range branchInfo.Refs {
				//fmt.Println(i, ".", j, ", ", repo.Name, ", ", ref.Name, ", ", ref.Creator.DisplayName)
				fmt.Printf("%d, %d, %s, %s, %s\n", i+1, j+1, repo.Name, ref.Name, ref.Creator.DisplayName)
			}
		}

	}
}

type AzureApiReturn struct {
	Repos []*Repository `json:"value"`
	Count int           `json:count`
}

type Repository struct {
	ID   string `json:id`
	Name string `json:name`
}

type RepoRef struct {
	Refs []*struct {
		ID      string `json:id`
		Name    string `json:name`
		Creator struct {
			DisplayName string `json:"displayName"`
		} `json:"creator"`
	} `json:"value"`

	Count int `json:count`

	//    Visibility string `jsonapi:"visibility"`
}
