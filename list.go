//still developing it

package list

import (
	"fmt"
	"sort"
)

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

func (ls *listStruct) getType() string {
	if len(ls.types) == 1 {
		for k := range ls.types {
			return k
		}
	}
	return "none"
}

func (ls *listStruct) Sort() {
	listType := ls.getType()
	switch listType {
	case "int":
		{
			sort.Slice((*ls).list, func(i, j int) bool {
				return ls.list[i].(int) < ls.list[j].(int)
			})
		}
	case "string":
		{
			sort.Slice((*ls).list, func(i, j int) bool {
				return ls.list[i].(string) < ls.list[j].(string)
			})
		}
	case "float64":
		{
			sort.Slice((*ls).list, func(i, j int) bool {
				return ls.list[i].(float64) < ls.list[j].(float64)
			})
		}
	case "float32":
		{
			sort.Slice((*ls).list, func(i, j int) bool {
				return ls.list[i].(float32) < ls.list[j].(float32)
			})
		}
	case "none":
		{
			panic("Can't convert this list elements must be from the same type")
		}
	}

}

func (ls *listStruct) Clear() {
	ls.list = nil
	ls.length = 0
	ls.types = nil

}

func (ls *listStruct) Reverse() {
	tmp := List()
	for i := range ls.list {
		tmp.Append(ls.list[ls.length-1-i])
	}
	ls.list = tmp.list
}

//this while testing the package
// func main() {
// 	ls := List(15, 16, 99, 5, 8, 37, "Yes")
// 	ls.Append(56, 98, 74, 6, 34, 5, true, 22, 99)
// 	fmt.Println(ls)
// 	fmt.Println(Len(ls))
// 	fmt.Println(ls.types)
// 	ls.DelByValue("Yes")
// 	ls.DelByValue(true)
// 	fmt.Println(ls)
// 	fmt.Println(ls.types)
// 	fmt.Println(ls.getType())
// 	fmt.Println(ls.Index(34))
// 	// ls.Clear()
// 	ls.Sort()
// 	fmt.Println(ls)
// 	ls.Reverse()
// 	fmt.Println(ls)
// }
