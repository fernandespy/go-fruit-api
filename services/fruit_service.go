package services

import (
	"encoding/json"
	"errors"
	"go-fruit-api/database"
	"go-fruit-api/models"
	"io/ioutil"
	"net/http"
)

type fruitAPIResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Family     string `json:"family"`
	Genus      string `json:"genus"`
	Order      string `json:"order"`
	Nutritions struct {
		Carbohydrates float64 `json:"carbohydrates"`
		Protein       float64 `json:"protein"`
		Fat           float64 `json:"fat"`
		Calories      float64 `json:"calories"`
		Sugar         float64 `json:"sugar"`
	} `json:"nutritions"`
}

func LoadFruitsFromAPI() error {
	resp, err := http.Get("https://www.fruityvice.com/api/fruit/all")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to obtain API data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data []fruitAPIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, f := range data {
		database.DB.Where(models.Fruit{Name: f.Name}).Assign(models.Fruit{
			FruitID:       f.ID,
			Family:        f.Family,
			Genus:         f.Genus,
			Order:         f.Order,
			Carbohydrates: f.Nutritions.Carbohydrates,
			Protein:       f.Nutritions.Protein,
			Fat:           f.Nutritions.Fat,
			Calories:      f.Nutritions.Calories,
			Sugar:         f.Nutritions.Sugar,
		}).FirstOrCreate(&models.Fruit{})
	}

	return nil
}
