package openai

var Gpt = NewGpt()

//TODO:每个新的会话都要一个新的gpt实例

func GetGpt() *GPT {
	return Gpt
}
