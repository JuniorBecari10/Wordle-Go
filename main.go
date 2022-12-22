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
  "time"
)

var (
  words []string
  sentWords []Word
  chosenWord string
  scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
)

const (
  gray byte = 0
  yellow byte = 1
  green byte = 2
)

// 0 - (gray - isn't in the word),
// 1 - (yellow - is in the word, but in the wrong place),
// 2 - (green - in the right place)
type Word struct {
  word string
  colors []byte
}

func main() {
  if len(os.Args) != 2 {
    color.Red("Usage: wordle <file>\n\n")
    
    color.Cyan("<file>: A file to read the words from.")
    os.Exit(0)
  }
  
  LoadWords(os.Args[1])
  
  rand.Seed(time.Now().UnixNano())
  ChooseWord()
  RunGame()
}

// --- Functions --- //

func RunGame() {
  for {
    Clear()
    PrintLogo()
    
    PrintWords()
    
    if Verify(scanner.Text()) {
      color.Green("Yeah! You hit the correct word!")
      os.Exit(0)
    }
    
    fmt.Printf("> ");
    scanner.Scan()
    word := scanner.Text()
    
    SendWord(word)
  }
}

func PrintLogo() {
  color.Green("#   #  ###  ###  ##   #    ###")
  color.Green("#   #  # #  # #  # #  #    #")
  color.Green("# # #  # #  ##   # #  #    ##")
  color.Green("## ##  # #  # #  # #  #    #")
  color.Green("#   #  ###  # #  ##   ###  ###")
}

func PrintWords() {
  
}

func SendWord(word string) {
  if len(word) == len(chosenWord) {
    wordAdd = Word{word, 0}
    
    
    
    sentWords = append(sentWords, word)
  }
}

func ChooseWord() {
  chosenWord = words[rand.Intn(len(words))]
}

func Verify(word string) bool {
  return word == chosenWord
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