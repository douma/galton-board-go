package galtonboard

type Bullet struct {
	position float32
}

func(b Bullet) DropLeft() Bullet {
	var newPosition float32 = b.position - 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) DropRight() Bullet {
	var newPosition float32 = b.position + 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) GetPosition() float32 {
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

	for i := 0; i<len(list); i++ {
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