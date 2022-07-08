package logger

type Text string

const reset = "\033[0m"

func (t *Text) Red() *Text          { *t = Text(string(*t) + "\033[31m "); return t }
func (t *Text) Green() *Text        { *t = Text(string(*t) + "\033[32m "); return t }
func (t *Text) Yellow() *Text       { *t = Text(string(*t) + "\033[33m "); return t }
func (t *Text) Blue() *Text         { *t = Text(string(*t) + "\033[34m "); return t }
func (t *Text) Purple() *Text       { *t = Text(string(*t) + "\033[35m "); return t }
func (t *Text) Cyan() *Text         { *t = Text(string(*t) + "\033[36m "); return t }
func (t *Text) White() *Text        { *t = Text(string(*t) + "\033[37m "); return t }
func (t *Text) Bold() *Text         { *t = Text(string(*t) + "\033[1m "); return t }
func (t *Text) Italic() *Text       { *t = Text(string(*t) + "\033[3m "); return t }
func (t *Text) Underline() *Text    { *t = Text(string(*t) + "\033[4m "); return t }
func (t *Text) Blink() *Text        { *t = Text(string(*t) + "\033[5m "); return t }
func (t *Text) Inverse() *Text      { *t = Text(string(*t) + "\033[7m "); return t }
func (t *Text) Hide() *Text         { *t = Text(string(*t) + "\033[8m "); return t }
func (t *Text) Strike() *Text       { *t = Text(string(*t) + "\033[9m "); return t }
func (t *Text) Frame() *Text        { *t = Text(string(*t) + "\033[51m "); return t }
func (t *Text) Overline() *Text     { *t = Text(string(*t) + "\033[53m "); return t }
func (t *Text) Encircled() *Text    { *t = Text(string(*t) + "\033[54m "); return t }
func (t *Text) Text(s string) *Text { *t = Text((string(*t)) + string(s) + reset); return t }
