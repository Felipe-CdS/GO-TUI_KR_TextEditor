package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

    file, err := os.Open("test_file")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    b1 := make([]byte, 500)
    n1, err := file.Read(b1)
    if err != nil {
        panic(err)
    }
    n1++;
    
    var s tcell.Screen = NewScreen();
    defer quit(s);

    for {
        var screenX, screenY int = s.Size();

        s.Clear()
        drawBorder(s)
        drawText(s, 1, 1, screenX, screenY-1, fmt.Sprintf(string(b1)))
        s.ShowCursor(cursorPosX, cursorPosY)
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
			if ev.Rune() == 'j' {
                if cursorPosY < screenY-2 {
                    cursorPosY++; 
                }
			}
			if ev.Rune() == 'k' {
                if cursorPosY > 1 {
                    cursorPosY--;
                }
			}
			if ev.Rune() == 'h' {
                if cursorPosX > 1 {
                    cursorPosX--;
                }
			}
			if ev.Rune() == 'l' {
                if cursorPosX < screenX-2 {
                    cursorPosX++;
                }
			}
        }
    }
}


