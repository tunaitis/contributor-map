package location

import (
	"testing"
)

func TestDbSearch(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{
			"Vilnius, Lithuania", "LT",
		},
		{
			"New York", "US",
		},
		{
			"New York City", "US",
		},
		{
			"Amsterdam, NL", "NL",
		},
		{
			"Amsterdam, Netherlands", "NL",
		},
		{
			"Honolulu, HI", "US",
		},
		{
			"Victoria, BC, Canada", "CA",
		},
		{
			"Victoria, British Columbia, Canada", "CA",
		},
		{
			"BC, Canada", "CA",
		},
		{
			"Richmond, CA", "US",
		},
		{
			"Paris", "FR",
		},
		{
			"tokyo", "JP",
		},
		{
			"Krakow/Wroclaw, Poland", "PL",
		},
		{
			"Bangalore", "IN",
		},
		{
			"SF", "US",
		},
	}

	db, err := NewDb()
	if err != nil {
		panic(err)
	}

	for _, table := range tables {
		result := db.Search(table.input)
		if result != table.expected {
			t.Errorf("with input '%s' expected '%s', got '%s'", table.input, table.expected, result)
		}
	}
}
