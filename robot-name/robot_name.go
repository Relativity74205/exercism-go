package robotname

import (
	"errors"
	"math/rand"
)

var names = getPossibleNames()

func getPossibleNames() []string {
	var names []string
	for c1 := 'A'; c1 <= 'Z'; c1++ {
		for c2 := 'A'; c2 <= 'Z'; c2++ {
			for c3 := '0'; c3 <= '9'; c3++ {
				for c4 := '0'; c4 <= '9'; c4++ {
					for c5 := '0'; c5 <= '9'; c5++ {
						names = append(names, string(c1)+string(c2)+string(c3)+string(c4)+string(c5))
					}
				}
			}

		}
	}
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	return names
}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() error {
	if len(names) == 0 {
		return errors.New("all names are taken")
	}
	r.name = names[0]
	names = names[1:]
	return nil
}
