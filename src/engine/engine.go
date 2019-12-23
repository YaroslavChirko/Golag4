package engine

import "fmt"

type Command interface {
	Execute(loop Handler)
}


type Handler interface {
	Post(cmd Command)
}

type EventLoop struct{
	mqArr []Command
	IsFull bool 
}

func (hand *EventLoop) Start() {
	hand.mqArr = make([]Command,0)
	for {
		if len(hand.mqArr)>0{
		hand.Pop()
		}else if(len(hand.mqArr)<1&& hand.IsFull == true){
			break
		}
	}
	
}

func (hand *EventLoop) Post(cmd Command) {
	hand.mqArr = append(hand.mqArr,cmd)
}

func (hand *EventLoop) Pop() {
	if(hand.mqArr[0]!=nil){
	hand.mqArr[0].Execute(hand)
	}
		hand.mqArr = hand.mqArr[1:] 
	
}
func (hand *EventLoop) AwaitFinish() {
    for(hand.IsFull ==true&&len(hand.mqArr)>0){}
}

type PrintCommand struct {
	Arg string
}

func (p PrintCommand) Execute(loop Handler) {
	fmt.Println(p.Arg)
}

type RevCommand struct {
	Arg string
}

func (rev RevCommand) Execute(loop Handler) {
 runes := []rune(rev.Arg)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    var tmp PrintCommand
    tmp.Arg = string(runes)
    loop.Post(tmp)

}