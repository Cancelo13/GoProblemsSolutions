// Package census simulates a system used to collect census data.
package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	}
	if r.Address["street"] == "" {
		return false
	}
	return true
}

func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

func Count(residents []*Resident) int {
	var counter int = 0
	for _, res := range residents {
		if res.HasRequiredInfo() {
			counter++
		}
	}
	return counter
}

