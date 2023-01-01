package simple_parser

type Env struct {
	table map[string]*Symbol
	prev  *Env
}

func NewEnv(p *Env) *Env {
	return &Env{
		table: make(map[string]*Symbol),
		prev:  p,
	}
}

func (e *Env) Put(s string, sym *Symbol) {
	e.table[s] = sym
}

func (e *Env) Get(s string) *Symbol {
	for env := e; env != nil; env = e.prev {
		found, ok := env.table[s]
		if ok {
			return found
		}
	}

	return nil
}
