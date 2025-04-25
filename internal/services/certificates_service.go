package services

import (
	"errors"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"time"

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

func (us *CertificatesService) AddCertificate(certificate *models.Certificates) error {
	if certificate == nil {
		return errors.New("certificate cannot be nil")
	}
	certificate.CreateTime = time.Now()
	certificate.UpdateTime = time.Now()
	return us.certificatesRepo.AddCertificate(certificate)
}
