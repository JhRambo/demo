package config

const COLLECTION = "testcollection"
const RECORDLOGS = "recordlogs"
const LOCKTTL = 1800

type Response struct {
	Code    int32  `json:"name:code"`
	Message string `json:"name:message"`
}
