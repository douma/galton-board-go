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
