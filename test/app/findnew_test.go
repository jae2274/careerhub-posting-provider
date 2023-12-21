package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/vars"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFindNew(t *testing.T) {

	quitChan := make(chan app.QuitSignal)
	src := jumpit.NewJumpitSource(3000, quitChan)
	jobRepo, ClosedQueue, findNewJobPostingApp := initFindNewComponents(t, src)

	allJpId, err := source.AllJobPostingIds(src)
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

	newJpIds, err := findNewJobPostingApp.Run()
	require.NoError(t, err)

	require.Equal(t, allJpId[20:], newJpIds)

	msgs := getClosedMessages(t, ClosedQueue)
	require.Len(t, msgs, 1)
	require.Len(t, msgs[0].JobPostingIds, len(closedJpIds))
	IsEqualClosedJobPostingIds(t, closedJpIds, msgs)

	allSavedJp, err := dynamo.GetAll(jobRepo, context.TODO())
	require.NoError(t, err)

	IsEqualSavedJobPostings(t, savedJpIds, allSavedJp) //savedJpIds와 closedJpIds 둘 다 DB에 저장되었으나 findNewJobPostingApp.Run() 실행 후 closedJpIds는 DB에서 삭제되었음
}

func initFindNewComponents(t *testing.T, src source.JobPostingSource) (*jobposting.JobPostingRepo, queue.Queue, *app.FindNewJobPostingApp) {
	envVars, err := vars.Variables()
	require.NoError(t, err)

	jobRepo := tinit.InitJobPostingRepo(t)
	ClosedQueue := tinit.InitSQS(t, envVars.ClosedQueue)

	return jobRepo, ClosedQueue, app.NewFindNewJobPostingApp(src, jobRepo, queue.NewClosedJobPostingQueue(ClosedQueue))
}

func getClosedMessages(t *testing.T, closedQ queue.Queue) []message_v1.ClosedJobPostings {
	result := make([]message_v1.ClosedJobPostings, 0)
	for {
		messages, err := closedQ.Recv()
		require.NoError(t, err)

		if len(messages) == 0 {
			break
		}

		closedMessages := make([]message_v1.ClosedJobPostings, len(messages))
		for i, message := range messages {
			err := proto.Unmarshal(message, &closedMessages[i])
			require.NoError(t, err)
		}

		result = append(result, closedMessages...)
	}
	return result
}

func IsEqualClosedJobPostingIds(t *testing.T, closedJpIds []jobposting.JobPostingId, closedMessages []message_v1.ClosedJobPostings) {
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

func IsEqualSavedJobPostings(t *testing.T, srcJpIds []*source.JobPostingId, savedJps []*jobposting.JobPosting) {
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
