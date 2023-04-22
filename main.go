package main

import (
	"database/sql"
	"github.com/SayatAbdikul/rest_api_for_startup/authorization"
	delete "github.com/SayatAbdikul/rest_api_for_startup/deleteRequests"
	get "github.com/SayatAbdikul/rest_api_for_startup/getRequests"
	patch "github.com/SayatAbdikul/rest_api_for_startup/patchRequests"
	post "github.com/SayatAbdikul/rest_api_for_startup/postRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	var err error
	server.DBConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/infomatrix_project")
	defer server.DBConn.Close()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/regStartup", post.RegStartup)
	http.HandleFunc("/regTeam", post.RegTeam)
	http.HandleFunc("/regAchievements", post.RegAchievements)
	http.HandleFunc("/get_startups", get.GetStartups)
	http.HandleFunc("/get_startup", get.GetStartup)
	http.HandleFunc("/patch_startup", patch.PatchStartup)
	http.HandleFunc("/patch_startup_description", patch.PatchStartupDescription)
	http.HandleFunc("/patch_team", patch.PatchTeam)
	http.HandleFunc("/patch_startup_achievement", patch.PatchStartupAchievement)
	http.HandleFunc("/delete_startup", delete.DeleteStartup)
	http.HandleFunc("/delete_achievement", delete.DeleteAchievement)
	http.HandleFunc("/delete_team_member", delete.DeleteTeamMember)
	http.HandleFunc("/auth_startup", authorization.StartupAuthorization)
	http.ListenAndServe(":9090", nil)
}
