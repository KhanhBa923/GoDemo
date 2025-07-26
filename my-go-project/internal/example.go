package internal

type Example struct {
    Name string
}

func NewExample(name string) *Example {
    return &Example{Name: name}
}

func (e *Example) DoSomething() string {
    return "Doing something with " + e.Name
}