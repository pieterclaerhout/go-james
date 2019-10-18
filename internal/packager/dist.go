package packager

import (
	"encoding/json"
)

type distribution struct {
	GOOS   string `json:"GOOS"`
	GOARCH string `json:"GOARCH"`
}

func (d distribution) ShouldBuild() bool {

	if d.GOOS != "linux" && d.GOOS != "darwin" && d.GOOS != "windows" {
		return false
	}

	if d.GOARCH != "386" && d.GOARCH != "amd64" {
		return false
	}

	return true

}

func (d distribution) String() string {
	return d.GOOS + "/" + d.GOARCH
}

func (d distribution) DirName() string {
	return d.GOOS + "_" + d.GOARCH
}

func (packager Packager) allDistributionsToBuild() ([]distribution, error) {

	buildDistributions := []distribution{}

	allDistributions, err := packager.allPossibleDistributions()
	if err != nil {
		return buildDistributions, err
	}

	for _, distribution := range allDistributions {
		if !distribution.ShouldBuild() {
			continue
		}
		buildDistributions = append(buildDistributions, distribution)
	}

	return buildDistributions, nil

}

func (packager Packager) allPossibleDistributions() ([]distribution, error) {

	allDistributions := []distribution{}

	cmd := []string{"go", "tool", "dist", "list", "--json"}

	output, err := packager.RunReturnOutput(cmd, "", map[string]string{})
	if err != nil {
		return allDistributions, err
	}

	if err := json.Unmarshal([]byte(output), &allDistributions); err != nil {
		return allDistributions, err
	}

	return allDistributions, nil

}
