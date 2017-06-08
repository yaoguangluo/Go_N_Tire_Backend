package delegateInterface

import (
	"net/http"
	"strconv"
	"time"

	"github.com/backend/code/rest/delegateDao"
	"github.com/backend/code/rest/delegateHandler"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
)

func BuildRestInterface(x delegateDao.Orm, m *martini.ClassicMartini) (int, string) {
	m.Get("/", func(r render.Render, req *http.Request) {
		if req.URL.Query().Get("wait") != "" {
			sleep, _ := strconv.Atoi(req.URL.Query().Get("wait"))
			time.Sleep(time.Duration(sleep) * time.Second)
		}
		r.HTML(200, "index", nil)
	})

	m.Get("/rest/:uname", func(r render.Render, params martini.Params, req *http.Request) {
		delegateHandler.HandlerRest(x, r, params)
	})

	return 200, ""
}
