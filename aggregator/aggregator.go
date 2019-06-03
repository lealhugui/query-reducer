package aggregator

func flatenRecord(connName string, record *map[string]interface{}) {
	(*record)["origin"] = connName
}

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
