package models

type TinyURL struct {
	ShortID string `json:"url_id,omitempty"`
	LongURL string `json:"long_url,omitempty"`
}

type TinyURLResponse struct {
	ShortID string `json:"url_id,omitempty"`
}

func (url TinyURL) ToResponse() *TinyURLResponse {
	return &TinyURLResponse{
		ShortID: url.ShortID,
	}
}
