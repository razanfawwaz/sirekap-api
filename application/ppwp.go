package application

import (
	"github.com/razanfawwaz/sirekap-api/config"
	"github.com/razanfawwaz/sirekap-api/domain"
	"github.com/razanfawwaz/sirekap-api/repository"
)

func GetAllData(limit int) ([]domain.ReportData, error) {
	db := config.GetDB()
	ppwpRepository := repository.NewPpwpRepository(db)
	ppwp, err := ppwpRepository.GetAllData(limit)
	if err != nil {
		return nil, err
	}
	return ppwp, nil
}
