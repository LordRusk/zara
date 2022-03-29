package modules

type Georona struct {
	ii, sig, pos int
}

func NewGeorona(interv, sig int) *Georona {
	return &Georona{
		ii:  interv,
		sig: sig,
	}
}

func (_ Georona) Run() (string, error) {

	return "", nil
}

func (g *Georona) Int() int {
	return g.ii
}

func (g *Georona) Sig() int {
	return g.sig
}

func (g *Georona) SetPos(pos int) {
	g.pos = pos
}

func (g *Georona) Pos() int {
	return g.pos
}

func (_ Georona) Name() string {
	return "Georona"
}
