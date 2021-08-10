package memory

type pointer int

type octet [8]int

type MemoryBlock struct {
	address   pointer
	container octet
	page      int
}

type MemoryPage struct {
	blocks  []MemoryBlock
	pageNum int
}

type Stack struct {
	container    [256]MemoryBlock
	currentIndex int
}

func (stack *Stack) PushStack(m MemoryBlock) {
	stack.container[stack.currentIndex] = m
	stack.currentIndex += 1
	if stack.currentIndex == 256 {
		stack.currentIndex = 0
	}
}

func (stack *Stack) PullStack() (x MemoryBlock) {
	ret := stack.container[stack.currentIndex]
	stack.currentIndex -= 1

	if stack.currentIndex == 0 {
		stack.currentIndex = 255
	}

	return ret
}

func (memBlock *MemoryBlock) ArithmeticShiftLeft() {
	memBlock.container[7] = memBlock.container[0]
	for i := 1; i < 7; i++ {
		memBlock.container[i-1] = memBlock.container[i]
	}

}

func (memBlock *MemoryBlock) ArithmeticShiftRight() {
	for i := 0; i < 6; i++ {
		memBlock.container[i+1] = memBlock.container[i]
	}
	memBlock.container[0] = memBlock.container[7]

}
