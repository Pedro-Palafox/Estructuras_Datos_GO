package main

import "fmt"

func main() {

	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack:", stack.items)

	valor, existe := stack.Peek()
	fmt.Println("Stack Peek:", valor, "Existe:", existe)

	valor, existe = stack.Pop()
	fmt.Println("Stack Pop:", valor, "Existe:", existe)

	fmt.Println("Stack después de Pop:", stack.items)

	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Queue:", queue.items)

	valor, existe = queue.Front()
	fmt.Println("Queue Front:", valor, "Existe:", existe)

	valor, existe = queue.Dequeue()
	fmt.Println("Queue Dequeue:", valor, "Existe:", existe)

	fmt.Println("Queue después de Dequeue:", queue.items)

	dictionary := Dictionary{
		items: make(map[string]int),
	}

	dictionary.Set("Pedro", 100)
	dictionary.Set("Juan", 95)

	fmt.Println("Dictionary:", dictionary.items)

	valor, existe = dictionary.Get("Pedro")
	fmt.Println("Dictionary Get:", valor, "Existe:", existe)

	fmt.Println("Contains Pedro:", dictionary.Contains("Pedro"))

	dictionary.Delete("Juan")

	fmt.Println("Dictionary después de borrarlo:", dictionary.items)

	fmt.Println("Dictionary IsEmpty:", dictionary.IsEmpty())
}
