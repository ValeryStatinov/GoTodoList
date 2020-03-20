package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type ProjectsManager struct {
	db *sql.DB
}

func (pm *ProjectsManager) GetByUserId(id int) (*[]models.Project, bool) {
	projects := make([]models.Project, 0)
	query := "select id, name from projects where userId=$1"

	rows, err := pm.db.Query(query, id)
	if err != nil {
		systemlogger.Log(err.Error(), query, string(id))
		return &projects, false
	}
	defer rows.Close()

	for rows.Next() {
		project := models.Project{}

		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			systemlogger.Log(err.Error(), query, string(id))
			return &projects, false
		}
		projects = append(projects, project)
	}

	return &projects, true
}

func (pm *ProjectsManager) GetById(id int) (*models.Project, bool) {
	project := &models.Project{}
	query := "select id, name, userId from projects where id=$1"

	err := pm.db.QueryRow(query, id).Scan(&project.Id, &project.Name, &project.UserId)
	if err != nil {
		systemlogger.Log(err.Error(), query, string(id))
		return project, false
	}

	return project, true
}
