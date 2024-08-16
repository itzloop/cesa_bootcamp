package main

import "errors"

func Add(a, b int) int {
	return a + b
}

type Person struct {
	username string
	password string
}

type Samane struct {
	users map[string]*Person
}

func (s *Samane) AddUser(person *Person) error {
	if person == nil {
		return errors.New("there is no user to add to samane")
	}
	if _, exists := s.users[person.username]; exists {
		return errors.New("the user has already been added to samane")
	}
	s.users[person.username] = person
	return nil
}

func NewPerson(username string, password string) *Person {
	return &Person{
		username: username,
		password: password,
	}
}
