package controllers

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/services"
	"github.com/gin-gonic/gin"
)

func HandleShowAllReports(g *gin.Context) {
	req := pb.GetRequest{
		Query:  g.DefaultQuery("query", "%"),
		Limit:  g.DefaultQuery("limit", "20"),
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
		Query:  g.DefaultQuery("username", "%"),
		Limit:  g.DefaultQuery("limit", "20"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := services.GetReportByUserName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
