package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/test/tinit"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {
	savedJpId := &jobposting.JobPostingId{Site: "jumpit", PostingId: "savedId"}
	savedJpId2 := &jobposting.JobPostingId{Site: "jumpit", PostingId: "savedId2"}
	savedJpId3 := &jobposting.JobPostingId{Site: "wanted", PostingId: "savedId3"}
	notExistedJpId := &jobposting.JobPostingId{Site: "notExistedSite", PostingId: "notExistedId"}
	notExistedJpId2 := &jobposting.JobPostingId{Site: "notExistedSite2", PostingId: "notExistedId2"}

	t.Run("SaveAndFind", func(t *testing.T) {
		jpRepo := tinit.InitJobPostingRepo(t)

		savedJp := jobposting.NewJobPosting(savedJpId.Site, savedJpId.PostingId)

		_, err := jpRepo.Save(savedJp)

		require.NoError(t, err)

		findedJp, err := jpRepo.Get(savedJpId)

		require.NoError(t, err)
		findedJp.CreatedAt = savedJp.CreatedAt //ignore createdAt
		require.Equal(t, savedJp, findedJp)

		ids, err := jpRepo.GetAllHiring(savedJpId.Site)
		require.NoError(t, err)
		require.Len(t, ids, 1)
		require.Equal(t, savedJpId.PostingId, ids[0].PostingId)
	})

	t.Run("FindNotExisted", func(t *testing.T) {
		jpRepo := tinit.InitJobPostingRepo(t)

		findedMatches, err := jpRepo.Get(notExistedJpId)

		require.NoError(t, err)
		require.Nil(t, findedMatches)

		ids, err := jpRepo.GetAllHiring(notExistedJpId.Site)
		require.NoError(t, err)
		require.Len(t, ids, 0)
	})

	t.Run("SaveAndFindAll", func(t *testing.T) {
		jpRepo := tinit.InitJobPostingRepo(t)

		savedJp := jobposting.NewJobPosting(savedJpId.Site, savedJpId.PostingId)
		savedJp2 := jobposting.NewJobPosting(savedJpId2.Site, savedJpId2.PostingId)
		savedJp3 := jobposting.NewJobPosting(savedJpId3.Site, savedJpId3.PostingId)
		savedJps := []*jobposting.JobPosting{savedJp, savedJp2, savedJp3}

		_, err := jpRepo.Save(savedJp)
		require.NoError(t, err)
		_, err = jpRepo.Save(savedJp2)
		require.NoError(t, err)
		_, err = jpRepo.Save(savedJp3)
		require.NoError(t, err)

		findedJps, err := jpRepo.Gets([]*jobposting.JobPostingId{savedJpId, notExistedJpId, savedJpId2, notExistedJpId2, savedJpId3})

		require.NoError(t, err)
		require.Len(t, findedJps, 3)

		for i, findedJp := range findedJps {
			savedJps[i].CreatedAt = findedJp.CreatedAt //ignore createdAt
		}

		require.True(t, isContain(findedJps, savedJp), "findedJps: %v, savedJps: %v", findedJps, savedJps)
		require.True(t, isContain(findedJps, savedJp2), "findedJps: %v, savedJps: %v", findedJps, savedJps)
		require.True(t, isContain(findedJps, savedJp3), "findedJps: %v, savedJps: %v", findedJps, savedJps)

		ids, err := jpRepo.GetAllHiring(savedJpId.Site)
		require.NoError(t, err)
		require.Len(t, ids, 2)
		require.True(t, isContainsId(ids, savedJpId))
		require.True(t, isContainsId(ids, savedJpId2))

		ids, err = jpRepo.GetAllHiring(savedJpId3.Site)
		require.NoError(t, err)
		require.Len(t, ids, 1)
		require.True(t, isContainsId(ids, savedJpId3))
	})
}

func isContain(jps []*jobposting.JobPosting, jp *jobposting.JobPosting) bool {
	for _, jp2 := range jps {
		if jp2.Site == jp.Site && jp2.PostingId == jp.PostingId {
			return true
		}
	}

	return false
}
