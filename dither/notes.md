# Notes — Dither Lab

Anotações soltas sobre o playground `static/dither-lab.html`.

## Estrutura do playground

- 8 presets salvos em código (`id: 1..8`) — esferas, donut, vórtice, lava,
  sonar, saturno, areia, cubo
- Cada preset define: `shape`, `dither`, `color`, `pxSize`, `tag`, `desc`
- Controles: shape, dither type, cor, pixel size, scale, speed

## Formas (shape)

- `1` — esfera 3D Edgar (modelo orbital usado em theedgar.dev)
- `2` — donut / torus via raymarching
- `3` — vórtice / swirl
- `4` — lava / simplex noise
- `5` — radar / sonar / ripple
- `6` — planeta com anéis
- `7` — areia granulada
- `8` — cubo minimalista

## Tipos de dither (dither)

- `1` — White noise (sin/cos hash)
- `2` — Bayer 2x2
- `3` — Bayer 4x4 (padrão)
- `4` — Bayer 8x8 (fino)

## Pegadinhas

- O `dpr` (device pixel ratio) é aplicado no `pxSize` para o pixel visual
  ficar consistente em retina / 4K
- O shader usa `u_mouse` em range `[-1, +1]` — precisa normalizar a posição
  do mouse em relação ao canvas, não à viewport
- O `u_color` vem como `vec4` mesmo sendo uma cor só, porque o alpha é
  multiplicado em `fragColor.a` para o canvas ter fundo transparente

## Bugs que já tive

- Borda serrilhada em canvas pequeno → solução: aplicar `image-rendering:
  pixelated` no CSS
- Mouse tracking travando após hover-out → esquecia de resetar
  `mouse.targetX/Y` no `mouseleave` (corrigido em `static/js/main.js:142-145`)
- Dither carregando em /work → `mainSphereCv` era `null` e o loop
  ainda chamava `render`; corrigido com guard `if (mainSphere)`
