package galtonboard

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)

type Bullet struct {
	position float64
}

func(b Bullet) DropLeft() Bullet {
	var newPosition float64 = b.position - 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) DropRight() Bullet {
	var newPosition float64 = b.position + 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) GetPosition() float64 {
	return b.position;
}

func BulletListFromlength(length int) []Bullet {
	var list []Bullet;

	for i := 0; i<length; i++ {
		list = append(list, Bullet{position: 0})
	}

	return list
}

type Tray struct {
	number int 
	bullets []Bullet
}

func TrayListFromlength(length int) []Tray {
	var list []Tray;

	for i := 0; i<length; i++ {
		list = append(list, Tray{number: i})
	}

	return list
}

func(t Tray) WithBullet(bullet Bullet) Tray {
	var list []Bullet;

	for i := 0; i<len(t.bullets); i++ {
		list = append(list, t.bullets[i])
	}
	list = append(list, bullet)
	return Tray{number: t.number, bullets: list}
}

func(t Tray) GetNumber() int {
	return t.number 
}

func(t Tray) GetNumberOfBullets() int {
	return len(t.bullets) 
}

type DropPolicyInterface interface {
    direction() string
}

type LeftOrientedDropPolicy struct {}
type RightOrientedDropPolicy struct {}
type RandomDropPolicy struct {}

func(p LeftOrientedDropPolicy) direction() string {
	return "left"
}

func(p RightOrientedDropPolicy) direction() string {
	return "right"
}

func(p RandomDropPolicy) direction() string {
	rand.Seed(time.Now().UnixNano())
	if(rand.Intn(2) == 1) {
		return "left"
	} else {
		return "right"
	}
}

type GaltonBoard struct {
	bullets []Bullet
	trays []Tray
	dropPolicy DropPolicyInterface 
}

func(g GaltonBoard) DropBullets() {
	for i := 0; i < len(g.bullets); i++ {
		var bullet Bullet = g.bullets[i]
		for x:= 0; x < g.getNumberOfDrops(); x++ {
			if g.dropPolicy.direction() == "right" {
				bullet = bullet.DropRight()
			} else {
				bullet = bullet.DropLeft()
			}
		}
		g.updateResultingTray(bullet)
	}
}

func(g GaltonBoard) getNumberOfDrops() int {
	return len(g.trays) -1;
}

func(g GaltonBoard) getMiddle() float64 {
	return float64(len(g.trays) / 2);
}

func(g GaltonBoard) updateResultingTray(bullet Bullet) {
	var middle float64 = g.getMiddle()
	var position float64 = middle + bullet.GetPosition()
	var resultTray int = int(math.Floor(position));

	for x := 0; x < len(g.trays); x++ {
		if g.trays[x].GetNumber() == resultTray {
			g.trays[x] = g.trays[x].WithBullet(bullet);		
		}
	}
}

func Run() {
	bullets := BulletListFromlength(200)
	trays := TrayListFromlength(10)
	dropPolicy := RandomDropPolicy{}
	board := GaltonBoard{bullets, trays, dropPolicy}
	board.DropBullets()

	for x:= 0; x < len(trays); x++ {	
		for a := 0; a < trays[x].GetNumberOfBullets(); a++ {
			fmt.Print("0")
		}
		fmt.Print("\n")
	}
}
