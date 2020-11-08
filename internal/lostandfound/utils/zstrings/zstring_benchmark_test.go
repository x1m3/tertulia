package zstrings

import (
	"testing"
)

// go test -bench . -benchtime 10s -benchmem
func BenchmarkZStringUncompressed(b *testing.B) {
	var s *ZString

	text := `Con cien cañones por banda,
		viento en popa a toda vela,
		no corta el mar, sino vuela,
		un velero bergantín:
		bajel pirata que llaman
		por su bravura el Temido,
		en todo mar conocido
		del uno al otro confín.


		La luna en el mar riela,
		en la lona gime el viento,
		y alza en blando movimiento
		olas de plata y azul;
		y ve el capitán pirata,
		cantando alegre en la popa,
		Asia a un lado, al otro Europa
		y allá a su frente Stambul.`

	for i := 0; i < b.N; i++ {
		s = NewZString(text)
		_ = s.Value()
	}
	_ = s
}

func BenchmarkZStringCompressed(b *testing.B) {
	var s *ZString

	text := `Con cien cañones por banda,
		viento en popa a toda vela,
		no corta el mar, sino vuela,
		un velero bergantín:
		bajel pirata que llaman
		por su bravura el Temido,
		en todo mar conocido
		del uno al otro confín.


		La luna en el mar riela,
		en la lona gime el viento,
		y alza en blando movimiento
		olas de plata y azul;
		y ve el capitán pirata,
		cantando alegre en la popa,
		Asia a un lado, al otro Europa
		y allá a su frente Stambul.`

	for i := 0; i < b.N; i++ {
		s = NewZStringCompressed(text)
		_ = s.Value()
	}
	_ = s
}



