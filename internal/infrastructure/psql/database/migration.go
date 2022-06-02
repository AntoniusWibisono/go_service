package database

import (
	"go_service/internal/domain"
	"go_service/internal/infrastructure/psql/seeders"

	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) (err error) {
	var member = new(domain.Members)
	var organizations = new(domain.Organizations)
	var comments = new(domain.Comments)

	if err = db.Migrator().DropTable(member); err != nil {
		return
	}
	if err = db.Migrator().DropTable(organizations); err != nil {
		return
	}
	if err = db.Migrator().DropTable(comments); err != nil {
		return
	}

	if err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return
	}

	if err = db.AutoMigrate(member); err != nil {
		return
	}
	if err = db.AutoMigrate(organizations); err != nil {
		return
	}
	if err = db.AutoMigrate(comments); err != nil {
		return
	}

	seedMember, seedOrganization, seedComment := seeders.CreateSeed()

	for _, seedOrganizationVal := range seedOrganization {
		db.Create(seedOrganizationVal)
	}
	for _, seedMemberVal := range seedMember {
		db.Create(seedMemberVal)
	}
	for _, seedCommentVal := range seedComment {
		db.Create(seedCommentVal)
	}

	return
}
