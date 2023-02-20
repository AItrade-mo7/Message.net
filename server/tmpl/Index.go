package tmpl

import (
	_ "embed"
)

//go:embed email-Sys.html
var SysEmail string

//go:embed email-Code.html
var CodeEmail string

//go:embed email-RegisterSucceed.html
var RegisterSucceedEmail string
