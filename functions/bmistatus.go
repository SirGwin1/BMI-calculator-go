package functions

func BMIStatus(bmi float64)string{
	switch{
	case bmi < 18.5:
		return "UnderWeight"
	case bmi < 25:
		return "Normal"
	case bmi < 30:
		return "OverWeight"
	default:
		return "Obessed"
	}
}