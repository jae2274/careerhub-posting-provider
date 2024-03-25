package source

import (
	"context"
	"slices"
	"testing"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/jumpit"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/wanted"

	"github.com/stretchr/testify/require"
	"gopkg.in/validator.v2"
)

func TestSource(t *testing.T) {
	callDelayMilis := int64(2000)

	t.Run("jumpit", func(t *testing.T) {
		ctx := context.Background()
		src := jumpit.NewJumpitSource(ctx, callDelayMilis)

		testSource(t, src)
	})

	t.Run("wanted", func(t *testing.T) {
		ctx := context.Background()
		src := wanted.NewWantedSource(ctx, callDelayMilis)

		testSource(t, src)
	})

	t.Run("GetMaxPrime", func(t *testing.T) {
		require.Equal(t, 2, source.GetMaxPrime(3))
		require.Equal(t, 3, source.GetMaxPrime(4))
		require.Equal(t, 3, source.GetMaxPrime(5))
		require.Equal(t, 5, source.GetMaxPrime(6))
		require.Equal(t, 7, source.GetMaxPrime(10))
		require.Equal(t, 13, source.GetMaxPrime(16))
		require.Equal(t, 17, source.GetMaxPrime(18))
		require.Equal(t, 19, source.GetMaxPrime(21))
		require.Equal(t, 23, source.GetMaxPrime(24))
	})
}

func testSource(t *testing.T, src source.JobPostingSource) {
	t.Run("list, detail, company", func(t *testing.T) {

		jobPostingIds, err := src.List(1, 10) //jumpit은 한 페이지당 최소 16개의 채용공고가 있음

		require.NoError(t, err)
		require.Len(t, jobPostingIds, 10)

		for _, jobPostingId := range jobPostingIds[0:2] {
			postingDetail, err := src.Detail(jobPostingId)
			require.NoError(t, err)

			require.NotNil(t, postingDetail)
			require.Equal(t, src.Site(), postingDetail.Site)
			require.NoError(t, validator.Validate(postingDetail))

			company, err := src.Company(postingDetail.CompanyId)

			require.NoError(t, err)
			require.NotNil(t, company)
			require.NoError(t, validator.Validate(company))
		}
	})

	t.Run("AllJobPostingIds", func(t *testing.T) {

		jobPostingIds, err := source.AllJobPostingIds(src)

		require.NoError(t, err)
		require.NotEmpty(t, jobPostingIds)

		page1, err := src.List(1, 10)
		require.NoError(t, err)

		last10 := jobPostingIds[len(jobPostingIds)-10:]
		slices.Reverse(last10)
		require.ElementsMatch(t, page1, last10) //가장 오래된 채용공고부터 정렬되어야 하므로, 마지막 10개는 첫 페이지와 같아야 한다.
	})
}
