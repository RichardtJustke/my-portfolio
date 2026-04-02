# richardtjustke.dev

> A minimalist portfolio built with Go — deliberately choosing the harder path to actually learn something.

![Portfolio Preview](./docs/preview.png)

---

## Demo

🔗 **[richardtjustke.dev](https://richardtjustke.dev)**

<!-- Add a short screen recording or GIF here -->
![Demo](./docs/demo.gif)

---

## Why Go instead of React?

I have a React background. Building this in React would have taken a weekend.

That's exactly why I didn't.

The goal wasn't to ship fast — it was to learn Go's `net/http`, understand how a server actually works, and stop relying on abstractions I don't fully understand. Every CSS file that didn't load, every template that broke, every middleware I had to wire up manually — that was the point.

---

## Stack

| Layer | Choice |
|---|---|
| Language | Go |
| Router | [Chi](https://github.com/go-chi/chi) |
| Templates | `html/template` (Go stdlib) |
| Markdown | [goldmark](https://github.com/yuin/goldmark) |
| Font | IBM Plex Mono |
| Weather | [wttr.in](https://wttr.in) API |
| Hosting | — |

No frameworks. No bundlers. No node_modules.

---

## Features

- **Dark minimalist design** — `#0c0c0c` background, IBM Plex Mono, blue accent `#3d6fa8`
- **Canvas dot-sphere animation** — generative animation rendered beside my name
- **Live temperature in footer** — fetched from `wttr.in` on each request
- **Borderless contact form** — clean, no distractions
- **/writing page** — markdown-powered, rendered server-side with goldmark

---

## Project Structure

```
my-portfolio/
├── main.go
├── handlers/
│   └── home.go
├── templates/
│   ├── base.html
│   ├── home.html
│   └── writing.html
└── static/
    ├── css/
    └── js/
```

---

## Running Locally

```bash
git clone https://github.com/RichardtJustke/my-portfolio
cd my-portfolio
go run main.go
```

Server starts at `http://localhost:3000`.

---

## Key Learnings

**Static file serving** requires explicit middleware in Go — `http.StripPrefix` is not optional.

```go
r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
```

Forgetting this was my first real stumbling block. Figuring it out was satisfying in a way that `import styles from './styles.css'` never is.

---

## License

MIT
