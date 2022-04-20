package controller

import "testing"

func TestAddition(t *testing.T) {
	// test case 1 : 2, 4 => 6
	if Addition(2, 4) != 6 {
		t.Error("2 + 4 Seharusnya 6")
	}

	// test case 2 : -2, 4 => 2
	if Addition(-2, 4) != 2 {
		t.Error("-2 + 4 Seharusnya 2")
	}

	// test case 3 : 1, 2 => 3
	if Addition(1, 2) != 3 {
		t.Error("1 + 2 Seharusnya 3")
	}

	//test case 4 : -1, -2 => -3
	if Addition(-1, -2) != -3 {
		t.Error("-1 + -2 Seharusnya -3")
	}
}

func TestSubtraction(t *testing.T) {
	// test case 1 : 3, 6 => -3
	if Subtraction(3, 6) != -3 {
		t.Error("3 - 6 Seharusnya -3")
	}

	// test case 2 : -1, 5 => -6
	if Subtraction(-1, 5) != -6 {
		t.Error("-1 - 5 Seharusnya -6")
	}

	// test case 3 : 4, -3 => 7
	if Subtraction(4, -3) != 7 {
		t.Error("4 - (-3) Seharusnya 7")
	}

	//test case 4 : -1, -6 => 5
	if Subtraction(-1, -6) != 5 {
		t.Error("-1 - (-6) Seharusnya 5")
	}
}

func TestDivision(t *testing.T) {
	// test case 1 : 6, 3 => 2
	if Division(6, 3) != 2 {
		t.Error("6 / 3 Seharusnya 2")
	}

	// test case 2 : -5, 5 => -1
	if Division(-5, 5) != -1 {
		t.Error("-5 / 5 Seharusnya -1")
	}

	// test case 3 : -1, -2 => 0.5
	if Division(-1, -2) != 0.5 {
		t.Error("-1 / -2 Seharusnya 0.5")
	}

	//test case 4 : 3, -4 => -0.75
	if Division(3, -4) != -0.75 {
		t.Error("3 / -4 Seharusnya -0.75")
	}
}

func TestMultiplication(t *testing.T) {
	// test case 1 : 2, 4 => 8
	if Multiplication(2, 4) != 8 {
		t.Error("2 * 4 Seharusnya 8")
	}

	// test case 2 : -3, 5 => -15
	if Multiplication(-3, 5) != -15 {
		t.Error("-3 * 5 Seharusnya -15")
	}

	// test case 3 : -1, -1 => 1
	if Multiplication(-1, -1) != 1 {
		t.Error("-1 * -2 Seharusnya 0.5")
	}

	//test case 4 : 6, -7 => -42
	if Multiplication(6, -7) != -42 {
		t.Error("6 * -7 Seharusnya -42")
	}
}
