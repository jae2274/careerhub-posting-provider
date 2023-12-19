package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/validator.v2"
)

func TestJumpitSource(t *testing.T) {
	callDelayMilis := int64(2000)

	t.Run("list, detail, company", func(t *testing.T) {
		source := jumpit.NewJumpitSource(callDelayMilis)
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
	})

	t.Run("AllJobPostingIds", func(t *testing.T) {
		src := jumpit.NewJumpitSource(callDelayMilis)
		src.Run(make(<-chan app.QuitSignal))

		jobPostingIds, err := source.AllJobPostingIds(src)

		require.NoError(t, err)
		require.NotEmpty(t, jobPostingIds)
		t.Log(len(jobPostingIds))

		for _, jpId := range jobPostingIds {
			_, ok := jpId.EtcInfo["jobCategory"]
			require.True(t, ok)
		}

		page1, err := src.List(1, 10)
		require.NoError(t, err)

		last10 := jobPostingIds[len(jobPostingIds)-10:]
		slices.Reverse(last10)
		require.ElementsMatch(t, page1, last10) //가장 오래된 채용공고부터 정렬되어야 하므로, 마지막 10개는 첫 페이지와 같아야 한다.
	})
}
