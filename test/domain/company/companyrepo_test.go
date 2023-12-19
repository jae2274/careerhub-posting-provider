package company

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/test/tinit"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompanyRepo(t *testing.T) {
	savedJpId := &company.CompanyId{Site: "jumpit", CompanyId: "savedId"}
	savedJpId2 := &company.CompanyId{Site: "jumpit", CompanyId: "savedId2"}
	savedJpId3 := &company.CompanyId{Site: "wanted", CompanyId: "savedId3"}
	notExistedJpId := &company.CompanyId{Site: "notExistedSite", CompanyId: "notExistedId"}
	notExistedJpId2 := &company.CompanyId{Site: "notExistedSite2", CompanyId: "notExistedId2"}

	t.Run("SaveAndFind", func(t *testing.T) {
		companyRepo := tinit.InitCompanyRepo(t)

		savedJp := company.NewCompany(savedJpId.Site, savedJpId.CompanyId)

		_, err := companyRepo.Save(savedJp)

		require.NoError(t, err)

		findedJp, err := companyRepo.Get(savedJpId)

		require.NoError(t, err)
		require.NotNil(t, findedJp)
		findedJp.CreatedAt = savedJp.CreatedAt //ignore createdAt
		require.Equal(t, savedJp, findedJp)
	})

	t.Run("FindNotExisted", func(t *testing.T) {
		companyRepo := tinit.InitCompanyRepo(t)

		findedMatches, err := companyRepo.Get(notExistedJpId)

		require.NoError(t, err)
		require.Nil(t, findedMatches)
	})

	t.Run("SaveAndFindAll", func(t *testing.T) {
		companyRepo := tinit.InitCompanyRepo(t)

		savedJp := company.NewCompany(savedJpId.Site, savedJpId.CompanyId)
		savedJp2 := company.NewCompany(savedJpId2.Site, savedJpId2.CompanyId)
		savedJp3 := company.NewCompany(savedJpId3.Site, savedJpId3.CompanyId)
		savedJps := []*company.Company{savedJp, savedJp2, savedJp3}

		_, err := companyRepo.Save(savedJp)
		require.NoError(t, err)
		_, err = companyRepo.Save(savedJp2)
		require.NoError(t, err)
		_, err = companyRepo.Save(savedJp3)
		require.NoError(t, err)

		findedJps, err := companyRepo.Gets([]*company.CompanyId{savedJpId, notExistedJpId, savedJpId2, notExistedJpId2, savedJpId3})

		require.NoError(t, err)
		require.Len(t, findedJps, 3)

		for i, findedJp := range findedJps {
			savedJps[i].CreatedAt = findedJp.CreatedAt //ignore createdAt
		}

		require.True(t, isContain(findedJps, savedJp), "findedJps: %v, savedJps: %v", findedJps, savedJps)
		require.True(t, isContain(findedJps, savedJp2), "findedJps: %v, savedJps: %v", findedJps, savedJps)
		require.True(t, isContain(findedJps, savedJp3), "findedJps: %v, savedJps: %v", findedJps, savedJps)
	})
}

func isContain(jps []*company.Company, jp *company.Company) bool {
	for _, jp2 := range jps {
		if jp2.Site == jp.Site && jp2.CompanyId == jp.CompanyId {
			return true
		}
	}

	return false
}

func isContainsId(ids []*company.CompanyId, idToFind *company.CompanyId) bool {
	for _, item := range ids {
		if item.Site == idToFind.Site && item.CompanyId == idToFind.CompanyId {
			return true
		}
	}
	return false
}
