package packager

import (
	"encoding/json"
)

type distributions []distribution

type distribution struct {
	GOOS   string `json:"GOOS"`
	GOARCH string `json:"GOARCH"`
}

func (d distribution) ShouldBuild() bool {

	// if d.GOOS == "aix" {
	// 	return false
	// }

	// if d.GOOS == "android" {
	// 	return false
	// }

	// if d.GOOS == "illumos" {
	// 	return false
	// }

	// if d.GOOS == "js" {
	// 	return false
	// }

	// if d.GOOS == "nacl" {
	// 	return false
	// }

	// if d.GOOS == "plan9" {
	// 	return false
	// }

	// if d.GOOS == "solaris" {
	// 	return false
	// }

	if d.GOOS != "linux" && d.GOOS != "darwin" && d.GOOS != "windows" {
		return false
	}

	if d.GOARCH != "386" && d.GOARCH != "amd64" {
		return false
	}

	// if strings.HasPrefix(d.GOARCH, "mips") {
	// 	return false
	// }

	// if strings.HasPrefix(d.GOARCH, "ppc") {
	// 	return false
	// }

	// if strings.HasPrefix(d.GOARCH, "s390x") {
	// 	return false
	// }

	return true
}

func (d distribution) String() string {
	return d.GOOS + "/" + d.GOARCH
}

func (d distribution) DirName() string {
	return d.GOOS + "-" + d.GOARCH
}

func (packager Packager) allPossibleDistributions() (distributions, error) {

	var allDistributions distributions

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
