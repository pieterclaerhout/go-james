package creator

const mainLibTemplate = `package {{.ShortPackageName}}
`

const mainCmdTemplate = `package main

import(
	"fmt"

	"{{.Project.Package}}/versioninfo"
)

func main() {
	fmt.Println("Project: "+ versioninfo.AppName)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)
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
