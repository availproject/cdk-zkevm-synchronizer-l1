package avail

import (
	"context"

	"github.com/availproject/cdk-avail-da-server/lib/avail"
	"github.com/ethereum/go-ethereum/common"
)

type Backend struct {
	*avail.AvailBackend
}

func New(l1RPCURL string, contractAddr common.Address, config avail.Config) (*Backend, error) {
	backend, err := avail.New(l1RPCURL, contractAddr, config)
	if err != nil {
		return nil, err
	}
	return &Backend{backend}, nil
}

// You can directly forward GetSequence if you want, since AvailBackend already has it
func (a *Backend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	return a.AvailBackend.GetSequence(ctx, batchHashes, dataAvailabilityMessage)
}
