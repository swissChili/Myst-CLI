package colors

func Set(target string, color string) string {
	Black := "\u001b[30m"
	Red := "\u001b[31m"
	Green := "\u001b[32m"
	Yellow := "\u001b[33m"
	Blue := "\u001b[34m"
	Magenta := "\u001b[35m"
	Cyan := "\u001b[36m"
	White := "\u001b[37m"
	Reset := "\u001b[0m"
	Bold := "\u001b[1m"

	switch target {
	case "black":
		return Black + color + Reset
	case "red":
		return Red + color + Reset
	
	case "green":
		return Green + color + Reset
	
	case "yellow":
		return Yellow + color + Reset
	
	case "blue":
		return Blue + color + Reset
	
	case "magenta":
		return Magenta + color + Reset
	
	case "cyan":
		return Cyan + color + Reset
	
	case "white":
		return White + color + Reset
	case "bold":
		return Bold + color + Reset
	}
	return color
}