# Dither

Anotações pessoais sobre dithering, ordered dithering e shaders WebGL2.
A ideia é juntar aqui a teoria, os snippets de GLSL e os presets que valem a pena lembrar.

## O que é dithering

Dithering é a técnica de simular gradações de cor usando apenas uma quantidade
limitada de cores (frequentemente só duas: preto e branco). O truque está em
variar a *densidade* dos pixels em uma região local — o olho faz a média.

Origens:

- Impressão (jornais, meio-tom / halftone)
- Hardware antigo (Game Boy, Mac 1-bit, monitores CGA)
- Computação gráfica moderna (rendering estilizado, post-process, pixel art)

## Ordered Dithering

Usa uma **matriz de limiar** (threshold map) em vez de ruído aleatório.
A mais comum é a **Bayer 8x8**:

```
0  32  8  40  2  34  10  42
48 16  56 24  50 18  58 26
12 44  4  36  14 46  6  38
60 28  52 20  62 30  54 22
3  35  11 43  1  33  9  41
51 19  59 27  49 17  57 25
15 47  7  39  13 45  5  37
63 31  55 23  61 29  53 21
```

Cada pixel da tela é indexado em uma posição da matriz, e o valor lá
(0..63) é usado como limiar: se o brilho calculado for maior, pinta; senão,
deixa em branco. O padrão resultante "engana" o olho simulando ~64 níveis
de cinza com apenas duas cores.

Vantagens sobre random dithering:

- Sem aleatoriedade — totalmente determinístico
- Padrão regular, agradável visualmente
- Barato de calcular (lookup + comparação)

## Como aplico no portfolio

A esfera do header (`#sphere`) e a mini-esfera do card do Dither Lab
(`#miniLabSphere`) são renderizadas com WebGL2 puro:

- **Forma**: esfera 3D calculada analiticamente (SDF implícita) a partir de
  `dot(uv, uv)` — sem malha, sem VBOs de geometria.
- **Iluminação**: Lambert com luz direcional que rastreia o mouse e gira
  com o tempo.
- **Dither**: matriz Bayer 8x8 como `const int[64]`, comparada com o
  cosseno do ângulo de luz normalizado.

Snippet GLSL completo: ver [`shaders/sphere.frag.glsl`](shaders/sphere.frag.glsl).

## Tipos de dither que testei

| Tipo       | Padrão                  | Quando uso                              |
| ---------- | ----------------------- | --------------------------------------- |
| Bayer 2x2  | Bloco grande, granulado | Estilo Game Boy, cubos minimalistas     |
| Bayer 4x4  | Bloco médio             | Padrão geral (bom custo-benefício)      |
| Bayer 8x8  | Pontilhado fino         | Esferas, formas orgânicas               |
| Blue noise | Pontilhado "areia"      | Superfícies planas, simulação de areia  |
| White noise| Pontilhado caótico       | Evitar (ruim visualmente, sem padrão)   |

## Referências

- [Bayer matrix (Wikipedia)](https://en.wikipedia.org/wiki/Ordered_dithering)
- `static/dither-lab.html` — playground interativo com 8 presets
- `static/js/main.js` — implementação em `createDitherInstance(...)`

## Snippets rápidos

- `shaders/sphere.frag.glsl` — fragment shader completo da esfera
- `shaders/dither-only.frag.glsl` — só o passo de dither, isolado
