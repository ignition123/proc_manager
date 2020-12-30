package lib

func GetTime(exchange string, segment string)(int, int, int, int){

	defer Handlepanic()

	// returned start hour, start min, end hour and end min

	// nse

	if exchange == "NSE" && segment == "CM"{

		return 9, 0, 15, 29

	} 

	if exchange == "NSE" && segment == "FO"{

		return 0, 15, 15, 29
	}

	if exchange == "NSE" && segment == "CD"{

		return 9, 0, 19, 29
	}

	// bse

	if exchange == "BSE" && segment == "CM"{

		return 9, 0, 15, 29
	}

	if exchange == "BSE" && segment == "FO"{

		return 9, 15, 15, 29
	}

	if exchange == "BSE" && segment == "CD"{

		return 9, 0, 19, 29
	}

	// mcx

	if exchange == "MCX" && segment == "FO"{

		return 9, 0, 23, 54
	}

	// ncdex

	if exchange == "NCDEX" && segment == "FO"{

		return 10, 0, 23, 29
	}

	return 0, 0, 0, 0
}