package member

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type MemberUsecase struct {
	MemberRepo IMemberRepository
}

func NewMemberUsecase(memberRepo IMemberRepository) *MemberUsecase {
	usecase := MemberUsecase{}
	usecase.MemberRepo = memberRepo
	return &usecase
}

func (u *MemberUsecase) RegisterMember(name string) (*member.Member, error) {
	member := member.NewMember(name)
	if err := u.MemberRepo.Save(member); err != nil {
		return nil, err
	}
	return member, nil
}

func (u *MemberUsecase) RegisterNewEmail(memberID member.MemberID, email member.Email) (*member.Member, error) {
	member, err := u.MemberRepo.Get(memberID)
	if err != nil {
		return nil, err
	}
	if err = member.RegisterEmail(email); err != nil {
		return nil, err
	}
	if err = u.MemberRepo.Save(member); err != nil {
		return nil, err
	}
	return member, nil
}
