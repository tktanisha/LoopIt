package models

type Society struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Pincode   string `json:"pincode"`
	CreatedAt string `json:"created_at"`
}
