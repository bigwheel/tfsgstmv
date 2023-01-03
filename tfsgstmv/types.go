package tfsgstmv

type GeneratedFile struct {
	Moveds []Moved `hcl:"moved,block"`
}

type Moved struct {
	From string `hcl:"from"`
	To   string `hcl:"to"`
}
