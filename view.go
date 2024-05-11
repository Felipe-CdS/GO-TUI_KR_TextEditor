package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var (
    screenStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
    cursorStyle = tcell.CursorStyleBlinkingBlock
    cursorPosX =  1
    cursorPosY =  1
)

func NewScreen() tcell.Screen {
    s, err := tcell.NewScreen();

    if err != nil {
        log.Fatalf("%+v", err)
    }
    if err:= s.Init(); err != nil {
        log.Fatalf("%+v", err)
    }

    s.SetStyle(screenStyle)
    s.SetCursorStyle(cursorStyle);
    s.Clear()

    return s;
}

func quit(s tcell.Screen) {
    maybePanic := recover()
    s.Fini()
    if maybePanic != nil {
        panic(maybePanic)
    }
}

func drawBorder(s tcell.Screen){
    var screenX, screenY int = s.Size();

    for i := 1; i < screenX-1; i++ {
        s.SetContent(i, 0, tcell.RuneHLine, nil, screenStyle);
    }

    for i := 1; i < screenY-1; i++ {
        s.SetContent(0, i, tcell.RuneVLine, nil, screenStyle);
        s.SetContent(screenX-1, i, tcell.RuneVLine, nil, screenStyle);
    }

    for i := 1; i < screenX-1; i++ {
        s.SetContent(i, screenY-1, tcell.RuneHLine, nil, screenStyle);
    }
    s.SetContent(0, 0, tcell.RuneULCorner, nil, screenStyle);
    s.SetContent(0, screenY-1, tcell.RuneLLCorner, nil, screenStyle);
    s.SetContent(screenX-1, 0, tcell.RuneURCorner, nil, screenStyle);
    s.SetContent(screenX-1, screenY-1, tcell.RuneLRCorner, nil, screenStyle);
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
        if(r == '\n'){
            row++;
            col = x1;
        }
		s.SetContent(col, row, r, nil, screenStyle)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
