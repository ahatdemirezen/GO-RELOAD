package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func hex(str string) string {
	values, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(values))
}

func bin(str string) string {
	values, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(values))
}

func exchange(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

// DuzenleMetin fonksiyonu, metni düzenler ve istenen kurallara göre noktalama işaretlerini ayarlar.
func noktalama(deger string) string {
	bos := []string{",", ".", "!", "?", ";", ":", "'", "\"", "(", ")", "-"}

	for _, k := range bos {
		deger = strings.Replace(deger, " "+k, k, -1)
		deger = strings.Replace(deger, k+" ", k, -1)
	}
	r := []rune(deger)
	dizi := []rune{}
	for i := 0; i < len(r)-1; i++ {
		if unicode.IsPunct(r[i]) && unicode.IsLetter(r[i+1]) {
			if r[i] == '\'' && unicode.IsLetter(r[i+1]) || r[i] == '-' && unicode.IsLetter(r[i+1]) || r[i] == '"' {
				dizi = append(dizi, r[i])
			} else {
				dizi = append(dizi, r[i], ' ')
			}
		} else if unicode.IsPunct(r[i]) && (r[i+1] == '\'' || (r[i+1]) == '"') {
			dizi = append(dizi, r[i], ' ')
		} else {
			dizi = append(dizi, r[i])
		}
	}
	dizi = append(dizi, r[len(r)-1])

	return string(dizi)
}

func main() {
	fileName := os.Args[1]
	toCreate := os.Args[2]

	// Dosyayı oku
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Dosyanın içeriğini satırlara ayır
	lines := strings.Split(string(content), "\n")
	var resultLines []string

	// satırları işle
	for i := 0; i <= len(lines)-1; i++ {
		line := lines[i]

		// Dosyanın içeriğini kelimelere ayır
		words := strings.Fields(line)
		res := ""

		words = exchange(words)

		// Kelimeleri tersten işle
		// harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat. '
		for j := len(words) - 1; j >= 0; j-- {
			if words[j] == "(up)" {
				words[j-1] = strings.ToUpper(words[j-1])
				j = j - 1
			} else if strings.Contains(words[j], ")") && words[j-1] == "(up," {
				nStr := strings.Trim(words[j], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[j-n+a-1] = strings.ToUpper(words[j-n+a-1])
				}
				j = j - 2
			} else if words[j] == "(low)" {
				words[j-1] = strings.ToLower(words[j-1])
				j = j - 1
			} else if strings.Contains(words[j], ")") && words[j-1] == "(low," {
				nStr := strings.Trim(words[j], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[j-n+a-1] = strings.ToLower(words[j-n+a-1])
				}
				j = j - 2
			} else if words[j] == "(cap)" {
				words[j-1] = strings.Title(words[j-1])
				j = j - 1
			} else if strings.Contains(words[j], ")") && words[j-1] == "(cap," {
				nStr := strings.Trim(words[j], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[j-n+a-1] = strings.Title(words[j-n+a-1])
				}
				j = j - 2
			} else if words[j] == "(hex)" {
				words[j-1] = hex(words[j-1])
				j = j - 1
			} else if strings.Contains(words[j], ")") && words[j-1] == "(hex," {
				nStr := strings.Trim(words[j], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[j-n+a-1] = hex(words[j-n+a-1])
				}
				j = j - 2
			} else if words[j] == "(bin)" {
				words[j-1] = bin(words[j-1])
				j = j - 1
			} else if strings.Contains(words[j], ")") && words[j-1] == "(bin," {
				nStr := strings.Trim(words[j], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[j-n+a-1] = bin(words[j-n+a-1])
				}
				j = j - 2
			}
			res = words[j] + " " + res
			if res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}
		}

		// Noktalama işlemleri ve düzenlemeler
		res = noktalama(res)
		resultLines = append(resultLines, res)
	}

	file, err := os.Create(toCreate)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(strings.Join(resultLines, "\n"))
	if err != nil {
		log.Fatal(err)
	}
}
