// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{
	{"DsaAKsMvZ6HrqhmbhLjV9qVbPkkzF5daowT", 5000 * 1e8},
}

// BlockOneLedgerTestNet2 is the block one output ledger for the 2nd test
// network.
var BlockOneLedgerTestNet2 = []*TokenPayout{
	{"TskYenJuHvH5ZLdtSEszWA33BTXYMnmu3ck", 100000 * 1e8},
	{"TsbmjFmnsv9G5gRvLXwFVnXm3uCQib1cRy4", 100000 * 1e8},
	{"Tsfg1wVYAc1b1U7yJVWvqXcxw7wVzb1nJ1F", 100000 * 1e8},
	{"TsSuQZMYVWF9d79R8qTJ2qdp596rshMHyVd", 100000 * 1e8},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network. See under "Decred organization related parameters" in params.go
// for information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"Ssqpw9GquM9gVjLrsbBP7KnNWq4pc1v95Gq", 100000 * 1e8},
}
