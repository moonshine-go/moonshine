package moonshine

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (p Params) Get(key string) (string, bool) {
	for _, param := range p {
		if param.Key == key {
			return param.Value, true
		}
	}
	return "", false
}

func (p Params) ByName(name string) (va string) {
	va, _ = p.Get(name)
	return
}
