package mcdn

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListViews(t *testing.T) {
	resp, err := DefaultInstance.ListViews()
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Countries)
}

func TestDescribeCdnRegionAndIsp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp()
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Countries)
	require.NotEmpty(t, resp.Result.Isps)
}
