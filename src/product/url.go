package product

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

type URL struct {
	gorm.Model
	Uuid      uuid.UUID `gorm:"type:char(36);primary_key"`
	Url       string    `gorm:"type:text"`
	ProductId string
}

func (u *URL) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.Url = parsed.String()
	u.Uuid = uuid.New()
	return nil
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", u.Url)), nil
}
