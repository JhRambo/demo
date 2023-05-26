package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	fileName := "example.go"
	content := `
	package main
	
	import "fmt"
	
	func main1() {
		fmt.Println("Hello, world!")
	}
	`

	// Check if file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Create new file
		err := ioutil.WriteFile(fileName, []byte(content), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// Replace contents of existing file
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Format the file
	cmd := exec.Command("go", "fmt", fileName)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File created and updated successfully.")
}
