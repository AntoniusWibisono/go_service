package database

import "gorm.io/gorm"

func MigrateAndSeed(db *gorm.DB) (err error) {
	var member = new(domain.Members)
	var organizations = new(domain.organizations)
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

	if err = db.AutoMigrate(member); err != nil {
		return
	}
	if err = db.AutoMigrate(*&organizations); err != nil {
		return
	}
	if err = db.AutoMigrate(comments); err != nil {
		return
	}

}
