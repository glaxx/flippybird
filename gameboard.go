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

import ()

type Gameboard struct {
	board     [][]byte
	Collision bool
}

func NewGameboard(x, y int) *Gameboard {
	board := make([][]byte, x)
	for i := 0; i != x; i++ {
		board[i] = make([]byte, y)
	}
	return &Gameboard{board, false}
}

func (g *Gameboard) Clear() {
	g.board = make([][]byte, X)
	for i := 0; i != X; i++ {
		g.board[i] = make([]byte, Y)
	}
}

func (g *Gameboard) SetPixel(x, y int, b byte) {
	if g.board[x][y] == 1 && b == 1 {
		g.Collision = true
	}
	g.board[x][y] = b
}

func (g *Gameboard) GetBoard() [][]byte {
	return g.board
}
