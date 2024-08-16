package main

import (
	"errors"
	"time"
)

type AccessType int

const (
	AccessTypeUnknown AccessType = iota
	AccessTypeMember
	AccessTypeAdmin
	AccessTypeOwner
)

// User is a struct
type User struct {
	id       uint
	username string
	password string
}

func CreateUser(id uint, username, password string) *User {
	return &User{
		id:       id,
		username: username,
		password: password,
	}
}

// Group is a struct
type Group struct {
	id          uint
	title       string
	adminUserId uint
	members     map[uint]*User
	accesses    map[uint]AccessType
	createdAt   time.Time
}

func CreateGroup(user *User, groupId uint, title string) (*Group, error) {
	if user == nil {
		return nil, errors.New("no user existed")
	}
	newGroup := &Group{
		id:          groupId,
		title:       title,
		adminUserId: user.id,
		members:     make(map[uint]*User),
		accesses:    make(map[uint]AccessType),
		createdAt:   time.Now(),
	}
	newGroup.members[user.id] = user
	newGroup.accesses[user.id] = AccessTypeOwner
	return newGroup, nil
}
func (g *Group) AddMember(adderUserId uint, newUserGroup *User) error {
	if newUserGroup == nil {
		return errors.New("user not exists")
	}
	if _, exists := g.members[adderUserId]; !exists {
		return errors.New("admin not found in group")
	}
	if access := g.accesses[adderUserId]; access != AccessTypeAdmin && access != AccessTypeOwner {
		return errors.New("no access for this user to add member")
	}

	if _, exists := g.members[newUserGroup.id]; exists {
		return errors.New("user is in the group")
	}
	g.members[newUserGroup.id] = newUserGroup
	g.accesses[newUserGroup.id] = AccessTypeMember
	return nil
}
func (g *Group) SetAccess(ownerUserId, userId uint, accessType AccessType) error {
	if _, exists := g.members[ownerUserId]; !exists {
		return errors.New("owner with given id is not in the group")
	}

	if g.adminUserId != ownerUserId {
		return errors.New("no access to change access of a user")
	}

	if _, exists := g.members[userId]; !exists {
		return errors.New("user with given user id is not available in the group")
	}
	g.accesses[userId] = accessType
	return nil
}
func (g *Group) RemoveMember(removerUserId uint, removedUserId uint) error {
	if _, exists := g.members[removerUserId]; !exists {
		return errors.New("can not remove member no user found")
	}
	if access := g.accesses[removerUserId]; access != AccessTypeAdmin && access != AccessTypeOwner {
		return errors.New("no access to remove member")
	}

	if _, exists := g.members[removedUserId]; exists {
		return errors.New("user is in the group")
	}

	if access := g.accesses[removedUserId]; access == AccessTypeOwner {
		return errors.New("can not remove owner of the group")
	}

	delete(g.members, removedUserId)
	delete(g.accesses, removedUserId)
	return nil
}
func (g *Group) LeaveGroup(userId uint) error {
	if _, exists := g.members[userId]; !exists {
		return errors.New("user not found in group")
	}
	delete(g.members, userId)
	delete(g.accesses, userId)
	return nil
}
