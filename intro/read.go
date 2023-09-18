package intro

import (
	"bufio"
	"log"
	"os"
)

func Read() ([]string, error) {

	file, err := os.Open("./intro/introduction.txt")
	if err != nil {
		log.Panicln("can't find introduction.txt.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text, nil
}
