package encoder

import (
	"encoding/json"
	"log"
)

func ResultSetEncoder(rs []map[string]interface{}) {
	b, _ := json.Marshal(rs)
	log.Fatal(string(b))
}
