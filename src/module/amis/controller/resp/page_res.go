package resp

type PageRes struct {
	Total int64                    `json:"total"`
	Rows  []map[string]interface{} `json:"rows"`
}
