package routes

import (
	"encoding/json"
	"github.com/lealhugui/query-reducer/aggregator"
	"github.com/lealhugui/query-reducer/db"
	"github.com/prometheus/common/log"
	"net/http"
)

type QueryRequestPayload struct {
	QueryText string
}

func QueryHandler (resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var qReq QueryRequestPayload
	if err := decoder.Decode(&qReq); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	if qReq.QueryText == "" {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Error("Empty Query String")
		return
	}
	qResult, err := db.Instance.Query(qReq.QueryText)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	rs := aggregator.AggregateResultSet(qResult)
	bytes, err := json.Marshal(rs)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Header().Add("Content-Type", "text/json")
	_, err = resp.Write(bytes)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

}
