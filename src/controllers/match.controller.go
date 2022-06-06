package controllers

import (
	"strconv"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/services"
	"github.com/gin-gonic/gin"
)

func HandleGetByMatchID(g *gin.Context) {
	page, err := strconv.Atoi(g.DefaultQuery("page", "1"))
	lib.CheckError(err)

	if page <= 0 {
		lib.NotFoundRequest(g, err)
		return
	}

	req := pb.GetRequest{
		Query:  g.Param("username"),
		Limit:  "10",
		Offset: strconv.Itoa((page - 1) * 10),
	}

	res, err := services.GetMatchByUsername(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
