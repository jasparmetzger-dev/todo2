package main

import (
	"errors"
	"fmt"
	"time"
)

var UserNotFound = errors.New("User not found")
var TaskNotFound = errors.New("Task not found")

type Store struct {
	UserMap    map[uint64]User
	TaskMap    map[uint64]Task
	nextUserId uint64
	nextTaskId uint64
}

func NewStore() *Store {
	return &Store{
		UserMap:    make(map[uint64]User),
		TaskMap:    make(map[uint64]Task),
		nextUserId: 1,
		nextTaskId: 1,
	}
}

//creating objects, this doesnt add them to the store, just creates them, you have to add them to the store with the add functions

func CreateUser(username string, password string) User {
	user := User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
	return user
}
func CreateTask(title string, description string, dueDate time.Time, userId uint64) Task {
	task := Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		UserId:      userId,
		CreatedAt:   time.Now(),
	}
	return task
}

// User CRUD operations
func (s *Store) AddUser(user User) {
	user.Id = s.nextUserId
	s.nextUserId++
	s.UserMap[user.Id] = user
}
func (s *Store) UpdateUser(id uint64, updatedUser User) error {
	if _, ok := s.UserMap[id]; ok {
		updatedUser.Id = id
		s.UserMap[id] = updatedUser
		return nil
	}
	return UserNotFound
}
func (s *Store) DeleteUser(id uint64) error {
	if _, ok := s.UserMap[id]; ok {
		delete(s.UserMap, id)
		return nil
	}
	return UserNotFound
}
func (s *Store) GetUser(id uint64) (User, error) {
	if user, ok := s.UserMap[id]; ok {
		return user, nil
	}
	return User{}, UserNotFound
}
func (s *Store) GetAllUsers() []User {
	users := make([]User, 0, len(s.UserMap))
	for _, user := range s.UserMap {
		users = append(users, user)
	}
	return users
}

//Task CRUD operations

func (s *Store) AddTask(t Task, userId uint64) error {

	t.Id = s.nextTaskId
	fmt.Println("looking for userId:", userId)
	fmt.Println("userMap:", s.UserMap)
	t.UserId = userId
	s.nextTaskId++
	s.TaskMap[t.Id] = t

	var u User = s.UserMap[userId]
	if u.Id == 0 {
		return UserNotFound
	}

	u.TaskIds = append(u.TaskIds, t.Id)
	s.UserMap[u.Id] = u

	return nil
}
func (s *Store) UpdateTask(id uint64, updatedTask Task) error {
	if _, ok := s.TaskMap[id]; ok {
		updatedTask.Id = id
		s.TaskMap[id] = updatedTask
		return nil
	}
	return TaskNotFound
}
func (s *Store) DeleteTask(id uint64) error {
	if _, ok := s.TaskMap[id]; ok {
		delete(s.TaskMap, id)
		return nil
	}
	return TaskNotFound
}
func (s *Store) GetTask(id uint64) (Task, error) {
	if task, ok := s.TaskMap[id]; ok {
		return task, nil
	}
	return Task{}, TaskNotFound
}
func (s *Store) GetAllTasks(uId uint64) []Task {
	tasks := make([]Task, 0)
	for _, task := range s.TaskMap {
		if task.UserId == uId {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// search user by taskId
func (s *Store) GetUserByTaskId(taskId uint64) (User, error) {
	if _, ok := s.TaskMap[taskId]; !ok {
		return User{}, TaskNotFound
	}

	for _, user := range s.UserMap {
		for _, tId := range user.TaskIds {
			if tId == taskId {
				return user, nil
			}
		}
	}
	return User{}, UserNotFound
}
