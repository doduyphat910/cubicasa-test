package presenter

type Response struct {
	Meta     Meta        `json:"meta,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Messages []string    `json:"messages,omitempty"`
}

type Meta struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Total   uint64 `json:"total,omitempty"`
}

type Paging struct {
	Size   uint64 `json:"size"`
	Number uint64 `json:"number"`
}
