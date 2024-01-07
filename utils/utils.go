package utils

import (
	"sort"
)

// map[key] -> [key]
func MapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapContains[M ~map[K]V, K comparable, V any](m M, key K) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// [d b a] [a b c] -> [a b c d]
func MergeUniqueSort(s1 []string, s2 []string) []string {
	// merge slices
	s3 := append(s1, s2...)
	if len(s3) == 0 {
		return s3
	}
	// sort strings
	sort.Strings(s3)
	// foreach compare to previous for delete duplicate element
	prev := 1
	for curr := 1; curr < len(s3); curr++ {
		if s3[curr-1] != s3[curr] {
			s3[prev] = s3[curr]
			prev++
		}
	}
	return s3[:prev]
}
