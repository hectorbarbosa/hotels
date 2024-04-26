package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetPrimesRange(t *testing.T) {
	s1 := []int{11, 13, 17, 19}
	s2, _ := GetPrimesRange([]int{11, 20})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestGetPrimesRange2(t *testing.T) {
	s1 := []int{11}
	s2, _ := GetPrimesRange([]int{11})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestGetPrimesRange3(t *testing.T) {
	s := []int{}
	_, err := GetPrimesRange(s)
	assert.Equal(t, err, fmt.Errorf("invalid input %v\n", s), "Должна быть ошибка")
}

func TestGetPrimesRange4(t *testing.T) {
	s1 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	s2, _ := GetPrimesRange([]int{2, 30})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestGetPrimesRange5(t *testing.T) {
	s := []int{-20, -10}
	_, err := GetPrimesRange(s)
	assert.Equal(t, err, fmt.Errorf("invalid input %v\n", s), "Должна быть ошибка")
}
