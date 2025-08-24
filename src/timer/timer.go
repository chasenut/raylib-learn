package timer

import rl "github.com/gen2brain/raylib-go/raylib"

type timerInterface interface {
	float32
	float32
	GetStatus() bool
	Init(float32)
}

type Timer struct {
	lifeTime float32
	initTime float32
}

type Repeater struct {
	lifeTime float32
	initTime float32
}

func (t *Timer) Init(lifeTime float32) {
	t.initTime = float32(rl.GetTime())
	t.lifeTime = lifeTime
}	

func (r *Repeater) Init(lifeTime float32) {
	r.initTime = float32(rl.GetTime())
	r.lifeTime = lifeTime
}	

func (t *Timer) GetStatus() bool {
	if float32(rl.GetTime()) - t.initTime > t.lifeTime {
		return true
	}
	return false
}

func (r *Repeater) GetStatus() bool {
	if float32(rl.GetTime()) - r.initTime > r.lifeTime {
		r.initTime = float32(rl.GetTime())
		return true
	}
	return false
}
