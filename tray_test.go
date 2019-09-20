package main

import (
	galtonboard "./galtonboard"
	"testing"
)

func TestShouldConstructRangeAndAssertIncrementalNumbering(t *testing.T) {
    
	var list []galtonboard.Tray = galtonboard.TrayListFromlength(10);

	if len(list) != 10 {
		t.Errorf("The list has an invalid length: %d, wanted: 10", len(list))
	}

	if(list[0].GetNumber() != 0) {
		t.Errorf("The list item has an invalid number %d, should be 0", list[0].GetNumber())
	}

	if(list[1].GetNumber() != 1) {
		t.Errorf("The list item has an invalid number %d, should be 1", list[1].GetNumber())
	}
}

func TestShouldAddBulletAndAddNewTray(t *testing.T) {
	var list []galtonboard.Tray = galtonboard.TrayListFromlength(10);
	var bullets []galtonboard.Bullet = galtonboard.BulletListFromlength(1);
	var tray galtonboard.Tray = list[0];
	var newTray galtonboard.Tray = tray.WithBullet(bullets[0]);

	if (newTray.GetNumberOfBullets() != 1) {
		t.Errorf("The new tray should have length 1, length %d given", newTray.GetNumberOfBullets())
	}

	if (tray.GetNumberOfBullets() != 0) {
		t.Errorf("The tray should have length 0, length %d given", tray.GetNumberOfBullets())
	}
}
