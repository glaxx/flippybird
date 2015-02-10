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
	//"log"
	"math/rand"
)

var X, Y int

type Game struct {
	passedTubes int
	nexttube    int
	gameObjs    []Gameobject
	gameboard   Gameboard
	Ended       bool
}

func NewGame(x, y int) *Game {
	gb := NewGameboard(x, y)
	X, Y = x, y
	return &Game{gameboard: *gb, nexttube: 5}
}

func (g *Game) AddGameObject(gobj Gameobject) {
	g.gameObjs = append(g.gameObjs, gobj)
}

func (g *Game) GetPassedTubes() int {
	return g.passedTubes
}

func (g *Game) Draw() {
	g.gameboard.Clear()
	if g.gameObjs[0].GetY() > Y {
		g.Ended = true
		return
	}
	for i := len(g.gameObjs) - 1; i >= 0; i-- {
		if g.gameObjs[i].GetY() < 0 {
			continue
		}
		spr := g.gameObjs[i].GetSprite()
		for x := 0; x != len(spr); x++ {
			for y := 0; y != len(spr[0]); y++ {
				if x+g.gameObjs[i].GetX() >= X || y+g.gameObjs[i].GetY() >= Y || x+g.gameObjs[i].GetX() <= 0 {
					continue
				}
				if spr[x][y] != 0 {
					g.gameboard.SetPixel(x+g.gameObjs[i].GetX(), y+g.gameObjs[i].GetY(), spr[x][y])
				}
			}
		}
	}
}

func (g *Game) GetBoard() [][]byte {
	return g.gameboard.GetBoard()
}

func (g *Game) Jump() {
	for _, obj := range g.gameObjs {
		obj.Jump()
	}
}

func (g *Game) Tick() {
	for i := 0; i != len(g.gameObjs); i++ {
		if g.gameboard.Collision == true {
			g.Ended = true
			return
		}
		g.gameObjs[i].Tick()
		/*
			if g.gameObjs[i].GetX() < -TUBEWIDTH {
				temp := append(g.gameObjs[i+1:])
				temp2 := append(g.gameObjs[:i])
				g.gameObjs = append(temp)
				for j := 0; j != len(temp2); j++ {
					g.gameObjs = append(g.gameObjs, temp2[j])
				}
				break
			}*/
	}
	if g.nexttube <= 0 {
		g.nexttube += 9 + rand.Intn(23) + TUBEWIDTH
		offset := g.gameObjs[len(g.gameObjs)-1].GetOffset()
		offset += rand.Intn(13) - 5 // Range -4 - 5
		if offset >= Y-5 || offset < 0 {
			offset = g.gameObjs[len(g.gameObjs)-1].GetOffset()
		}
		g.AddGameObject(NewTube(offset))
	}
	g.nexttube--
	for i := 1; i != len(g.gameObjs); i++ {
		if g.gameObjs[i].GetX() == (X/2)-(TUBEWIDTH/2) {
			g.passedTubes++
		}
	}
}

type Gameobject interface {
	Tick()
	GetSprite() [][]byte
	GetX() int
	GetY() int
	Jump()
	GetOffset() int
}
