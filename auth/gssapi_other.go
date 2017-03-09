//+build gssapi,!windows

package auth

import "fmt"

// GSSAPI is the mechanism name for GSSAPI.
const GSSAPI = "GSSAPI"

func newGSSAPIAuthenticator(cred *Cred) (Authenticator, error) {
	return nil, fmt.Errorf("GSSAPI support not enabled for non-windows builds")
}