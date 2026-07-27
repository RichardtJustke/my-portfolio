# Registro de Alterações

## 2026-07-27 - Criação do NewSite com Fiber v2.0

### O que foi feito
Criação completa do portfólio versão 2.0 usando Fiber como framework web, do zero.

### Motivo
Substituição do router Chi pelo Fiber para simplificar o código, melhorar performance e facilitar manutenção futura. O Fiber oferece uma API mais expressiva e integração direta com templates.

### Arquivos impactados
- `main.go` — novo arquivo principal com Fiber (rotas, logger, serve de estáticos embutidos)
- `handlers/handlers.go` — configuração de templates e função `renderTemplate` para Fiber
- `handlers/home.go` — handler Home adaptado para Fiber
- `handlers/work.go` — handler Work com dados completos (stack, experiências, projetos) adaptado para Fiber
- `handlers/resume.go` — handler Resume adaptado para Fiber
- `templates/` — copiados do my-portfolio (base.html, home.html, work.html, resume.html)
- `static/` — copiados do my-portfolio (css/style.css, js/main.js, sw.js, manifest.json)
- `assets/` — copiados do my-portfolio (favicon, imagens, gifs, fontes)
- `go.mod` — módulo iniciado com dependência `github.com/gofiber/fiber/v2`