package collatzconjecture

import "errors"

func CollatzConjecture(input int) (output int, err error) {
	for i := 0; input >= 1; i++ {

		if input == 1 {
			return i, nil
		}
		if input%2 == 0 {
			input = input / 2

		} else {
			input = input*3 + 1
		}

	}
	return 0, errors.New("native value is error")
}
