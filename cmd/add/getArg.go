package add

func GetArg(catFlag, noteFlag *bool) string {

	prompt := ""

	if *catFlag { // these flags won't collide;
		prompt += "> category" // there's a flag check in root

	} else if *noteFlag {
		prompt += "> note"

	} else {
		// which note are we adding this tag to?
		//
		prompt += "> tag" // TODO if tag, read the above
	}

	return "placeholder"
}
