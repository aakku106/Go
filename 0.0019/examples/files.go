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
		if thers error, we simpally return it (I only assumed the only eooe we will get is file dont exist in this example)
		and wrirte defer func to close that file,
		that function will called just before return
	*/
	return nil
}

// What we could also do is
func ReadFile2(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}
