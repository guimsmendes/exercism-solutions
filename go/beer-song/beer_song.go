package beer

import (
	"errors"
	"fmt"
)

const lastLine = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
const lastBeer = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
const singleBeer = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"

func Verse(i int) (string, error) {

	switch {
	case i < 0 || i > 99:
		return "", errors.New("Out of range.")
	case i == 0:
		return lastLine, nil
	case i == 1:
		return lastBeer, nil
	case i == 2:
		return singleBeer, nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", i, i, i-1), nil
	}

}

func Verses(start int, end int) (out string, err error) {
	if start < end {
		err = errors.New("Out of range.")
		return
	}
	for i := start; i >= end; i-- {
		line, e := Verse(i)
		if e != nil {
			err = e
			return
		}
		out += line + "\n"
	}
	return
}

func Song() string {
	out, _ := Verses(99, 0)
	return out
}
