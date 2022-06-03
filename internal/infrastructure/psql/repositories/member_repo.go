package repositories

import "go_service/internal/domain"

func (r *DBRepository) GetMemberData(organizationId string) (member []domain.Members, err error) {
	tx := r.db.Model(domain.Members{}).
		Select("id, username, avatar_url, followers, following").
		Where("organization_id =?", organizationId).Order("followers desc")

	if err = tx.Find(&member).Error; err != nil {
		return
	}

	return
}
