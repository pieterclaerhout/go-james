package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/config"
)

const dockerfileTemplate = `FROM golang:alpine AS mod-download

RUN apk update && apk add git && rm -rf /var/cache/apk/*
RUN go get -u github.com/pieterclaerhout/go-james/cmd/go-james

RUN mkdir -p /app

ADD go.mod /app
ADD go.sum /app

WORKDIR /app

RUN go mod download


FROM mod-download AS builder

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -a --ldflags '-extldflags -static -s -w' -o "{{outputPath}}" "{{mainPackage}}"


FROM scratch

COPY --from=builder "/app/{{outputPath}}" /

ENTRYPOINT ["/{{outputPath}}"]`

type dockerFile struct {
	text string
}

func newDockerFile(cfg config.Config) dockerFile {

	text := dockerfileTemplate
	text = strings.ReplaceAll(text, "{{outputPath}}", cfg.Project.Name)
	text = strings.ReplaceAll(text, "{{mainPackage}}", cfg.Project.MainPackage)

	return dockerFile{
		text: text,
	}
}

func (g dockerFile) string() string {
	return g.text
}
