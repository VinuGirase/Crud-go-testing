package models

type Cohort struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
}
