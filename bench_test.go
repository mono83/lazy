package lazy

import (
	"fmt"
	"testing"
)

func supply() int { return 1 }

func BenchmarkConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Const(1)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(supply)
	}
}

func BenchmarkAll(b *testing.B) {
	funcs := []string{"const", "new"}
	reads := []int{1, 100, 1000, 10000}

	for _, cnt := range reads {
		for _, funcName := range funcs {
			var producer func() func() int
			switch funcName {
			case "const":
				producer = func() func() int { return Const(10) }
			case "new":
				producer = func() func() int { return New(supply) }
			default:
				panic(funcName)
			}
			l := producer()
			b.Run(
				fmt.Sprintf("Read-%s-%dtimes", funcName, cnt),
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						for j := 0; j < cnt; j++ {
							l()
						}
					}
				},
			)
		}
	}
}
