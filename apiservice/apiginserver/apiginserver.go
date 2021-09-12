package apiginserver

import (
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"log"

)

const (
	SERVER_NAME = "demo-http" // server name
)

func BuildServer() (srv server.Server){

	srv = httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8080"),
	)

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := newDemo()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	return
}

//demo
type demo struct{}

func newDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demoPOST)
	router.GET("/demo", a.demoGET)
}

func (a *demo) demoGET(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call GET go-micro v3 http server success"})
}

func (a *demo) demoPOST(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call POST go-micro v3 http server success"})
}
