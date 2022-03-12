package config

import (
	"net/http"

	"github.com/fortinetdev/forti-sdk-go/fortianalyzer/auth"
)

// Config provides configuration to a FAZ client instance
// It saves authentication information and a http connection
// for FAZ Client instance to create New connction to FAZ
// and Send data to FAZ,  etc. (needs to be extended later.)
type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}
