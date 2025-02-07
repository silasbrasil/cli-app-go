package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	questions "github.com/silasbrasil/cli-app-go/questions"
)

func main() {
	var envModel questions.EnvModel
	envModel.Title = "Em qual ambiente você gostaria de trabalhar?"
	envModel.Choices = []string{"local", "staging", "prod"}
	envModel.Selected = make(map[int]struct{})

	p := tea.NewProgram(envModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Há algum erro: %v", err)
		os.Exit(1)
	}

	for k := range envModel.Selected {
		fmt.Printf("%s \n", envModel.Choices[k])
	}
}
