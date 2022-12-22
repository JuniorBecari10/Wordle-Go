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
    color.Red("Usage: wordle <file> [-l length] [-a attempts]\n\n")
    
    color.Cyan("<file>: A file to read the words from.")
    color.Cyan("[-l length]: The desired length of the word.")
    color.Cyan("[-a attempts]: Option to set the allowed number of attempts. Set to -1 for unlimited attempts.")
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
    
    col := color.New(color.FgCyan)
    col.Printf("The word has %d letters.\n\n", len(chosenWord))
    
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
  fmt.Println()
}

func PrintWords() {
  if len(sentWords) > 0 {
    for _, w := range sentWords {
      for ii, c := range w.word {
        col := color.New()
        
        if w.colors[ii] == 0 {
          col = col.Add(color.FgWhite)
        } else if w.colors[ii] == 1 {
          col = col.Add(color.FgYellow)
        } else if w.colors[ii] == 2 {
          col = col.Add(color.FgGreen)
        }
        
        col.Printf("%s ", strings.ToUpper(string(c)))
      }
      
      fmt.Println()
    }
    
    fmt.Println()
  }
}

func SendWord(word string) {
  if len(word) == len(chosenWord) {
    wordAdd := Word { word, make([]byte, len(word)) }
    
    for i, c := range word {
      co := 0
      
      if c == rune(chosenWord[i]) {
        co = 2
      } else if strings.Contains(chosenWord, string(c)) {
        co = 1
      }
      
      wordAdd.colors[i] = byte(co)
    }
    
    sentWords = append(sentWords, wordAdd)
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