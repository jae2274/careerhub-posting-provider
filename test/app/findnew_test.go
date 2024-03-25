package app

import (
	"fmt"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/app"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/jumpit"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/wanted"

	// "github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/queue"
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"
	"github.com/jae2274/careerhub-posting-provider/test/tinit"

	"github.com/stretchr/testify/require"
)

type MockSource struct {
	originSrc source.JobPostingSource
	pageCache map[string][]*jobposting.JobPostingId
}

func NewMockSource(originSrc source.JobPostingSource) *MockSource {
	return &MockSource{
		originSrc: originSrc,
		pageCache: make(map[string][]*jobposting.JobPostingId),
	}
}

func pageKey(page, size int) string { return fmt.Sprintf("%d_%d", page, size) }

func (s *MockSource) Site() string     { return s.originSrc.Site() }
func (s *MockSource) MaxPageSize() int { return s.originSrc.MaxPageSize() }
func (s *MockSource) List(page, size int) ([]*jobposting.JobPostingId, error) { //호출할 때마다 같은 결과가 보장되지 않아 테스트 시엔 cache를 사용하여 같은 결과를 보장
	if list, ok := s.pageCache[pageKey(page, size)]; ok {
		return list, nil
	}

	list, err := s.originSrc.List(page, size)
	s.pageCache[pageKey(page, size)] = list
	return list, err
}
func (s *MockSource) Detail(jpId *jobposting.JobPostingId) (*jobposting.JobPostingDetail, error) {
	return s.originSrc.Detail(jpId)
}
func (s *MockSource) Company(companyId string) (*company.CompanyDetail, error) {
	return s.originSrc.Company(companyId)
}

func TestFindNew(t *testing.T) {
	testFindNew := func(t *testing.T, ctx context.Context, originSrc source.JobPostingSource) {
		src := NewMockSource(originSrc)

		allJpId, err := source.AllJobPostingIds(src)
		require.NoError(t, err)

		grpcService, findNewJobPostingApp := initFindNewComponents(t, src)
		for _, jpId := range allJpId[:3] {
			detail, err := src.Detail(jpId)
			require.NoError(t, err)
			err = grpcService.RegisterJobPostingInfo(ctx, detail)
			require.NoError(t, err)
		}

		closedJobPostingIds := []*jobposting.JobPostingId{
			{Site: "jumpit", PostingId: "closed_1"},
			{Site: "jumpit", PostingId: "closed_2"},
			{Site: "jumpit", PostingId: "closed_3"},
		}

		for _, jpId := range closedJobPostingIds {
			err = grpcService.RegisterJobPostingInfo(ctx, dummyJobPosting(jpId))
			require.NoError(t, err)
		}

		newJpIds, err := findNewJobPostingApp.Run(ctx)
		require.NoError(t, err)
		require.Equal(t, len(allJpId), newJpIds.TotalCount)
		require.Equal(t, allJpId[3:], newJpIds.NewPostingIds)
		require.Equal(t, closedJobPostingIds, newJpIds.ClosePostingIds)
	}

	t.Run("jumpit", func(t *testing.T) {
		ctx := context.Background()
		src := jumpit.NewJumpitSource(ctx, 3000)
		testFindNew(t, ctx, src)
	})

	t.Run("wanted", func(t *testing.T) {
		ctx := context.Background()
		src := wanted.NewWantedSource(ctx, 3000)
		testFindNew(t, ctx, src)
	})
}

func initFindNewComponents(t *testing.T, src source.JobPostingSource) (provider_grpc.ProviderGrpcService, *app.FindNewJobPostingApp) {

	grpcClient := tinit.InitGrpcClient(t)
	grpcService := provider_grpc.NewProviderGrpcService(grpcClient)

	return grpcService, app.NewFindNewJobPostingApp(src, grpcService)
}

func dummyJobPosting(jpId *jobposting.JobPostingId) *jobposting.JobPostingDetail {
	return &jobposting.JobPostingDetail{
		Site:        jpId.Site,
		PostingId:   jpId.PostingId,
		CompanyId:   "dummy",
		CompanyName: "dummy",
		JobCategory: []string{"dummy"},
		MainContent: jobposting.MainContent{
			PostUrl:        "dummy",
			Title:          "dummy",
			Intro:          "dummy",
			MainTask:       "dummy",
			Qualifications: "dummy",
			Preferred:      "dummy",
			Benefits:       "dummy",
		},
		RequiredSkill: []string{"dummy"},
		Tags:          []string{"dummy"},
		RequiredCareer: jobposting.Career{
			Min: new(int32),
			Max: new(int32),
		},
		PublishedAt:   new(int64),
		ClosedAt:      new(int64),
		Address:       []string{"dummy"},
		ImageUrl:      new(string),
		CompanyImages: []string{"dummy"},
	}
}
