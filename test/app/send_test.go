package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

	"github.com/jae2274/goutils/cchan"
	"github.com/stretchr/testify/require"
)

func TestSendJobPostingApp(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		ctx := context.Background()
		src := jumpit.NewJumpitSource(ctx, 3000)

		grpcService, sendJobApp := initComponents(t, src)

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
			isRegistered, err := grpcService.IsCompanyRegistered(context.TODO(), &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})
			require.NoError(t, err)
			require.True(t, isRegistered)
		}
	})
}

func initComponents(t *testing.T, src source.JobPostingSource) (provider_grpc.ProviderGrpcService, *app.SendJobPostingApp) {
	grpcClient := tinit.InitGrpcClient(t)
	grpcService := provider_grpc.NewProviderGrpcService(grpcClient)

	return grpcService, app.NewSendJobPostingApp(src, grpcService)
}
