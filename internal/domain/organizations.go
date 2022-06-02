package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/plugin/soft_delete"
)

type Organizations struct {
	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null;index" json:"name"`
	CreatedAt time.Time `gorm:"default:now();index" json:"-"`
	UpdatedAt null.Time `gorm:"default:null;index" json:"-"`
	DeletedAt soft_delete.DeletedAt
}

type OrganizationsRepository interface {
	GetOrganizationData(organizationName string) (organization Organizations, err error)
}
