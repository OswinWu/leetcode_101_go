package countStudents

func countStudents(students []int, sandwiches []int) int {
	stuLen := len(students)
	sandwLen := len(sandwiches)
	lastStuLen := stuLen
	for i := 0; i < sandwLen; i++ {
		leftStu := []int{}
		for j := 0; j < stuLen; j++ {
			if sandwiches[i] == students[j] {
				stuLen = stuLen - 1
				students = append(students[j+1:], leftStu...)
				break
			} else {
				leftStu = append(leftStu, students[j])
			}
		}
		if stuLen == 0 {
			break
		}
		if lastStuLen == stuLen {
			break
		} else {
			lastStuLen = stuLen
		}
	}
	return stuLen
}
