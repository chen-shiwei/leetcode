package _542

type location struct {
	x int
	y int
}

func updateMatrix(matrix [][]int) [][]int {
	q := NewQueue()
	maxY := len(matrix)
	maxX := len(matrix[0])
	for i, nums := range matrix {
		for j, num := range nums {
			if num == 0 {
				q.Push(location{
					x: j,
					y: i,
				})
			} else {
				matrix[i][j] = -1
			}
		}

	}

	for !q.Empty() {
		// 下左上右 x 横 y 竖
		dx := []int{0, -1, 0, 1}
		dy := []int{-1, 0, 1, 0}
		l := q.Pop().(location)

		for i := 0; i < 4; i++ {
			newX := l.x + dx[i]
			newY := l.y + dy[i]
			if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY && matrix[newY][newX] == -1 {
				matrix[newY][newX] = matrix[l.y][l.x] + 1
				q.Push(location{
					x: newX,
					y: newY,
				})
			}

		}
	}

	return matrix
}

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Pop() interface{} {

	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	s.data = s.data[1:]
	return popVal
}

func (s *Queue) Peek() interface{} {
	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	return popVal
}

func (s *Queue) Push(x interface{}) {
	s.data = append(s.data, x)
	return
}

func (s *Queue) Empty() bool {
	return len(s.data) == 0
}

func (s Queue) Len() int {
	return len(s.data)
}
