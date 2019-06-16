# Introduction
Miscellaneous utilities for Azure DevOps

# Getting Started
## A. Pre-requisites
* Golang is installed on your system

## B. How to run?
### 1. Configure environment

export AZURE_DEVOPS_ORG="<Azure DevOps Org>"
export AZURE_DEVOPS_PROJECT="<Azure DevOps projectname>"
export AZURE_DEVOPS_TOKEN="<Azure DevOps Token, generate this from your [security settings](https://dev.azure.com/nationwide-sl/_usersSettings/tokens)>"

### 2. Run commands

Run the following command to get list of git branches based on the specified environment variables
```
go run inspiredbytech/azure-devops-utils/cmd/gitutils -cmd=git-branches
```

# Contribute
Anyone and everyone is welcome to contribute new charts or updates. Please the below process:

1. Create a branch
2. Raise a pull request
3. Send it to contributors for approval