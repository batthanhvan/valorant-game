package controllers

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/service"
	"github.com/gin-gonic/gin"
)

func HandleGetByUserName(g *gin.Context) {
	req := pb.GetRequest{
		Query:  g.DefaultQuery("query", g.Param("username")),
		Limit:  g.DefaultQuery("limit", "20"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := service.GetByUserName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func HandlePostModifyUser(g *gin.Context) {
	req := pb.PostModifyUserReq{
		Username:   g.DefaultQuery("username", g.Param("username")),
		Playername: g.DefaultQuery("playername", g.Param("playername")),
		Tagline:    g.DefaultQuery("tagline", g.Param("tagline")),
	}

	res, err := service.ModifyUser(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
