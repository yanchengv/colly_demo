package pachong

import "time"

type Article struct {
	ID        uint
	Title     string
	Subtitle  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
