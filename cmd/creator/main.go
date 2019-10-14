package creator

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// NewCmd defines the new command
var NewCmd = climax.Command{
	Name:  "new",
	Brief: "Create a new Go app or library",
	Help:  "Create a new Go app or library",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoNew()
	},
}

// InitCmd defines the init command
var InitCmd = climax.Command{
	Name:  "init",
	Brief: "Create a new Go app or library in an existing directory",
	Help:  "Create a new Go app or library in an existing directory",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoInit()
	},
}
