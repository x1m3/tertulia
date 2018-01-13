package model

type tertuliaRepos struct {
	Topics  ItopicsCRUD
	Persons IPersonsCRUD
}

type Tertulia struct {
	Topics *Topics
	//Persons *Persons
	Repos tertuliaRepos
}

func NewTertulia(topicsRepo ItopicsCRUD, personsRepo IPersonsCRUD) *Tertulia {
	return &Tertulia{
		Topics: nil,
		Repos: tertuliaRepos{
			Topics:  topicsRepo,
			Persons: personsRepo,
		},
	}
}
