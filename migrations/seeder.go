package migrations

import "test-api/app/model"

var (
	role     model.Role
	category model.Category
	brand    model.Brand
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		role.Seed(),
		category.Seed(),
		brand.Seed(),
	}
}
