package projects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func newProject(id int, name string) Project {
	project := Project{id, name}
	return project
}

func HandleProjects(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		rows, err := GetPreparedGetProjectsStmt().Query()

		if err != nil {
			fmt.Println("error query")
			return
		}
		defer rows.Close()

		if err = rows.Err(); err != nil {
			fmt.Println("error rows")
			return
		}

		projects := make([]Project, 0)
		for rows.Next() {
			var id int
			var name string

			err := rows.Scan(&id, &name)
			if err != nil {
				fmt.Println("error scan")
				return
			}

			projects = append(projects, newProject(id, name))
		}

		js, err := json.Marshal(projects)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(js)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	}

	if request.Method == "POST" {
		var project Project
		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}
		err = json.Unmarshal(body, &project)
		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}

		_, err = GetPreparedPostProjectStmt().Exec(project.Name)

		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}

		writer.WriteHeader(200)
	}
}
