package central

type HelathResponse struct {
	Code        int     `json:"code"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Message     *string `json:"message"`
}
