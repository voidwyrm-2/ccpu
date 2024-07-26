package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
)

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func writeBytesFile(filename string, data []byte) error {
	// Open the file with write permissions, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the bytes to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

var showTokens bool = false

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'casm [files...]'")
		return
	}

	if ind := slices.Index(os.Args, "-h"); ind != -1 {
		fmt.Println("usage: casm [files...] [-o [path]] [-t]")
		return
	} else if ind := slices.Index(os.Args, "--help"); ind != -1 {
		fmt.Println("usage: casm [files...] [-o [path]] [-t]")
		return
	}

	os.Args = os.Args[1:]

	outFile := ""

	if ind := slices.Index(os.Args, "-o"); ind != -1 {
		if ind == len(os.Args)-1 {
			fmt.Println("expected '-o [output filename]'")
			return
		}
		outFile = os.Args[ind+1]
		os.Args = slices.Delete(os.Args, ind, ind+2)
	} else {
		outFile = os.Args[0]
	}

	if ext := path.Ext(outFile); ext != ".bin" {
		outFile = outFile[:len(outFile)-len(ext)] + ".bin"
	}

	if ind := slices.Index(os.Args, "-t"); ind != -1 {
		showTokens = true
		os.Args = slices.Delete(os.Args, ind, ind+1)
	}

	srcAcc := ""
	for _, f := range os.Args {
		content, err := readFile(f)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if srcAcc == "" {
			srcAcc += content
		} else {
			srcAcc += "\n" + content
		}
	}

	lexer := NewLexer(srcAcc)
	tokens, lexerErr := lexer.Lex()
	if lexerErr != nil {
		fmt.Println(lexerErr.Error())
		return
	}
	if showTokens {
		for _, t := range tokens {
			fmt.Println(t.Tostr())
		}
	}

	bytes, interpreterErr := interpret(tokens)
	if interpreterErr != nil {
		fmt.Println(interpreterErr.Error())
		return
	}

	writeBytesFile(outFile, bytes)
}
