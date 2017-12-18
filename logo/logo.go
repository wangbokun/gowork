package logo

import "github.com/wangbokun/gowork/color"

var (
	// Version is the default version of SKM

	logo = `			OpsManager
			`
	Version = `							Version 0.0.1												`
	Frame   = "--------------------------------------------------------------------------"
)

func DisplayLogo() {
	color.Blue2(Frame)
	color.Yellow(logo)
	color.Blue(Version)
	color.Blue2(Frame)
}
