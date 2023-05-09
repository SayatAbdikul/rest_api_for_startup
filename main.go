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
	http.HandleFunc("/reg_startup", post.RegStartup)
	http.HandleFunc("/reg_team", post.RegTeam)
	http.HandleFunc("/reg_achievement", post.RegAchievements)
	http.HandleFunc("/reg_investor", post.RegInvestor)
	http.HandleFunc("/reg_cases", post.RegCases)
	http.HandleFunc("/reg_favourite_startup", post.RegFavouriteStartup)
	http.HandleFunc("/reg_favourite_investor", post.RegFavouriteInvestor)
	http.HandleFunc("/get_startups", get.GetStartups)
	http.HandleFunc("/get_investors", get.GetInvestors)
	http.HandleFunc("/get_investor", get.GetInvestor)
	http.HandleFunc("/get_startup", get.GetStartup)
	http.HandleFunc("/patch_startup", patch.PatchStartup)
	http.HandleFunc("/patch_investor", patch.PatchInvestor)
	http.HandleFunc("/patch_startup_description", patch.PatchStartupDescription)
	http.HandleFunc("/patch_investor_description", patch.PatchInvestorDescription)
	http.HandleFunc("/patch_team", patch.PatchTeam)
	http.HandleFunc("/patch_startup_achievement", patch.PatchStartupAchievement)
	http.HandleFunc("/patch_case", patch.PatchCase)
	http.HandleFunc("/delete_startup", delete.DeleteStartup)
	http.HandleFunc("/delete_investor", delete.DeleteInvestor)
	http.HandleFunc("/delete_achievement", delete.DeleteAchievement)
	http.HandleFunc("/delete_case", delete.DeleteCase)
	http.HandleFunc("/delete_favourite_startup", delete.DeleteFavStartup)
	http.HandleFunc("/delete_favourite_investor", delete.DeleteFavInvestor)
	http.HandleFunc("/delete_team_member", delete.DeleteTeamMember)
	http.HandleFunc("/auth_startup", authorization.StartupAuthorization)
	http.HandleFunc("/auth_investor", authorization.InvestorAuthorization)
	http.ListenAndServe(":9001", nil)
}
