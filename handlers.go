package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//handlers

func Profile(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//return user profile
		user, err := myUser(c, s)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to get user"})
			return
		}

		c.JSON(200, gin.H{"user": user})
	}
}
func UpdateProfile(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}
		if body.Username == "" && body.Password == "" {
			c.JSON(400, gin.H{"error": "at least one field is required"})
			return
		}

		//get user
		user, err := myUser(c, s)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to get user"})
			return
		}

		//update user
		if body.Username != "" {
			for _, takenUser := range s.UserMap {
				if takenUser.Username == body.Username {
					c.JSON(400, gin.H{"error": "username taken"})
					return
				}
			}
			user.Username = body.Username
		}
		if body.Password != "" {
			user.Password = body.Password
		}
		err = s.UpdateUser(user.Id, user)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to update user"})
			return
		}
		c.JSON(200, gin.H{"message": "user updated successfully"})
	}
}

func ListTodos(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//list all todos
		var todoList []Task = s.GetAllTasks(c.MustGet("user_id").(uint64))
		c.JSON(200, gin.H{"todos": todoList})
	}
}
func CreateTodo(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			DueDate     string `json:"due_date"` // ISO8601 format
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		//parse due date
		dueDate, err := time.Parse(time.RFC3339, body.DueDate)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid due date format", "message": "a due date in the RFC3339 format is required"})
			return
		}
		//create task
		uId := c.MustGet("user_id").(uint64)
		fmt.Println("uId from c.MustGet: /v", uId)
		var task Task = CreateTask(body.Title, body.Description, dueDate, uId)
		err = s.AddTask(task, uId)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create task", "message": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "task created successfully", "task": task})
	}
}
func ListFulfilled(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//list all todos
		var todoList []Task = s.GetAllTasks(c.MustGet("user_id").(uint64))
		//return fulfilled todos
		fulfilled := make([]Task, len(todoList)/2)
		for _, t := range todoList {
			if t.Completed {
				fulfilled = append(fulfilled, t)
			}
		}
		c.JSON(200, gin.H{"fulfilled tasks": fulfilled})
	}
}

func UpdateTodo(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Completed   bool   `json:"completed"`
			DueDate     string `json:"due_date"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		//get task
		task, err := myTask(c, s)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to get task"})
			return
		}

		//update task
		if body.Title != "" {
			task.Title = body.Title
		}
		if body.Description != "" {
			task.Description = body.Description
		}
		if body.Completed != false && task.Completed != body.Completed {
			task.Completed = body.Completed
		}
		if body.DueDate != "" {
			dueDate, err := time.Parse(time.RFC3339, body.DueDate)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid due date format", "message": "a due date in the RFC3339 format is required"})
				return
			}
			task.DueDate = dueDate
		}

		err = s.UpdateTask(task.Id, task)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to update task"})
			return
		}
		c.JSON(200, gin.H{"message": "task updated successfully", "task": task})
	}
}
func DeleteTodo(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get task
		task, err := myTask(c, s)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to get task"})
			return
		}

		//delete task
		err = s.DeleteTask(task.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to delete task"})
			return
		}
		c.JSON(200, gin.H{"message": "task deleted successfully"})
	}
}
func FulfillTodo(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var body struct {
			Completed bool `json:"completed"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		//get task
		task, err := myTask(c, s)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to get task"})
			return
		}

		//update task
		task.Completed = body.Completed
		err = s.UpdateTask(task.Id, task)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to update task"})
			return
		}
		c.JSON(200, gin.H{"message": "task updated successfully", "task": task})
	}

}

//helpers

func myUser(c *gin.Context, s *Store) (User, error) {
	UserId, ok := c.MustGet("user_id").(uint64)
	if !ok {
		return User{}, errors.New("user_id not found in context")
	}
	user, err := s.GetUser(UserId)
	return user, err
}
func myTask(c *gin.Context, s *Store) (Task, error) {
	id_str := c.Param("id")
	Id, err := strconv.ParseUint(id_str, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid task id"})
		return Task{}, err
	}
	task, err := s.GetTask(Id)
	if err != nil {
		c.JSON(404, gin.H{"error": "task not found"})
		return Task{}, err
	}
	return task, nil
}
