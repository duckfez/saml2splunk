package commands

import (
	"github.com/duckfez/saml2splunk/v2/helper/credentials"
	"github.com/duckfez/saml2splunk/v2/helper/linuxkeyring"
)

func init() {
	if keyringHelper, err := linuxkeyring.NewKeyringHelper(); err == nil {
		credentials.CurrentHelper = keyringHelper
	}
}
