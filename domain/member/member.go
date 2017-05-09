package member

import (
	"errors"
	"regexp"

	uuid "github.com/satori/go.uuid"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/common"
)

type MemberID string
type Email string

var (
	emailRegex          = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	ErrorDuplicateEmail = errors.New("Duplicate email")
	ErrorEmailInvalid   = errors.New("Email invalid")
)

func (e Email) Equal(test Email) bool {
	return e == test
}

func IsEmailValid(email Email) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(string(email))
}

type Member struct {
	common.BaseEntity
	ID     MemberID
	Name   string
	Emails []Email
}

func NewMember(name string) *Member {
	member := Member{}
	uid := uuid.NewV4().String()
	member.ID = MemberID(uid)
	member.Name = name
	member.Emails = []Email{}
	return &member
}

func (m *Member) IsContainEmail(email Email) bool {
	for _, e := range m.Emails {
		if e.Equal(email) {
			return true
		}
	}
	return false
}

func (m *Member) RegisterEmail(email Email) error {
	if !IsEmailValid(email) {
		return ErrorEmailInvalid
	}
	if m.IsContainEmail(email) {
		return ErrorDuplicateEmail
	}
	m.Emails = append(m.Emails, email)
	return nil
}
