package characters

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Reputation struct {
	value float32
}

type inventory struct {
	inventory [5]string //Placeholder for more complex stuff in future
	hand      int       //Reputation/Suspicion only changes if this variable points to contraband while NPC is nearby
}

type clothing struct {
	name string //TODO: Include powerups later
}

type being struct {
	name       string
	reputation Reputation
	position   rl.Vector2
}

type Player struct {
	being
	inventory
	clothing
}

type NPC struct {
	being
	routine []string
	isAlive bool
}

type Rival struct {
	NPC //TODO: Include more complex behavior later
}

type Senpai struct {
	NPC
	suspicion float32
}

func (r *Reputation) Value() float32 {
	return r.value
}

func (r *Reputation) ChangeReputation(delta float32) {
	r.value += delta
	if r.value >= 1 {
		r.value = 1
	} else if r.value < -1 {
		r.value = -1
	}
}

func (i *inventory) WhatsInHand() string {
	if i.hand < 5 && i.hand >= 0 {
		return i.inventory[i.hand]
	} else {
		return ""
	}
}

func (s *Senpai) GetSuspicious(delta float32) {
	s.suspicion += delta
	if s.suspicion > 1 {
		s.suspicion = 1
	} else if s.suspicion <= 0 {
		s.suspicion = 0
	}
}
