package handler

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	//"github.com/wzzlYwzzl/httpdatabase/client"
	"github.com/wzzlYwzzl/httpdatabase/resource/users"
	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type ApiHandler struct {
	userns users.Users
	dbconf *sqlop.MysqlCon
}

// Web-service filter function used for request and response logging.
func wsLogger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf(FormatReqestLog(req))
	chain.ProcessFilter(req, resp)
	log.Printf(FormatResponseLog(resp, req))
}

// FormatRequestLog formats request log string.
// TODO(maciaszczykm): Display request body.
func FormatRequestLog(req *restful.Request) string {
	reqURI := ""
	if req.Request.URL != nil {
		reqURI = req.Request.URL.RequestURI()
	}

	return fmt.Sprintf(RequestLogString, req.Request.Proto, req.Request.Method,
		reqURI, req.Request.RemoteAddr)
}

// FormatResponseLog formats response log string.
// TODO(maciaszczykm): Display response content.
func FormatResponseLog(resp *restful.Response, req *restful.Request) string {
	return fmt.Sprintf(ResponseLogString, req.Request.RemoteAddr, resp.StatusCode())
}

func CreateApiHandler(dbconf *sqlop.MysqlCon) http.Handler {
	wsContainer := restful.NewContainer()

	apiHandler := new(ApiHandler)
	apiHandler.dbconf = dbconf

	usersWs := new(restful.WebService)
	usersWs.Filter(wsLogger)
	usersWs.Path("/api/v1/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	usersWs.Route(usersWs.GET("/{name}").
		To(apiHandler.judgeUser).
		Writes(users.Users{}))
	usersWs.Route(usersWs.POST("/api/v1/{name}/{namespaces}").
		To(apiHandler.createNS))
}

func (apiHandler *ApiHandler) judgeUser(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	apiHandler.userns.Name = name

	b, err := apiHandler.userns.JudgeExist(apiHandler.dbconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

}

// Handler that writes the given error to the response and sets appropriate HTTP status headers.
func handleInternalError(response *restful.Response, err error) {
	log.Print(err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusInternalServerError, err.Error()+"\n")
}
