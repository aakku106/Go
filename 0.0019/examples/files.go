package examples

import (
	"fmt"
	"io"
	"os"
)

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

/*
This func just read all content of file and simplly prints it on consile/terminal
The thing to notice here is the way i placed defer func between opeaning a file and reading a file

1. Why not before opeaning file,
because opean func may throw errors like file dont exist, closing it make no scence and add up
more error liek ErrInvalid    = errors.New("invalid argument")

2. why before read, why not after:
	Well i see no concreat reasons here,but its good pratice to defer close file just after checking error for file opean

here close function will be called/executed despite of success or falure of readAll func.

*/
