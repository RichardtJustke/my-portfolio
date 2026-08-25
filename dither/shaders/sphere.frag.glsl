#version 300 es
precision highp float;

uniform vec2  u_resolution;
uniform float u_time;
uniform float u_pxSize;
uniform vec4  u_color;
uniform vec2  u_mouse;
uniform float u_scale;
uniform float u_shape;     // 1 = sphere, 2 = torus (placeholder)
uniform int   u_ditherType; // 1 = noise, 2 = bayer2, 3 = bayer4, 4 = bayer8

out vec4 fragColor;

const int bayer8x8[64] = int[64](
   0, 32,  8, 40,  2, 34, 10, 42,
  48, 16, 56, 24, 50, 18, 58, 26,
  12, 44,  4, 36, 14, 46,  6, 38,
  60, 28, 52, 20, 62, 30, 54, 22,
   3, 35, 11, 43,  1, 33,  9, 41,
  51, 19, 59, 27, 49, 17, 57, 25,
  15, 47,  7, 39, 13, 45,  5, 37,
  63, 31, 55, 23, 61, 29, 53, 21
);

const int bayer4x4[16] = int[16](
   0,  8,  2, 10,
  12,  4, 14,  6,
   3, 11,  1,  9,
  15,  7, 13,  5
);

const int bayer2x2[4] = int[4](
  0, 2,
  3, 1
);

float getBayer8(ivec2 p) {
  int x = (p.x % 8 + 8) % 8;
  int y = (p.y % 8 + 8) % 8;
  return float(bayer8x8[y * 8 + x]) / 64.0;
}

float getBayer4(ivec2 p) {
  int x = (p.x % 4 + 4) % 4;
  int y = (p.y % 4 + 4) % 4;
  return float(bayer4x4[y * 4 + x]) / 16.0;
}

float getBayer2(ivec2 p) {
  int x = (p.x % 2 + 2) % 2;
  int y = (p.y % 2 + 2) % 2;
  return float(bayer2x2[y * 2 + x]) / 4.0;
}

void main() {
  float pxSize = max(1.0, u_pxSize);
  vec2  pxCoord   = (gl_FragCoord.xy - 0.5 * u_resolution) / pxSize;
  vec2  gridCoord = floor(pxCoord);
  ivec2 iGrid     = ivec2(gridCoord);

  // Normaliza uv em torno do centro, escalado por u_scale
  vec2 uv = (gridCoord * pxSize) / (0.5 * min(u_resolution.x, u_resolution.y));
  uv /= max(0.001, u_scale);

  // Forma: esfera 3D via distância implícita
  float shape = 0.0;
  if (u_shape < 1.5) {
    float d = 1.0 - dot(uv, uv);
    if (d > 0.0) {
      vec3 norm = vec3(uv, sqrt(d));
      float t = u_time;
      vec3 lightDir = normalize(vec3(
        cos(1.5 * t) + u_mouse.x * 1.5,
        0.8 - u_mouse.y * 1.5,
        sin(1.25 * t)
      ));
      shape = clamp(0.5 + 0.5 * dot(lightDir, norm), 0.0, 1.0);
    }
  }

  // Dither: lookup na matriz apropriada
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
  vec3  fgColor = u_color.rgb * u_color.a;

  fragColor = vec4(fgColor * result, u_color.a * result);
}
