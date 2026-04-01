package handlers

import (
	"net/http"
)

type StackItem struct {
	Name string
	Icon string
}

type Experience struct {
	Company string
	Role    string
	Desc    string
	Period  string
	Tag     string
}

type Project struct {
	ID        string
	Name      string
	Desc      string
	Status    string
	Tags      []string
	About     string
	Challenge string
	Tech      []string
	Link      string
	Label     string
}

type WorkData struct {
	Page        string
	Stack       []StackItem
	Experiences []Experience
	Projects    []Project
}

func Work(w http.ResponseWriter, r *http.Request) {
	data := WorkData{
		Page: "work",
		Stack: []StackItem{
			{Name: "JS", Icon: "javascript/javascript-original.svg"},
			{Name: "TS", Icon: "typescript/typescript-original.svg"},
			{Name: "React", Icon: "react/react-original.svg"},
			{Name: "Go", Icon: "go/go-original-wordmark.svg"},
			{Name: "HTML", Icon: "html5/html5-original.svg"},
			{Name: "CSS", Icon: "css3/css3-original.svg"},
			{Name: "Node", Icon: "nodejs/nodejs-original.svg"},
			{Name: "Docker", Icon: "docker/docker-original.svg"},
			{Name: "Git", Icon: "git/git-original.svg"},
			{Name: "Figma", Icon: "figma/figma-original.svg"},
			{Name: "Tailwind", Icon: "tailwindcss/tailwindcss-original.svg"},
			{Name: "Arduino", Icon: "arduino/arduino-original.svg"},
		},
		Experiences: []Experience{
			{
				Company: "EasyPlan",
				Role:    "Estagiário · Desenvolvedor de Software",
				Desc:    "Sistema de vendas criado do zero. Hoje é inteiramente delegado a mim.",
				Period:  "2024 — hoje",
				Tag:     "atual",
			},
			{
				Company: "Ketsu",
				Role:    "Fundador · Designer · Developer",
				Desc:    "Startup de serviços digitais. 3 anos em design e desenvolvimento.",
				Period:  "2022 — hoje",
				Tag:     "founder",
			},
			{
				Company: "Eletronorte",
				Role:    "Assistente Administrativo",
				Desc:    "Telecom. Automação de planilhas, design e suporte técnico.",
				Period:  "anterior",
				Tag:     "",
			},
		},
		Projects: []Project{
			{
				ID:        "justkeos",
				Name:      "JustkeOS",
				Desc:      "Portfólio como sistema operacional. Boot no pendrive, entra no portfólio.",
				Status:    "em breve",
				Tags:      []string{"C", "kernel"},
				Label:     "PROJETO PESSOAL",
				About:     "Um amigo me desafiou a fazer um portfólio melhor que o dele — feito em Assembly. Minha resposta foi entregar um .iso: você coloca no pendrive, dá boot, e entra direto no meu portfólio em formato de sistema operacional.",
				Challenge: "Mexer no kernel é território novo. Vai exigir aprender a fundo como um SO funciona desde o bootloader até a camada de exibição.",
				Tech:      []string{"C", "Assembly", "Kernel", "Bootloader"},
				Link:      "",
			},
			{
				ID:        "justkecli",
				Name:      "JustkeCLI",
				Desc:      "Portfólio interativo no terminal via npm.",
				Status:    "em andamento",
				Tags:      []string{"Node.js", "npm"},
				Label:     "PROJETO PESSOAL",
				About:     "Portfólio interativo via linha de comando. Instala pelo npm e navega pelo meu portfólio direto no terminal.",
				Challenge: "Aplicar JavaScript na prática — tenho o conhecimento teórico, mas construir uma CLI de verdade está mostrando onde estão as lacunas reais.",
				Tech:      []string{"Node.js", "npm", "JavaScript"},
				Link:      "https://github.com/RichardtJustke",
			},
			{
				ID:        "tonho",
				Name:      "Ketsu Tonho",
				Desc:      "E-commerce com orçamento de evento automatizado.",
				Status:    "em andamento",
				Tags:      []string{"React", "Supabase"},
				Label:     "KETSU · CLIENT WORK",
				About:     "E-commerce onde o cliente sai com o orçamento completo de um evento em mãos. Remodelação completa do site com foco em simplificar o processo de cotação.",
				Challenge: "Entregar uma experiência que realmente resolve o problema de orçamento de forma intuitiva.",
				Tech:      []string{"React", "Node.js", "Supabase", "TypeScript"},
				Link:      "",
			},
			{
				ID:        "reboot",
				Name:      "Reboot",
				Desc:      "Robô seguidor de linha autônomo. Meu primeiro projeto real.",
				Status:    "concluído",
				Tags:      []string{"Arduino", "C++"},
				Label:     "PROJETO PESSOAL · PRIMEIRO PROJETO",
				About:     "Meu robô seguidor de linha. Construído durante o curso de robótica, foi onde a programação deixou de ser teoria e virou algo físico funcionando na minha frente.",
				Challenge: "Aprender programação do zero enquanto construía algo concreto.",
				Tech:      []string{"Arduino", "C++", "Robótica", "Eletrônica"},
				Link:      "https://github.com/RichardtJustke",
			},
		},
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
