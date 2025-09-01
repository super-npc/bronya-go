package resp

type SiteResp struct {
	Pages []Module `json:"pages"`
}

type Module struct {
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	Order    int    `json:"order"`
	Label    string `json:"label"`
	Menu     []Menu `json:"children"`
}

type Menu struct {
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	Order    int    `json:"order"`
	Label    string `json:"label"`
	Icon     string `json:"icon"`
	Children []Leaf `json:"children"`
}

type Leaf struct {
	Id        string `json:"id"`
	ParentId  string `json:"parentId"`
	Order     int    `json:"order"`
	Label     string `json:"label"`
	Icon      string `json:"icon"`
	SchemaApi string `json:"schemaApi"`
	Url       string `json:"url"`
}
