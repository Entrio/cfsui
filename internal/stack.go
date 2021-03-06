package internal

import "sync"

type (
	Stack struct {
		containers []stackContainer
		head       *stackContainer
		tail       *stackContainer
		mutex      sync.Mutex
	}

	stackContainer struct {
		element interface{}
		next    *stackContainer
		prev    *stackContainer
	}
)

func NewStack() Stack {
	return Stack{
		containers: make([]stackContainer, 0),
		head:       nil,
		tail:       nil,
		mutex:      sync.Mutex{},
	}
}

// Push adds a new container to the top of the stack and returns the pointer to the element
func (s *Stack) Push(data interface{}) interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	newContainer := stackContainer{
		element: data,
		prev:    nil,
		next:    nil,
	}

	if s.head == nil {
		s.head = &newContainer
		s.tail = &newContainer
		s.containers = append(s.containers, newContainer)
		return &newContainer.element
	}

	if s.tail.next == nil {
		s.tail.next = &newContainer
	}

	s.head.next = &newContainer
	newContainer.prev = s.head
	s.head = &newContainer

	s.containers = append(s.containers, newContainer)

	return &newContainer.element
}

// Pop removes the "topmost" container from the stack and returns the element
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.head == s.tail {
		ele := s.tail.element
		s.tail = nil
		s.head = nil
		s.containers = make([]stackContainer, 0)
		return ele
	}

	ele := s.head
	s.head = ele.prev
	s.head.next = nil

	return ele.element
}

// Peek returns the "topmost" container element of the stack without popping it
func (s *Stack) Peek() interface{} {
	return s.head.element
}

// Count returns the current number of containers on the stack
func (s *Stack) Count() int {
	return len(s.containers)
}
