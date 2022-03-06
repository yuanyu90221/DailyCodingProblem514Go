package findlcn

var itemExists = struct{}{}

type HashSet struct {
	Items map[int]struct{}
}

func NewSet(values ...int) *HashSet {
	set := &HashSet{Items: make(map[int]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *HashSet) Add(items ...int) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

func (set *HashSet) Remove(items ...int) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

func (set *HashSet) Contains(items ...int) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

func (set *HashSet) Size() int {
	return len(set.Items)
}

func (set *HashSet) Clear() {
	set.Items = make(map[int]struct{})
}

func (set *HashSet) Values() []int {
	values := make([]int, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}
