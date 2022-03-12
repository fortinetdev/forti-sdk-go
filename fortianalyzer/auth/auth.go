package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for FortiAnalyzer
type Auth struct {
	Hostname    string
	User        string
	Passwd      string
	CABundle    string
	Insecure    *bool
	Refresh     bool

	Session     string
}

// NewAuth inits Auth object with the given metadata
func NewAuth(hostname, user, passwd, cabundle, session string) *Auth {
	return &Auth{
		Hostname:   hostname,
		User:       user,
		Passwd:     passwd,
		CABundle:   cabundle,

		Session:    session,
	}
}

// GetEnvHostname gets FortiAnalyzer hostname from OS environment
// It returns the hostname
func (m *Auth) GetEnvHostname() (string, error) {
	h := os.Getenv("FORTIANALYZER_ACCESS_HOSTNAME")

	if h == "" {
		return h, fmt.Errorf("GetEnvHostname error")
	}

	m.Hostname = h

	return h, nil
}

// GetEnvUsername gets FortiAnalyzer hostname from OS environment
// It returns the username
func (m *Auth) GetEnvUsername() (string, error) {
	h := os.Getenv("FORTIANALYZER_ACCESS_USERNAME")

	if h == "" {
		return h, fmt.Errorf("GetEnvUsername error")
	}

	m.User = h

	return h, nil
}

// GetEnvPassword gets FortiAnalyzer hostname from OS environment
// It returns the password
func (m *Auth) GetEnvPassword() (string, error) {
	h := os.Getenv("FORTIANALYZER_ACCESS_PASSWORD")

	if h == "" {
		return h, fmt.Errorf("GetEnvPassword error")
	}

	m.Passwd = h

	return h, nil
}

// GetEnvCABundle gets CA Bundle file from OS environment
// It returns the CA Bundle file
func (m *Auth) GetEnvCABundle() (string, error) {
	c := os.Getenv("FORTIANALYZER_CA_CABUNDLE")

	if c == "" {
		return c, nil
	}

	m.CABundle = c

	return c, nil
}

// GetEnvInsecure gets Insecure value from OS environment
// It returns the insecure value
func (m *Auth) GetEnvInsecure() (bool, error) {
	c := os.Getenv("FORTIANALYZER_INSECURE")

	if c == "true" {
		return true, nil
	}

	return false, nil
}
