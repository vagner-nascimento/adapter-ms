package infra

type consumerResponseDetails struct {
	HasErrs bool `json:"hasErrors"`
}

type consumerResponse struct {
	Id      string                    `json:"requestId"`
	Details []consumerResponseDetails `json:"responses"`
}
