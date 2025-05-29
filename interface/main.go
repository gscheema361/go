package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

//	type outputable interface {
//		Save() error
//		Display()
//	}
type outputable interface {
	saver
	displayer
}

func main() {
	// printSomething(1)
	// printSomething(2.5)
	// printSomething("hello")
	// printSomething(true)
	// printSomething([]int{1, 2, 3})
	// printSomething(map[string]int{"a": 1, "b": 2})
	printSomething(1)
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	// todo.Display()
	// err = saveData(todo)
	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)

}

func printSomething(value interface{}) {

	intVal, ok := value.(int)
	if ok {
		fmt.Println("int:", intVal)
	}
	strVal, ok := value.(string)
	if ok {
		fmt.Println("str:", strVal)
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float:", floatVal)
	} else {
		fmt.Println("float: not found")
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("int:", value)
	// case float64:
	// 	fmt.Println("float64:", value)
	// case string:
	// 	fmt.Println("string:", value)
	// case bool:
	// 	fmt.Println("bool:", value)
	// default:
	// 	fmt.Println("unknown type:", reflect.TypeOf(value))
	// }

}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
