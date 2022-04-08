package api

import (
	"github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/service"
	"github.com/gin-gonic/gin"
)

func HandleGetByUserName(g *gin.Context) {
	req := pb.PlayerGetRequest{
		Query:  g.DefaultQuery("query", "%"),
		Limit:  g.DefaultQuery("limit", "5"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := service.GetByUserName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
