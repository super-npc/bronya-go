package resp

type AppInfoResp struct {
	AppName           string `json:"appName"`
	GitBaseVersion    string `json:"gitBaseVersion"`
	GitProjectVersion string `json:"gitProjectVersion"`
}
