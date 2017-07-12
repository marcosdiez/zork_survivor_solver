package main

import "testing"

func TestSurvivorPointList(t *testing.T) {

	theList := NewSurvivorPointList()

	theList.AddPoint(NewSurvivorPointUnsafe(1, 2))
	theList.AddPoint(NewSurvivorPointUnsafe(3, 1))
	theList.AddPoint(NewSurvivorPointUnsafe(1, 00))
	theList.AddPoint(NewSurvivorPointUnsafe(0, 1))
	theList.AddPoint(NewSurvivorPointUnsafe(1, 1))
	theList.AddPoint(NewSurvivorPointUnsafe(0, 0))

	if theList.Length() != 6 {
		t.Errorf("Error, the list should have the size of 3, but it's size is %v", theList.Length())
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(1, 2)) {
		t.Errorf("Missing point in list")
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(3, 1)) {
		t.Errorf("Missing point in list")
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(1, 0)) {
		t.Errorf("Missing point in list")
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(0, 1)) {
		t.Errorf("Missing point in list")
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(1, 1)) {
		t.Errorf("Missing point in list")
	}

	if !theList.PointExists(NewSurvivorPointUnsafe(0, 0)) {
		t.Errorf("Missing point in list")
	}
}

func TestSurvivorPointListAddingPointsTwice(t *testing.T) {

	theList := NewSurvivorPointList()

	if theList.AddPoint(NewSurvivorPointUnsafe(1, 2)) != nil {
		t.Error("Error inserting point to list")
	}

	if theList.AddPoint(NewSurvivorPointUnsafe(3, 2)) != nil {
		t.Error("Error inserting point to list")
	}

	if theList.AddPoint(NewSurvivorPointUnsafe(0, 0)) != nil {
		t.Error("Error inserting point to list")
	}

	if theList.AddPoint(NewSurvivorPointUnsafe(1, 2)) == nil {
		t.Error("Error inserting point to list")
	}
}

func TestAvailableSteps(t *testing.T) {
	sb := NewSurvivorBoard()

	as := sb.AvailableSteps(NewSurvivorPointUnsafe(0, 0))

	if as.Length() != 2 {
		t.Error("Error in list size")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(1, 2)) {
		t.Error("Expected point not in list")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(2, 1)) {
		t.Error("Expected point not in list")
	}
}

func TestAvailableStepsB(t *testing.T) {
	sb := NewSurvivorBoard()

	as := sb.AvailableSteps(NewSurvivorPointUnsafe(1, 1))

	if as.Length() != 4 {
		t.Error("Error in list size")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(0, 3)) {
		t.Error("Expected point not in list")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(2, 3)) {
		t.Error("Expected point not in list")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(3, 0)) {
		t.Error("Expected point not in list")
	}

	if !as.PointExists(NewSurvivorPointUnsafe(3, 2)) {
		t.Error("Expected point not in list")
	}

}
