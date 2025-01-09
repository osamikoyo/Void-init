package components

import "os"

const MainGoData = `
package main

import "github.com/osamikoyo/void"

func main(){
	cli := void.NewCLI()
}

`

func WriteMaingo(path string) error {
	file, err := os.Open(path)
	if err != nil{
		return err
	}

	if _, err = file.WriteString(MainGoData);err != nil{
		return err
	}

	return nil
}