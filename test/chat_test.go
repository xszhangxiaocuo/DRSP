package test

import (
	"DRSP/internal/openai"
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestChat(t *testing.T) {
	gpt := openai.NewGPT()
	var question string
	reader := bufio.NewReader(os.Stdin)
	for {
		question, _ = reader.ReadString('\n')
		question = strings.Replace(question, "\n", "", -1)
		answer := gpt.Chat(question)
		fmt.Println(answer)
	}

}
