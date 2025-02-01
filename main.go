package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, "Crie um roteiro de estudo para iniciantes em desenvolvimento de software. "+
		"O roteiro deve incluir:\n"+
		"1. Conceitos básicos de programação.\n"+
		"2. Linguagens de programação recomendadas para iniciantes.\n"+
		"3. Boas práticas e ferramentas essenciais.\n"+
		"4. Projetos práticos para consolidar o aprendizado."+
		"#Saida em português",
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = completion
}
