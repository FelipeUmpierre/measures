package domain

type (
	// User holds the domain data
	User struct {
		ID   int    `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}
)
