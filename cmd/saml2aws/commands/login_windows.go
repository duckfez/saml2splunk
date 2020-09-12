package commands

import (
	"github.com/duckfez/saml2splunk/v2/helper/credentials"
	"github.com/duckfez/saml2splunk/v2/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}
