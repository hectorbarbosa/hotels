package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpellComputers(t *testing.T) {
	assert.Equal(t, SpellComputers(-10), "invalid input", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(0), "0 компьютеров", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(1), "1 компьютер", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(3), "3 компьютера", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(15), "15 компьютеров", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(25), "25 компьютеров", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(2000), "2000 компьютеров", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(41), "41 компьютер", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(512), "512 компьютеров", "Строки должны быть идентичны")
	assert.Equal(t, SpellComputers(27511), "27511 компьютеров", "Строки должны быть идентичны")
}
