package packager

import (
	"strconv"

	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// PackageCmd defines the package command
var PackageCmd = climax.Command{
	Name:  "package",
	Brief: "Build the main executable of your project for windows / linux / mac",
	Help:  "Build the main executable of your project for windows / linux / mac",
	Flags: []climax.Flag{
		{
			Name:     "verbose",
			Short:    "v",
			Usage:    `--verbose`,
			Help:     `print the names of packages as they are compiled.`,
			Variable: false,
		},
		{
			Name:     "concurrency",
			Short:    "c",
			Usage:    `--concurrency`,
			Help:     `how many package processes can run simultaneously (defaults to the number of CPUs).`,
			Variable: false,
		},
	},
	Handle: func(ctx climax.Context) int {

		verbose := ctx.Is("verbose")
		concurrency, _ := ctx.Get("concurrency")
		concurrencyAsInt, _ := strconv.Atoi(concurrency)

		executor := internal.NewExecutor("")
		return executor.DoPackage(verbose, concurrencyAsInt)

	},
}
