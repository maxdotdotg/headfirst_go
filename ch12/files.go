package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/* not final version, but useful for rememering the call stack
func scanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// errors get propogated up the call chain (or dir tree, in this case)
		fmt.Printf("returning err from scanDir(\"%s\") call\n", path)
		return err
	}

	for _, file := range files {
		// make the file path
		filePath := filepath.Join(path, file.Name())

		// do that recursion, yo
		// if it's a dir, start from the top
		if file.IsDir() {
			err := scanDir(filePath)
			if err != nil {
				fmt.Printf("returning err from scanDir(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	// complete successfully, return err == nil
	return nil
}
*/

// rewritten scanDir using panic to avoid walking up the call stack
func scanDir(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDir(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func reportPanic() {
	p := recover()

	// if recover returns nil, there's no panic to catch
	if p == nil {
		return
	}

	// catch the error thrown in panic and print it
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// if the panic value isn't an error,
		// resume panic-ing with the same value?
		panic(p)
	}
}

func main() {
	defer reportPanic()
	panic("some other err")
	scanDir("/usr/local/bin/")
}
