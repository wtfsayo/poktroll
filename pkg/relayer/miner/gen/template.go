package main

import "text/template"

var (
	relayFixtureLineFmt   = "\t\t\"%x\","
	relayFixturesTemplate = template.Must(
		template.New("relay_fixtures_test.go").Parse(
			`// To regenerate all fixtures, use make go_fixturegen; to regenerate only this
// test's fixtures run go generate ./pkg/relayer/miner/miner_test.go.
package miner_test

var (
	// marshaledMinableRelaysHex are the hex encoded strings of serialized
	// relayer.MinedRelays which have been pre-mined to difficulty 2 by
	// populating the signature with random bytes. It is intended for use
	// in tests.
	marshaledMinableRelaysHex = []string{
{{.MarshaledMinableRelaysHex}}
	}

	// marshaledUnminableRelaysHex are the hex encoded strings of serialized
	// relayer.MinedRelays which have been pre-mined to **exclude** relays with
	// difficulty 2 (or greater). Like marshaledMinableRelaysHex, this is done
	// by populating the signature with random bytes. It is intended for use in
	// tests.
	marshaledUnminableRelaysHex = []string{
{{.MarshaledUnminableRelaysHex}}
	}
)
`,
		),
	)
)
