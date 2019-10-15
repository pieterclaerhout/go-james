package creator

const mainLibTemplate = `package {{.ShortPackageName}}
`

const mainCmdTemplate = `package main

import(
	"fmt"
)

func main() {
	fmt.Println("{{.Project.Name}}")
}
`

const versionInfoTemplate = `package versioninfo

// AppName contains the name of the app
var AppName string

// Revision will be injected with the current commit hash
var Revision string

// Branch will be injected with the current branch name
var Branch string
`
