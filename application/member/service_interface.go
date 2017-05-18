package member

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IMemberRepository interface {
	GetAll() ([]*member.Member, error)
	Get(memberID member.MemberID) (*member.Member, error)
	Save(member *member.Member) error
}
