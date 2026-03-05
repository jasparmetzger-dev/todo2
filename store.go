package main

import (
	
)

var IdCounter int = 1

type Store struct {
	Users [int]
	User{}
	Tasks []Task
}

func NewStore() *Store {
	return &Store{
		Users: []User{},
		Tasks: []Task{},
	}
}

func (s *Store) AddUser(user User) {
	user.Id = IdCounter
	IdCounter++
	s.Users = append(s.Users, user)
}

func (s *Store) UpdateUser(id int, updatedUser User) {
	if 


