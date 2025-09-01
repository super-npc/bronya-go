package resp

type TopRightHeaderResp struct {
	SubSystemButtons []SubSystemButtons `json:"subSystemButtons"`
	LoginUser        LoginUser          `json:"loginUser"`
}

type SubSystemButtons struct {
	Label string `json:"label"`
	Level string `json:"level"`
}

type LoginUser struct {
	Id         string `json:"id"`
	UserAvatar string `json:"userAvatar"`
	UserName   string `json:"userName"`
	Roles      []Role `json:"roles"`
}

type Role struct {
	Id       string `json:"id"`
	RoleName string `json:"roleName"`
}
