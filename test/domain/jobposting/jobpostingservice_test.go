package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/test/tinit"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobPostingService(t *testing.T) {
	jumpitSite := "jumpit"
	wantedSite := "wanted"
	notSavedJumpitJps := []*jobposting.JobPostingId{
		{Site: jumpitSite, PostingId: "posting1"},
		{Site: jumpitSite, PostingId: "posting2"},
		{Site: jumpitSite, PostingId: "posting3"},
	}

	savedWantedJps := []*jobposting.JobPostingId{
		{Site: wantedSite, PostingId: "posting1"},
		{Site: wantedSite, PostingId: "posting2"},
		{Site: wantedSite, PostingId: "posting3"},
	}

	t.Run("Hiring JobPostings are all new", func(t *testing.T) {
		service := initService(t)

		separateIds, err := service.SeparateIds(jumpitSite, notSavedJumpitJps)
		require.NoError(t, err)
		require.Len(t, separateIds.NewPostingIds, 3)
		require.Len(t, separateIds.ClosePostingIds, 0)
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[0]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[1]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[2]))
	})

	t.Run("Hiring JobPostings are all saved", func(t *testing.T) {
		service := initService(t)

		for _, jpId := range savedWantedJps {
			err := service.Save(jpId)
			require.NoError(t, err)
		}

		separateIds, err := service.SeparateIds(wantedSite, savedWantedJps)
		require.NoError(t, err)
		require.Len(t, separateIds.NewPostingIds, 0)
		require.Len(t, separateIds.ClosePostingIds, 0)
	})

	t.Run("All saved JobPostings are closed", func(t *testing.T) {
		service := initService(t)

		for _, jpId := range savedWantedJps {
			err := service.Save(jpId)
			require.NoError(t, err)
		}

		separateIds, err := service.SeparateIds(wantedSite, notSavedJumpitJps)

		require.NoError(t, err)
		require.Len(t, separateIds.NewPostingIds, 3)
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[0]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[1]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[2]))

		require.Len(t, separateIds.ClosePostingIds, 3)
		require.True(t, isContainsId(separateIds.ClosePostingIds, savedWantedJps[0]))
		require.True(t, isContainsId(separateIds.ClosePostingIds, savedWantedJps[1]))
		require.True(t, isContainsId(separateIds.ClosePostingIds, savedWantedJps[2]))
	})
}

func initService(t *testing.T) *jobposting.JobPostingService {
	repo := tinit.InitJobPostingRepo(t)
	return jobposting.NewJobPostingService(repo)
}

func isContainsId(ids []*jobposting.JobPostingId, idToFind *jobposting.JobPostingId) bool {
	for _, item := range ids {
		if item.Site == idToFind.Site && item.PostingId == idToFind.PostingId {
			return true
		}
	}
	return false
}
