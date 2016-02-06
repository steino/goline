package goline

type CompletionHandler func(line []rune) [][]rune

type Completion struct {
	handler  CompletionHandler
	curIndex int
	curLine []rune
}

func (c *Completion) StartCompletion(l *GoLine) (bool, error) {
	c.curIndex = 0
	c.curLine = l.CurLine[:l.Len]
	l.AddHandler(CHAR_TAB, c.NextCompletion)
	c.NextCompletion(l)
	return false, nil
}

func (c *Completion) NextCompletion(l *GoLine) (bool, error) {
	completions := c.handler(c.curLine)
	if len(completions) == 0 {
		return false, nil
	}
	var line []rune
	if c.curIndex > len(completions)-1 {
		c.curIndex = 0
		line = c.curLine
	} else {
		line = completions[c.curIndex]
		c.curIndex++
	}

	out := make([]rune, MAX_LINE-len(line))
	l.CurLine = append(line, out...)
	l.Position = len(line)
	l.Len = len(line)
	return false, nil
}

func (l *GoLine) ResetCompletion() {
	l.AddHandler(CHAR_TAB, l.completion.StartCompletion)
}

func (l *GoLine) SetCompletionHandler(handler CompletionHandler) {
	l.completion.handler = handler
	l.AddHandler(CHAR_TAB, l.completion.StartCompletion)
}

func SetupCompletion(l *GoLine) {
	c := Completion{}

	var defaultCompletionHandler CompletionHandler = func(line []rune) [][]rune {
		return make([][]rune, 0)
	}

	c.handler = defaultCompletionHandler
	l.completion = c
	l.AddHandler(CHAR_TAB, c.StartCompletion)
}
