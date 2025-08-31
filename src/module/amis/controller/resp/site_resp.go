package resp

type T struct {
	Pages []struct {
		Id       string `json:"id"`
		ParentId string `json:"parentId"`
		Order    int    `json:"order"`
		Label    string `json:"label"`
		Children []struct {
			Id       string `json:"id"`
			ParentId string `json:"parentId"`
			Order    int    `json:"order"`
			Label    string `json:"label"`
			Icon     string `json:"icon"`
			Children []struct {
				Id        string `json:"id"`
				ParentId  string `json:"parentId"`
				Order     int    `json:"order"`
				Label     string `json:"label"`
				Icon      string `json:"icon"`
				SchemaApi string `json:"schemaApi"`
				Url       string `json:"url"`
			} `json:"children"`
		} `json:"children"`
	} `json:"pages"`
}
