package db

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lealhugui/query-reducer/config"
	_ "github.com/lib/pq" //postgres driver
	"github.com/prometheus/common/log"
)

//Instance is a pointer for the global Instance
var Instance *Controller

//ConnInstance represents one instance of a DBConnection
type ConnInstance struct {
	Name string
	Conn *sqlx.DB
}

//Query is the main method for querying data
func (c ConnInstance) query(sql string) ([]map[string]interface{}, error) {
	rows, err := c.Conn.Query(sql)
	if err != nil {

		return nil, err
	}

	defer rows.Close()

	result := []map[string]interface{}{}

	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		queryScanner := map[string]interface{}{}
		for i, col := range columns {
			queryScanner[col.Name()] = new(GenericScanner)
			values[i] = queryScanner[col.Name()]

		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}
		resultObj := map[string]interface{}{}
		for _, col := range columns {
			resultObj[col.Name()] =
				queryScanner[col.Name()].(*GenericScanner).value

		}
		result = append(result, resultObj)
	}
	return result, nil
}

func (c ConnInstance) queryAsync(sql string, flowCtrl chan queryChanResult) {
	result, err := c.query(sql)
	if err != nil {
		log.Error(err)
	} else {
		flowCtrl <- queryChanResult{c.Name, result}
	}
}

//Controller represents a set of ConnInstances
type Controller struct {
	Connections map[string]ConnInstance
}

type queryChanResult struct {
	ConnName  string
	ResultSet []map[string]interface{}
}

//Query is the main method for querying data over a set of dbs
func (ctr Controller) Query(sql string) (result map[string]interface{}, e error) {
	queryResults := make(chan queryChanResult, len(ctr.Connections))
	for _, conn := range ctr.Connections {
		go conn.queryAsync(sql, queryResults)
	}
	result = map[string]interface{}{}

	i := 0
	for r := range queryResults {
		i++
		result[r.ConnName] = r.ResultSet
		if i == len(ctr.Connections) {
			close(queryResults)
			break
		}
	}
	return result, e
}

//AppendConn appends one ConnInstance on Controller.Connections if it doesn't exists (search by name)
func (ctr Controller) AppendConn(dbConn config.DbConfig) {
	connName := dbConn.ConnectionName
	if connName == "" {
		connName = fmt.Sprintf("%s@%s", dbConn.DbName, dbConn.Host)
	}
	if _, exits := ctr.Connections[connName]; exits {
		return
	}

	newConn, err := aquireConn(dbConn)
	if err != nil {
		log.Error(err)
		return
	}

	ctr.Connections[connName] = ConnInstance{
		Name: connName,
		Conn: newConn,
	}

}

//New is the factory method for the Controller
func New(dbs []config.DbConfig) (*Controller, error) {
	if Instance != nil {
		return Instance, errors.New("unable to reinitialize controller")
	}
	Instance = &Controller{}
	Instance.Connections = make(map[string]ConnInstance)
	if len(dbs) > 0 {
		for _, dbConf := range dbs {
			Instance.AppendConn(dbConf)
		}
	}

	return Instance, nil
}

//aquireConn creates a DB connection
func aquireConn(dbConf config.DbConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.DbName,
	)
	conn, err := sqlx.Connect("postgres", connStr)
	return conn, err
}
