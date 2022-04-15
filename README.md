## Setup

### Adjust go settings
```
# .bashrc/.zshrc
export GO111MODULE="auto"
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### Install gotests
```sh
go get -u github.com/cweill/gotests/...
```

## Add a new challenge
```sh
./newChallenge.sh
```