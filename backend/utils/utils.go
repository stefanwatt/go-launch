package utils

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func ParseCommand(commandWithArgs string) (string, []string) {
	parts := strings.Fields(commandWithArgs)
	command := parts[0]
	args := parts[1:]
	return command, args
}

func Find[T any](arr []T, f func(T) bool) (T, error) {
	var zero T
	for _, value := range arr {
		if f(value) {
			return value, nil
		}
	}
	return zero, fmt.Errorf("no match found")
}

func MapArray[T any, U any](arr []T, f func(T) U) []U {
	var result []U
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}

func Filter[T any](arr []T, f func(T) bool) []T {
	var result []T
	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func Flatten[T any](slice [][]T) []T {
	var flatSlice []T
	for _, innerSlice := range slice {
		flatSlice = append(flatSlice, innerSlice...)
	}
	return flatSlice
}

type MatchWithDistance[T any] struct {
	Value    T
	Distance int
}

func FuzzyFindObj[T any](searchTerm string, arr []T, propNames []string) []T {
	var matchStructs []MatchWithDistance[T]

	for _, prop := range propNames {
		props := MapArray(arr, func(obj T) string {
			v := reflect.ValueOf(obj)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				v = v.Elem() // This "converts" the pointer to the value it points to
			}
			return v.FieldByName(prop).String()
		})
		propMatches := fuzzy.RankFindNormalizedFold(searchTerm, props)
		for _, match := range propMatches {
			matchStructs = append(matchStructs, MatchWithDistance[T]{arr[match.OriginalIndex], match.Distance})
		}
	}

	sort.Slice(matchStructs, func(i, j int) bool {
		return matchStructs[i].Distance < matchStructs[j].Distance
	})

	var sortedMatches []T
	for _, matchStruct := range matchStructs {
		sortedMatches = append(sortedMatches, matchStruct.Value)
	}

	return sortedMatches
}
