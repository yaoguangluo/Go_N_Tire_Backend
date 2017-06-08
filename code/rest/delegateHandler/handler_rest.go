package delegateHandler

import (
	//"crypto/md5"
	//"encoding/base64"
	//"encoding/json"
	"fmt"
	//"io"
	//"log"
	//"net/http"
	"github.com/backend/code/rest/delegateDao"
	"github.com/martini-contrib/render"
	//"github.com/martini-contrib/sessions"
	"github.com/go-martini/martini"
)

func HandlerRest(x delegateDao.Orm, r render.Render, params martini.Params) (int, string) {

	fmt.Println("Hello World!1")
	str := params["uname"]
	obj := map[string]interface{}{
		"error": "10001",
		"msg":   x.SelectValuebyName(str)}

	r.JSON(200, map[string]interface{}{
		"error": "10001",
		"msg":   obj})
	return 200, ""
}
