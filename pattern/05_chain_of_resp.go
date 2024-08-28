package pattern

import "fmt"

// Processor — интерфейс обработчика.
type Processor interface {
	Process(Request)
	SetNext(Processor)
}

type Kind int

const (
	Urgent Kind = 1 << iota
	Special
	Valuable
)

// Request описывает поля запроса.
type Request struct {
	Kind Kind
	Data string
}

// Printer — обработчик.
type Printer struct {
	next Processor
}

func (p *Printer) Process(r Request) {
	fmt.Printf("Printer: %s\n", r.Data)
	if p.next != nil {
		p.next.Process(r)
	}
}

func (p *Printer) SetNext(next Processor) {
	p.next = next
}

// Saver — обработчик.
type Saver struct {
	next Processor
}

func (s *Saver) Process(r Request) {
	// обрабатывает не все запросы
	if r.Kind&(Valuable|Special) != 0 {
		fmt.Printf("Saver: %s\n", r.Data)
		// сохраняем состояние
		// ...
	}
	if s.next != nil {
		s.next.Process(r)
	}
}

func (s *Saver) SetNext(next Processor) {
	s.next = next
}

// Logger — обработчик.
type Logger struct {
	next Processor
}

func (l *Logger) Process(r Request) {
	if r.Kind&Urgent != 0 {
		fmt.Printf("Logger: %s\n", r.Data)
		// записываем в лог
		// ...
	}
	if l.next != nil {
		l.next.Process(r)
	}
}

func (l *Logger) SetNext(next Processor) {
	l.next = next
}

func main() {
	p := new(Printer)
	l := new(Logger)
	l.SetNext(p)
	s := new(Saver)
	s.SetNext(l)
	s.Process(Request{0, "Average"})
	s.Process(Request{Valuable, "Do not forget"})
	s.Process(Request{Urgent | Special, "Alert!!!"})
}
