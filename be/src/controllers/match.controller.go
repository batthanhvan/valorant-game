package controllers

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/services"
	"github.com/gin-gonic/gin"
)

func HandleGetByMatchID(g *gin.Context) {

	req := pb.GetRequest{
		Query:  g.Param("username"),
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := services.GetMatchByUsername(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
