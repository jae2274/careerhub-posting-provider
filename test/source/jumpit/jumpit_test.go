package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/validator.v2"
)

func TestJumpitSource(t *testing.T) {
	source := jumpit.NewJumpitSource(1000)
	source.Run(make(<-chan app.QuitSignal))

	jobPostingIds, err := source.List(1, 10) //jumpit은 한 페이지당 최소 16개의 채용공고가 있음

	require.NoError(t, err)
	require.Len(t, jobPostingIds, 10)

	for _, jobPostingId := range jobPostingIds[0:2] {
		postingDetail, err := source.Detail(jobPostingId)
		require.NoError(t, err)

		require.NotNil(t, postingDetail)
		require.Equal(t, "jumpit", postingDetail.Site)
		require.NoError(t, validator.Validate(postingDetail))

		company, err := source.Company(postingDetail.CompanyId)

		require.NoError(t, err)
		require.NotNil(t, company)
		require.NoError(t, validator.Validate(company))
	}
}
