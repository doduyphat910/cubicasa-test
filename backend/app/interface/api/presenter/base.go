package presenter

type Response struct {
	Meta   Meta        `json:"meta,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
}

type Meta struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Total   uint64 `json:"total,omitempty"`
}

type Error struct {
	Code     int    `json:"code"`
	Messages string `json:"message"`
}
