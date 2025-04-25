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

// GetCertificatesByPage 分页查询未删除的证书
func (r *CertificatesRepository) GetCertificatesByPage(page, pageSize int) ([]*models.Certificates, error) {
	var certificates []*models.Certificates
	err := r.engine.
		Table(models.Certificates{}.TableName()).
		Where("is_deleted = 1").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&certificates).Error
	return certificates, err
}

func (r *CertificatesRepository) AddCertificate(certificate *models.Certificates) error {
	err := r.engine.Table(models.Certificates{}.TableName()).Create(certificate).Error
	return err
}

// UpdateCertificate 根据证书 ID 更新证书信息
func (r *CertificatesRepository) UpdateCertificate(certificate *models.Certificates) error {
	// 这里假设使用 ID 作为唯一标识进行更新
	err := r.engine.Table(models.Certificates{}.TableName()).
		Where("id = ?", certificate.ID).
		Updates(certificate).
		Error
	return err
}

// DeleteCertificate 根据证书 ID 标记证书为已删除
func (r *CertificatesRepository) DeleteCertificate(id interface{}) error {
	updateData := map[string]interface{}{
		"is_deleted": 0,
	}
	err := r.engine.Table(models.Certificates{}.TableName()).
		Where("id = ?", id).
		Updates(updateData).
		Error
	return err
}
