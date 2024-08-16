package main

import (
	"context"
	"errors"
)

const UserId = "user_id"

type Messenger interface {
	SignUp(username, password string) (*User, error)
	CreateGroup(title string) (*Group, error)
	LeaveGroup(groupId uint) error
	AddMember(groupId, newUserId uint, accessType AccessType) error
	RemoveMember(groupId, removedUserId uint) error
}
type MessengerModule struct {
	// UserRepository
	groups []*Group
	users  []*User
	ctx    context.Context
}

func NewMessengerModule() Messenger {
	return &MessengerModule{
		groups: make([]*Group, 0),
		users:  make([]*User, 0),
		ctx:    context.Background(),
	}
}
func (m *MessengerModule) contextWithValue(userId uint) context.Context {
	ctx := context.WithValue(m.ctx, UserId, userId)
	return ctx
}

func (m *MessengerModule) SignUp(username, password string) (*User, error) {
	for _, u := range m.users {
		if username == u.username {
			return nil, errors.New("this user signed up before")
		}
	}
	newUser := CreateUser(uint(len(m.users)+1), username, password)
	m.ctx = m.contextWithValue(newUser.id)
	m.users = append(m.users, newUser)
	return newUser, nil
}
func (m *MessengerModule) CreateGroup(title string) (*Group, error) {
	userId := m.ctx.Value(UserId).(uint)
	var user *User
	for _, u := range m.users {
		if u.id == userId {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	group, err := CreateGroup(user, uint(len(m.groups)+1), title)
	if err != nil {
		return nil, err
	}
	m.groups = append(m.groups, group)
	return group, nil
}
func (m *MessengerModule) LeaveGroup(groupId uint) error {
	userId := m.ctx.Value(UserId).(uint)

	var group *Group
	for _, g := range m.groups {
		if g.id == groupId {
			group = g
			break
		}
	}
	if group == nil {
		return errors.New("no group with given id found")
	}

	err := group.LeaveGroup(userId)
	return err
}

// default accessType is unknown set if its not default
func (m *MessengerModule) AddMember(groupId, newUserId uint, accessType AccessType) error {
	// TODO: YOU SHOULD IMPLEMENT
	return nil
}
func (m *MessengerModule) RemoveMember(groupId, removedUserId uint) error {
	// TODO: YOU SHOULD IMPLEMENT
	return nil
}
