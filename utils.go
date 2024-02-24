package main

import (
	"fmt"
	"strings"
)

func parseCommand(commandWithArgs string) (string, []string) {
	parts := strings.Fields(commandWithArgs)
	command := parts[0]
	args := parts[1:]
	return command, args
}

func find[T any](arr []T, f func(T) bool) (T, error) {
	var zero T
	for _, value := range arr {
		if f(value) {
			return value, nil
		}
	}
	return zero, fmt.Errorf("no match found")
}

func mapArray[T any, U any](arr []T, f func(T) U) []U {
	var result []U
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}

func flatten[T any](slice [][]T) []T {
	var flatSlice []T
	for _, innerSlice := range slice {
		for _, value := range innerSlice {
			flatSlice = append(flatSlice, value)
		}
	}
	return flatSlice
}
