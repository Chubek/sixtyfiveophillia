package memory

type Pointer int

type Octet [8]int

type MemoryBlock struct {
	Address   Pointer
	Container Octet
	Page      int
}

type MemoryPage struct {
	Blocks  []MemoryBlock
	PageNum int
}

type Stack struct {
	Container    [256]MemoryBlock
	CurrentIndex int
}

func (stack *Stack) PushStack(m MemoryBlock) {
	stack.Container[stack.CurrentIndex] = m
	stack.CurrentIndex += 1
	if stack.CurrentIndex == 256 {
		stack.CurrentIndex = 0
	}
}

func (stack *Stack) PullStack() (x MemoryBlock) {
	ret := stack.Container[stack.CurrentIndex]
	stack.CurrentIndex -= 1

	if stack.CurrentIndex == 0 {
		stack.CurrentIndex = 255
	}

	return ret
}

func (memBlock *MemoryBlock) ArithmeticShiftLeft() {
	memBlock.Container[7] = memBlock.Container[0]
	for i := 1; i < 7; i++ {
		memBlock.Container[i-1] = memBlock.Container[i]
	}

}

func (memBlock *MemoryBlock) ArithmeticShiftRight() {
	for i := 0; i < 6; i++ {
		memBlock.Container[i+1] = memBlock.Container[i]
	}
	memBlock.Container[0] = memBlock.Container[7]

}
