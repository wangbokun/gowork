package logo

import "github.com/wangbokun/gowork/color"

var (
	// Version is the default version of SKM
	Version = "0.1"
	logo    = "
	,        ,
	/(        )`
	\ \___   / |
	/- _  `-/  '
   (/\/ \ \   /\
   / /   | `    \
   O O   ) /    |
   `-^--'`<     '  System Version %s
  (_.)  _  )   /  Time %s
   `.___/`    /  Cpu: 2.49GHz Intel Pentium Xeon Processor, 1GB RAM
	 `-----' /   Mem: Bogomips Total
<----.     __ / __   \   
<----|====O)))==) \) /====
<----'    `--' `.__,' \
	 |        |
	  \       /       /\
 ______( (_  / \______/
,'  ,-----'   |
`--{__________)
	"
)

func displayLogo() {
	color.Yellow(logo)
	color.Blue(Version)
}
