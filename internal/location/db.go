package location

import (
	"encoding/json"
	"os"
	"path"
	"strings"
)

type city struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionCode  string `json:"region_code"`
	RegionName  string `json:"region_name"`
}

type region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"state_code"`
}

type country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"iso2"`
}

type synonym struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Db struct {
	cities    []city
	regions   []region
	countries []country
	synonyms  []synonym
}

func NewDb() (*Db, error) {
	db := Db{}

	err := db.load()
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *Db) load() error {
	err := db.loadCountries()
	if err != nil {
		return err
	}

	err = db.loadRegions()
	if err != nil {
		return err
	}

	err = db.loadCities()
	if err != nil {
		return err
	}

	err = db.loadSynonyms()
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) Search(query string) string {

	query = strings.Trim(query, "  ")
	for _, s := range db.synonyms {
		query = strings.Replace(query, s.Key, s.Value, 1)
	}

	var first, second, third string
	if strings.Contains(query, ",") {
		splitQuery := strings.Split(query, ",")
		switch len(splitQuery) {
		case 2:
			first = strings.Trim(splitQuery[0], " ")
			second = strings.Trim(splitQuery[1], " ")
		case 3:
			first = strings.Trim(splitQuery[0], " ")
			second = strings.Trim(splitQuery[1], " ")
			third = strings.Trim(splitQuery[2], " ")
		}
	} else {
		first = query
	}

	if first != "" && second != "" && third != "" {
		for _, c := range db.cities {
			if strings.EqualFold(first, c.Name) &&
				(strings.EqualFold(second, c.RegionCode) ||
					strings.EqualFold(second, c.RegionName)) &&
				(strings.EqualFold(third, c.CountryCode) ||
					strings.EqualFold(third, c.CountryName)) {
				return c.CountryCode
			}
		}
		for _, c := range db.cities {
			if first == c.Name &&
				(third == c.CountryCode || third == c.CountryName) {
				return c.CountryCode
			}
		}
	}

	if first != "" && second != "" && len(second) == 2 {
		for _, c := range db.cities {
			if strings.EqualFold(first, c.Name) &&
				strings.EqualFold(second, c.RegionCode) {
				return c.CountryCode
			}
		}
	}

	if first != "" && second != "" {
		for _, c := range db.cities {
			if strings.EqualFold(first, c.Name) &&
				(strings.EqualFold(second, c.RegionCode) ||
					strings.EqualFold(second, c.RegionName)) {
				return c.CountryCode
			}

			if strings.EqualFold(first, c.Name) &&
				(strings.EqualFold(second, c.CountryCode) ||
					strings.EqualFold(second, c.CountryName)) {
				return c.CountryCode
			}

			if (strings.EqualFold(first, c.RegionCode) ||
				strings.EqualFold(first, c.RegionName)) &&
				(strings.EqualFold(second, c.CountryCode) ||
					strings.EqualFold(second, c.CountryName)) {
				return c.CountryCode
			}
		}

		for _, c := range db.cities {
			if strings.EqualFold(second, c.CountryName) {
				return c.CountryCode
			}
		}

		for _, c := range db.cities {
			if strings.EqualFold(second, c.Name) {
				return c.CountryCode
			}
		}
	}

	if first != "" {
		for _, c := range db.cities {
			if strings.EqualFold(first, c.Name) ||
				strings.EqualFold(first, c.RegionName) ||
				strings.EqualFold(first, c.CountryName) {
				return c.CountryCode
			}
		}
	}

	return ""
}

func (db *Db) loadCities() error {
	file, err := os.Open(path.Join("data", "cities.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&db.cities)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) loadRegions() error {
	file, err := os.Open(path.Join("data", "states.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&db.regions)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) loadCountries() error {
	file, err := os.Open(path.Join("data", "countries.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&db.countries)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) loadSynonyms() error {
	file, err := os.Open(path.Join("data", "synonyms.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&db.synonyms)
	if err != nil {
		return err
	}

	return nil
}
