package goline

import "fmt"

type History struct {
	history  [][]rune
	curIndex int
}

func (h *History) PreviousHistory(l *GoLine) (bool, error) {
	if h.curIndex > 0 {
		l.CurLine = h.history[h.curIndex-1]
		l.Position = len(l.CurLine)
		l.Len = len(l.CurLine)
		h.curIndex--
	}
	return false, nil
}

func (h *History) NextHistory(l *GoLine) (bool, error) {
	if h.curIndex < len(h.history)-1 {
		l.CurLine = h.history[h.curIndex+1]
		l.Position = len(l.CurLine)
		l.Len = len(l.CurLine)
		h.curIndex++
	} else if h.curIndex == len(h.history)-1 {
		return DeleteLine(l)
	}
	return false, nil
}

func (h *History) AddLine(line []rune) {
	h.history = append(h.history, line)
	h.curIndex = len(h.history)
	fmt.Println(len(h.history))
}

func (h *History) HistoryFinish(l *GoLine) (bool, error) {
	h.AddLine(l.CurLine[:l.Len])
	return Finish(l)
}

func SetupHistory(l *GoLine) {
	h := History{}

	l.AddHandler(CHAR_CTRLP, h.PreviousHistory)
	l.AddHandler(CHAR_CTRLN, h.NextHistory)
	l.AddHandler(ESCAPE_UP, h.PreviousHistory)
	l.AddHandler(ESCAPE_DOWN, h.NextHistory)

	// Overwrite any previous definition
	l.AddHandler(CHAR_ENTER, h.HistoryFinish)
}