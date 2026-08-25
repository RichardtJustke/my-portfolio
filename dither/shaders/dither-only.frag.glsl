#version 300 es
precision highp float;

uniform vec2  u_resolution;
uniform float u_pxSize;
uniform int   u_ditherType; // 1 = noise, 2 = bayer2, 3 = bayer4, 4 = bayer8
uniform vec4  u_color;

in vec2 v_uv;     // 0..1
out vec4 fragColor;

// Mesmas matrizes do sphere.frag.glsl (omitidas aqui por brevidade)
float getBayer8(ivec2 p);
float getBayer4(ivec2 p);
float getBayer2(ivec2 p);

void main() {
  // Converte uv normalizado em coordenadas de pixel "dithered"
  vec2  pxCoord   = v_uv * u_resolution / max(1.0, u_pxSize);
  vec2  gridCoord = floor(pxCoord);
  ivec2 iGrid     = ivec2(gridCoord);

  // "shape" simulado por gradiente linear — útil pra ver o dither isolado
  float shape = v_uv.x * 0.7 + v_uv.y * 0.3;

  float dither = 0.0;
  if (u_ditherType == 1) {
    dither = fract(sin(dot(gridCoord, vec2(12.9898, 78.233))) * 43758.5453) - 0.5;
  } else if (u_ditherType == 2) {
    dither = getBayer2(iGrid) - 0.5;
  } else if (u_ditherType == 4) {
    dither = getBayer8(iGrid) - 0.5;
  } else {
    dither = getBayer4(iGrid) - 0.5;
  }

  float result = step(0.5, shape + dither);
  fragColor = vec4(u_color.rgb * result, result);
}
