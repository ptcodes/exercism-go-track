package beer

import (
	"errors"
	"fmt"
)

func Verse(n int) (string, error) {
	if n > 99 {
		return "", errors.New("must be less or equal to 99")
	}
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
}

func Verses(hi, lo int) (string, error) {
	if hi > 99 || lo < 0 || hi < lo {
		return "", errors.New("must be within a correct range")
	}
	verses := ""
	for i := hi; i >= lo; i-- {
		verse, _ := Verse(i)
		verses += verse + "\n"
	}
	return verses, nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
