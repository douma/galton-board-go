package main

import (
	galtonboard "./galtonboard"
	"testing"
)

func TestBulletShouldBeConstructedFromLength(t *testing.T) {
    
	var list []galtonboard.Bullet = galtonboard.ListFromlength(10);

	if len(list) != 10 {
		t.Errorf("The list has an invalid length: %d, wanted: 10", len(list))
	}
}

func TestShouldDropLeftAndReturnNewBullet(t *testing.T) {
	var list []galtonboard.Bullet = galtonboard.ListFromlength(1);
	var bullet galtonboard.Bullet = list[0];
	var newBullet galtonboard.Bullet = bullet.DropLeft();

	if(newBullet.GetPosition() == bullet.GetPosition()) {
		t.Errorf("Bullets should defer")
	}
	if(newBullet.GetPosition() != -0.5) {
		t.Errorf("Incorrect value %f, should be -0.5", newBullet.GetPosition())
	}
}

func TestShouldDropRightAndReturnNewBullet(t *testing.T) {
	var list []galtonboard.Bullet = galtonboard.ListFromlength(1);
	var bullet galtonboard.Bullet = list[0];
	var newBullet galtonboard.Bullet = bullet.DropRight();

	if(newBullet.GetPosition() == bullet.GetPosition()) {
		t.Errorf("Bullets should defer")
	}
	if(newBullet.GetPosition() != 0.5) {
		t.Errorf("Incorrect value %f, should be -0.5", newBullet.GetPosition())
	}
}