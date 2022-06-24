package builder

import (
	"fmt"
	"redis_client_example/commands"
	"strings"
)

func BuildCommandtring(texts ...string) string {

	commandLength := len(texts)

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*%d\r\n", commandLength))

	for _, text := range texts {
		sb.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(text), text))
	}

	return sb.String()
}

func BuildArrayString(listName string, texts ...interface{}) string {

	commandList := []string{
		commands.RPUSH,
		listName,
	}

	for _, text := range texts {
		str := fmt.Sprintf("%s", text)
		elements := strings.Split(str[1:len(str)-1], " ")

		for _, element := range elements {
			commandList = append(commandList, element)
		}
	}

	commandLength := len(commandList)

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*%d\r\n", commandLength))

	for _, command := range commandList {
		sb.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(command), command))

	}

	return sb.String()
}
