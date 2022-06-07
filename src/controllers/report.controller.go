package controllers

import (
	"net/http"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db/reports"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/services"
	"github.com/gin-gonic/gin"
)

func HandleShowAllReports(g *gin.Context) {

	req := pb.GetRequest{
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := services.GetAllReports(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func HandleGetReportByUsername(g *gin.Context) {

	req := pb.GetRequest{
		Query:  g.Param("username"),
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := services.GetReportByUserName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func getReportQuestions() []string {
	return []string{"WHAT IS THE USERNAME OF THE PLAYER YOU ARE REPORTING?", "WHAT TYPE OF BEHAVIOR ARE YOU REPORTING?",
		"PASTE THE ID OF THE MATCH YOU'RE REFERRING TO HERE", "DESCRIPTION"}
}

func getReportCategories() []string {
	return []string{"AFK", "Assisting Enemy", "Cheating", "Communication Abuse - Text",
		"Communication Abuse - Voice", "Negative Attitude", "Offensive Name", "Threats"}
}

func HandleGetReportOptions(g *gin.Context) {
	res := pb.ReportOptionsGetResponse_Data{
		Questions:  getReportQuestions(),
		Categories: getReportCategories(),
	}

	lib.Success(g, &res)
}

func HandlePostReport(c *gin.Context) {
	var input reports.PostReport

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, httpStatus := services.PostNewReport(input.Username, input.ReportCategory, input.MatchID, input.ReportDetail)
	c.JSON(httpStatus, gin.H{"status": res})
}
