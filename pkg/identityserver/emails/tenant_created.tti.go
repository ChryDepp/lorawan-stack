// Copyright © 2019 The Things Industries B.V.

package emails

import (
	"strings"

	"go.thethings.network/lorawan-stack/pkg/version"
)

// TenantCreated is the email that is sent when a new tenant has been created.
type TenantCreated struct {
	Data
	GlobalNetworkName string

	InitialPassword string
}

func (TenantCreated) TTESVersion() string {
	return strings.TrimSuffix(version.TTN, "-dev")
}

// TemplateName returns the name of the template to use for this email.
func (TenantCreated) TemplateName() string { return "tenant_created" }

const tenantCreatedSubject = `Welcome to {{.GlobalNetworkName}}`

const tenantCreatedText = `Dear {{.User.Name}},

Your tenant "{{.Entity.ID}}" on {{.GlobalNetworkName}} is ready.

The username of your admin user is "{{.User.ID}}".

{{- if .InitialPassword}}

Before you get started, please reset your password by following the link below:

{{.Network.IdentityServerURL}}/update-password?user={{.User.ID}}&current={{.InitialPassword}}

{{- end }}

To access the console, you can use the URL {{.Network.ConsoleURL}}.

You can read how to get started with the console in the Getting Started guide:
https://enterprise.thethingsstack.io/{{.TTESVersion}}/guides/getting-started/console/

It is also possible to use the command-line interface (CLI).
You can read how to get started with the CLI in the Getting Started guide:
https://enterprise.thethingsstack.io/{{.TTESVersion}}/guides/getting-started/cli/

Other users of your tenant can register accounts, but they can not use it until
the account is approved by an admin user. Approving accounts is currently only
possible with the CLI. Once logged in with the CLI, you can approve other users
using the following command:

$ ttn-lw-cli users update [user-id] --state APPROVED

If at any point you have questions, you can reach out to us using the support
button that is shown in the console.
`

// DefaultTemplates returns the default templates for this email.
func (TenantCreated) DefaultTemplates() (subject, html, text string) {
	return tenantCreatedSubject, "", tenantCreatedText
}
