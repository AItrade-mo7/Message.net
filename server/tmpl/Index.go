package tmpl

import (
	_ "embed"
)

//go:embed email-sys.html
var SysEmail string

//go:embed email-code.html
var CodeEmail string

//go:embed email-register.html
var RegisterSucceedEmail string
