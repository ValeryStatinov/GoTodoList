package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type ProjectsManager struct {
	db *sql.DB
}

func (pm *ProjectsManager) GetByUserId(id int) (*[]models.ProjectWithTasksCount, bool) {
	projects := make([]models.ProjectWithTasksCount, 0)
	query := "select id, name from projects where userId=$1"

	rows, err := pm.db.Query(query, id)
	if err != nil {
		systemlogger.Log(err.Error(), query, string(id))
		return &projects, false
	}
	defer rows.Close()

	for rows.Next() {
		project := models.ProjectWithTasksCount{}

		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			systemlogger.Log(err.Error(), query, string(id))
			return &projects, false
		}
		projects = append(projects, project)
	}

	for index, project := range projects {
		var count int
		query = "select count(*) from projects a inner join tasks b on a.id = b.projectId where a.id=$1"
		err = pm.db.QueryRow(query, project.Id).Scan(&count)
		if err != nil {
			systemlogger.Log(err.Error(), query, string(id))
			return &projects, false
		}
		projects[index].TasksCount = count
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

func (pm *ProjectsManager) Create(p *models.Project, u *models.User) bool {
	query := "insert into projects (name, userId) values ($1, $2) returning id"
	err := pm.db.QueryRow(query, p.Name, u.Id).Scan(&p.Id)
	if err != nil {
		systemlogger.Log(err.Error(), query, p.Name, string(u.Id))
		return false
	}

	return true
}
