#!/bin/sh -eu

# Put installed packages into ./bin
export GOBIN=$PWD/`dirname $0`/bin

if [ -d ".git" ] 
then
    export BUILD=`git rev-parse --short HEAD || ""`
    export BRANCH=`(git symbolic-ref --short HEAD | tr -d \/ ) || ""`
    if [ "$BRANCH" = master ]
    then
        export BRANCH=""
    fi

    export FLAGS="-X github.com/sam-blackfly/dabba/internal.Branch=$BRANCH -X github.com/sam-blackfly/dabba/internal.Build=$BUILD"
else
    export FLAGS=""
fi

mkdir -p bin

CGO_ENABLED=1 go build -trimpath -ldflags "$FLAGS" -v -o "bin/" ./cmd/...
