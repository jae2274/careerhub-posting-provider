package appfunc

import (
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
)

type SeparatedIds struct {
	TotalCount      int
	NewPostingIds   []*jobposting.JobPostingId
	ClosePostingIds []*jobposting.JobPostingId
}

func SeparateIds(savedJpIds []*jobposting.JobPostingId, hiringJpIds []*jobposting.JobPostingId) *SeparatedIds {

	savedJpMap := make(map[jobposting.JobPostingId]interface{})
	for _, jp := range savedJpIds {
		savedJpMap[*jp] = false
	}

	newJpIds := make([]*jobposting.JobPostingId, 0)
	for _, srcJpId := range hiringJpIds {
		jpId := jobposting.JobPostingId{Site: srcJpId.Site, PostingId: srcJpId.PostingId}
		if _, ok := savedJpMap[jpId]; ok {
			delete(savedJpMap, jpId)
		} else {
			newJpIds = append(newJpIds, srcJpId)
		}
	}

	closeJpIds := make([]*jobposting.JobPostingId, 0)
	for savedJpId := range savedJpMap {
		func(jpId jobposting.JobPostingId) {
			closeJpIds = append(closeJpIds, &jpId)
		}(savedJpId)
	}

	return &SeparatedIds{
		TotalCount:      len(hiringJpIds),
		NewPostingIds:   newJpIds,
		ClosePostingIds: closeJpIds,
	}
}
