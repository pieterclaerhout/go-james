package creator

import (
	"github.com/pieterclaerhout/go-james/internal/config"
)

const gitlabCITemplate = `image: golang:latest

variables:
  REPO_NAME: gitlab.com/$CI_PROJECT_PATH

before_script:
  - mkdir -p .go 
  - mkdir -p $GOPATH/src/$(dirname $REPO_NAME)
  - ln -svf $CI_PROJECT_DIR $GOPATH/src/$REPO_NAME
  - cd $GOPATH/src/$REPO_NAME
  - export PATH=$PATH:$GOPATH/bin

stages:
  - Build
  - Test
  - Publish

.go-cache:
  variables:
      GOPATH: $CI_PROJECT_DIR/.go
  cache:
    paths:
      - .go/pkg
      - .go/bin
    key: 
      files:
        - go.sum
      prefix: $CI_JOB_IMAGE

Build:
  stage: Build
  script:
    - go get -u github.com/pieterclaerhout/go-james/cmd/go-james
    - go get -v -t -d $(go list ./... )
    - go-james build
  extends:
    - .go-cache

Test:
  stage: Test
  script:
    - go-james test
  extends:
    - .go-cache

Staticcheck:
  stage: Test
  script:
    - go-james staticcheck
  extends:
    - .go-cache

Pack and Publish:
  stage: Publish
  script:
    - go-james package
  artifacts:
    name: $CI_PROJECT_NAME-$CI_SHORT_SHA-$CI_COMMIT_REF_SLUG
    paths:
      - build/*.*
    when: always
  extends:
    - .go-cache`

type gitlabCI struct {
	text string
}

func newGitlabCI(cfg config.Config) gitlabCI {
	return gitlabCI{
		text: gitlabCITemplate,
	}
}

func (g gitlabCI) string() string {
	return g.text
}
