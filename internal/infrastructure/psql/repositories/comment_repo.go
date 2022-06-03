package repositories

import "go_service/internal/domain"

func (r *DBRepository) GetCommentByOrganizationId(organizationId string) (comments []domain.Comments, err error) {
	tx := r.db.Model(domain.Comments{}).
		Select("comments.comment").
		Where("organization_id =?", organizationId)

	if err = tx.Find(&comments).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) SoftDeleteCommentData(organizationId string) (err error) {
	tx := r.db.Model(domain.Comments{}).Where("organization_id =?", organizationId).Update("DeletedAt", 1)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) PostNewComment(organizationId, memberId, comment string) (comments domain.Comments, err error) {
	tx := r.db.Model(domain.Comments{}).Create(map[string]interface{}{
		"Comment":        comment,
		"MemberId":       memberId,
		"OrganizationId": organizationId,
		"DeletedAt":      0,
	})

	if err = tx.Find(&comments, map[string]interface{}{
		"comment":         comment,
		"member_id":       memberId,
		"organization_id": organizationId,
	}).Error; err != nil {
		return
	}

	return
}
