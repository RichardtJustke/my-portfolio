# Shaders

Fragment shaders GLSL relacionados ao dither.

- [`sphere.frag.glsl`](sphere.frag.glsl) — shader completo da esfera 3D dither
  usada no header (`#sphere`) e no card do Dither Lab (`#miniLabSphere`).
  Inclui: cálculo analítico da esfera, iluminação Lambert com tracking de
  mouse, e 4 tipos de dither (noise, Bayer 2/4/8).

- [`dither-only.frag.glsl`](dither-only.frag.glsl) — versão isolada do
  passo de dither, aplicada sobre um gradiente linear. Útil pra debugar
  o padrão sem a forma 3D no caminho.

## Uniformes principais

| Uniform        | Tipo   | Função                                    |
| -------------- | ------ | ----------------------------------------- |
| `u_resolution` | vec2   | Tamanho do canvas em pixels               |
| `u_time`       | float  | Tempo em segundos desde o start           |
| `u_pxSize`     | float  | Tamanho do "pixel" dithered               |
| `u_color`      | vec4   | RGBA da cor (alpha é multiplicado)        |
| `u_mouse`      | vec2   | Posição normalizada do mouse [-1, +1]     |
| `u_scale`      | float  | Escala da forma (menor = maior)           |
| `u_shape`      | float  | ID da forma (1 = sphere, ...)             |
| `u_ditherType` | int    | 1 noise, 2 bayer2, 3 bayer4, 4 bayer8     |
