package aggregator

func flatenRecord(connName string, record *map[string]interface{}) {
	(*record)["origin"] = connName
}

/*AggregateResultSet aggregates a series of resultsets.
It expects a map where the key is the origin, and the value a slice of map[string]interface{}
The return value is the slice of map[string]interface{} with the key "origin" injected on each item
*/
func AggregateResultSet(rs map[string]interface{}) []map[string]interface{} {
	if len(rs) == 0 {
		return nil
	}

	result := []map[string]interface{}{}

	// for each resultSet
	for k, v := range rs {
		for _, r := range v.([]map[string]interface{}) {
			flatenRecord(k, &r)
			result = append(result, r)
		}
	}
	return result
}
