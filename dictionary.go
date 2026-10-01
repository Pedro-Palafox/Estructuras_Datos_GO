package main

type Dictionary struct {
	items map[string]int
}

func (d *Dictionary) Set(clave string, valor int) {
	d.items[clave] = valor
}

func (d *Dictionary) Get(clave string) (int, bool) {
	valor, existe := d.items[clave]
	return valor, existe
}

func (d *Dictionary) Delete(clave string) {
	delete(d.items, clave)
}

func (d *Dictionary) Contains(clave string) bool {
	_, existe := d.items[clave]
	return existe
}

func (d *Dictionary) IsEmpty() bool {
	return len(d.items) == 0
}
