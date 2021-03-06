package validation

import (
	"testing"

	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/legacy"
)

func dummyValidateTx(*bc.Tx) error {
	return nil
}

func generate(tb testing.TB, prev *bc.Block) *bc.Block {
	b := &legacy.Block{
		BlockHeader: legacy.BlockHeader{
			Version:           1,
			Height:            prev.Height + 1,
			PreviousBlockHash: prev.ID,
			TimestampMS:       prev.TimestampMs + 1,
			BlockCommitment:   legacy.BlockCommitment{},
		},
	}

	var err error
	b.TransactionsMerkleRoot, err = bc.MerkleRoot(nil)
	if err != nil {
		tb.Fatal(err)
	}

	return legacy.MapBlock(b)
}
