package main

type Dictionary struct {
	items map[string]int
}

// Agrega o modifica un par clave valor
func (d *Dictionary) Set(clave string, valor int) {
	d.items[clave] = valor
}

// Busca una clave y devuelve su valor
func (d *Dictionary) Get(clave string) (int, bool) {
	valor, existe := d.items[clave]
	return valor, existe
}

// Elimina una clave del Dictionary
func (d *Dictionary) Delete(clave string) {
	delete(d.items, clave)
}

// Comprueba si una clave existe
func (d *Dictionary) Contains(clave string) bool {
	_, existe := d.items[clave]
	return existe
}

// Comprueba si el Dictionary está vacío
func (d *Dictionary) IsEmpty() bool {
	return len(d.items) == 0
}
