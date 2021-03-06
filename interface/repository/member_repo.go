package repository

import (
	"sync"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

var (
	defaultMemberDB = map[member.MemberID]*member.Member{
		member.MemberID("1"): &member.Member{
			ID:   member.MemberID("1"),
			Name: "Tuan",
			Emails: []member.Email{
				member.Email("shiro@zabuton.co.jp"),
			},
		},
		member.MemberID("2"): &member.Member{
			ID:   member.MemberID("2"),
			Name: "Tue",
			Emails: []member.Email{
				member.Email("kenji@zabuton.co.jp"),
			},
		},
	}
)

type InMemMemberRepo struct {
	mtx    sync.RWMutex
	member map[member.MemberID]*member.Member
}

func NewInMemMemberRepo() *InMemMemberRepo {
	return &InMemMemberRepo{
		member: defaultMemberDB,
	}
}

func (r *InMemMemberRepo) GetAll() ([]*member.Member, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	res := []*member.Member{}
	for _, val := range r.member {
		res = append(res, val)
	}
	return res, nil
}
func (r *InMemMemberRepo) Get(memberID member.MemberID) (*member.Member, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.member[memberID]; ok {
		return val, nil
	}
	return nil, member.ErrorMemberNotFound
}
func (r *InMemMemberRepo) Save(member *member.Member) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.member[member.ID] = member
	return nil
}
