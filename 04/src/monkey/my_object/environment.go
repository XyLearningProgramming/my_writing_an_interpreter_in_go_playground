package my_object

type Environment struct {
	values map[string]Object
	outie  *Environment
}

func NewEnvironment() *Environment { return &Environment{values: map[string]Object{}, outie: nil} }

func NewEnclosedEnvironment(outie *Environment) *Environment {
	return &Environment{values: map[string]Object{}, outie: outie}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.values[name]
	if !ok && e.outie != nil {
		obj, ok = e.outie.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, value Object) Object {
	e.values[name] = value
	return value
}
