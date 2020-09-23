package seed

import "github.com/MHNassar1/Auth/api/models"

var propSeeds = &ISeed{
	RunSeed: false,
	Items: models.Properties{
		Name:     "personal_info",
		ShowName: "Personal Information",
		ParentId: 0,
	},
}
