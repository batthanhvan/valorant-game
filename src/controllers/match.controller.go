package controllers

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/service"
	"github.com/gin-gonic/gin"
)

func HandleGetByMatchID(g *gin.Context) {
	req := pb.GetRequest{
		Query:  g.DefaultQuery("query", "%"),
		Limit:  g.DefaultQuery("limit", "20"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := service.GetByMatchID(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
