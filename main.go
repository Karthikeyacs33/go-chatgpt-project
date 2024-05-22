package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main(){
	godotenv.Load()

	apikey:= os.Getenv("API_KEY")
	if apikey ==""{
		log.Fatalln("API_KEY is not set")
	}

	ctx:= context.Background()
	client:= openai.NewClient(apikey)
	response,err:= client.CreateChatCompletion(ctx,openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: "You are a helpful assistant.",
			},
		},
	})
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println(response.Choices[0].Message.Content)
}