type Node struct{
	Content string
	Next *node
}

func PrintList (first *Node){
	var curs *Node
	curs = first
	for curs != nil{
		
		fmt.Println(*curs.Content) // in C: curs->Content
		curs = curs.Next
	}
}

func main(){
	var frst, seconf, third *Noode
	first = new(Node)
	second = new(Node)
	third = new(Node)
	first.Content = "Primo"
	second.Content = "Secondo"
	third.Content = "Terzo"
	first.Next = second
	third.Next = nil
}