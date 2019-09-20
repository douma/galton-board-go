package main

import (
	galtonboard "./galtonboard"
	"testing"
)

func TestDropAndReturnAllBulletsInLeftTray(t *testing.T) {
    
	bullets := galtonboard.BulletListFromlength(10)
	trays := galtonboard.TrayListFromlength(10)
	dropPolicy := galtonboard.LeftOrientedDropPolicy{}
	board := galtonboard.GaltonBoard{bullets, trays, dropPolicy}
	board.DropBullets()

	if trays[0].GetNumberOfBullets() != 10 {
		t.Errorf("The list has an invalid length: %d, wanted: 10", trays[0].GetNumberOfBullets())
	}
}

func TestDropAndReturnAllBulletsInRightTray(t *testing.T) {
    
	bullets := galtonboard.BulletListFromlength(10)
	trays := galtonboard.TrayListFromlength(10)
	dropPolicy := galtonboard.RightOrientedDropPolicy{}
	board := galtonboard.GaltonBoard{bullets, trays, dropPolicy}
	board.DropBullets()

	if trays[9].GetNumberOfBullets() != 10 {
		t.Errorf("The list has an invalid length: %d, wanted: 10", trays[0].GetNumberOfBullets())
	}
}

func TestDropAndReturnAllBulletsInMiddleTray(t *testing.T) {
    
	bullets := galtonboard.BulletListFromlength(10)
	trays := galtonboard.TrayListFromlength(3)
	dropPolicy := galtonboard.LeftRightSwitchDropPolicy{}
	board := galtonboard.GaltonBoard{bullets, trays, dropPolicy}
	board.DropBullets()

	if trays[1].GetNumberOfBullets() != 10 {
		t.Errorf("The list has an invalid length: %d, wanted: 10", trays[1].GetNumberOfBullets())
	}
}
