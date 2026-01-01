package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

)


func main()  {
	//this first 3 section allow us to use the program everywhere we are
	// we also need to run go build -o mycli to create an executable, and then move it to /usr/local/bin/
	// get command file path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory")
		os.Exit(1)
	}

	
	// creates .mycli directory if it does not exist
	configDir:= homeDir + "/.mycli"
	err = os.MkdirAll(configDir, 0755)
	if err != nil{
		fmt.Println("Error creating config Directory:", err)
		os.Exit(1)
	}

	// file path used every where
	filename := configDir + "/commands.txt"

	// what we have done:
	// homeDir         = /home/yourname
	//configDir       = /home/yourname/.mycli
	//filename        = /home/yourname/.mycli/commands.txt



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
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil{
			fmt.Println("Error while opening file:", err)
			os.Exit(1)
		}
		// defer allows me to not have to write file.Colose() everywhere
		defer file.Close()

		// write entry to the file
		_,err = file.WriteString(entry + "\n")
		if err != nil{
			fmt.Println("Error writing to file", err)
			os.Exit(1)
		}
		fmt.Println("Command added succesfully!")




	case "search":
		if len(os.Args)<3{
			fmt.Println("you did not provide any word")
			fmt.Println("type: mycli search 'word'")
			os.Exit(1)

		}
		searchTerm:= strings.ToLower(os.Args[2])



		// lets read the entire file
		content, err:= os.ReadFile(filename)
		if err!= nil{
			fmt.Println("Error while reading file,", err)
			os.Exit(1)
		}

		// converting bytes in to string and split by new lines
		lines:= strings.Split(string(content), "\n")

		//Display header
		fmt.Println("\nSearch results:")
		fmt.Println("----------------")


		// lets loop through the lines while ignoring spaces
		count := 1
		found := false
		for _, line := range lines{
			if line ==""{
				continue
			}
		// Split each line by | to get command and description
			parts := strings.Split(line, " | ")
			if len(parts)== 2{
				//Check if search term is in command OR description (case-insensitive)
				commandLower := strings.ToLower(parts[0])
				descriptionLower:= strings.ToLower(parts[1])

				if strings.Contains(commandLower, searchTerm) || strings.Contains(descriptionLower, searchTerm){
					fmt.Printf("[%d] %s\n", count, parts[0])
					fmt.Printf("    Descritpion: %s\n\n", parts[1])
					count++
					found = true
				}
			}
		}

		if !found {
			fmt.Println("No command found matching:", os.Args[2])
		}












	case "list":		
		//checks if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err){
			fmt.Println("This file does not exist, use add command to create file")
			os.Exit(0)
		}


		//reading entire file
		content, err := os.ReadFile(filename)
		if err != nil{
			fmt.Println("Error reding file:", err)
			os.Exit(1)
		}


		//converintg bytes to string and split by new line
		lines:= strings.Split(string(content), "\n")

		// Display each command with a number
		fmt.Println("\nYour saved commands:")
		fmt.Println("-------------------")
		count := 1
		for _, line := range lines{
			if line == ""{
				continue
			}
			//split each line by | to get command and descritpion
			parts := strings.Split(line, " | ")
			if len(parts) == 2{
				fmt.Printf("[%d] %s\n", count, parts[0])
				fmt.Printf("    Descritpion: %s\n\n", parts[1])
				count++
			}
		}



	case "delete":
		if len(os.Args) < 3{
			fmt.Println("No number provided")
			fmt.Println("Usage : mycli delete <number>")
			os.Exit(1)
		}


		// convering string to number. Ex mycli delete "3" this is a string not an int
		deleteNum, err := strconv.Atoi(os.Args[2])
		if err != nil{
			fmt.Println("Invalid number. Please provide a number")
			os.Exit(1)
		}


		content, err := os.ReadFile(filename)
		if err != nil{
			fmt.Println("Error while reding file")
		os.Exit(1)
		}



		lines:=strings.Split(string(content), "\n")


		//build a list of non empty lines
		var validLines []string
		for _,line := range lines {
			if line != ""{
				validLines = append(validLines, line)
			}
		}

		// check if number is valid
		if deleteNum < 1 || deleteNum > len(validLines){
			fmt.Println("Invalid number. You have", len(validLines), "commands.")
			os.Exit(1)
		}

		// remove line
		numToDelete := deleteNum - 1
		validLines = append(validLines[:numToDelete], validLines[numTorelete+1:]...)


		// Join lines back together with newlines
		newContent := strings.Join(validLines, "\n") + "\n"

		// overriting file complitely
		// 0644 means 6 owner can read 4 + 2 write =6 
		// group 4 = can only read
		// others  4 = can only read


		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil{
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
		fmt.Println("Command deleted succesfully")


	default:
		fmt.Println("Uknown command: %s\n", command)
		os.Exit(1)
	}


}
