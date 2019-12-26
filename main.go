package main

import(
	"os"
	"bufio"
	"strings"
	"engine"
)





func parse(word string) engine.Command{
	if strings.HasPrefix(word,"print"){
		var cmd engine.PrintCommand
			word = strings.TrimSpace(strings.Replace(word,"print","",1))
			cmd.Arg = word
		return cmd
		}else if strings.HasPrefix(word,"reverse"){
			var cmd engine.RevCommand
			word = strings.TrimSpace(strings.Replace(word,"reverse","",1))
			cmd.Arg = word
		return cmd
		}
		return nil
}


func main(){
eventLoop := new(engine.EventLoop)
eventLoop.IsFull = false
eventLoop.WaitNx = make(chan string)
go eventLoop.Start()

if input, err := os.Open("nw1.txt"); err == nil {
 defer input.Close()
 scanner := bufio.NewScanner(input)
 for scanner.Scan() {
 commandLine := scanner.Text()
 cmd := parse(commandLine)
 if cmd!=nil{
 	eventLoop.Post(cmd)
 	eventLoop.WaitNx<-"go rout"
 	}
 }
eventLoop.IsFull = true
close(eventLoop.WaitNx)
}
eventLoop.AwaitFinish()
}