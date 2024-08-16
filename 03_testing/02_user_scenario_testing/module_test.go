package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUp(t *testing.T) {
	m := NewMessengerModule()

	// Successful sign up
	user, err := m.SignUp("amir2134", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Failed sign up
	user, err = m.SignUp("amir2134", "pass1234")
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestCreatedGroup(t *testing.T) {
	m := NewMessengerModule()

	user, err := m.SignUp("amir2134", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	expected := "amir_owner_of_group"
	group, err := m.CreateGroup(expected)
	assert.NotNil(t, group)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), group.id)
	assert.EqualValues(t, expected, group.title)
	assert.Len(t, group.members, 1)
	assert.Equal(t, AccessTypeOwner, group.accesses[user.id])
	assert.Equal(t, user.id, group.adminUserId)
}

func TestLeaveGroup(t *testing.T) {
	m := NewMessengerModule()

	user, err := m.SignUp("amir2134", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	expected := "amir_owner_of_group"
	group, err := m.CreateGroup(expected)
	assert.NotNil(t, group)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), group.id)
	assert.EqualValues(t, expected, group.title)
	assert.Len(t, group.members, 1)
	assert.Equal(t, AccessTypeOwner, group.accesses[user.id])
	assert.Equal(t, user.id, group.adminUserId)

	err = m.LeaveGroup(group.id)
	assert.Nil(t, err)
	assert.Len(t, group.members, 0)
}

// TestAddMemberWithDifferentAccesses
// TestRemoveMemberWithDifferentAccesses (if owner fails otherwise remove)
// TODO: Implement if owner leave what should we do
//...
