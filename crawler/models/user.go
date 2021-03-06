package models

import (
	"time"
)

type User struct {
	Id                     int64     `db:"id" json:"id"`
	IdStr                  string    `db:"id_str" json:"id_str"`
	ScreenName             string    `db:"screen_name" json:"screen_name"`
	Name                   string    `db:"name" json:"name"`
	CreatedAt              string    `db:"created_at" json:"created_at"`
	FollowersCount         int64     `db:"followers_count" json:"followers_count"`
	FriendsCount           int64     `db:"friends_count" json:"friends_count"`
	Verified               bool      `db:"verified" json:"verified"`
	AdditionalData         *string   `db:"additional_data,omitempty" json:"-"`
	NextCursor             int64     `db:"next_cursor" db:"next_cursor"`
	NextCursorStr          string    `db:"next_cursor_str" json:"next_cursor_str"`
	AreFollowersDownloaded bool      `db:"are_followers_downloaded" json:"-"`
	DateLastChange         time.Time `db:"date_last_change" json:"-"`
	Protected              *bool     `db:"protected" json:"protected"`
	Location               *string   `db:"location" json:"location"`
	Description            string    `db:"-" json:"description"`
}

func minInt(ints ...int) int {
	if len(ints) == 0 {
		return 0
	}
	min := ints[0]
	for _, num := range ints {
		if num < min {
			min = num
		}
	}
	return min
}

func BatchUsersArray(arr []*User, batchSize int) [][]*User {
	batchCount := len(arr)/batchSize + 1
	batched := make([][]*User, 0, batchCount)
	for i := 0; i < batchCount; i++ {

		lowerBound := i * batchSize
		upperBound := minInt((i+1)*batchSize, len(arr))
		batch := arr[lowerBound:upperBound]
		batched = append(batched, batch)
	}

	return batched
}
