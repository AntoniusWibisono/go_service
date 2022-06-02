package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/plugin/soft_delete"
)

type Comments struct {
	ID             string        `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	OrganizationId string        `gorm:"type:uuid;default:null" json:"organizationId"`
	Organizations  Organizations `gorm:"foreignKey:OrganizationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MemberId       string        `gorm:"type:uuid;default:null" json:"memberId"`
	Members        Members       `gorm:"foreignKey:MemberId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment        string        `gorm:"type:varchar(150);not null;index" json:"comment"`
	CreatedAt      time.Time     `gorm:"default:now();index" json:"-"`
	UpdatedAt      null.Time     `gorm:"default:null;index" json:"-"`
	DeletedAt      soft_delete.DeletedAt
}

type CommentsRepository interface {
	GetListComments(searchBy, searchValue, sortBy, sortType string, page, perPage int64) (comments []Comments, count int64, err error)
	GetCommentByOrganizationId(organizationId string) (comments []Comments, err error)
	PostNewComment(organizationId, memberId, comment string) (comments Comments, err error)
	SoftDeleteCommentData(organizationId string) (comments Comments, err error)
}
