package result

type Result struct {
	Result_des  string                 `json:"result_des"`
	Result_code string         `json:"result_code"`
	Data        string           `json:"data,omitempty"`
}
