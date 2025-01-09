package components

import "os"

const RootData = `
package cmd



`

func WriteRoot(path string) error {
	file, err := os.Open(path)
	if err != nil{
		return err
	}

	if _, err = file.WriteString(RootData);err != nil{
		return err
	}

	return nil
}