package models

type Ipconfig struct {
	IP       string `json:"ip" db:"ip"`
	Hostfqdn string `json:"hostfqdn" db:"hostfqdn"`
	Active   string `json:"active" db:"active"`
}

// TableName overrides the table name used by Pop.
func (a Ipconfig) TableName() string {
	return "ipconfigs"
}
