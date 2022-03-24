#! /bin/sh
set -e

PROJ="pathao" # your repo name
ORG_PATH="github.com/suaas21"
REPO_PATH="${ORG_PATH}/${PROJ}"

echo "checking go installed"

if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit 1
fi

echo "checking git installed"

if ! [ -x "$(command -v git)" ]; then
    echo "git is not installed"
    exit 1
fi

echo "setting GOPATH"

if [ -z "${GOPATH}" ]; then
    echo "set GOPATH"
    exit 1
fi


PATH="${PATH}:${GOPATH}/bin"
COMMIT=`git rev-parse --short HEAD`
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)

if [ -z "${TAG}" ]; then
    VERSION=${COMMIT}
else
    VERSION=${TAG}
fi

if [ -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ]; then
    VERSION="${VERSION}-dirty"
fi

echo "go installing....."

export GO111MODULE=on

go mod verify
go mod vendor
go fmt ./...
go install -v -ldflags="-X ${REPO_PATH}/version.Version=${VERSION}" ./cmd/...