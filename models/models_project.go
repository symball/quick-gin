// ******** Connects to Postgresql DB to extract and modify data in DB tables

package models

import (
	"log"
	"time"

	"github.com/pledgecamp/pledgecamp-oracle/constants"
	"github.com/symball/quick-gin/connect"
)

const (
	projectTable = "project"
)

// Project struct
type Project struct {
	ID          int                     `db:"project_id,omitempty"`
	CreatedAt   time.Time               `db:"project_created_at"`
	Status      constants.ProjectStatus `db:"project_status"`
	Title       string                  `db:"project_title"`
	Description string                  `db:"project_description"`
}

var projects []Project
var project Project

// ProjectInsert - insert new project entry
func ProjectInsert(project Project) (Project, error) {
	dbConnection := connect.Postgres()
	defer dbConnection.Close()

	projectCollection := dbConnection.Collection(projectTable)
	log.Print("Inside project model ", project)
	_, err := projectCollection.Insert(project)
	if err != nil {
		log.Println(err)
		return project, err
	}
	return project, nil
}

// ProjectsFetch - Get project entries
func ProjectsFetch() ([]Project, error) {
	dbConnection := connect.Postgres()
	defer dbConnection.Close()
	projectCollection := dbConnection.Collection(projectTable)
	res := projectCollection.Find()
	log.Print(res)
	err := res.All(&projects)
	if err != nil {
		log.Println(err)
		return projects, err
	}
	return projects, nil
}
