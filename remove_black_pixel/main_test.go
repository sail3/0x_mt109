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
		t.Error("remove black pixels2 dont passed")
		t.Fail()
	}
}
func TestRemoveBlackPixels3(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	output := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	result := RemoveBlackPixels(input)
	if !reflect.DeepEqual(output, result) {
		t.Error("remove black pixels3 dont passed")
		t.Fail()
	}
}
func TestRemoveBlackPixels4(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1},
	}

	output := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1},
	}
	result := RemoveBlackPixels(input)
	if !reflect.DeepEqual(output, result) {
		t.Error("remove black pixels4 dont passed")
		t.Fail()
	}
}
func TestRemoveBlackPixels5(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1},
	}
	output := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1},
	}
	result := RemoveBlackPixels(input)
	if !reflect.DeepEqual(output, result) {
		t.Error("remove black pixels4 dont passed")
		t.Fail()
	}
}
