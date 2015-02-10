/*
 *    Copyright (C) 2015 Stefan Luecke
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors: Stefan Luecke <glaxx@glaxx.net>
 */

package main

import (
	"math/rand"
)

type Player struct {
	y      int
	speed  int
	sprite [][]byte
}

func NewPlayer(y int, t bool) *Player {
	spr := [][]byte{
		{1},
	}
	if t == true {
		spr = [][]byte{
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{1, 1, 1, 1},
		}

	}

	return &Player{y, 1, spr}
}

func (p *Player) Tick() {
	p.y++
	if (p.y - p.speed) >= 0 {
		p.y -= p.speed
	}

	if p.speed >= -1 {
		p.speed--
	}
}

func (p *Player) Jump() {
	p.speed = 3 + rand.Intn(1)
}

func (p *Player) GetSprite() [][]byte {
	return p.sprite
}

func (p *Player) GetX() int {
	return (X / 2)
}

func (p *Player) GetY() int {
	return p.y
}

func (p *Player) GetOffset() int {
	return 0
}
