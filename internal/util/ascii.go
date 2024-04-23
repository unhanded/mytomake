package util

import "fmt"

func Banner() error {
	lines := []string{
		" __   __  __   __  _______  _______  __   __  _______  ___   _  _______",
		"|  |_|  ||  | |  ||       ||       ||  |_|  ||   _   ||   | | ||       |",
		"|       ||  |_|  ||_     _||   _   ||       ||  |_|  ||   |_| ||    ___|",
		"|       ||       |  |   |  |  | |  ||       ||       ||      _||   |___ ",
		"|       ||_     _|  |   |  |  |_|  ||       ||       ||     |_ |    ___|",
		"| ||_|| |  |   |    |   |  |       || ||_|| ||   _   ||    _  ||   |___ ",
		"|_|   |_|  |___|    |___|  |_______||_|   |_||__| |__||___| |_||_______|",
	}

	fmt.Print("\n")

	for _, line := range lines {
		fmt.Print(line + "\n")
	}

	fmt.Print("\n")

	fmt.Print("Fictious data generation\n")
	fmt.Print("---   ---   ---  --  --  --  -- - - -\n\n")

	return nil
}
