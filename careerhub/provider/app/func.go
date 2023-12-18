package app

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/source"
)

type SeparatedIds struct {
	NewPostingIds   []*source.JobPostingId
	ClosePostingIds []*jobposting.JobPostingId
}

func SeparateIds(savedJpIds []*jobposting.JobPostingId, hiringJpIds []*source.JobPostingId) *SeparatedIds {

	savedJpMap := make(map[jobposting.JobPostingId]interface{})
	for _, jp := range savedJpIds {
		savedJpMap[*jp] = false
	}

	newJpIds := make([]*source.JobPostingId, 0)
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
		NewPostingIds:   newJpIds,
		ClosePostingIds: closeJpIds,
	}
}
