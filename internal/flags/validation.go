package flags

func validateFlags(f *Flags) {
	opCount := 0
	if f.Add != "" {
		opCount++
	}
	if f.List != "" {
		opCount++
	}
	if f.Del > 0 {
		opCount++
	}
	if f.Mut > 0 {
		opCount++
	}
	if f.ListLists {
		opCount++
	}
	if f.AddList != "" {
		opCount++
	}
	if f.DelList != 0 {
		opCount++
	}

	formatCount := 0
	if f.Red {
		formatCount++
	}
	if f.Green {
		formatCount++
	}
	if f.Blue {
		formatCount++
	}
	if f.Yellow {
		formatCount++
	}

	mutCount := 0
	if f.Cross {
		mutCount++
	}
	if f.Pin {
		mutCount++
	}

	if opCount > 1 {
		fmt.Println("Error: only one of -l, -a, -d, -m, -ls, -la, -ld can be used at a time.")
		os.Exit(1)
	}

	if formatCount > 1 {
		fmt.Println("Error: only one of -r, -g, -b, -y can be used at a time.")
		os.Exit(1)
	}

	if mutCount > 1 {
		fmt.Println("Error: only one of -p, -c can be used at a time.")
		os.Exit(1)
	}

	if (f.Red || f.Green || f.Blue || f.Yellow || f.Cross || f.Pin) && (f.Add == "" && f.Mut == 0) {
		fmt.Println("Error: mutating requires an operation like -a or -m.")
		os.Exit(1)
	}
}
