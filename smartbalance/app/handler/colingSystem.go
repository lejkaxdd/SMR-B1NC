package handler

import (
	"context"
	"strconv"
	// "strings"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	// "regexp"
	"net/http"
	"smb/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func (h *Handler) collingData(c *gin.Context) {

	if c.Request.Method == "GET" { 
		
		c.HTML(http.StatusOK, "coolingInsert.html", nil)

		// Insert
	} else if c.Request.Method == "POST" && c.Request.FormValue("operation") == ("Insert") {

		err := c.Request.ParseForm()
		if err != nil {
			return
		}

		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
		}

		c_grpc := api.NewSmartBalanceServiceClient(conn)
		req := &api.CoolingSystemRequest{Info: &api.CoolingSystem{Coolinglevel: c.PostForm("coolLevel"), Coolingfrequency: c.PostForm("coolFreq"), Coolingtype: c.PostForm("coolType")}}
		res, err := c_grpc.CoolingSystem(context.Background(), req)
	
		if err != nil {
			log.Println(err)
		}

		defer conn.Close()

		log.Println(res)

		c.JSON(http.StatusOK, gin.H{
			"id": res.GetRecord(),
		})

		// Check
	} else {

		id := c.Request.Form.Get("id")


		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
		}

		c_grpc := api.NewSmartBalanceServiceClient(conn)
		req := &api.CoolingSystemGetRequest{Record: id}
		res, err := c_grpc.CoolingSystemCheck(context.Background(), req)
	
		if err != nil {
			log.Println(err)
		}

		defer conn.Close()


		// if strings.TrimSpace(id) == "" {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "empty input id"})
		// 	return
		// }
		// stm := fmt.Sprintf(`{Record\s*=\s*(%s)\s*[^}]+}`, id)
		// re := regexp.MustCompile(stm)

		// myfile, err := os.ReadFile("/application/cooling_audit.txt")
		// if err != nil {
		// 	fmt.Print(err)
		// }

		// str := string(myfile)
		// match := re.FindString(str)
		// fmt.Println(match)

		val := rand.Intn(4) + 1
		// file :=
		data, err := os.ReadFile(fmt.Sprintf("/application/assets/img/col%s.png", strconv.Itoa(val)))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "coolingCheck.html", gin.H{
			"res":   res,
			"image": base64.StdEncoding.EncodeToString(data),
		})
	}
}

func (h *Handler) cooling(c *gin.Context) {

	filename := c.Query("filename")
	data, err := os.ReadFile("/application/assets/img/" + filename)
	// fmt.Println(string(data))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "image.html", gin.H{"image": base64.StdEncoding.EncodeToString(data)})
}
