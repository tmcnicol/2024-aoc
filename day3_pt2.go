//go:generate stringer -type=itemType
package aoc

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func day3Part2(input string) int {
	data := readfile(input)
	raw := parseData(data)

	result := 0
	doState := true

	l := newLexer(raw)
	go l.run()

	for i := range l.items {
		switch i.typ {
		case itemDo:
			doState = true
		case itemDont:
			doState = false
		case itemMul:
			if !doState {
				continue
			}
			result += evalMul(i)

		default:
			panic(i)
		}

	}
	return result
}

func evalMul(i item) int {
	if i.typ != itemMul {
		panic(i.typ)
	}

	operands := strings.TrimPrefix(i.value, "mul")
	operands = strings.Replace(operands, "(", "", -1)
	operands = strings.Replace(operands, ")", "", -1)
	args := strings.Split(operands, ",")

	if len(args) != 2 {
		panic(i)
	}

	arg1, arg2 := must(strconv.Atoi(args[0])), must(strconv.Atoi(args[1]))
	return arg1 * arg2
}

type itemType int

type item struct {
	typ   itemType
	value string
}

func (i item) String() string {
	return i.typ.String() + " " + i.value
}

const (
	itemMul itemType = iota
	itemDo
	itemDont
)

const eof = -1

type lexer struct {
	input string
	start int
	pos   int
	width int
	items chan item
}

func newLexer(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}

	return l
}

func (l *lexer) run() {
	for state := lexOP; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.items <- item{
		typ:   t,
		value: l.input[l.start:l.pos],
	}
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, width := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = width
	l.pos += l.width

	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

type stateFn func(*lexer) stateFn

func lexOP(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			return nil
		case r == 'm':
			l.backup()
			return lexMul
		case r == 'd':
			l.backup()
			return lexDo
		default:
			l.ignore()
		}
	}
}

func lexMul(l *lexer) stateFn {
	// cheating with regex
	if !strings.HasPrefix(l.input[l.pos:], "mul") {
		l.next() // consume the character so we don't loop forever
		l.ignore()
		return lexOP
	}
	l.pos += len("mul")

	// read the next digit
OPERANDS:
	for {
		switch r := l.next(); {
		case r == eof:
			return nil
		case '0' <= r && r <= '9':
			continue
		case r == ',' || r == '(':
			continue
		case r == ')':
			break OPERANDS
		default:
			l.ignore()
			return lexOP
		}
	}

	l.emit(itemMul)

	return lexOP
}

func lexDo(l *lexer) stateFn {
	if strings.HasPrefix(l.input[l.pos:], "don't") {
		l.pos += len("don't")
		l.emit(itemDont)
		return lexOP
	}

	if strings.HasPrefix(l.input[l.pos:], "do") {
		l.pos += len("do")
		l.emit(itemDo)
		return lexOP
	}
	l.ignore()
	return lexOP
}
