package commands

import (
	"github.com/duckfez/saml2splunk/v2/helper/credentials"
	"github.com/duckfez/saml2splunk/v2/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}
