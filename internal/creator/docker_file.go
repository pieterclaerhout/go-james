package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/config"
)

const dockerfileTemplate = `FROM golang:alpine AS go-james

RUN apk update && apk add git && rm -rf /var/cache/apk/*
RUN GO111MODULE=on go get -u github.com/pieterclaerhout/go-james/cmd/go-james


FROM go-james AS mod-download

RUN mkdir -p /app

ADD go.mod /app
ADD go.sum /app

WORKDIR /app


RUN go mod download

FROM mod-download AS builder

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 go-james build -v


FROM scratch

COPY --from=builder "/app/build/{{outputPath}}" /

ENTRYPOINT ["/{{outputPath}}"]`

type dockerFile struct {
	text string
}

func newDockerFile(cfg config.Config) dockerFile {

	text := dockerfileTemplate
	text = strings.ReplaceAll(text, "{{outputPath}}", cfg.Project.Name)

	return dockerFile{
		text: text,
	}
}

func (g dockerFile) string() string {
	return g.text
}
