package controller

func Addition(operand_1, operand_2 int) int {
	result := operand_1 + operand_2
	return result
}

func Subtraction(operand_1, operand_2 int) int {
	result := operand_1 - operand_2
	return result
}

func Division(operand_1, operand_2 int) float32 {
	result := float32(operand_1) / float32(operand_2)
	return result
}

func Multiplication(operand_1, operand_2 int) int {
	result := operand_1 * operand_2
	return result
}
