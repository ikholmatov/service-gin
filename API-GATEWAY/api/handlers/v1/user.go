package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/venomuz/project4/API-GATEWAY/genproto"
	l "github.com/venomuz/project4/API-GATEWAY/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	body := pb.Useri{}
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by id
// route /v1/users/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetByID(
		ctx, &pb.GetIdFromUser{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
type  response struct{
	aa string
	as string
}

func (h *handlerV1) Param(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	mail := c.Query("mail")
	password := c.Query("password")
	fmt.Println(mail,password)
	_, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	ff := response{aa: mail, as: password}
	fmt.Print(ff)
	c.JSON(http.StatusOK, gin.H{
		"gg": mail,
		"ssd": password,
	})
}
