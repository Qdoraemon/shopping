package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type CertificatesRepository struct {
	engine *gorm.DB
}

func NewCertificatesRepository(engine *gorm.DB) *CertificatesRepository {
	return &CertificatesRepository{engine: engine}
}

// TODO 分页查询
func (r *CertificatesRepository) GetCertificatesByPage(page, pageSize int) ([]*models.Certificates, error) {
	var certificates []*models.Certificates
	err := r.engine.
		Table(models.Certificates{}.TableName()).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&certificates).Error
	return certificates, err
}
