package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(12, 5)
	expected := 17
	if result != expected {
		t.Fatalf("the actual result is %d, which expected value was %d", result, expected)
	}
	t.Log(fmt.Sprint("test is passed"))
}

func TestAddUserToSamane(t *testing.T) {
	samane := &Samane{
		users: make(map[string]*Person),
	}
	newUser := NewPerson("amir1234", "password1234")
	assert.NotNil(t, newUser)

	err := samane.AddUser(newUser)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, len(samane.users), 1)
	assert.Len(t, samane.users, 1)

	err = samane.AddUser(newUser)
	assert.Error(t, err)
	assert.NotNil(t, err)

	newUser = NewPerson("amir12345", "password1234")
	assert.NotNil(t, newUser)

	err = samane.AddUser(newUser)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Len(t, samane.users, 2)
}
