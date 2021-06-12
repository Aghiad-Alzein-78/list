//still developing it

package main

import "fmt"

//Make it as a struct to add more functionality
//types will be used later in Sorting
//list won't be able to sort unless it is
//from the same type
type listStruct struct {
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
func (ls *listStruct) Append(elements ...interface{}) {
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
func List(a ...interface{}) listStruct {
	ls := listStruct{}
	ls.Append(a...)
	return ls
}

//get the length of the list
func Len(ls listStruct) int {
	return ls.length
}

//use the Stringer interface to print the list
func (ls listStruct) String() string {
	str := "["
	for i := 0; i < Len(ls); i++ {
		if i == Len(ls)-1 {
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
func (ls *listStruct) DelByIndex(index int) {
	if index >= Len(*ls) {
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
func (ls *listStruct) Index(val interface{}) int {
	for idx, item := range ls.list {
		if item == val {
			return idx
		}
	}
	return -1
}

//delete the desired element "Just Once"
func (ls *listStruct) DelByValue(val interface{}) {
	pos := ls.Index(val)
	if pos != -1 {
		ls.DelByIndex(pos)
	}
}
func main() {
	ls := List(1, 2, 3, 4, 5, "Yes")
	ls.Append(56, 74, 34, true, 22)
	fmt.Println(ls)
	fmt.Println(Len(ls))
	fmt.Println(ls.types)
	ls.DelByIndex(5)
	ls.DelByValue(true)
	fmt.Println(ls)
	fmt.Println(ls.types)

}
