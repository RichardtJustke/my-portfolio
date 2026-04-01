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
	Image     string
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
				Role:    "Intern · Software Developer",
				Desc:    "Sales system built from scratch. Today it is entirely delegated to me.",
				Period:  "2024 — present",
				Tag:     "current",
			},
			{
				Company: "Ketsu",
				Role:    "Founder · Designer · Developer",
				Desc:    "Digital services startup. 3 years of design and development experience.",
				Period:  "2022 — present",
				Tag:     "founder",
			},
			{
				Company: "Eletronorte",
				Role:    "Administrative Assistant",
				Desc:    "Telecom. Spreadsheet automation, design and IT support.",
				Period:  "former",
				Tag:     "",
			},
		},
		Projects: []Project{
			{
				ID:        "justkeos",
				Name:      "JustkeOS",
				Desc:      "Portfolio as an operating system. Boot from a pendrive and enter the portfolio.",
				Status:    "soon",
				Tags:      []string{"C", "kernel"},
				Label:     "PERSONAL PROJECT",
				About:     "A friend challenged me to build a better portfolio than his — which was made in Assembly. My response was to deliver an .iso: you put it on a flash drive, boot it up, and enter straight into my portfolio in the form of an OS.",
				Challenge: "Messing with the kernel is new territory. It requires deeply learning how an OS works from the bootloader to the display layer.",
				Tech:      []string{"C", "Assembly", "Kernel", "Bootloader"},
				Link:      "https://github.com/RichardtJustke/justke-os",
				Image:     "",
			},
			{
				ID:        "justkecli",
				Name:      "JustkeCLI",
				Desc:      "Interactive portfolio in the terminal via npm.",
				Status:    "in progress",
				Tags:      []string{"Node.js", "npm"},
				Label:     "PERSONAL PROJECT",
				About:     "Interactive portfolio via command line. Install through npm and navigate my portfolio directly in your terminal.",
				Challenge: "Applying JavaScript in practice — I have the theoretical knowledge, but building a real CLI shows where the actual gaps are.",
				Tech:      []string{"Node.js", "npm", "JavaScript"},
				Link:      "https://github.com/RichardtJustke/justke-cli",
				Image:     "/assets/cli.mp4",
			},
			{
				ID:        "tonho",
				Name:      "Ketsu Tonho",
				Desc:      "E-commerce with automated event budgeting.",
				Status:    "in progress",
				Tags:      []string{"React", "Supabase"},
				Label:     "KETSU · CLIENT WORK",
				About:     "An e-commerce where the client leaves with a complete event budget in hand. A total website remodel focused on simplifying the quoting process.",
				Challenge: "Delivering an experience that genuinely solves the budgeting problem intelligently.",
				Tech:      []string{"React", "Node.js", "Supabase", "TypeScript"},
				Link:      "https://github.com/RichardtJustke/Ketsu-Tonho",
				Image:     "/assets/tonho.jpg",
			},
			{
				ID:        "chicas",
				Name:      "Chicas Eventos",
				Desc:      "Event planning & management e-commerce platform.",
				Status:    "in progress",
				Tags:      []string{"React", "Supabase", "E-commerce"},
				Label:     "KETSU · CLIENT WORK",
				About:     "A comprehensive event planning platform for Chicas Eventos. Complete e-commerce integration with event budgeting and management tools.",
				Challenge: "Building a scalable solution that handles complex event management workflows while maintaining excellent user experience.",
				Tech:      []string{"React", "TypeScript", "Supabase", "Tailwind CSS"},
				Link:      "https://github.com/RichardtJustke/Ketsu-Chicas",
				Image:     "/assets/chicas.jpg",
			},
			{
				ID:        "reboot",
				Name:      "Reboot",
				Desc:      "Autonomous line-follower robot. My first real project.",
				Status:    "done",
				Tags:      []string{"Arduino", "C++"},
				Label:     "PERSONAL PROJECT · FIRST PROJECT",
				About:     "My line-follower robot. Built during a robotics course, it was where programming stopped being theory and became something physical working right in front of me.",
				Challenge: "Learning programming from scratch while building something concrete at the same time.",
				Tech:      []string{"Arduino", "C++", "Robotics", "Electronics"},
				Link:      "https://github.com/RichardtJustke",
				Image:     "/assets/Reboot4.jpg",
			},
		},
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
