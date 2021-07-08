package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestIntSlice(t *testing.T) {
	var ids []int
	for i := 1; i < 150000; i++ {
		ids = append(ids, i)
	}
	defaultFor(ids)
	linqFor(ids)
}

const (
	size = 1000000
)

func Test111(t *testing.T) {

	type AutoGenerated struct {
		Age   int    `json:"age"`
		Name  string `json:"name"`
		Child []int  `json:"child"`
	}

	jsonStr1 := `{"age": 14,"name": "potter", "child":[1,2,3]}`
	a := AutoGenerated{}
	json.Unmarshal([]byte(jsonStr1), &a)
	aa := a.Child
	fmt.Println(aa)
	jsonStr2 := `{"age": 12,"name": "potter", "child":[3,4,5,7,8,9]}`
	json.Unmarshal([]byte(jsonStr2), &a)
	fmt.Println(aa)
}

func TestJSON(t *testing.T) {
	list1 := GetCompanyByCountry("USA")
	t.Log(fmt.Printf("%x", &list1))
	t.Log(list1)

	j, _ := jsoniter.MarshalToString(&list1)
	t.Log(j)

	var list11 []Company
	jsoniter.UnmarshalFromString(j, &list11)
	t.Log(list11)
}

func TestQueryCompany(t *testing.T) {

	list1 := GetCompanyByCountry("USA")
	t.Log(fmt.Printf("%x", &list1))
	t.Log(list1)

	list2 := GetCompanyByCountry("USA")
	t.Log(fmt.Printf("%x", &list2))
	t.Log(list2)

	name1 := GetCompanyByName("Microsoft")
	t.Log(fmt.Printf("%x", &name1))
	t.Log(name1)

	name2 := GetCompanyByName("Microsoft")
	t.Log(fmt.Printf("%x", &name2))
	t.Log(name2)
}

func BenchmarkSelectWhereFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		linq.Range(1, size).Select(func(i interface{}) interface{} {
			return -i.(int)
		}).Where(func(i interface{}) bool {
			return i.(int) > -100
		}).First()
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		input     interface{}
		predicate func(interface{}) bool
		expected  int
	}{
		{
			input: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i interface{}) bool {
				return i.(int) == 3
			},
			expected: 2,
		},
		{
			input: "sstr",
			predicate: func(i interface{}) bool {
				return i.(rune) == 'r'
			},
			expected: 3,
		},
		{
			input: "gadsgsadgsda",
			predicate: func(i interface{}) bool {
				return i.(rune) == 'z'
			},
			expected: -1,
		},
	}

	for _, test := range tests {
		index := linq.From(test.input).IndexOf(test.predicate)
		if index != test.expected {
			t.Errorf("From(%v).IndexOf() expected %v received %v", test.input, test.expected, index)
		}

		index = linq.From(test.input).IndexOfT(test.predicate)
		if index != test.expected {
			t.Errorf("From(%v).IndexOfT() expected %v received %v", test.input, test.expected, index)
		}
	}
}

type MyQuery linq.Query

func (q MyQuery) GreaterThan(threshold int) linq.Query {
	return linq.Query{
		Iterate: func() linq.Iterator {
			next := q.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if item.(int) > threshold {
						return
					}
				}
				return
			}
		},
	}
}
