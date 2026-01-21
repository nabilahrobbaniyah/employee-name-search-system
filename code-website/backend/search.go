package main

import "strings"

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

func sequentialIterative(data []Employee, target string) *Employee {
	for i := 0; i < len(data); i++ {
		if strings.EqualFold(data[i].Name, target) {
			return &data[i]
		}
	}
	return nil
}

func sequentialRecursive(data []Employee, target string, index int) *Employee {
	if index >= len(data) {
		return nil
	}
	if strings.EqualFold(data[index].Name, target) {
		return &data[index]
	}
	return sequentialRecursive(data, target, index+1)
}
