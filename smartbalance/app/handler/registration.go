package handler

import (
	"context"
	"fmt"
	"net/http"
	"smb/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)


func (h *Handler) regpage(c *gin.Context){
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *Handler) regpage_auth(c *gin.Context){

	if c.Request.Form.Get("password") == c.Request.Form.Get("confirmpassword"){
		
		conn, err := grpc.Dial("172.26.0.4:50051", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
		}	

		fmt.Println( c.Request.Form.Get("username") , c.Request.Form.Get("password"))
		c_grpc := api.NewSmartBalanceServiceClient(conn)
		req := &api.CreateUserRequest{Info: &api.User{Username: c.PostForm("username"), Password: c.PostForm("password")}}
		res, err := c_grpc.CreateUser(context.Background(), req)
		if err != nil{
			log.Println(err)
		}

		c.HTML(http.StatusOK, "wellregister.html", gin.H{
			"confirm": res.GetConfirm(),
		  })

	}else {
		c.HTML(http.StatusOK, "wellregister.html", gin.H{
			"confirm": "Password doesn't match",
		  })
	}

}
