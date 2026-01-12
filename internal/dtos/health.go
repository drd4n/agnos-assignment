package dtos

type HealthCheckResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}
