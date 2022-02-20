package http

type httpError struct {
	Message string `json:"message"`
}

type authResponse struct {
	Token    string `json:"access_token"`
	Type     string `json:"token_type"`
	Expires  int16  `json:"expires_in"`
	Instance string `json:"rest_instance_url"`
}
