package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type ProjectsManager struct {
	db *sql.DB
}

func (pm *ProjectsManager) GetAll() []models.Project {
	projects := make([]models.Project, 0)

	rows, err := pm.db.Query("select * from projects")
	if err != nil {
		systemlogger.Log(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		project := models.Project{}

		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			systemlogger.Log(err.Error())
		}

		projects = append(projects, project)
	}

	return projects
}

func (pm *ProjectsManager) GetByUserId(id int) []models.Project {
	projects := make([]models.Project, 0)

	rows, err := pm.db.Query("select id, name from projects where id=$1", id)
	if err != nil {
		systemlogger.Log(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		project := models.Project{}

		rows.Scan(&project.Id, &project.Name)
		projects = append(projects, project)
	}

	return projects
}
