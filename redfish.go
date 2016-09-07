package redfish

import (
	// import go-swagger so it can generate the bindings
	_ "github.com/go-swagger/go-swagger"

	// import the generated go-swagger bindings
	_ "github.com/codedellemc/gorackhd-redfish/client"
	_ "github.com/codedellemc/gorackhd-redfish/models"
)
