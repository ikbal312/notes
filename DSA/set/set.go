package set

type Set struct {
	integerMap map[int]bool
}

// contstructor
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// add element
func (set *Set) AddElement(element int) {
	if !set.ContainElement(element) {
		set.integerMap[element] = true
	}
}

// contain element in set
func (set *Set) ContainElement(element int) bool {
	_, exist := set.integerMap[element]
	return exist
}

// delete element
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// intersectSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()
	for value := range set.integerMap {
		if anotherSet.ContainElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet

}

// Union
func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	for value := range set.integerMap {
		unionSet.AddElement(value)
	}
	for value := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
