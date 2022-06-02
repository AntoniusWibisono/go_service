package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/plugin/soft_delete"
)

type Members struct {
	ID             string        `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	OrganizationId string        `gorm:"type:uuid;default:null" json:"organizationId"`
	Organizations  Organizations `gorm:"foreignKey:OrganizationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Username       string        `gorm:"type:varchar(150);not null;index" json:"username"`
	Password       string        `gorm:"type:varchar(150);not null;index" json:"password"`
	AvatarUrl      string        `gorm:"type:varchar(150);not null;index" json:"avatarUrl"`
	Followers      int64         `gorm:"type:int;not null;" json:"followers"`
	Following      int64         `gorm:"type:int;not null;" json:"following"`
	CreatedAt      time.Time     `gorm:"default:now();index" json:"-"`
	UpdatedAt      null.Time     `gorm:"default:null;index" json:"-"`
	DeletedAt      soft_delete.DeletedAt
}

type MembersRepository interface {
	GetMemberData(organizationId string) (member Members, err error)
}
