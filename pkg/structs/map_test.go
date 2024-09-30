package structs

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_IntToString(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []string{"1", "2", "3", "4"}
	result := Map(input, func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	assert.Equal(t, expected, result)
}

func TestMap_StringToInt(t *testing.T) {
	input := []string{"1", "2", "3", "4"}
	expected := []int{1, 2, 3, 4}
	result := Map(input, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
	assert.Equal(t, expected, result)
}

func TestMap_EmptyInput(t *testing.T) {
	input := []int{}
	expected := []string{}
	result := Map(input, func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	assert.Equal(t, expected, result)
}
