package examples

import "os"

func ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	/*
	what we did is opeaned a file, check for error (may be file dont exist)
	*/
	return nil
}
