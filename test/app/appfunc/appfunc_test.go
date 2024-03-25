package appfunc

import (
	"testing"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/app/appfunc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"

	"github.com/stretchr/testify/require"
)

func TestSeparateIds(t *testing.T) {
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

		separateIds := appfunc.SeparateIds(make([]*jobposting.JobPostingId, 0), convertJpIdsToSourceJpIds(notSavedJumpitJps))

		require.Len(t, separateIds.NewPostingIds, 3)
		require.Len(t, separateIds.ClosePostingIds, 0)
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[0]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[1]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[2]))
	})

	t.Run("Hiring JobPostings are all saved", func(t *testing.T) {

		separateIds := appfunc.SeparateIds(savedWantedJps, convertJpIdsToSourceJpIds(savedWantedJps))

		require.Len(t, separateIds.NewPostingIds, 0)
		require.Len(t, separateIds.ClosePostingIds, 0)
	})

	t.Run("All saved JobPostings are closed and Hiring JobPostings are all new", func(t *testing.T) {

		separateIds := appfunc.SeparateIds(savedWantedJps, convertJpIdsToSourceJpIds(notSavedJumpitJps))

		require.Len(t, separateIds.NewPostingIds, 3)
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[0]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[1]))
		require.True(t, isContainsId(separateIds.NewPostingIds, notSavedJumpitJps[2]))

		require.Len(t, separateIds.ClosePostingIds, 3)
		require.True(t, isContainsId(convertJpIdsToSourceJpIds(separateIds.ClosePostingIds), savedWantedJps[0]))
		require.True(t, isContainsId(convertJpIdsToSourceJpIds(separateIds.ClosePostingIds), savedWantedJps[1]))
		require.True(t, isContainsId(convertJpIdsToSourceJpIds(separateIds.ClosePostingIds), savedWantedJps[2]))
	})
}

func convertJpIdsToSourceJpIds(jpIds []*jobposting.JobPostingId) []*jobposting.JobPostingId {
	srcJpIds := make([]*jobposting.JobPostingId, 0)
	for _, jpId := range jpIds {
		srcJpIds = append(srcJpIds, &jobposting.JobPostingId{Site: jpId.Site, PostingId: jpId.PostingId})
	}
	return srcJpIds
}

func isContainsId(ids []*jobposting.JobPostingId, idToFind *jobposting.JobPostingId) bool {
	for _, item := range ids {
		if item.Site == idToFind.Site && item.PostingId == idToFind.PostingId {
			return true
		}
	}
	return false
}
