package source

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"slices"
)

type JobPostingSource interface {
	Site() string
	MaxPageSize() int
	List(page, size int) ([]*jobposting.JobPostingId, error) //가장 최신의 채용공고부터 정렬되도록 반환
	Detail(*jobposting.JobPostingId) (*jobposting.JobPostingDetail, error)
	Company(companyId string) (*company.CompanyDetail, error)
}

// src에서 모든 채용공고id를 가져온다.
// 가장 오래된 채용공고부터 정렬되도록 반환되어야 한다.
func AllJobPostingIds(src JobPostingSource) ([]*jobposting.JobPostingId, error) {
	maxPageSize := src.MaxPageSize()

	jobPostingIds := make([]*jobposting.JobPostingId, 0)

	page := 0
	alreadyListed := make(map[string]bool)
	for {
		page++
		ids, err := src.List(page, maxPageSize)
		if err != nil {
			return nil, err
		}

		if len(ids) == 0 {
			break
		}

		jobPostingIds = appendJobPostingIds(jobPostingIds, ids, alreadyListed)

		//다음 페이지를 호출 전에 채용공고 추가 또는 제거로 인해 누락된 채용공고가 있는지 확인한다.
		nextOffset := page * maxPageSize

		supplePageSize := GetMaxPrime(maxPageSize)
		supplePage := nextOffset / supplePageSize

		ids, err = src.List(supplePage, supplePageSize)
		if err != nil {
			return nil, err
		}
		jobPostingIds = appendJobPostingIds(jobPostingIds, ids, alreadyListed)
	}
	slices.Reverse(jobPostingIds)
	return jobPostingIds, nil
}

func GetMaxPrime(maxPageSize int) int {
	primeNumbers := []int{47, 43, 41, 37, 31, 29, 23, 19, 17, 13, 11, 7, 5, 3, 2}

	for _, prime := range primeNumbers {
		if maxPageSize > prime && maxPageSize%prime != 0 {
			return prime
		}
	}
	return 2
}

func appendJobPostingIds(jobPostingIds []*jobposting.JobPostingId, newIds []*jobposting.JobPostingId, alreadyListed map[string]bool) []*jobposting.JobPostingId {
	for _, id := range newIds {
		if _, ok := alreadyListed[id.PostingId]; !ok {
			alreadyListed[id.PostingId] = true
			jobPostingIds = append(jobPostingIds, id)
		}
	}

	return jobPostingIds
}
