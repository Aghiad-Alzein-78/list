//still developing it

package main

import "fmt"

//Make it as a struct to add more functionality
//types will be used later in Sorting
//list won't be able to sort unless it is
//from the same type
type List struct {
	list   listAyuda
	length int
	types  typeMap
}

type typeMap map[string]int

//interface slice will help to get multi types
//like the Python list
type listAyuda []interface{}

//check if the type map has a type inside
func (tm typeMap) contains(item string) bool {
	for k := range tm {
		if k == item {
			return true
		}
	}
	return false
}

//append to the list
//will append multiple elements in one command
func (ls *List) Append(elements ...interface{}) {
	if ls.length == 0 {
		ls.types = make(map[string]int)
	}
	for _, item := range elements {
		ls.list = append(ls.list, item)
		ls.length++
		entryType := fmt.Sprintf("%T", item)
		if ls.types.contains(entryType) {
			ls.types[entryType]++
		} else {
			ls.types[entryType] = 1
		}
	}
}

//initialize list
func list(a ...interface{}) List {
	ls := List{}
	ls.Append(a...)
	return ls
}

//get the length of the list
func len(ls List) int {
	return ls.length
}

//use the Stringer interface to print the list
func (ls List) String() string {
	str := "["
	for i := 0; i < len(ls); i++ {
		if i == len(ls)-1 {
			str += fmt.Sprintf("%v", ls.list[i])
		} else {
			str += fmt.Sprintf("%v,", ls.list[i])
		}
	}
	str += "]"
	return str
}

//delete element from the list by index
//will decrease the size
//and check the types
func (ls *List) delByIndex(index int) {
	if index >= len(*ls) {
		panic(fmt.Sprintln("The index is bigger than the list size"))
	}
	typeToRemove := fmt.Sprintf("%T", ls.list[index])
	ls.list = append(ls.list[:index], ls.list[index+1:]...)
	ls.length--
	if ls.types[typeToRemove] == 1 {
		delete(ls.types, typeToRemove)
	} else {
		ls.types[typeToRemove]--
	}
}

//get the index of an element
func (ls *List) findIndex(val interface{}) int {
	for idx, item := range ls.list {
		if item == val {
			return idx
		}
	}
	return -1
}

//delete the desired element "Just Once"
func (ls *List) delByValue(val interface{}) {
	pos := ls.findIndex(val)
	if pos != -1 {
		ls.delByIndex(pos)
	}
}
func main() {
	ls := list(1, 2, 3, 4, 5, "Yes")
	ls.Append(56, 74, 34, true, 22)
	fmt.Println(ls)
	fmt.Println(len(ls))
	fmt.Println(ls.types)
	ls.delByIndex(5)
	ls.delByValue(true)
	fmt.Println(ls)
	fmt.Println(ls.types)

}
