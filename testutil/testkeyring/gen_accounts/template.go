//go:build ignore

package main

import "text/template"

var (
	preGeneratedAccountLineFmt = "\t\tmustParsePreGeneratedAccount(%q),"
	accountsTableTemplate      = template.Must(
		template.New("accounts_table.go").Parse(
			`// DO NOT EDIT. This Code is generated by gen_accounts/gen.go,
// changes will be overwritten upon regeneration.
//
// To regenerate this file, use make go_testgen_accounts or go generate ./testutil/testkeyring/keyring.go.

package testkeyring

import "github.com/pokt-network/poktroll/cmd/poktrolld/cmd"

var preGeneratedAccounts *PreGeneratedAccountIterator

func init() {
	cmd.InitSDKConfig()

	preGeneratedAccounts = NewPreGeneratedAccountIterator(
{{.newPreGeneratedAccountIteratorArgLines}}
	)
}
`,
		),
	)
)
