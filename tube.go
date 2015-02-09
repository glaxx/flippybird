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
 	//"math/rand"
)

const (
	TUBEWIDTH = 4
)

type Tube struct {
	x      int
	offset int
	sprite [][]byte
}

func NewTube(offset int) *Tube {
	spr := make([][]byte, TUBEWIDTH)
	for x := 0; x != len(spr); x++ {
		spr[x] = make([]byte, Y)
		for y := 0; y != len(spr[0]); y++ {
			if y == offset || y == offset + 5 {
				spr[x][y] = 1
			} else if x == 0 || x == 3 || y == offset + 1 || y == offset + 2 || y == offset + 3 || y == offset + 4 {
				spr[x][y] = 0
			} else {
				spr[x][y] = 1
			}
		}
	}
	return &Tube{X, offset, spr}
}

func (t *Tube) Jump() {
	return
}

func (t *Tube) Tick() {
	t.x--
}

func (t *Tube) GetX() int {
	return t.x
}

func (t *Tube) GetY() int {
	return 0
}

func (t *Tube) GetOffset() int {
	return t.offset
}

func (t *Tube) GetHit() bool {
	return false
}

func (t *Tube) GetSprite() [][]byte {
	return t.sprite
}
