package model

type ConfigurationModel struct {
	ServiceName string `json:ServiceName`
	HashKey     string `json:"Hashkey"`
	HashValue   string `json:"Hashvalue"`
}
