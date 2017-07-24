package multiplestructs

import "time"

// Team - the main team model that we pass around
type Team struct {
	ID                       int
	Name                     string
	Subdomain                string
	Logo                     string
	OwnerID                  int
	IsDeleted                bool
	IsDeletedChangedByUserID int
	IsDeletedChangedAt       *time.Time
	IsLocked                 bool
	IsLockedChangedByUserID  int
	IsLockedChangedAt        *time.Time
	CreatedAt                *time.Time
	UpdatedAt                *time.Time
	EndedAt                  *time.Time
}

// TeamDB - The model how its constructed in the database
type TeamDB struct {
	ID                       int        `db:"id"`
	Name                     string     `db:"name"`
	Subdomain                string     `db:"subdomain"`
	Logo                     string     `db:"logo"`
	OwnerID                  int        `db:"ownerID"`
	IsDeleted                bool       `db:"isDeleted"`
	IsDeletedChangedByUserID int        `db:"isDeletedChangedByUserID"`
	IsDeletedChangedAt       *time.Time `db:"isDeletedChangedAt"`
	IsLocked                 bool       `db:"isLocked"`
	IsLockedChangedByUserID  int        `db:"isLockedChangedByUserID"`
	IsLockedChangedAt        *time.Time `db:"isLockedChangedAt"`
	CreatedAt                *time.Time `db:"createdAt"`
	UpdatedAt                *time.Time `db:"updatedAt"`
}

// TeamAPIReturn - The model how we want to return the json formatted data
type TeamAPIReturn struct {
	ID                       int        `json:"id,omitempty"`
	Name                     string     `json:"name"`
	Subdomain                string     `json:"subdomain"`
	Logo                     string     `json:"logo"`
	OwnerID                  int        `json:"ownerID"`
	IsDeleted                bool       `json:"isDeleted"`
	IsDeletedChangedByUserID int        `json:"isDeletedChangedByUserID,omitempty"`
	IsDeletedChangedAt       *time.Time `json:"isDeletedChangedAt,omitempty"`
	IsLocked                 bool       `json:"isLocked"`
	IsLockedChangedByUserID  int        `json:"isLockedChangedByUserID,omitempty"`
	IsLockedChangedAt        *time.Time `json:"isLockedChangedAt,omitempty"`
	CreatedAt                *time.Time `json:"createdAt"`
	UpdatedAt                *time.Time `json:"updatedAt"`
	EndedAt                  *time.Time `json:"endedAt,omitempty"`
}
