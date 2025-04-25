package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type CertificatesService struct {
	certificatesRepo *repositories.CertificatesRepository
}

func NewCertificatesService(engine *gorm.DB) *CertificatesService {
	return &CertificatesService{certificatesRepo: repositories.NewCertificatesRepository(engine)}
}

func (us *CertificatesService) GetCertificatesByPage(page, pageSize int) ([]*models.Certificates, error) {
	return us.certificatesRepo.GetCertificatesByPage(page, pageSize)
}
