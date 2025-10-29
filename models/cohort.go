package models

type Cohort struct {
	Name        string `json:"name"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
}
