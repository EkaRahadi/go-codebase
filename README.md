## Setup Linter - Mandatory

### Windows
### binary will be $(go env GOPATH)/bin/golangci-lint
> 1. curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.0
> 2. golangci-lint --version

### MacOS
> `brew install golangci-lint && brew upgrade golangci-lint`

Or go to this [link](https://golangci-lint.run/welcome/install/#local-installation)

## Setup Git Hooks - Mandatory
1. run this on terminal `go install github.com/automation-co/husky@latest`
2. run this on terminal `go install golang.org/x/tools/cmd/goimports@latest`

## Extra
> All available commands is in `makefile`

## Run the app in dev environment

1. Ensure that `golang` is already installed in your system
2. Ensure docker is running
3. Clone the project and go to the directory
4. Type `go mod tidy` to install the packages
5. Type `docker compose up` to run the http server

## Create a branch when developing a feature

1. Create new branch from `dev`
2. Branch name format: "{type}/{ticket-number}-{feature-name}"
3. Type: feat, fix, chore
4. Example: "feat/SEAL-42-user–registration"
5. Make sure to pull from `dev` before requesting to merge

> Don't forget to `git pull` from the staging branch when you're developing a feature

## Commit convention

1. Commit convention, we are following https://www.conventionalcommits.org/en/v1.0.0
2. Commit message format: "{type}({scope}): {message}"
3. Type: feat, fix, test, chore, style
4. Example: "feat(registration): create registration page"
