package models

import "gorm.io/gorm"

type Fruit struct {
	gorm.Model
	FruitID       int     `json:"id"`
	Name          string  `json:"name"`
	Family        string  `json:"family"`
	Genus         string  `json:"genus"`
	Order         string  `json:"order"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
	Fat           float64 `json:"fat"`
	Calories      float64 `json:"calories"`
	Sugar         float64 `json:"sugar"`
}
