package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
}

type NestedIterator struct {
	elems []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var result []int
	dfs := func(nestedList []*NestedInteger) {
		for _, item := range nestedList {
			if item.IsInteger() {
				result = append(result, item.GetInteger())
			} else {
				dfs(item.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{result}
}

func (this *NestedIterator) Next() int {
	r := this.elems[0]
	this.elems = this.elems[1:]
	return r
}

func (this *NestedIterator) HasNext() bool {
	if len(this.elems) > 0 {
		return true
	} else {
		return false
	}
}

func main() {

}
