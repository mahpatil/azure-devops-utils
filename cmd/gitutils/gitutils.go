package main

import (
	"fmt"
	"inspiredbytech/azure-devops-utils/api"
	"log"
	"os"
)

func main() {
	var url = os.Getenv("AZURE_URL")
	var token = os.Getenv("AZURE_TOKEN")
	var org = os.Getenv("AZURE_ORG")
	var project = os.Getenv("AZURE_PROJECT")
	log.Print(token)

	url = fmt.Sprintf(url, org, project)
	log.Println(url)
	var client = api.NewApiClient(token)
	var retVal = AzureApiReturn{}
	err := client.Get(url, &retVal)
	if err != nil {
		log.Print(err)
	}

	log.Printf("Received %d repositories", retVal.Count)

	var url2 = "https://dev.azure.com/%s/%s/_apis/git/repositories/%s/refs"
	var url3 = ""
	var branchInfo = RepoRef{}
	fmt.Printf("Repo#, Branch#, Repo Name, Branch Name, Creator\n")
	for i, repo := range retVal.Repos {
		url3 = fmt.Sprintf(url2, org, project, repo.Name)
		//log.Print(url3)
		branchInfo = RepoRef{}
		err := client.Get(url3, &branchInfo)
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
