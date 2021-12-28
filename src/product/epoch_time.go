package product

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// EpochTime allows the proper un/marshaling of a seconds-since-epoch value from JSON.
type EpochTime struct {
	gorm.Model
	Uuid uuid.UUID `gorm:"type:char(36);primary_key"`
	time.Time
	CreatedTimeId      string
	LastImageId        string
	LastModificationId string
	CompletedId        string
}

func (t *EpochTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	t.Uuid = uuid.New()
	secs, err := strconv.ParseInt(string(b), 0, 64)
	if err != nil {
		t.Time = time.Time{}
		return err
	}

	t.Time = time.Unix(secs, 0)
	return nil
}

func (t *EpochTime) MarshalJSON() ([]byte, error) {
	if !t.isSet() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", t.Time.UnixNano()/1000000)), nil
}

var nilTime = (time.Time{}).UnixNano()

// isSet allows testing if the value was set. If it is empty, it will return false.
func (t *EpochTime) isSet() bool {
	return t.UnixNano() != nilTime
}
