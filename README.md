# Leetcode, challenges, misc stuff

## Setup

### Adjust go settings
```
# .bashrc/.zshrc
export GO111MODULE="auto"
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

## Dev

### Install gotests
```sh
go get -u github.com/cweill/gotests/...
```

### Test
```sh
make test
```

### Add challenge wizard
```sh
make new
```

### Generate test cases
- Enter the challenge folder
- Write your function signature in ${challenge}.go
- Generate:
```sh
gotests --all --parallel -w ${challenge}.go
```
