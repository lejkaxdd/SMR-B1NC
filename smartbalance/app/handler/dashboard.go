package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func (h *Handler) dashboard(c *gin.Context){
	
	cookie, err := c.Cookie("Cookie")
	
	log.Println(cookie)

	if err != nil {
		cookie = "NotSet"
		c.HTML(http.StatusOK,"index.html", gin.H{
			"result": "Unable to get Cookie",
		  })
	}

	verify, err := verifyJWT(cookie)
	if err != nil {
		log.Println(err)
	}

	log.Println("TOKEN SIGN", verify)
	
	if verify == "administrator" {
		c.HTML(http.StatusOK, "admin.html", nil)
	}else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"result": "You don't have permission to look this page",
		  })
	}

}