package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lealhugui/query-reducer/aggregator"
	"github.com/lealhugui/query-reducer/db"
	"github.com/prometheus/common/log"
)

//QueryRequestPayload is the payload definition for the "/query" route
type QueryRequestPayload struct {
	QueryText string
}

//QueryHandler is the handler for the "/query" route
func QueryHandler(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var qReq QueryRequestPayload
	if err := decoder.Decode(&qReq); err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if qReq.QueryText == "" {
		err := errors.New("Empty Query String")
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	qResult, err := db.Instance.Query(qReq.QueryText)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rs := aggregator.AggregateResultSet(qResult)
	//bytes, err := json.Marshal(rs)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, rs)

}
