package queue

import (
	"net/http"

	"github.com/gorilla/mux"
	web_broker "github.com/phpCoder88/web-broker/pkg/web-broker"
)

type queueSvc interface {
	Put(req *web_broker.PutValueReq) error
	Get(req *web_broker.GetValueReq) (*web_broker.GetValueResp, error)
}

func RegisterPublicHTTP(queueSvc queueSvc) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{queue}", putToQueue(queueSvc)).Methods(http.MethodPut)
	r.HandleFunc("/{queue}", getFromQueue(queueSvc)).Methods(http.MethodGet)

	return r
}
