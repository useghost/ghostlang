package main

import (
	"fmt"
	"os"
	"ghostlang/encoding"
)

func writeBytesToFile(filePath string, bytes []byte) error {
	file, err := os.Create(filePath);

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes);
	if err != nil {
		return err
	}
	return nil
}

func readBytesFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath);
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 4)
	_, err = file.Read(buf);
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func main() {
	bytes := []byte{255,255,200, 244};
	err := writeBytesToFile("output.txt", bytes);
	if err != nil {
		fmt.Println("error")
	}

	readBytes, err := readBytesFromFile("output.txt")
	fmt.Println("read as:", readBytes, encoding.ReadI32LE(readBytes))
}