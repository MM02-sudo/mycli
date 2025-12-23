package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"

)


func main()  {
	//this checks if user wrote nothing after the program name
	//And shows how to use the program
	
	if len(os.Args)<2{
		fmt.Println("Usage: mycli <command>")
		fmt.Println("command:")
		fmt.Println("  add    - Add new command")
		fmt.Println("write description")
		fmt.Println("  search - Search for commands")
		fmt.Println("  list   - List all commands")
		fmt.Println("  delete - Delete command")

		// Exit code 1 stops the program,since we did no enter
		// a correct command
		os.Exit(1)
	}

	command := os.Args[1]
	fmt.Printf("You ran: %s\n", command)

	// here we see if the command is valid
	switch command {
	case "add":
		if len(os.Args)<3{
			fmt.Println("Usage: mycli add <command-text>")
			fmt.Println("Example: mycli add \"find . -name '*.go'\"")
			os.Exit(1)
		}

		//we get evrything after the add and join it in to a single string
		commandText:= strings.Join(os.Args[2:], " ")

		//asking for description for command
		fmt.Printf("Enter a descritpion for the command: ")
		var description string

		//this reads the entire line encluding the spaces
		scanner:= bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			description = scanner.Text()
			
		}
		// formating command | description
		entry := commandText + " | " + description

		//open or create a file for appending command
		filename:= "command.txt"
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil{
			fmt.Println("Error while opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// write entry to the file
		_,err = file.WriteString(entry + "\n")
		if err != nil{
			fmt.Println("Error writing to file", err)
			os.Exit(1)
		}
		fmt.Println("Command added succesfully!")

	case "search":
		fmt.Println("Search command - not implemented yet")


	case "list":
		fmt.Println("List command - not implemented yet")







	case "delete":
		fmt.Println("delete command - not implemented yet")
	default:
		fmt.Println("Uknown command: %s\n", command)
		os.Exit(1)
	}


}
