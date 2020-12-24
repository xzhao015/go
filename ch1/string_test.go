package ch1

import (
	"regexp"
	"strings"
	"testing"
)

//ReformatDate is a function.
func ReformatDate(date string) string {
	monthMap := map[string]string{
		"Jan": "1",
		"Feb": "2",
		"Mar": "3",
		"Apr": "4",
		"May": "5",
		"Jun": "6",
		"Jul": "7",
		"Aug": "8",
		"Sep": "9",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12"}
	arr := strings.Fields(date)
	var day, month, year = arr[0], arr[1], arr[2]
	dayReg := regexp.MustCompile("[0-9]{1,2}")
	dayString := dayReg.Find([]byte(day))
	outputList := []string{year, monthMap[month], string(dayString)}
	output := strings.Join(outputList, "-")
	return output
}

func TestString(t *testing.T) {
	string1 := "hello, I am a string"
	string2 := "hello, I am another string"
	//compare return 0/-1/1
	t.Log(strings.Compare(string1, string2))
	t.Log(strings.Compare("a", "a"))
	t.Log(strings.Compare("a", "ab"))
	t.Log(strings.Compare("b", "a"))
	//contains return true/false
	t.Log(strings.Contains("abc", "ab"))
	t.Log(strings.Contains("a", "ab"))
	//ContainsAny return true/false
	t.Log(strings.ContainsAny("team", "i"))
	t.Log(strings.ContainsAny("fail", "ui"))
	t.Log(strings.ContainsAny("ure", "ui"))
	t.Log(strings.ContainsAny("failure", "ui"))
	t.Log(strings.ContainsAny("foo", ""))
	t.Log(strings.ContainsAny("", ""))
	//ContainsRune
	t.Log(strings.ContainsRune("aardvark", 97))
	t.Log(strings.ContainsRune("timeout", 97))
	//Count
	t.Log(strings.Count("cheese", "e"))
	t.Log(strings.Count("five", "iv"))
	t.Log(strings.Count("ababababababa", "ab"))
	//Fields, split the string
	t.Logf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//HasPrefix return true/false
	t.Log(strings.HasPrefix("Gopher", "Go")) //true
	t.Log(strings.HasPrefix("Gopher", "C"))  //false
	t.Log(strings.HasPrefix("Gopher", ""))   //true
	//HasSuffix
	t.Log(strings.HasSuffix("Amigo", "go"))  //true
	t.Log(strings.HasSuffix("Amigo", "O"))   //false
	t.Log(strings.HasSuffix("Amigo", "Ami")) //false
	t.Log(strings.HasSuffix("Amigo", ""))    //true
	//Index return int
	t.Log(strings.Index("chicken", "ken")) //4
	t.Log(strings.Index("chicken", "dmr")) //-1
	//IndexByte return int
	t.Log(strings.IndexByte("golang", 'g'))  //0
	t.Log(strings.IndexByte("gophers", 'h')) //3
	t.Log(strings.IndexByte("golang", 'x'))  //-1
	//Join return sring
	s := []string{"foo", "bar", "baz"}
	t.Log(strings.Join(s, ", ")) //foo, bar, baz
	//LastIndex
	t.Log(strings.Index("go gopher", "go"))         //0
	t.Log(strings.LastIndex("go gopher", "go"))     //3
	t.Log(strings.LastIndex("go gopher", "rodent")) //-1
	//map
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	t.Log(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	//Repeat
	t.Log("ba" + strings.Repeat("na", 2))
	//Replace
	t.Log(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	t.Log(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
	//ReplaceAll
	t.Log(strings.ReplaceAll("oink oink oink", "oink", "moo")) //moo moo moo
	//Split return []string
	t.Logf("%q\n", strings.Split("a,b,c", ",")) //[]string{a, b, c}
	t.Logf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	t.Logf("%q\n", strings.Split(" xyz ", ""))
	t.Logf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//SplitAfter
	t.Logf("%q\n", strings.SplitAfter("a,b,c", ",")) //["a," "b," "c"]
	//SplitAfterN
	t.Logf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) //["a," "b,c"]
	//SplitN
	t.Logf("%q\n", strings.SplitN("a,b,c", ",", 2)) //["a" "b,c"]
	t.Logf("%q\n", strings.SplitN("a,b,c", ",", 1)) // ["a,b,c"]
	t.Log(strings.SplitN("a,b,c", ",", 0))          //[]
	//Title
	t.Logf(strings.Title("her royal highness")) //Her Royal Highness
	t.Logf(strings.Title("loud noises"))
	t.Logf(strings.Title("хлеб"))
	//ToTitle
	t.Logf(strings.ToTitle("her royal highness")) //HER ROYAL HIGHNESS
	t.Logf(strings.ToTitle("loud noises"))
	t.Logf(strings.ToTitle("хлеб"))
	//ToLower
	t.Log(strings.ToLower("Gopher")) //gopher
	//ToUpper
	t.Log(strings.ToUpper("Gopher")) //GOPHER
	//Trim like strip() in python
	t.Log(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))      //Hello, Gophers
	t.Log(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))  //Hello, Gophers!!!
	t.Log(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) //¡¡¡Hello, Gophers
	s1 := "¡¡¡Hello, Gophers!!!"
	s1 = strings.TrimPrefix(s1, "¡¡¡Hello, ")
	s1 = strings.TrimPrefix(s1, "¡¡¡Howdy, ")
	t.Log(s1) //Gophers!!!

	t.Log(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) //Hello, Gophers
	s1 = "¡¡¡Hello, Gophers!!!"
	s1 = strings.TrimSuffix(s1, ", Gophers!!!")
	s1 = strings.TrimSuffix(s1, ", Marmots!!!")
	t.Log(s1) //¡¡¡Hello
}
