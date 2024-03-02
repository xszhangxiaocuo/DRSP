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
		text = fmt.Sprintf("我将提供一段由ocr识别生成的食品信息，由于ocr识别可能存在错误，分析过程中请你请你注意配料以及过敏原的合理性，接下来根据以下文本提取出食品的配料并分析配料中可能存在的过敏原，如果有危害人体健康对配料也请指出：%s", question)
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
