package mongo

// Media define
type Media struct {
	PublicID string `json:"public_id" bson:"public_id"`
	Width    int    `json:"width" bson:"width"`
	Height   int    `json:"height" bson:"height"`
	Format   string `json:"format" bson:"format"`
	Bytes    int    `json:"bytes" bson:"bytes"`
	URL      string `json:"byte" bson:"url"`
}

func (p *Media) GetPublicID() string {
	if p.PublicID == "" {
		return ""
	}
	return p.PublicID
}
func (p *Media) GetWidth() int {
	if p.Width == 0 {
		return 0
	}
	return p.Width
}
func (p *Media) GetHeight() int {
	if p.Height == 0 {
		return 0
	}
	return p.Height
}
func (p *Media) GetFormat() string {
	if p.Format == "" {
		return ""
	}
	return p.Format
}
func (p *Media) GetBytes() int {
	if p.Bytes == 0 {
		return 0
	}
	return p.Bytes
}
func (p *Media) GetURL() string {
	if p.URL == "" {
		return ""
	}
	return p.URL
}
