package goztl

type EnrollmentSecret struct {
	ID                 int      `json:"id"`
	Secret             string   `json:"secret"`
	MetaBusinessUnitID int      `json:"meta_business_unit"`
	TagIDs             []int    `json:"tags"`
	SerialNumbers      []string `json:"serial_numbers"`
	UDIDs              []string `json:"udids"`
	Quota              *int     `json:"quota"`
	RequestCount       int      `json:"request_count"`
}

type EnrollmentSecretRequest struct {
	MetaBusinessUnitID int      `json:"meta_business_unit"`
	TagIDs             []int    `json:"tags"`
	SerialNumbers      []string `json:"serial_numbers"`
	UDIDs              []string `json:"udids"`
	Quota              *int     `json:"quota"`
}
