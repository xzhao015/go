package ch1

import (
	"reflect"
	"testing"
)

// TestVariable is for variable to test
func TestHello(t *testing.T) {
	want := "Hello, World!"
	if got := Hello(); got != want {
		t.Errorf("not equal")
	}

}

func TestVariable(t *testing.T) {
	var a *int
	if a != nil {
		t.Errorf("*int is not nil")
	}
	var b []int
	if b != nil {
		t.Errorf("[]]int is not nil")
	}
	var c map[string]int
	if c != nil {
		t.Errorf("map[string]int is not nil")
	}
	var d chan int
	if d != nil {
		t.Errorf("chan int is not nil")
	}
	var e func(string) int
	if e != nil {
		t.Errorf("func(string) int is not nil")
	}
	var j error
	if j != nil {
		t.Errorf("error int is not nil")
	}
	var g int
	if g != 0 {
		t.Errorf("int is not 0")
	}
	var h string
	if h != "" {
		t.Errorf("string init is not \"\"")
	}

}

func TestCont(t *testing.T) {
	const (
		a = iota
		b
		c
		d = "hello"
		f
		g
	)
	t.Log(a, b, c, d, f, g)
}

func TestList(t *testing.T) {
	var list1 [10]int
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	list1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	list1[0] = 1
	list1[3] = 2
	t.Log(list1)
	t.Log(balance)
	for index, value := range list1 {
		t.Log(index, value)
	}
	t.Log(len(list1))
	t.Log(list1[2:6])
}

func TestSlice(t *testing.T) {
	var slice1 []int
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(len(slice1))
	t.Log(slice1[2:6])
	t.Log(cap(slice1))
	slice1 = append(slice1, 2)
	t.Log(cap(slice1))
	t.Log(len(slice1))
	t.Log(slice1[len(slice1)-1])
	t.Log(cap(slice1))

	slice2 := make([]int, 12)
	t.Log(slice2)
	t.Log(cap(slice2))

	slice3 := make([]int, len(slice1), (cap(slice1))*2)
	copy(slice3, slice1)
	t.Log(slice1)
	t.Log(slice3)
	t.Log(cap(slice1))
	t.Log(cap(slice3))

}

func TestStruct(t *testing.T) {
	type structName struct {
		a int
		b int
		c string
		d []int
	}
	struct1 := structName{a: 4, b: 5, c: "123", d: []int{1, 2, 3}}
	t.Log(struct1)
	struct2 := structName{4, 5, "123", []int{1, 2, 3}}
	t.Log(struct2)
	t.Log(struct2.a)
	struct2.d = []int{3, 4, 5, 6}
	struct2.d = append(struct2.d, 7, 8, 9, 0)
	t.Log(struct2)
}

func TestMap(t *testing.T) {
	var map1 map[string]int
	map2 := make(map[string]int)
	if map2 == nil {
		t.Log("map2 is nil")
	} else {
		t.Log("map2 is not nil")
	}
	if map1 == nil {
		t.Log("map1 is nil")
	} else {
		t.Log("map1 is not nil")
	}

	//map1["France"] = 1  not init, cannot assign value

	map2["France"] = 1
	map2["Italy"] = 2
	map2["Japan"] = 3
	map2["India"] = 4
	t.Log(map2)
	//map loop
	for country, index := range map2 {
		t.Log(country, index)
	}
	//map check element
	index, ok := map2["UK"]
	if ok {
		t.Log(index)
	} else {
		t.Log("cannot get the index of UK")
	}
	//delete element
	delete(map2, "Japan")
	t.Log(map2)
	//map is referent
	map3 := map2
	map3["china"] = 100
	t.Log(map2)
	t.Log(map3)
	//map with struct
	type User struct {
		name       string
		occupation string
	}
	u1 := User{
		name:       "John Doe",
		occupation: "gardener",
	}
	u2 := User{
		name:       "Richard Roe",
		occupation: "driver",
	}
	u3 := User{
		name:       "Lucy Smith",
		occupation: "teacher",
	}
	users := map[int]User{
		1: u1,
		2: u2,
		3: u3,
	}
	t.Log(users)
	u1.name = "hello"
	t.Log(u1)
	t.Log(users[1].name)
	users[1] = u1
	t.Log(users)

}

func TestMap1(t *testing.T) {
	list1 := [3]int{1, 2, 3}
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := make(map[string]interface{})
	map2["France"] = 1
	map2["Italy"] = 2
	map2["Japan"] = 3
	map2["India"] = "abc"
	map2["China"] = []int{1, 2, 4}
	map2["abc"] = &list1
	t.Log(map2)
	map2["abc"] = map1
	t.Log(reflect.TypeOf(map1))
	map2["abc"].(map[string]int)["a"] = 8
	t.Log(map2)
	t.Log(list1)

}
