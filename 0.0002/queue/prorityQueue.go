package main

import (
	"fmt"
	"os"
)

func prorityQueue() {
	fmt.Println("Accessing Protity Queue.....")
	fmt.Println("Protuty Queue Accesed")
	for {
		fmt.Println("Choose Between")
		fmt.Println("1.To Enqueue")
		fmt.Println("2. To Deque")
		fmt.Println("3. Peek")
		fmt.Println("-1 To back to Queue")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case -1:
			return
		case quit:
			os.Exit(0)
		case enque:
			enquedValue, prorityOfEnqueuedValue, wasEnqueueSuccess := EnqueueInProrityQueue()
			if wasEnqueueSuccess {
				fmt.Printf("-->\t%d\tEnqueued with prority:\t%d\n", enquedValue, prorityOfEnqueuedValue)
			}
		case deque:
			dequeuedValue, prorityOfDequeuedValue, wasProrityQueueEmpty := DequeueFromProrityQueue()
			if wasProrityQueueEmpty {
				fmt.Println("The Prority Queue was Empty!!!")
			} else {
				fmt.Printf("-->\t%d\tEnqueued with prority:\t%d\n", dequeuedValue, prorityOfDequeuedValue)
			}
		case peek:
			peekIntoProrityQueue()
		default:
			fmt.Println("Choose between 1,2,3 or 0 to exit program")
		}
	}
}
func EnqueueInProrityQueue() (int, int8, bool) {
a:
	fmt.Println("Enter the value to Enqueue in Prority Queue")
	fmt.Println("In format: value,prority")
	fmt.Println("Note: Prority shall be from 0 to 4, lower the value, higher the prority")

	// fmt.Printf("\tlen:\t%d\tcap:\t%d\n", len(prorityQueueStoragePlace), cap(prorityQueueStoragePlace))
	// fmt.Printf("\t%d\n", (len(prorityQueueStoragePlace[0]) + len(prorityQueueStoragePlace[1]) + len(prorityQueueStoragePlace[2]) + len(prorityQueueStoragePlace[3]) + len(prorityQueueStoragePlace[4])))
	var (
		value   int
		prority int8
	)
	fmt.Scanf("%d,%d", &value, &prority)
	if prority >= 5 || prority < 0 {
		fmt.Println("choose Between 0,1,2,3,4")
		goto a
	}
	prorityQueueStoragePlace[prority] = append(prorityQueueStoragePlace[prority], value)

	//	switch prority {
	//	case 0:
	//		prorityQueueStoragePlace[0] = append(prorityQueueStoragePlace[0], value)
	//	case 1:
	//		prorityQueueStoragePlace[1] = append(prorityQueueStoragePlace[1], value)
	//	case 2:
	//		prorityQueueStoragePlace[2] = append(prorityQueueStoragePlace[2], value)
	//	case 3:
	//		prorityQueueStoragePlace[3] = append(prorityQueueStoragePlace[3], value)
	//	case 4:
	//		prorityQueueStoragePlace[4] = append(prorityQueueStoragePlace[4], value)
	//	default:
	//		fmt.Println("Cloose between: 0,1,2,3,4")
	//		return 0, 0, false
	//	}

	//	fmt.Printf("\n%v\n", prorityQueueStoragePlace)

	return value, prority, true
}
func DequeueFromProrityQueue() (int, int8, bool) {
	if isProrityQueueEmpty() {
		return 0, 0, true
	}
	for i := range 5 {
		if len(prorityQueueStoragePlace[i]) != 0 {
			value := prorityQueueStoragePlace[i][0]
			prorityQueueStoragePlace[i] =
				prorityQueueStoragePlace[i][1:]
			return value, int8(i), false
		}
	}
	return 0, 0, true
}
func peekIntoProrityQueue() {
	if isProrityQueueEmpty() {
		fmt.Println("The Prority Queue is Empty")
		return
	}
	for i := range 5 {
		if len(prorityQueueStoragePlace[i]) != 0 {
			fmt.Printf("---->\t%d\t<---- cumming out next\n", prorityQueueStoragePlace[i][0])
			return
		}
	}
}
func isProrityQueueEmpty() bool {
	if (len(prorityQueueStoragePlace[0]) + len(prorityQueueStoragePlace[1]) + len(prorityQueueStoragePlace[2]) + len(prorityQueueStoragePlace[3]) + len(prorityQueueStoragePlace[4])) <= 0 {
		return true
	}
	return false
}

/*
Lest see how we make  prority queue:
1. 1st it shall follow fifo rele: first in 1st out,which means i have to 1st make a normal linear queue(not circular queue cause i have a idea to use slice like a ummm some shord of garbage collector thing like, when lest say slice is larger than 50, ie. the front is larger than 50 than only do [front:], and front=0; i think it should be ~ O(1)may be lest cee)
2. i need a prority, so lest decleare a var called prority, where we can give values from 0 to 4, lower the value higher the prority, I think this is good idea for now



the main problem here is how do i manage the prority,
i have to store value & prority in same index of array, which isen't possible(without usign may be like struct), we have ot use linkedlist, which i dont want now,
so i have 2 ideas to implement this prorityQueue
1. make 2 array, one to store value and another to store prority.
2. make 5 array, 0,1,2,3,4 and store value in respective prority arrays accorting to their prority value

i kind of like option no 2, cause th elogic will be cleaner and easier
cause in option no 1, i guss we have to short array in basis of prority each time, & aslso we have to read values from 2 arry, make the logic to shor the array with values and short prority array accordingly will that be O(n) or n^2 i wonder,

But i am still unsure about my 2 iseas and i think 2nd idea will take more space, but i belive we can do operation on O(1) in that queue

my idea of ques in2 in option

var(
 zero []int
one []int
.
.
four []int
)

but thers a another way in Go to do same thing but in more cleanear and scalable way
prorityQueue[i][] int
 this way we can change the value i as our linkin gin our case it will be 5.

and another benifit will be fast memory access, cause way 1 all array or rather say slice are store in random place in memory, and in way 2 they are stored continously
when we have 5 array it dosent really matter, but if we have 5millon array than it will matter, it will be chasing friendly , cpu have to do less work less seeking time in memory
and if we are making games, high frequency trading or kernal sheduler than it really massters
it dosent matter in our case, but will i will use way 2

### now you might think one thing,i told i will make it like garbacecollection and clear dequeuede eleent menually, but i used inbuild go slices insted of front, because in go slice dont oun array its pointer to array with length and capacity, so its basically same as i told earlier, but still i will emplement that gc like behavior in prorityQueue.go in 0.0003 folder
*/
