package req

type PageReq struct {
	OrderBy     string `json:"orderBy"`
	OrderDir    string `json:"orderDir"`
	Page        int    `json:"page"`
	PerPage     int    `json:"perPage"`
	One2ManyReq struct {
		Entity         string `json:"entity"`
		EntityField    string `json:"entityField"`
		EntityFieldVal string `json:"entityFieldVal"`
	} `json:"One2ManyReq"`
	BindMiddleChild struct {
		Entity       string `json:"entity"`
		SelfField    string `json:"selfField"`
		SelfFieldVal string `json:"selfFieldVal"`
		JoinField    string `json:"joinField"`
	} `json:"BindMiddleChild"`
}
