
# Text Modifier Tool
This tool was developed using the Go programming language and is designed to modify text files according to certain rules. Features included include various changes such as converting hexadecimal and binary numbers to decimal, converting text to uppercase, lowercase, or initials to uppercase, editing punctuation marks, and managing apostrophes.

## Features
● Hexadecimal and Binary Conversion: Converts hexadecimal and binary numbers to decimal equivalents.

● Text Case Modification: Change text to uppercase, lowercase, or capitalize specified words or portions of text.

● Punctuation Handling: Ensure correct placement of punctuation marks and handle groups of punctuation appropriately.

● Apostrophe Placement: Correctly place apostrophes around words enclosed in single quotation marks.

● Indefinite Article Correction: Automatically change the indefinite article 'a' to 'an' when preceding words starting with a vowel or 'h'.

## Usage

```bash
$ go run . sample.txt result.txt
```
- The tool takes an input file (sample.txt) containing the text to be modified and an output file (result.txt) where the modified text will be saved.

#### Example

Input (sample.txt)
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
Output (result.txt)
```
Simply add 66 and 2 and you will see the result is 68.
```
Input (sample.txt)
```
Punctuation tests are ... kinda boring ,don't you think !?
```
Output (result.txt)
```
Punctuation tests are... kinda boring, don't you think!?
```

## How It Works
The tool parses the input text file, applies the specified modifications, and writes the modified text to the output file. It utilizes Go's file system (fs) API for file operations and string manipulation functions for text modifications.

#### Author
[Ahat Demirezen](https://github.com/ahatdemirezen)

## License

[MIT](https://choosealicense.com/licenses/mit/)




