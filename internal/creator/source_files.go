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
