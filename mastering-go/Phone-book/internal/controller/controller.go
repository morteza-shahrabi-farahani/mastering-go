package controller

import (
	"fmt"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
)

func CommandLineHandler(arguments []string) {
	if err := checkArgumentsLength(arguments); err != nil {
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Please provide a search term")
			return
		}

		usersList, appErr := phonebook.GetList()
		if appErr != nil {
			fmt.Println(appErr.Message)
			return
		}

		result, err := phonebook.Serach(usersList, arguments[2])
		if err != nil {
			fmt.Println(err.Message)
			return
		}

		fmt.Println(result)

	case "list":
		usersList, err := phonebook.GetList()
		if err != nil {
			fmt.Println(err.Message)
			return
		}

		fmt.Println(usersList)

	case "insert":
		if err := validateInsert(arguments); err != nil {
			fmt.Println(err)
			return
		}

		id, err := phonebook.Insert(&phonebook.Entry{Name: arguments[2], Surname: arguments[3], PhoneNumber: arguments[4]})
		if err != nil {
			fmt.Println(err.Message)
			return
		}

		fmt.Printf("successfully inserted with id = %d \n", id)

	case "delete":
		if err := validateDelete(arguments); err != nil {
			fmt.Println(err)
			return
		}

		err := phonebook.Delete(arguments[2])
		if err != nil {
			fmt.Println(err.Message)
			return
		}

		fmt.Println("successfully deleted")

	default:
		fmt.Println("not valid option")
	}
}

func validateDelete(arguments []string) error {
	if len(arguments) != 3 {
		return fmt.Errorf("not enought arguments for delete")
	}

	return nil
}

func validateInsert(arguments []string) error {
	if len(arguments) != 5 {
		return fmt.Errorf("not enought arguments for insert")
	}

	return nil
}

func checkArgumentsLength(arguments []string) error {
	if len(arguments) == 1 {
		return fmt.Errorf("Please enter required arguments!!")
	}

	return nil
}
