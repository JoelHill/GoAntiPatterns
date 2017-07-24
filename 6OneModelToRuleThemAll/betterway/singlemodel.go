package singlemodel

// Team - the team model
import "time"

// Team - Single model that represents the json, database, and internal team model.
type Team struct {
	ID   int    `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name"`
	// if subdomain is derived and is never received nor sent back, we can simply ignore the json value.
	Subdomain                string     `json:"-"`
	Logo                     string     `db:"logo" json:"logo"`
	OwnerID                  int        `db:"ownerID" json:"ownerID"`
	IsDeleted                bool       `db:"isDeleted" json:"isDeleted"`
	IsDeletedChangedByUserID int        `db:"isDeletedChangedByUserID" json:"isDeletedChangedByUserID,omitempty"`
	IsDeletedChangedAt       *time.Time `db:"isDeletedChangedAt" json:"isDeletedChangedAt,omitempty"`
	IsLocked                 bool       `db:"isLocked" json:"isLocked"`
	IsLockedChangedByUserID  int        `db:"isLockedChangedByUserID" json:"isLockedChangedByUserID,omitempty"`
	IsLockedChangedAt        *time.Time `db:"isLockedChangedAt" json:"isLockedChangedAt,omitempty"`
	CreatedAt                *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt                *time.Time `db:"updatedAt" json:"updatedAt"`
	EndedAt                  *time.Time `json:"endedAt,omitempty"`
}

// If you need to manipulate the return json, you can do this:

// IO input / output represents an API request or response body
type IO struct {
	Team            *Team `json:"team"`
	IsInVisionAdmin bool  `json:"isInVisionAdmin,omitempty"`
}

// {
//     "team": {
//         "id": 1,
//         "name": "invision"
//     }
// }
