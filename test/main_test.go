package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReset(t *testing.T) {
	// tinit.InitCompanyRepo(t)
	// tinit.InitJobPostingRepo(t)

	// tinit.InitGrpcClient(t)
	max, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.000+00:00")
	require.NoError(t, err)

	t.Log(max.UnixMilli())
}
