package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/source/wanted"
	"fmt"

	// "careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

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

		jobRepo, ClosedQueue, findNewJobPostingApp := initFindNewComponents(t, src)

		require.NoError(t, err)

		savedJpIds := allJpId[:20]
		closedJpIds := []jobposting.JobPostingId{
			{Site: src.Site(), PostingId: "closed_1"},
			{Site: src.Site(), PostingId: "closed_2"},
			{Site: src.Site(), PostingId: "closed_3"},
		}

		for _, jpId := range savedJpIds {
			_, err = jobRepo.Save(jobposting.NewJobPosting(jpId.Site, jpId.PostingId))
			require.NoError(t, err)
		}
		for _, jpId := range closedJpIds {
			_, err = jobRepo.Save(jobposting.NewJobPosting(jpId.Site, jpId.PostingId))
			require.NoError(t, err)
		}

		newJpIds, err := findNewJobPostingApp.Run(ctx)
		require.NoError(t, err)

		require.Equal(t, allJpId[20:], newJpIds)

		msgs := getClosedMessages(t, ClosedQueue)
		require.Len(t, msgs, 1)
		require.Len(t, msgs[0].JobPostingIds, len(closedJpIds))
		IsEqualClosedJobPostingIds(t, closedJpIds, msgs)

		allSavedJp, err := jobRepo.GetAllHiring(src.Site())
		require.NoError(t, err)

		IsEqualSavedJobPostings(t, savedJpIds, allSavedJp) //savedJpIds와 closedJpIds 둘 다 DB에 저장되었으나 findNewJobPostingApp.Run() 실행 후 closedJpIds는 DB에서 삭제되었음
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

func initFindNewComponents(t *testing.T, src source.JobPostingSource) (*jobposting.JobPostingRepo, tinit.MockGrpcClient, *app.FindNewJobPostingApp) {

	jobRepo := tinit.InitJobPostingRepo(t)
	grpcClient := tinit.InitGrpcClient(t)

	return jobRepo, grpcClient, app.NewFindNewJobPostingApp(src, jobRepo, grpcClient)
}

func getClosedMessages(t *testing.T, grpcClient tinit.MockGrpcClient) []*provider_grpc.JobPostings {

	return grpcClient.GetClosedJpIds()
}

func IsEqualClosedJobPostingIds(t *testing.T, closedJpIds []jobposting.JobPostingId, closedMessages []*provider_grpc.JobPostings) {
	require.Len(t, closedMessages, 1)
	require.Len(t, closedMessages[0].JobPostingIds, len(closedJpIds))
Outer:
	for _, closedMessage := range closedMessages[0].JobPostingIds {
		for _, closedJpId := range closedJpIds {
			if closedMessage.Site == closedJpId.Site && closedMessage.PostingId == closedJpId.PostingId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", closedMessage.Site, closedMessage.PostingId)
		t.FailNow()
	}
}

func IsEqualSavedJobPostings(t *testing.T, srcJpIds []*jobposting.JobPostingId, savedJps []*jobposting.JobPostingId) {
	require.Len(t, savedJps, len(srcJpIds))

Outer:
	for _, savedJp := range savedJps {
		for _, srcJpId := range srcJpIds {
			if savedJp.Site == srcJpId.Site && savedJp.PostingId == srcJpId.PostingId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", savedJp.Site, savedJp.PostingId)
		t.FailNow()
	}
}
