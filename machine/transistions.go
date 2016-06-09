package machine

//Transistions represents allowed transistions
type Transistions map[string]map[string]bool //map[State(string)]map[State(string)]bool

//Allowed is the transistion allowed
func (t Transistions) Allowed(from, to string) bool {
	v := t[from][to]
	return v
}

//MapTransistion map transistion set to allowed state set
func (t Transistions) MapTransistion() {
	//map from > to > bool

}
