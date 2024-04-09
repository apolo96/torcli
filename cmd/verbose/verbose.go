package verbose

var FlagV bool

func Show(msg string) {
	if FlagV {
		println("Verbose: " + msg)
	}
}
