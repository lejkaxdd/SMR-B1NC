package handler

import (
	"context"
	"net"
	"net/http"
	"smb/pkg/api"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func (h *Handler) loginpage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) loginpage_auth(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		return
	}

	conn, err := grpc.Dial("172.26.0.4:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	c_grpc := api.NewSmartBalanceServiceClient(conn)
	req := &api.CheckUserRequest{Info: &api.User{Username: c.PostForm("username"), Password: c.PostForm("password")}}
	res, err := c_grpc.CheckUser(context.Background(), req)
	if err != nil {
		log.Println(err)
	}

	if res.GetToken() == "1" {
		token, err := generateJWT(c.PostForm("username"))

		if err != nil{
			log.Printf("Failed to generate token for user:%s", c.PostForm("username"))
		}
		
		ip := string(GetOutboundIP())
		c.SetCookie("Cookie", token, 3600, "/", ip, false, true)
		c.HTML(http.StatusOK, "index.html", nil)
	}else {
		log.Printf("Failed to generate token for user:%s", c.PostForm("username"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}