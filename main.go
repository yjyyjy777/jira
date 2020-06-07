package main

import (
	"bufio"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"os"
	"strings"
)

var EI []string

func file() {
	fileName := "./1.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		EI = append(EI, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
	}
}

func main() {
	file()
	tp := jira.BasicAuthTransport{
		Username: "yangjinyuan",
		Password: "xxxxx",
	}

	client, _ := jira.NewClient(tp.Client(), "https://support.nqsky.com")
	for i := 0; i < len(EI); i++ {
		issue, _, _ := client.Issue.Get(EI[i], nil)
		//fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
		//fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
		//fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
		//fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
		fmt.Printf("%s: 状态: %s\n", issue.Key, issue.Fields.Status.Name)

	}

}
