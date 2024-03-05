package openai

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"log"
)

type GPT struct {
	Client   *openai.LLM
	messages []llms.MessageContent
}

func NewGpt() *GPT {
	llm, err := openai.New(openai.WithBaseURL("https://api.openai-proxy.org/v1"))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &GPT{Client: llm}
}

func (gpt *GPT) Chat(message string) string {
	ctx := context.Background()
	var text string
	if len(gpt.messages) == 0 {
		text = fmt.Sprintf("我将提供一段由ocr识别生成的食品信息，由于ocr识别可能存在错误，分析过程中请你请你注意配料以及过敏原的合理性，接下来根据以下文本提取出食品的配料并分析配料中可能存在的过敏原，如果有危害人体健康对配料也请指出：%s", message)
	} else {
		text = message
	}
	gpt.messages = append(gpt.messages, llms.TextParts(schema.ChatMessageTypeHuman, text))

	resp, err := gpt.Client.GenerateContent(ctx, gpt.messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	respContent := resp.Choices[0].Content
	gpt.messages = append(gpt.messages, llms.TextParts(schema.ChatMessageTypeSystem, respContent))

	fmt.Println(resp)
	return respContent
}
