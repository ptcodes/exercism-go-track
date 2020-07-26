package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var nameCache = make(map[string]bool)
var nameLimit = 26 * 26 * 10 * 10 * 10
var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func letter() string {
	return string(alphabet[rand.Intn(26)])
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s%03d", letter(), letter(), rand.Intn(1000))
}

func (r *Robot) Name() (string, error) {
	if len(nameCache) == nameLimit {
		return "", errors.New("nameCache has reached its maximum capacity")
	}
	if r.name == "" {
		for {
			r.name = generateName()
			if !nameCache[r.name] {
				nameCache[r.name] = true
				break
			}
		}
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
