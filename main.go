package main

import (
  "fmt"
  "os"
  "os/exec"
  "github.com/fatih/color"
  "runtime"
  "strings"
  "bufio"
  "math/rand"
)

var (
  words []string
  sentWords []string
  chosenWord string
  scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
  if len(os.Args) != 2 {
    color.Red("Usage: wordle <file>\n\n")
    
    color.Cyan("<file>: A file to read the words from.")
    os.Exit(0)
  }
  
  LoadWords(os.Args[1])
  RunGame()
}

// --- Functions --- //

func RunGame() {
  ChooseWord()
  
  for {
    Clear()
    color.Green("  W O R D L E\n\n")
    
    PrintWords()
    
    fmt.Printf("> ");
    scanner.Scan()
    word := scanner.Text()
    
    SendWord(word)
  }
}

func PrintWords() {
  
}

func SendWord(word string) {
  if len(word)
}

func ChooseWord() {
  
}

func Clear() {
  if (runtime.GOOS == "windows") {
    cmd := exec.Command("cmd", "/c", "cls")
    
    cmd.Stdout = os.Stdout
    cmd.Run()
  } else {
    cmd := exec.Command("clear")
    
    cmd.Stdout = os.Stdout
    cmd.Run()
  }
}

func LoadWords(file string) {
  b, err := os.ReadFile(file)
  
  if err != nil {
    color.Red("An error occurred while reading the file.")
    color.Red("Perhaps the file doesn't exist.")
    os.Exit(0)
  }
  
  str := string(b)
  words = strings.Split(str, " ")
}