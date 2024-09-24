package source

import (
	"errors"
	"fmt"
	"slices"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/jjson"
	"github.com/jae2274/goutils/terr"
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

func CallApi[RESULT any](aActor *apiactor.ApiActor, request *apiactor.Request) (*RESULT, error) {
	request.Header.Add("X-Crawler-Message", "Please_allow_crawling")
	request.Header.Add("X-Crawler-Start-Time", "04_09_00")
	request.Header.Add("X-Crawler-Target", "Job_postings_and_company_information")
	request.Header.Add("X-Crawler-Purpose", "Introduction_and_linking_to_job_postings_to_external_sites")
	request.Header.Add("X-Crawler-Benefit", "Increased_job_applications_through_external_links")

	rc, err := aActor.Call(request)
	if err != nil {
		return nil, terr.Wrap(errors.Join(err, fmt.Errorf("\turl: %s", request.Url)))
	}

	result, err := jjson.UnmarshalReader[RESULT](rc)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return result, nil
}
