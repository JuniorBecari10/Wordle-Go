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
  "strconv"
)

const (
  gray byte = 0
  yellow byte = 1
  green byte = 2
  
  attemptsDefault int = 6
)

var (
  words []string
  sentWords []Word
  chosenWord string
  scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
  
  length int = 0
  attempts int = attemptsDefault
  dict bool = true
  showDict bool = false
)

// 0 - (gray - isn't in the word),
// 1 - (yellow - is in the word, but in the wrong place),
// 2 - (green - in the right place)
type Word struct {
  word string
  colors []byte
}

func main() {
  if len(os.Args) < 2 {
    color.Red("Usage: wordle <file> [-l length] [-a attempts] [-d]\n\n")
    
    color.Cyan("<file>: A file to read the words from.")
    color.Cyan("[-l length]: The desired length of the word.")
    color.Cyan("[-a attempts]: Option to set the allowed number of attempts. Set to -1 for unlimited attempts.")
    color.Cyan("[-d]: Option to define restriction to only words of the dictionary (inside <file>).")
    os.Exit(0)
  }
  
  ReadArgs()
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
    
    color.Yellow("Hint: If you want to exit, type !exit, and then confirm.\n")
    
    if dict && showDict {
      colr := color.New(color.FgRed)
      colr.Printf("The word %s is not in the dictionary.\n", scanner.Text())
    }
    
    fmt.Println()
    
    col := color.New(color.FgCyan)
    
    if len(chosenWord) > 1 {
      col.Printf("The word has %d letters.\n", len(chosenWord))
    } else {
      col.Printf("The word has %d letter.\n", len(chosenWord))
    }
    
    if attempts > 0 {
      if attempts > 1 {
        col.Printf("You have %d remaining attempts.\n", attempts)
      } else {
        col.Printf("You have %d remaining attempt.\n", attempts)
      }
    } else if attempts == 0 {
      col.Printf("Your attempts have ended.\n")
    } else {
      col.Printf("You have unlimited attempts.\n")
    }
    
    if !dict {
      col.Println("You can type words outside the dictionary.")
    }
    
    if attempts == 0 {
      fmt.Println()
      PrintWords()
      
      color.Red("Oh no.")
      color.Red("The word was %s.", chosenWord)
      
      os.Exit(0)
    }
    
    fmt.Println()
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
  //color.Green("#   #  ###  ###  ##   #    ###")
  //color.Green("#   #  # #  # #  # #  #    #")
  //color.Green("# # #  # #  ##   # #  #    ##")
  //color.Green("## ##  # #  # #  # #  #    #")
  //color.Green("#   #  ###  # #  ##   ###  ###")
  
  color.Green(" __          ______  _____  _____  _      ______ ")
  color.Green(" \\ \\        / / __ \\|  __ \\|  __ \\| |    |  ____|")
  color.Green("  \\ \\  /\\  / / |  | | |__) | |  | | |    | |__   ")
  color.Green("   \\ \\/  \\/ /| |  | |  _  /| |  | | |    |  __|  ")
  color.Green("    \\  /\\  / | |__| | | \\ \\| |__| | |____| |____ ")
  color.Green("     \\/  \\/   \\____/|_|  \\_\\_____/|______|______|")
  
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
  showDict = false
  
  if word == "!exit" {
    col := color.New(color.FgRed)
    col.Printf("Are you sure? (y/n) ")
    
    scanner.Scan()
    ans := scanner.Text()
    
    if strings.ToLower(ans) == "y" {
      col = color.New(color.FgCyan)
      col.Printf("The word was %s.", chosenWord)
      
      os.Exit(0)
    }
  }
  
  if len(word) == len(chosenWord) {
    if dict && !ContainsArray(words, word) {
      showDict = true
      return
    }
    
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
    
    if attempts > 0 {
      attempts--
    }
  }
}

func ChooseWord() {
  chosenWord = words[rand.Intn(len(words))]
  
  if length != 0 {
    for len(chosenWord) != length {
      chosenWord = words[rand.Intn(len(words))]
    }
  }
}

func Verify(word string) bool {
  return strings.ToLower(word) == strings.ToLower(chosenWord)
}

func ReadArgs() {
  for i, arg := range os.Args {
    if i == 0 {
      continue
    }
    
    a := strings.ToLower(string(arg))
    
    if i < len(os.Args) - 1 {
      if a == "-l" {
        ll, err := strconv.Atoi(os.Args[i + 1])
        
        if err != nil || ll < 1 {
          color.Red("Invalid length.")
          os.Exit(0)
        }
        
        length = ll
      } else if a == "-a" {
        aa, err := strconv.Atoi(os.Args[i + 1])
        
        if err != nil || aa == 0 {
          color.Red("Invalid number of attempts.")
          os.Exit(0)
        }
        
        attempts = aa
      }
    }
    
    if a == "-d" {
      dict = false
    }
  }
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

func ContainsArray(arr []string, s string) bool {
  for _, v := range arr {
    if strings.ToLower(v) == strings.ToLower(s) {
      return true
    }
  }
  
  return false
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