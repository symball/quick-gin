package models

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var testProjectId int
var testId int
var testVoteId int

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Main - No .env file found")
	}
}

// Tests for models_project.go
func TestProjectInsert(t *testing.T) {
	log.Println("********************************* TestProjectInsert() **************************************")
	var testProject Project
	testProject.CreatedAt = time.Now()
	testProject.Status = 1
	testProject.Title = "Test Project 1"
	testProject.Description = "Test Project 1 Description"
	_, err := ProjectInsert(testProject)
	if err != nil {
		t.Error("Could not insert project")
	} else {
		fmt.Printf("The project %v was inserted successfully. \n", testProject)
		log.Println("********************************* End TestProjectInsert() **************************************")
	}
}

func TestProjectsFetch(t *testing.T) {
	log.Println("********************************* TestProjectsFetch() **************************************")
	var testProject Project
	testProject.CreatedAt = time.Now()
	testProject.Status = 1
	testProject.Title = "Test Project 1"
	testProject.Description = "Test Project 1 Description"
	_, err := ProjectInsert(testProject)

	var testProject2 Project
	testProject2.CreatedAt = time.Now()
	testProject2.Status = 1
	testProject2.Title = "Test Project 1"
	testProject2.Description = "Test Project 1 Description"
	_, err = ProjectInsert(testProject2)

	testProjectSet, err := ProjectsFetch()
	if err != nil {
		t.Error("Could not get current projects")
	}

	if len(testProjectSet) == 0 {
		t.Error("Could not fetch projects")
	}
	log.Println("********************************* End TestProjectFetchCurrent() **************************************")
}
