package app

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/app"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/jumpit"
	"github.com/jae2274/careerhub-posting-provider/test/tinit"

	"github.com/jae2274/goutils/cchan"
	"github.com/stretchr/testify/require"
)

func TestSendJobPostingApp(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		ctx := context.Background()
		src := jumpit.NewJumpitSource(ctx, 3000)

		jobPostingClient, reviewClient := tinit.InitGrpcClient(t)
		grpcService := provider_grpc.NewProviderGrpcService(jobPostingClient, reviewClient)
		sendJobApp := app.NewSendJobPostingApp(src, grpcService)

		jpIds, err := src.List(1, 3)
		require.NoError(t, err)

		processedChan, errChan := sendJobApp.Run(ctx, jpIds)

		require.NoError(t, err)

		results := cchan.WaitClosed(processedChan)
		require.Len(t, results, len(jpIds))
		if len(errChan) > 0 {
			for {
				select {
				case err := <-errChan:
					t.Log(err)
				default:
					t.FailNow()
				}
			}
		}
		savedIds, err := grpcService.GetAllHiring(ctx, src.Site())
		require.NoError(t, err)
		require.Equal(t, jpIds, savedIds)
		// IsEqualSavedJobPostingIds(t, jpIds, savedIds)

		for _, jpId := range jpIds {
			detail, err := src.Detail(jpId)
			require.NoError(t, err)
			require.Equal(t, detail, jobPostingClient.GetJobPosting(jpId))

			companyId := &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			}

			isRegistered, err := grpcService.IsCompanyRegistered(ctx, companyId)
			require.NoError(t, err)
			require.True(t, isRegistered)

			companyDetail, err := src.Company(companyId.CompanyId)
			require.NoError(t, err)
			require.Equal(t, companyDetail, jobPostingClient.GetCompany(companyId))
			require.NotNil(t, reviewClient.GetCrawlingTask(detail.CompanyName))

		}
	})
}
