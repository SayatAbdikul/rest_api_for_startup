package main

import (
	"database/sql"
	"net/http"

	"github.com/SayatAbdikul/rest_api_for_startup/authorization"
	delete "github.com/SayatAbdikul/rest_api_for_startup/deleteRequests"
	get "github.com/SayatAbdikul/rest_api_for_startup/getRequests"
	patch "github.com/SayatAbdikul/rest_api_for_startup/patchRequests"
	post "github.com/SayatAbdikul/rest_api_for_startup/postRequests"
	"github.com/SayatAbdikul/rest_api_for_startup/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	server.DBConn, err = sql.Open("mysql", "admin_root:BMKX55Rnt3MECAHB@tcp(31.172.67.121:3306)/infomatrix_project")
	defer server.DBConn.Close()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/api/reg_startup", post.RegStartup)
	http.HandleFunc("/api/reg_team", post.RegTeam)
	http.HandleFunc("/api/reg_achievement", post.RegAchievements)
	http.HandleFunc("/api/reg_investor", post.RegInvestor)
	http.HandleFunc("/api/reg_cases", post.RegCases)
	http.HandleFunc("/api/reg_favourite_startup", post.RegFavouriteStartup)
	http.HandleFunc("/api/reg_favourite_investor", post.RegFavouriteInvestor)
	http.HandleFunc("/api/get_startups", get.GetStartups)
	http.HandleFunc("/api/get_investors", get.GetInvestors)
	http.HandleFunc("/api/get_investor", get.GetInvestor)
	http.HandleFunc("/api/get_startup", get.GetStartup)
	http.HandleFunc("/api/patch_startup", patch.PatchStartup)
	http.HandleFunc("/api/patch_investor", patch.PatchInvestor)
	http.HandleFunc("/api/patch_startup_description", patch.PatchStartupDescription)
	http.HandleFunc("/api/patch_investor_description", patch.PatchInvestorDescription)
	http.HandleFunc("/api/patch_team", patch.PatchTeam)
	http.HandleFunc("/api/patch_startup_achievement", patch.PatchStartupAchievement)
	http.HandleFunc("/api/patch_case", patch.PatchCase)
	http.HandleFunc("/api/delete_startup", delete.DeleteStartup)
	http.HandleFunc("/api/delete_investor", delete.DeleteInvestor)
	http.HandleFunc("/api/delete_achievement", delete.DeleteAchievement)
	http.HandleFunc("/api/delete_case", delete.DeleteCase)
	http.HandleFunc("/api/delete_favourite_startup", delete.DeleteFavStartup)
	http.HandleFunc("/api/delete_favourite_investor", delete.DeleteFavInvestor)
	http.HandleFunc("/api/delete_team_member", delete.DeleteTeamMember)
	http.HandleFunc("/api/auth_startup", authorization.StartupAuthorization)
	http.HandleFunc("/api/auth_investor", authorization.InvestorAuthorization)
	http.ListenAndServe(":9001", nil)
}
