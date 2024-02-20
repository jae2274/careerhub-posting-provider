package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/test/tinit"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		findedJp.ID = primitive.NilObjectID
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

	t.Run("Delete Chunk 25 size", func(t *testing.T) { //25개 이상의 데이터를 삭제시엔 25개씩 삭제해야함
		jpRepo := tinit.InitJobPostingRepo(t)

		savedSize := 60
		savedJps := make([]*jobposting.JobPosting, savedSize)
		savedJpIds := make([]*jobposting.JobPostingId, savedSize)
		for i := 0; i < savedSize; i++ {
			savedJpIds[i] = &jobposting.JobPostingId{Site: "jumpit", PostingId: fmt.Sprintf("savedId%d", i)}
			savedJps[i] = jobposting.NewJobPosting(savedJpIds[i].Site, savedJpIds[i].PostingId)
		}

		for i := 0; i < savedSize; i++ {
			_, err := jpRepo.Save(savedJps[i])
			require.NoError(t, err)
		}

		ids, err := jpRepo.GetAllHiring(savedJpIds[0].Site)
		require.NoError(t, err)
		require.Len(t, ids, savedSize)

		err = jpRepo.DeleteAll(savedJpIds[0:55]) // 25 + 25 + 5로, 3번에 걸쳐 삭제
		require.NoError(t, err)

		ids, err = jpRepo.GetAllHiring(savedJpIds[0].Site)
		require.NoError(t, err)
		require.Len(t, ids, 5)
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

func isContainsId(ids []*jobposting.JobPostingId, idToFind *jobposting.JobPostingId) bool {
	for _, item := range ids {
		if item.Site == idToFind.Site && item.PostingId == idToFind.PostingId {
			return true
		}
	}
	return false
}
