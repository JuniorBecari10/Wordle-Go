# Wordlen

A recreation of the game Wordle in Go.

## How to play

The game will choose a random word from the specified file and you have to guess it. <br />
You can start typing any word of the same length as the chosen one (the game will tell you) <br />
And as you type, the word you typed will be colored as:

### Colors

`Green` - The letter is in the correct place.
`Yellow` - The letter is in the word, but in another place.
`White` - The letter isn't in the word.

### Options

You can launch the game with some options too.

`<file>` **Obligatory.** Here you specify the file containing the words for the game (the dictionary). This must be the first argument.
`[-l length]` **Optional.** Here you can specify the length of the chosen word instead of a random length.
`[-a attempts]` **Optional.** Here you can specify the amount of attempts you have. If you type a negative number, you can have unlimited attempts! But you can't type `0` here. **Default value: 6.**
`[-d]` **Optional.** Flag to specify if you can type only words of the dictionary or not. Typing this will disable the option. **Default value: true.**

#### Examples:

```
wordlen words/english.txt -l 5 -a 10
```

This command will launch the game with the file `words/english.txt` as dictionary, with the word of 5 letters and 10 attempts.

```
wordlen words/portugues.txt -a -1 -d
```

This command will launch the game with the file `words/portugues.txt` as dictionary, with random length, unlimited attempts and you can type words outside the dictionary.

> Note: you can swap the arguments, so you can specify the attempts before the length and vice versa.

### Exit

If you want to exit the game, but you want to see the chosen word, type `!exit` while playing the game. <br />
You will be prompted if you are sure, if you are, type `y`, and then the game will close.