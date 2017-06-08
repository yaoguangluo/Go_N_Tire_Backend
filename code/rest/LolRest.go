package main

import (
	"github.com/backend/code/rest/delegateDao"
	//"github.com/backend/code/rest/delegateHandler"
	"github.com/backend/code/rest/delegateInterface"
	"github.com/backend/code/rest/lib"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"

	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"syscall"

	"github.com/martini-contrib/render"
)

var (
	ac    ApiConfig
	x     delegateDao.Orm
	m     *martini.Martini
	fconf *string = flag.String("c", "./config.ini", "config file")
)

func main() {
	err := lib.DBC.Read(*fconf)
	orm, err := lib.DBC.InitOrm()
	x = delegateDao.Orm{DB: orm}
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Hello World!")
	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "templates",
		},
	))
	delegateInterface.BuildRestInterface(x, m)
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go http.Serve(listener, m)
	log.Println("Listening on 0.0.0.0:" + port)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	listener.Close()
}
