package models

type Pod struct {
	Name   string `json:"name"`
	HostIP string `json:"hostIp"`
	PodIP  string `json:"podIp"`
}
