# Leetcode, challenges, misc stuff

## Python

### Dev
```sh
cd py
```


## Go

### Setup

#### Adjust go settings
```
# .bashrc/.zshrc
export GO111MODULE="auto"
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### Dev
```sh
cd go
```

#### Install gotests
```sh
go get -u github.com/cweill/gotests/...
```

#### Test

- Test all
```sh
make test
```
- Test specific challenge
```sh
go test -v github.com/alainrk/l337C0d3/challenge
```

#### Add challenge wizard
```sh
make new
```

#### Generate test cases
- Enter the challenge folder
- Write your function signature in ${challenge}.go
- Generate:
```sh
gotests --all --parallel -w ${challenge}.go
```
