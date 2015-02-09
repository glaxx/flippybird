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
	"flag"
	"github.com/nsf/termbox-go"
	"log"
	"net"
	"time"
	"sync"
//	"fmt"
)

var wg sync.WaitGroup
var CONN *net.UDPConn

func main() {
	err := termbox.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer termbox.Close()
	var x = flag.Int("x", 80, "Lenght of your flipdot display")
	var y = flag.Int("y", 16, "Height of your flipdot display")
	var addr = flag.String("addr", "localhost:2323", "Address + Port of your flipdot display")
	flag.Parse()

	udpaddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Panicln(err)
	}
	conn, err := net.DialUDP(udpaddr.Network(), nil, udpaddr)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	CONN = conn
	send_buffer(GetLogo())
	time.Sleep(3 * time.Second)
	g := NewGame(*x, *y)
	g.AddGameObject(NewPlayer(8, false))
	//g.AddGameObject(NewTube(5))
	ch := make(chan int)
	wg.Add(1)
	go game(ch, g)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				select {
				case ch <- 0:
					wg.Wait()
					send_buffer(InttoMsg(*x,*y,g.GetPassedTubes()))
					return

				default:
					wg.Wait()
					send_buffer(InttoMsg(*x,*y,g.GetPassedTubes()))
					return 

				}
			case termbox.KeySpace:
				if (g.Ended == false) {
					ch <- 1
				}
				
			}
		case termbox.EventError:
			panic(ev.Err)
		default:
			if(g.Ended == true) {
				wg.Wait()
				send_buffer(InttoMsg(*x,*y,g.GetPassedTubes()))
				return
			}
		}
		if (g.Ended == true) {
			wg.Wait()
			send_buffer(InttoMsg(*x,*y,g.GetPassedTubes()))
			return
		}
	}
}

func game(msg chan int, g *Game) {
	defer wg.Done()
	for {
		time.Sleep(1000 / 13 * time.Millisecond) //TODO
		select {
		case val := <- msg:
			switch(val) {
			case 0:
				return
			case 1:
				g.Jump()
			}
		default:

		}
		if (g.Ended == true) {
			return
		}
		g.Draw()
		g.Tick()
		send_buffer(g.GetBoard())

	}
	
}

func send_buffer(buffer [][]byte) {
	/*for y := 0; y != len(buffer[0]); y++ {
		for x := 0; x != len(buffer); x++ {
			buffer[x][y] = 1
		}
	}*/
	msg := make([]byte, (len(buffer) * len(buffer[0]))/8)
	cnt := 0
	for y := 0; y != len(buffer[0]); y++ {
		for x := 0; x != len(buffer); x++ {
			msg[cnt/8] = msg[cnt/8] << 1
			msg[cnt/8] |= buffer[x][y]
			
			cnt++
		}
	}
	_, err := CONN.Write(msg)
	if err != nil {
		log.Panic(err)
	}
}
/*
func dbg(msg []byte) {
	for i := 0; i != len(msg); i++ {
		fmt.Printf("%b", msg[i])
	}
	fmt.Printf("\n")
}
*/