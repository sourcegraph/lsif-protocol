package reader

import (
	"strconv"
	"testing"
)

func TestInterner(t *testing.T) {
	testCases := [][][]byte{
		{
			[]byte(`1`),
			[]byte(`2`),
			[]byte(`3`),
			[]byte(`4`),
			[]byte(`5`),
			[]byte(`100`),
			[]byte(`200`),
			[]byte(`300`),
			[]byte(`400`),
			[]byte(`500`),
		},
		{
			[]byte(`"1"`),
			[]byte(`"2"`),
			[]byte(`"3"`),
			[]byte(`"4"`),
			[]byte(`"5"`),
			[]byte(`"100"`),
			[]byte(`"200"`),
			[]byte(`"300"`),
			[]byte(`"400"`),
			[]byte(`"500"`),
		},
		{
			[]byte(`"17f5d4ea-b851-4189-9de7-736002d52d05"`),
			[]byte(`"dc916a1f-c34b-45f0-80ce-e1fc00c019d5"`),
			[]byte(`"46a0ca88-4abc-4180-bc52-3745c3414b6a"`),
			[]byte(`"ae581041-3ed5-444f-8dab-d1e2363cd936"`),
			[]byte(`"da74139d-1403-4e76-b5be-3fe9ce04ecf8"`),
		},
		{
			[]byte(`"rectangle"`),
			[]byte(`"america"`),
			[]byte(`"megaphone"`),
			[]byte(`"monday"`),
			[]byte(`"the next word"`),
		},
	}

	for _, bytes := range testCases {
		interner := NewInterner()
		returned := map[string]int{}

		for _, b := range bytes {
			v, err := interner.Intern(b)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if _, ok := returned[string(b)]; ok {
				t.Fatalf("duplicate id")
			}

			returned[string(b)] = v
		}

		for _, b := range bytes {
			v, err := interner.Intern(b)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if v != returned[string(b)] {
				t.Fatalf("id does not match existing value")
			}
		}
	}
}

// da global variable
var i = 0

func BenchmarkInterner(b *testing.B) {
	/* testCases := [][]byte{
		[]byte(`"rectangle"`),
		[]byte(`"america"`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`"rectangle"`),
		[]byte(`"america"`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`keno`),
		[]byte(`kindergartener`),
		[]byte(`newscasting`),
		[]byte(`malignly`),
		[]byte(`metallophone`),
		[]byte(`pantagraph`),
		[]byte(`remuda`),
		[]byte(`demagogue`),
		[]byte(`immobile`),
		[]byte(`militarised`),
		[]byte(`monument`),
		[]byte(`nonvictory`),
		[]byte(`compossible`),
		[]byte(`valence`),
		[]byte(`dragonlike`),
		[]byte(`ecocide`),
		[]byte(`southwester`),
		[]byte(`fellation`),
		[]byte(`naseby`),
		[]byte(`anglicisation`),
		[]byte(`bacteriostat`),
		[]byte(`bouclï¿¥ï¾½`),
		[]byte(`chemokinetic`),
		[]byte(`unleasable`),
		[]byte(`silviculturist`),
		[]byte(`grishun`),
		[]byte(`fugitiveness`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`keno`),
		[]byte(`kindergartener`),
		[]byte(`newscasting`),
		[]byte(`bouclï¿¥ï¾½`),
		[]byte(`chemokinetic`),
		[]byte(`unleasable`),
		[]byte(`malignly`),
		[]byte(`metallophone`),
		[]byte(`pantagraph`),
		[]byte(`remuda`),
		[]byte(`demagogue`),
		[]byte(`immobile`),
		[]byte(`militarised`),
		[]byte(`monument`),
		[]byte(`nonvictory`),
	} */

	b.StopTimer()
	interner := NewInterner()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//returned := []int{}
			for n := 0; n < 50000000; n++ {
				//for _, by := range testCases {

				str := []byte(`"` + strconv.Itoa(n) + `"`)
				i, _ = interner.Intern(str)

				//returned = append(returned, i)
			}
			for n := 0; n < 50000000; n++ {
				//for _, by := range testCases {

				str := []byte(`"` + strconv.Itoa(n) + `"`)
				i, _ = interner.Intern(str)

				//returned = append(returned, i)
			}
		}
	})
}

func BenchmarkSyncInterner(b *testing.B) {
	/* testCases := [][]byte{
		[]byte(`"rectangle"`),
		[]byte(`"america"`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`"rectangle"`),
		[]byte(`"america"`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`keno`),
		[]byte(`kindergartener`),
		[]byte(`newscasting`),
		[]byte(`malignly`),
		[]byte(`metallophone`),
		[]byte(`pantagraph`),
		[]byte(`remuda`),
		[]byte(`demagogue`),
		[]byte(`immobile`),
		[]byte(`militarised`),
		[]byte(`monument`),
		[]byte(`nonvictory`),
		[]byte(`compossible`),
		[]byte(`valence`),
		[]byte(`dragonlike`),
		[]byte(`ecocide`),
		[]byte(`southwester`),
		[]byte(`fellation`),
		[]byte(`naseby`),
		[]byte(`anglicisation`),
		[]byte(`bacteriostat`),
		[]byte(`bouclï¿¥ï¾½`),
		[]byte(`chemokinetic`),
		[]byte(`unleasable`),
		[]byte(`silviculturist`),
		[]byte(`grishun`),
		[]byte(`fugitiveness`),
		[]byte(`"megaphone"`),
		[]byte(`"monday"`),
		[]byte(`"the next word"`),
		[]byte(`keno`),
		[]byte(`kindergartener`),
		[]byte(`newscasting`),
		[]byte(`bouclï¿¥ï¾½`),
		[]byte(`chemokinetic`),
		[]byte(`unleasable`),
		[]byte(`malignly`),
		[]byte(`metallophone`),
		[]byte(`pantagraph`),
		[]byte(`remuda`),
		[]byte(`demagogue`),
		[]byte(`immobile`),
		[]byte(`militarised`),
		[]byte(`monument`),
		[]byte(`nonvictory`),
	} */

	b.StopTimer()
	interner := NewLockfreeInterner()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//returned := []int{}
			for n := 0; n < 50000000; n++ {
				//for _, by := range testCases {

				str := []byte(`"` + strconv.Itoa(n) + `"`)
				i, _ = interner.Intern(str)

				//returned = append(returned, i)
			}
			for n := 0; n < 50000000; n++ {
				//for _, by := range testCases {

				str := []byte(`"` + strconv.Itoa(n) + `"`)
				i, _ = interner.Intern(str)

				//returned = append(returned, i)
			}
		}
	})
}
