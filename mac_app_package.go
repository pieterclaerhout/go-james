package james

import (
	"image"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/JackMordaunt/icns"
	"github.com/groob/plist"
)

type MacAppPackage struct {
	ExecutablePath string
	IconPath       string
}

func NewMacAppPackage(executablePath string, iconPath string) *MacAppPackage {
	return &MacAppPackage{
		ExecutablePath: executablePath,
		IconPath:       iconPath,
	}
}

func (macAppPackage *MacAppPackage) Create() error {

	dstPath := strings.TrimSuffix(macAppPackage.ExecutablePath, path.Ext(macAppPackage.ExecutablePath)) + ".app"

	contentsPath := filepath.Join(dstPath, "Contents")
	if err := os.MkdirAll(contentsPath, 0755); err != nil {
		return err
	}

	macosPath := filepath.Join(contentsPath, "MacOS")
	if err := os.MkdirAll(macosPath, 0755); err != nil {
		return err
	}

	resourcesPath := filepath.Join(contentsPath, "Resources")
	if err := os.MkdirAll(resourcesPath, 0755); err != nil {
		return err
	}

	dstExecutablePackage := filepath.Join(macosPath, filepath.Base(macAppPackage.ExecutablePath))
	if err := os.Rename(macAppPackage.ExecutablePath, dstExecutablePackage); err != nil {
		return err
	}

	iconPath := filepath.Join(resourcesPath, "Icon.icns")
	if err := macAppPackage.createIcon(iconPath); err != nil {
		return err
	}

	infoPlistPath := filepath.Join(contentsPath, "Info.plist")
	if err := macAppPackage.createInfoPlist(infoPlistPath); err != nil {
		return err
	}

	return nil

}

func (macAppPackage MacAppPackage) createIcon(iconPath string) error {

	pngf, err := os.Open(macAppPackage.IconPath)
	if err != nil {
		return err
	}
	defer pngf.Close()

	srcImg, _, err := image.Decode(pngf)
	if err != nil {
		return err
	}

	dest, err := os.Create(iconPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	return icns.Encode(dest, srcImg)
}

func (macAppPackage MacAppPackage) createInfoPlist(infoPlistPath string) error {

	type infoPlist struct {
		CFBundleExecutable     string          `plist:"CFBundleExecutable"`
		CFBundleGetInfoString  string          `plist:"CFBundleGetInfoString"`
		CFBundleIconFile       string          `plist:"CFBundleIconFile"`
		CFBundleIdentifier     string          `plist:"CFBundleIdentifier"`
		CFBundleName           string          `plist:"CFBundleName"`
		CFBundlePackageType    string          `plist:"CFBundlePackageType"`
		NSAppTransportSecurity map[string]bool `plist:"NSAppTransportSecurity"`
	}

	info := infoPlist{
		CFBundleExecutable:    filepath.Base(macAppPackage.ExecutablePath),
		CFBundleGetInfoString: filepath.Base(macAppPackage.ExecutablePath),
		CFBundleIconFile:      "Icon.icns",
		CFBundleIdentifier:    filepath.Base(macAppPackage.ExecutablePath),
		CFBundleName:          filepath.Base(macAppPackage.ExecutablePath),
		CFBundlePackageType:   "APPL",
		NSAppTransportSecurity: map[string]bool{
			"NSAllowsArbitraryLoads": true,
		},
	}

	file, _ := os.Create(infoPlistPath)
	defer file.Close()

	encoder := plist.NewEncoder(file)
	return encoder.Encode(info)

}
