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

func ListFromlength(length int) []Bullet {
	var list []Bullet;

	for i := 1; i<=length; i++ {
		list = append(list, Bullet{position: 0})
	}

	return list
}
