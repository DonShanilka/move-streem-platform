package Models

type Movie struct {
	//gorm.Model // 👈 adds ID, CreatedAt, UpdatedAt, DeletedAt automatically

	Title       string
	Description string
	ReleaseYear int
	Language    string
	Duration    int
	Rating      string
	AgeRating   string
	Country     string

	Thumbnail []byte `gorm:"type:longblob"`
	Banner    []byte `gorm:"type:longblob"`
	Trailer   []byte `gorm:"type:longblob"`

	MovieURL string
}
