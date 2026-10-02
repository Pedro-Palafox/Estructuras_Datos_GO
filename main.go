package main

import "fmt"

func main() {

	//use stack agregamos elementos y comprobamos LIFO
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack:", stack.items)

	// Consultamos y eliminamos el elemento del tope
	valor, existe := stack.Peek()
	fmt.Println("Stack Peek:", valor, "Existe:", existe)

	valor, existe = stack.Pop()
	fmt.Println("Stack Pop:", valor, "Existe:", existe)

	fmt.Println("Stack después de Pop:", stack.items)

	// agregamos elementos y comprobamos FIFO
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Queue:", queue.items)

	// Consultamos y eliminamos el primer elemento
	valor, existe = queue.Front()
	fmt.Println("Queue Front:", valor, "Existe:", existe)

	valor, existe = queue.Dequeue()
	fmt.Println("Queue Dequeue:", valor, "Existe:", existe)

	fmt.Println("Queue después de Dequeue:", queue.items)

	// agregamos y consultamos pares clave-valor
	dictionary := Dictionary{
		items: make(map[string]int),
	}

	dictionary.Set("Pedro", 100)
	dictionary.Set("Juan", 95)

	fmt.Println("Dictionary:", dictionary.items)

	// Buscamos una clave y comprobamos si existe
	valor, existe = dictionary.Get("Pedro")
	fmt.Println("Dictionary Get:", valor, "Existe:", existe)

	fmt.Println("Contains Pedro:", dictionary.Contains("Pedro"))

	// Eliminamos una clave y comprobamos si está vacío
	dictionary.Delete("Juan")

	fmt.Println("Dictionary después de borrarlo:", dictionary.items)

	fmt.Println("Dictionary IsEmpty:", dictionary.IsEmpty())
}
