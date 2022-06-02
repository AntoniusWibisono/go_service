package repositories

import "go_service/internal/domain"

func (r *DBRepository) GetOrganizationData(organizationName string) (organization domain.Organizations, err error) {
	tx := r.db.Model(domain.Organizations{}).
		Select("organizations.id").
		Where("name =?", organizationName)

	if err = tx.Find(&organization).Error; err != nil {
		return
	}

	return
}
