package wanted

import (
	"careerhub-dataprovider/careerhub/provider/source/wanted"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	categoryMap, err := wanted.GetJobCategoryMap()
	require.NoError(t, err)

	require.NotEmpty(t, categoryMap)
}
