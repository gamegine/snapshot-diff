package utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestMapKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	want := []string{"a", "b", "c"}
	var got = MapKeys(m)
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMergeUniqueSort(t *testing.T) {
	a := []string{"a", "c", "e"}
	b := []string{"a", "b", "d", "e"}
	want := []string{"a", "b", "c", "d", "e"}
	var got = MergeUniqueSort(a, b)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMergeUniqueSortEmpty(t *testing.T) {
	a := []string{}
	b := []string{}
	want := []string{}
	var got = MergeUniqueSort(a, b)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
