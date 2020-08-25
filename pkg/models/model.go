package models

import(
	"time"
)

type UserData struct {
	ID string //primary key not null
	video string //will have the path location to users video -- not null (if nothing uploaded the default path will have a default video maybe if not photo)
	photo string //will have the path location to users photo -- not null (if nothing uploaded the default path will have a default photo)
	Created time.Time
	Expires time.Time
}