package mysqlstore

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelayReport(t *testing.T) {
	ctx := context.Background()
	res, err := vendorTest.GetVendorsTotalDelay(ctx)

	_ = res
	require.NoError(t, err)
}
