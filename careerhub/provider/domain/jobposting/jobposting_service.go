package jobposting

type JobPostingService struct {
	repo *JobPostingRepo
}

func NewJobPostingService(repo *JobPostingRepo) *JobPostingService {
	return &JobPostingService{
		repo: repo,
	}
}

type SeparateIds struct {
	NewPostingIds   []*JobPostingId
	ClosePostingIds []*JobPostingId
}

func (s *JobPostingService) Save(jpId *JobPostingId) error {
	jp := NewJobPosting(jpId.Site, jpId.PostingId)
	_, err := s.repo.Save(jp)
	return err
}

func (s *JobPostingService) SeparateIds(site string, hiringJpIds []*JobPostingId) (*SeparateIds, error) {
	savedJps, err := s.repo.GetAllHiring(site)
	if err != nil {
		return nil, err
	}

	savedJpMap := make(map[JobPostingId]bool)
	for _, jp := range savedJps {
		savedJpMap[JobPostingId{Site: jp.Site, PostingId: jp.PostingId}] = false
	}

	newJpIds := make([]*JobPostingId, 0)
	for _, jpId := range hiringJpIds {
		if _, ok := savedJpMap[*jpId]; ok {
			delete(savedJpMap, *jpId)
		} else {
			newJpIds = append(newJpIds, jpId)
		}
	}

	closeJpIds := make([]*JobPostingId, 0)
	for savedJpId := range savedJpMap {
		func(jpId JobPostingId) {
			closeJpIds = append(closeJpIds, &jpId)
		}(savedJpId)
	}

	return &SeparateIds{
		NewPostingIds:   newJpIds,
		ClosePostingIds: closeJpIds,
	}, nil
}
