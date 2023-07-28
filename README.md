### Training

This repository contains dataset information, and questions list.


#### Pre-commit

We are using pre-commit for linting and code styling enforcement. So whenever you are commit the code pre-commit hook ensure code styling, and consistancy.

```shell
# Pre-requisite for pre-commit
# golangci-lint installation
$> go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

# gogetsum installation
$> go get gotest.tools/gotestsum

# gocyclo installation
$> go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

# goimport installation
$> go install golang.org/x/tools/cmd/goimports@latest

# gocritic installation
$> go install -v github.com/go-critic/go-critic/cmd/gocritic@latest

# Pre-commit installation
$> pip install pre-commit

# Initialize pre-commit
$> pre-commit
```


### After command created

```Shell
#run the main file
$> go run main.go

#to run particular command and flag
$> go run main.go cmdName --flagname[it can be multiple] 

```
