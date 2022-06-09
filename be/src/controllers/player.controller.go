package controllers

import (
	"net/http"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db/users"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/services"
	"github.com/batthanhvan/src/utils/token"
	"github.com/gin-gonic/gin"
)

func HandleGetByUserName(g *gin.Context) {
	req := pb.GetRequest{
		Query: g.Param("username"),
	}

	res, err := services.GetByUserName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func HandlePostModifyUser(c *gin.Context) {
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := users.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.PostModifyUserReq{
		Username:   user.Username,
		Playername: c.Param("playername"),
		Tagline:    c.Param("tagline"),
	}

	res, err := services.ModifyUser(&req)
	if err != nil {
		lib.BadRequest(c, err)
		return
	}
	lib.Success(c, res)
}
