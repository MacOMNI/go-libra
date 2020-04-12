package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/the729/go-libra/generated/pbtypes"
	"github.com/the729/go-libra/types"
)

// QueryEventsByAccessPath queries list of events by access path does necessary crypto verifications.
func (c *Client) QueryEventsByAccessPath(ctx context.Context, ap *types.AccessPath, start uint64, ascending bool, limit uint64) ([]*types.ProvenEvent, error) {
	c.accMu.RLock()
	frozenSubtreeRoots := cloneSubtrees(c.acc.FrozenSubtreeRoots)
	numLeaves := c.acc.NumLeaves
	c.accMu.RUnlock()

	resp, err := c.ac.UpdateToLatestLedger(ctx, &pbtypes.UpdateToLatestLedgerRequest{
		ClientKnownVersion: numLeaves - 1,
		RequestedItems: []*pbtypes.RequestItem{
			{
				RequestedItems: &pbtypes.RequestItem_GetEventsByEventAccessPathRequest{
					GetEventsByEventAccessPathRequest: &pbtypes.GetEventsByEventAccessPathRequest{
						AccessPath: &pbtypes.AccessPath{
							Address: ap.Address[:],
							Path:    ap.Path,
						},
						StartEventSeqNum: start,
						Ascending:        ascending,
						Limit:            limit,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	pli, err := c.verifyLedgerInfoAndConsistency(resp, numLeaves, frozenSubtreeRoots)
	if err != nil {
		return nil, err
	}

	resp1 := resp.ResponseItems[0].GetGetEventsByEventAccessPathResponse()
	if resp1 == nil {
		return nil, errors.New("nil response")
	}
	// b, err := json.MarshalIndent(resp1, "", "    ")
	// log.Printf("resp1: %s", string(b))

	if len(resp1.EventsWithProof) < int(limit) {
		//TODO: verify resp1.ProofOfLatestEvent
	}
	pevs := make([]*types.ProvenEvent, 0, len(resp1.EventsWithProof))
	for _, pbev := range resp1.EventsWithProof {
		ev := &types.EventWithProof{}
		if err := ev.FromProto(pbev); err != nil {
			return nil, fmt.Errorf("event from protobuf error: %v", err)
		}
		pev, err := ev.Verify(pli)
		if err != nil {
			return nil, fmt.Errorf("event verification error: %v", err)
		}
		pevs = append(pevs, pev)
	}

	return pevs, nil
}
