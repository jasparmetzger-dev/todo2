package main

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Tasks     []Task    `json:"tasks"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
