package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Metadata struct {
	// meta
	Timestamp time.Time `validate:"required"`

	// base
	NameJa      string `validate:"required"`
	NameEn      string `validate:"required"`
	JoinYear    int    `validate:"required"`
	Description string `validate:"required"`

	// pictures
	PicturePath string `validate:"required"`

	// social
	GitHub  string
	Twitter string
	Website string
}

type Member struct {
	Metadata Metadata
	Body     string
}

type AppContext struct {
	PicturesDirectory string
	OutDirectory      string
	Since             time.Time
}

func MemberPreprocess(member *Member) error {
	if ValidateMember(*member) != nil {
		return fmt.Errorf("invalid member: %w", ValidateMember(*member))
	}

	member.Metadata.GitHub = strings.TrimPrefix(member.Metadata.GitHub, "https://github.com/")
	member.Metadata.Twitter = strings.TrimPrefix(member.Metadata.Twitter, "https://x.com/")
	member.Metadata.Twitter = strings.TrimPrefix(member.Metadata.Twitter, "@")
	return nil
}
func ValidateMember(member Member) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(member); err != nil {
		return err
	}
	return nil
}
