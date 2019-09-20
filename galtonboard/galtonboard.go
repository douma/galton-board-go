package galtonboard

type Bullet struct {
	position float32
}

func(b Bullet) dropLeft() Bullet {
	var newPosition float32 = b.position - 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) dropRight() Bullet {
	var newPosition float32 = b.position + 0.5
	return Bullet{position: newPosition}
}

func(b Bullet) getPosition() float32 {
	return b.position;
}

func ListFromlength(length int) []Bullet {
	var list []Bullet;

	for i := 1; i<=length; i++ {
		list = append(list, Bullet{position: 0})
	}

	return list
}
