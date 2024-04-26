package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonDivisors(t *testing.T) {
	s1 := []int{1}
	_, err := CommonDivisors(s1)
	assert.Equal(t, err, fmt.Errorf("invalid input %v\n", s1), "Должна быть ошибка")
}

func TestCommonDivisors2(t *testing.T) {
	s1 := []int{2, 3, 6}
	s2, _ := CommonDivisors([]int{42, 12, 18})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestCommonDivisors3(t *testing.T) {
	s1 := []int{2, 5, 10, 25, 50}
	s2, _ := CommonDivisors([]int{500, 50, 250, 1000})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestCommonDivisors4(t *testing.T) {
	s1 := []int{500, 50, -250, 1000}
	_, err := CommonDivisors(s1)
	assert.Equal(t, err, fmt.Errorf("invalid input %v\n", s1), "Должна быть ошибка")
}

func TestCommonDivisors5(t *testing.T) {
	s1 := []int{2}
	s2, _ := CommonDivisors([]int{2})
	if reflect.DeepEqual(s1, s2) != true {
		t.Errorf("Ожидалось %v, получили %v", s1, s2)
	}
}

func TestCommonDivisors6(t *testing.T) {
	s1 := []int{}
	_, err := CommonDivisors(s1)
	assert.Equal(t, err, fmt.Errorf("invalid input %v\n", s1), "Должна быть ошибка")
}
