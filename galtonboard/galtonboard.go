package galtonboard

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)

type Bullet struct {
	Position float64
}

func(b Bullet) DropLeft() Bullet {
	var newPosition float64 = b.Position - 0.5
	return Bullet{Position: newPosition}
}

func(b Bullet) DropRight() Bullet {
	var newPosition float64 = b.Position + 0.5
	return Bullet{Position: newPosition}
}

func(b Bullet) GetPosition() float64 {
	return b.Position;
}

func BulletListFromlength(length int) []Bullet {
	var list []Bullet;

	for i := 0; i<length; i++ {
		list = append(list, Bullet{Position: 0})
	}

	return list
}

type Tray struct {
	Number int 
	Bullets []Bullet
}

func TrayListFromlength(length int) []Tray {
	var list []Tray;

	for i := 0; i<length; i++ {
		list = append(list, Tray{Number: i})
	}

	return list
}

func(t Tray) WithBullet(bullet Bullet) Tray {
	var list []Bullet;

	for i := 0; i<len(t.Bullets); i++ {
		list = append(list, t.Bullets[i])
	}
	list = append(list, bullet)
	return Tray{Number: t.Number, Bullets: list}
}

func(t Tray) GetNumber() int {
	return t.Number 
}

func(t Tray) GetNumberOfBullets() int {
	return len(t.Bullets) 
}

type DropPolicyInterface interface {
    direction() string
}

type LeftOrientedDropPolicy struct {}
type RightOrientedDropPolicy struct {}
type RandomDropPolicy struct {}
type LeftRightSwitchDropPolicy struct {
	LastPosition string
}

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

func(p *LeftRightSwitchDropPolicy) direction() string {
	if(p.LastPosition == "right") {
		p.LastPosition = "left";
		return "left"
	} else {
		p.LastPosition = "right";
		return "right"
	}
}

type GaltonBoard struct {
	Bullets []Bullet
	Trays []Tray
	DropPolicy DropPolicyInterface 
}

func(g GaltonBoard) DropBullets() {
	for i := 0; i < len(g.Bullets); i++ {
		var bullet Bullet = g.Bullets[i]
		for x:= 0; x < g.getNumberOfDrops(); x++ {
			if g.DropPolicy.direction() == "right" {
				bullet = bullet.DropRight()
			} else {
				bullet = bullet.DropLeft()
			}
		}
		g.updateResultingTray(bullet)
	}
}

func(g GaltonBoard) getNumberOfDrops() int {
	return len(g.Trays) -1;
}

func(g GaltonBoard) getMiddle() float64 {
	return float64(len(g.Trays) / 2);
}

func(g GaltonBoard) updateResultingTray(bullet Bullet) {
	var middle float64 = g.getMiddle()
	var position float64 = middle + bullet.GetPosition()
	var resultTray int = int(math.Floor(position));

	for x := 0; x < len(g.Trays); x++ {
		if g.Trays[x].GetNumber() == resultTray {
			g.Trays[x] = g.Trays[x].WithBullet(bullet);		
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
