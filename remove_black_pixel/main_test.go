package main

import (
	"reflect"
	"testing"
)

func TestRemoveBlackPixels(t *testing.T) {
	input := [][]int{
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
	output := [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	result := RemoveBlackPixels(input)
	if !reflect.DeepEqual(output, result) {
		t.Error("remove black pixels dont passed")
		t.Fail()
	}
}

func TestRemoveBlackPixels2(t *testing.T) {
	input := [][]int{
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
	output := [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1},
	}

	result := RemoveBlackPixels(input)
	if !reflect.DeepEqual(output, result) {
		t.Error("remove black pixels dont passed")
		t.Fail()
	}
}
