package expenses

import "fmt"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var rec []Record = []Record{}
	for _, v := range in {
		if predicate(v) {
			rec = append(rec, v)
		}
	}
	return rec
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64 = 0
	var res []Record = Filter(in, ByDaysPeriod(p))
	for _, v := range res {
		sum += v.Amount
	}
	return sum
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var res []Record = Filter(in, ByCategory(c))
	if len(res) == 0{
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return TotalByPeriod(res, p), nil
}
