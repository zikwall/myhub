package domain

import "github.com/google/uuid"

type Channel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Timezone struct {
	ID      uuid.UUID `json:"id"`
	Zone    int16     `json:"zone"` // оставляем тип int16 для соответствия базе данных
	ZoneRFC string    `json:"zone_rfc"`
	Name    string    `json:"name"`
}

type Country struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Iso33661 string    `json:"iso3366_1"`
}

type Tag struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type ChannelUser struct {
	ID         uuid.UUID   `json:"id"`
	ChannelID  uuid.UUID   `json:"channel_id"`
	ZoneID     uuid.UUID   `json:"zone_id"`
	CategoryID uuid.UUID   `json:"category_id"`
	UserID     uuid.UUID   `json:"user_id"`
	Name       string      `json:"name"`
	URL        string      `json:"url"`
	Countries  []uuid.UUID `json:"countries"`
	Tags       []uuid.UUID `json:"tags"`
}
