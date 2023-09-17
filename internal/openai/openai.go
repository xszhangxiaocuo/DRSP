package openai

import (
	"DRSP/config"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type GPT struct {
	Client   *openai.Client
	messages []openai.ChatCompletionMessage
}

func NewGPT() *GPT {
	return &GPT{
		Client: openai.NewClient(config.AppConf.OpenaiConfig.Key),
	}
}

func (gpt *GPT) Chat(question string) string {
	var text string
	if len(gpt.messages) == 0 {
		text = fmt.Sprintf("我想向你咨询有关食品相关的问题，现在我有一个食品，它的配料表如下：%s，请你帮我分析一下这个食物的配料表并指出其中的致敏物质，有害物质以及对人类身体健康有害的食品添加剂", question)
	} else {
		text = question
	}
	gpt.messages = append(gpt.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})

	resp, err := gpt.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: gpt.messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err.Error()
	}

	content := resp.Choices[0].Message.Content
	gpt.messages = append(gpt.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content
}
