package zstrings

import (
	"testing"
	"runtime"
	"fmt"
)

func TestZString(t *testing.T) {

	strings := [...]string{
		"lala",
		"El perro de san roque no tiene rabo",
		`Con cien cañones por banda,
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
		y allá a su frente Stambul.


		«Navega, velero mío,
		sin temor,
		que ni enemigo navío,
		ni tormenta, ni bonanza,
		tu rumbo a torcer alcanza,
		ni a sujetar tu valor.


		«Veinte presas
		hemos hecho
		a despecho
		del inglés,
		y han rendido
		sus pendones
		cien naciones
		a mis pies.


		«¿Qué es mi barco? Mi tesoro.
		¿Qué es mi Dios? La libertad.
		¿Mi ley? ¡La fuerza y el viento!
		¿Mi única patria? ¡La mar!


		«Allá muevan feroz guerra
		ciegos reyes
		por un palmo más de tierra:
		que yo tengo aquí por mío
		cuanto abarca el mar bravío,
		a quien nadie impuso leyes.


		«Y no hay playa
		sea cual quiera,
		ni bandera
		de esplendor,
		que no sienta
		mi derecho
		y dé pecho
		a mi valor.


		«¿Qué es mi barco? Mi tesoro.
		¿Qué es mi Dios? La libertad.
		¿Mi ley? ¡La fuerza y el viento!
		¿Mi única patria? ¡La mar!


		«A la voz de «¡barco viene!»
		Es de ver
		cómo vira y se previene
		a todo trapo a escapar:
		que yo soy el rey del mar,
		y mi furia es de temer.


		«En las presas
		yo divido
		lo cogido
		por igual:
		sólo quiero
		por riqueza
		la belleza
		sin rival.

		«¿Qué es mi barco? Mi tesoro.
		¿Qué es mi Dios? La libertad.
		¿Mi ley? ¡La fuerza y el viento!
		¿Mi única patria? ¡La mar!


		«¡Sentenciado estoy a muerte!
		Yo me río:
		no me abandone la suerte,
		y al mismo que me condena,
		colgaré de alguna antena,
		quizá en su propio navío.


		«Y si caigo,
		¿qué es la vida?
		Por perdida
		ya la di
		cuando el yugo
		del esclavo,
		como un bravo,
		sacudí.

		«¿Qué es mi barco? Mi tesoro.
		¿Qué es mi Dios? La libertad.
		¿Mi ley? ¡La fuerza y el viento!
		¿Mi única patria? ¡La mar!


		«Son mi música mejor
		aquilones;
		el estrépito y temblor
		de los cables sacudidos,
		del negro mar los bramidos
		y el rugir de mis cañones.


		«Y del trueno
		al son violento,
		y del viento
		al rebramar,
		yo me duermo
		sosegado.
		Arrullado
		por el mar.


		«¿Qué es mi barco? Mi tesoro.
		¿Qué es mi Dios? La libertad.
		¿Mi ley? ¡La fuerza y el viento!
		¿Mi única patria? ¡La mar! `,
	}


	for _, i := range strings {
		s := NewZString(i)
		sz := NewZStringCompressed(i)
		if s.Value() != i {
			t.Errorf("Expecting <%s>, got <%s>", i, s.Value())
		}
		if sz.Value() != i {
			t.Errorf("Expecting <%s>, got <%s>", i, sz.Value())
		}
		if sz.Value() != s.Value() {
			t.Errorf("Compressed and uncompressed values differ <%v> versus <%v>", s.Value(), sz.Value())
		}

		fmt.Println(len(sz.value), len(s.value))
		/*if len(sz.value) > len(s.value) {
			t.Errorf("Bad bussiness. Compresed string is bigger than uncompressed <%d> versus <%d>", len(sz.value), len(s.value))
		}*/
	}
}

func TestMemory(t *testing.T) {
	var lista1 [10000]*ZString
	var lista2 [10000]*ZString
	var m1, m2, m3 runtime.MemStats
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

	runtime.ReadMemStats(&m1)
	for i := 0; i < 10000; i++ {
		lista1[i] = NewZString(text)
	}
	runtime.ReadMemStats(&m2)
	for i := 0; i < 10000; i++ {
		lista2[i] = NewZStringCompressed(text)
	}
	runtime.ReadMemStats(&m3)
	fmt.Printf("Bytes allocated uncompressed: %v\n", m2.Alloc - m1.Alloc)
	fmt.Printf("Bytes allocated compressed: %v\n", m3.Alloc - m2.Alloc)
	fmt.Printf("Memory saved: %d %\n", 100 * (m3.Alloc - m2.Alloc) / (m2.Alloc - m1.Alloc))
}
