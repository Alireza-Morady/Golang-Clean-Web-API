package logging

func mapToZapParams(Extra map[ExtraKey]interface{}) []interface{}{
	params := make([]interface{},0)
	for k,v := range Extra{
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}