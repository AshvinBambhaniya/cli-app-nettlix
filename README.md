### Training

This repository contains dataset information, and questions list.

We are using for asdf(https://asdf-vm.com/) (https://wiki.improwised.com/s/e03445b6-2fb5-4352-b50e-dcd56aa103ab) for managing multiple runtime versions.

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

### CLI Excercise

In this CLI excerise you will create the commands for below questions. We are using https://github.com/spf13/cobra package for making the command. CLI questions found in Issue.

Dataset: https://www.kaggle.com/datasets/dgoenrique/netflix-movies-and-tv-shows

Please refer https://github.com/Improwised/phoenix-shailja-goswami-backend/issues/1 for CLI excerice.

Please refer https://github.com/Improwised/phoenix-shailja-goswami-backend/issues/2 for API excercise.
