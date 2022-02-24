package models

type Page struct {
	PageNumber int64 `json:"PageNumber,omitempty"`
	PageSize   int64 `json:"PageSize,omitempty"`
}

type KVPair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
