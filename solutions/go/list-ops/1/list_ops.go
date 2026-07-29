package listops

type IntList []int

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = s.Append(lst)
	}
	return s
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var lst IntList = []int{}
	for _, v := range s{
		if fn(v){
			lst = append(lst, v)
		}
	}
	return lst
}

func (s IntList) Length() int {
	var len int = 0
	for _, v := range s{
		len++
		v = v
	}
	return len
}

func (s IntList) Map(fn func(int) int) IntList {
	var lst IntList = []int {}
	for _,v := range s{
		lst = append(lst, fn(v))
	}
	return lst
}

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	var acc int = initial
	for i := 0 ; i < s.Length() ; i++ {
		acc = fn(acc, s[i])
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	var acc int = initial
	for i := s.Length() - 1 ; i >= 0 ; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Reverse() IntList {
	var sz int = s.Length()
	var lst IntList = []int{}
	for i := sz - 1 ; i >= 0 ; i--{
		lst = append(lst, s[i])
	}
	return lst
}