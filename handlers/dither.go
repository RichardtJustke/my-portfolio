package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ShaderExample struct {
	Name        string
	Code        string
	Description string
}

type DitherType struct {
	Name       string
	Desc       string
	Pattern    string
	WhenToUse  string
}

type DitherData struct {
	Page           string
	Intro          string
	BayerMatrix    string
	ShaderExamples []ShaderExample
	DitherTypes    []DitherType
}

func Dither(c *fiber.Ctx) error {
	data := DitherData{
		Page: "dither",
		Intro: `Dithering é a técnica de simular gradações de cor usando apenas uma quantidade
limitada de cores (frequentemente só duas: preto e branco). O truque está em
variar a densidade dos pixels em uma região local — o olho faz a média.`,
		BayerMatrix: `0  32  8  40  2  34  10  42
48 16  56 24  50 18  58 26
12 44  4  36  14 46  6  38
60 28  52 20  62 30  54 22
3  35  11 43  1  33  9  41
51 19  59 27  49 17  57 25
15 47  7  39 13 45  5  37
63 31  55 23  61 29  53 21`,
		ShaderExamples: []ShaderExample{
			{
				Name: "Bayer 8x8 Lookup",
				Code: `float getBayer8(ivec2 p) {
  int x = (p.x % 8 + 8) % 8;
  int y = (p.y % 8 + 8) % 8;
  return float(bayer8x8[y * 8 + x]) / 64.0;
}`,
				Description: "Função que faz lookup na matriz Bayer 8x8. Cada pixel da tela é indexado em uma posição da matriz, e o valor lá (0..63) é usado como limiar.",
			},
			{
				Name: "Sphere SDF",
				Code: `float d = 1.0 - dot(uv, uv);
if (d > 0.0) {
  vec3 norm = vec3(uv, sqrt(d));
  shape = clamp(0.5 + 0.5 * dot(lightDir, norm), 0.0, 1.0);
}`,
				Description: "Forma de esfera 3D calculada analiticamente (SDF implícita) a partir de dot(uv, uv) — sem malha, sem VBOs de geometria.",
			},
			{
				Name: "Dither Step",
				Code: `float result = step(0.5, shape + dither);
fragColor = vec4(fgColor * result, u_color.a * result);`,
				Description: "Compara o valor da forma com o limiar do dither. Se for maior, pinta; senão, deixa em branco.",
			},
		},
		DitherTypes: []DitherType{
			{Name: "Bayer 2x2", Desc: "Bloco grande, granulado", Pattern: "4x4", WhenToUse: "Estilo Game Boy, cubos minimalistas"},
			{Name: "Bayer 4x4", Desc: "Bloco médio", Pattern: "16x16", WhenToUse: "Padrão geral (bom custo-benefício)"},
			{Name: "Bayer 8x8", Desc: "Pontilhado fino", Pattern: "64x64", WhenToUse: "Esferas, formas orgânicas"},
			{Name: "Blue noise", Desc: "Pontilhado areia", Pattern: "Aleatório", WhenToUse: "Superfícies planas, simulação de areia"},
			{Name: "White noise", Desc: "Pontilhado caótico", Pattern: "Aleatório", WhenToUse: "Evitar (ruim visualmente)"},
		},
	}
	return renderTemplate(c, data)
}
