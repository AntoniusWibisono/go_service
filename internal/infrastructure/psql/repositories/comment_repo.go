package repositories

import "go_service/internal/domain"

func (r *DBRepository) GetListComments(searchBy, searchValue, sortBy, sortType string, page, perPage int64) (comments []domain.Comments, count int64, err error) {
	tx := r.db.Model(domain.Comments{}).
		Select("comments.id, comments.comment")

	if err = tx.Count(&count).Error; err != nil {
		return
	}

	_generatePaginationQueries(tx, page, perPage)
	if err = tx.Find(&comments).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) GetCommentByOrganizationId(organizationId string) (comments []domain.Comments, err error) {
	tx := r.db.Model(domain.Comments{}).
		Select("comments.comment").
		Where("organization_id =?", organizationId)

	if err = tx.Find(&comments).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) SoftDeleteCommentData(organizationId string) (comments domain.Comments, err error) {
	r.db.Model(domain.Comments{}).Where("organization_id =?", organizationId).Update("DeletedAt", 1)

	return
}

func (r *DBRepository) PostNewComment(organizationId, memberId, comment string) (comments domain.Comments, err error) {
	r.db.Model(domain.Comments{}).Create(map[string]interface{}{
		"Comment":        comment,
		"MemberId":       memberId,
		"OrganizationId": organizationId,
		"DeletedAt":      0,
	})

	return
}
