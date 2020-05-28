package creator

import (
	"github.com/pieterclaerhout/go-james/internal/config"
)

const githubActionTemplate = `name: Build and Publish

on: [push]

jobs:

  build-test-staticcheck:
    name: Build, Test and Check
    runs-on: ubuntu-latest
    steps:

    - name: Set up Go 1.14
      uses: actions/setup-go@v1
      with:
        go-version: 1.14
      id: go

    - name: Environment Variables
      uses: FranzDiebold/github-env-vars-action@v1.0.0

    - name: Check out code into the Go module directory
      uses: actions/checkout@v2
      with:
        lfs: true

    - name: Restore Cache
      uses: actions/cache@preview
      id: cache
      with:
        path: ~/go/pkg
        key: 1.14-${{ runner.os }}-${{ hashFiles('**/go.sum') }}

    - name: Get go-james
      run: |
        go get -u github.com/pieterclaerhout/go-james/cmd/go-james

    - name: Get dependencies
      run: |
        go get -v -t -d ./...

    - name: Build
      run: |
        export PATH=${PATH}:` + "`" + `go env GOPATH` + "`" + `/bin
        go-james build

    - name: Test
      run: |
        export PATH=${PATH}:` + "`" + `go env GOPATH` + "`" + `/bin
        go-james test

    - name: Staticcheck
      run: |
        export PATH=${PATH}:` + "`" + `go env GOPATH` + "`" + `/bin
        go-james staticcheck

    - name: Package
      run: |
        export PATH=${PATH}:` + "`" + `go env GOPATH` + "`" + `/bin
        go-james package

    - uses: actions/upload-artifact@v2
      name: Publish
      with:
        name: ${{ env.GITHUB_REPOSITORY_NAME }}-${{ env.GITHUB_SHA_SHORT }}-${{ env.GITHUB_REF_NAME }}
        path: build/*.*`

type githubAction struct {
	text string
}

func newGithubAction(cfg config.Config) githubAction {
	return githubAction{
		text: githubActionTemplate,
	}
}

func (g githubAction) string() string {
	return g.text
}
