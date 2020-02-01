package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Serial struct {
	Url  string
	Name string
}

func SerialArray(fileName string) []Serial {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error!!!\n___________________ File \"url.txt\" not found\n%s\n", err)
	}

	serialLine := strings.Split(string(file), "\r\n")
	serials := make([]Serial, 0)

	for i := 0; i < len(serialLine); i += 2 {
		serials = append(serials, Serial{Url: serialLine[i], Name: serialLine[i+1]})
	}

	return serials
}

func Exit() {
	fmt.Println("Press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}

func main() {

	serials := SerialArray("url.txt")

	for i := 0; i < len(serials); i++ {

		resp, err := http.Get(serials[i].Url)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var regMask *regexp.Regexp
		var serialStr []string
		// it's mask for new seria from site https://www.toramp.com
		regMask = regexp.MustCompile(`<td id="not-air" class="air-date"><span title="[A-Za-zА-Яа-я0-9 ]+">[0-9а-я ]+`)

		serialStr = strings.Split(regMask.FindString(string([]byte(body))), ">")

		fmt.Printf("Следующая серия сериала \"%v\" выйдет %v\n", serials[i].Name, serialStr[2])
	}

	Exit()
}
