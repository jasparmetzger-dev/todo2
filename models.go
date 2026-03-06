package main

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	TaskIds   []uint64  `json:"task_ids"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	UserId      uint64    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
