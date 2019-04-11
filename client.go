package skrill

import (
	"context"
)

type Client interface {
	PrepareSale(ctx context.Context)
	PrepareTransfer(ctx context.Context)
	PrepareRefund(ctx context.Context)
	PrepareOnDemand(ctx context.Context)
	ExecuteTransfer(ctx context.Context)
	ExecuteRefund(ctx context.Context)
	ExecuteOnDemand(ctx context.Context)
	History(ctx context.Context)
}
