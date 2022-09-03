package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{Name: "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Chris", City: "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{Name: "Chris", Age: 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Nested fields",
			Input: Person{
				Name:    "Chris",
				Profile: Profile{Age: 33, City: "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Slice",
			Input: []Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "Arrays",
			Input: [2]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(s string) {
				got = append(got, s)
			})
			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}
