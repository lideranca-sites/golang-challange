package calculator 

type Calculator interface {
	MakeCalculation(int, int) int
}

type Add struct {
}

func (a Add) MakeCalculation(x int, y int) int {
	return x + y
}

type Subtract struct {
}

func (s Subtract) MakeCalculation(x int, y int) int {
	return x - y
}

func GetCalculator(operator string) Calculator {
	switch operator {
	case "add":
		return Add{}
	case "subtract":
		return Subtract{}
	default:
		return Add{}
	}
}
