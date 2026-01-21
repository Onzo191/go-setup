package seeds

import "gorm.io/gorm"

// Seeder is an interface that all seeders must implement
type Seeder interface {
	Name() string
	Run(db *gorm.DB) error
}
