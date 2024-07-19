package pkg

// nb (sienna): these progressions came from Claude 3.5 Sonnet and I did not hand check them. YMMV

var DiscoProgressions = ProgressionType{
	Name: DiscoKind.String(),
	Progressions: []Progression{
		{
			Name: "I-V-vi-IV (Basic Disco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "i-iv-V-i (Minor Disco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Jazz-influenced Disco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "I-bVII-IV-V (Funk-inspired Disco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "Im7-bVII7-bVI7-V7 (Chromatic Disco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-IV-V-IV (Chic - Le Freak)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V (Donna Summer - I Feel Love)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV-I-ii-V (ABBA - Dancing Queen)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V (Sister Sledge - We Are Family)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Im7-IV7-bVII7 (Bee Gees - Stayin' Alive)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-bVII-IV (KC and the Sunshine Band - That's the Way (I Like It))",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-IV-V-i (Gloria Gaynor - I Will Survive)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Earth, Wind & Fire - September)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
	},
}

var SoulProgressions = ProgressionType{
	Name: SoulKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V (Basic Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "ii-V-I (Jazz-influenced Soul)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "I-vi-IV-V (Doo-wop Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-iv-V (Minor Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-iii-IV-V (Gospel-influenced Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-V-vi-IV (Modern Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
			},
		},
		{
			Name: "I-bVII-IV-V (Otis Redding - (Sittin' On) The Dock of the Bay)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (Sam Cooke - A Change Is Gonna Come)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "i-bIII-IV-iv (Aretha Franklin - I Never Loved a Man (The Way I Love You))",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-IV-V-IV (Al Green - Let's Stay Together)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
			},
		},
		{
			Name: "i-iv-VII-VI (Marvin Gaye - What's Going On)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Maj9Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Stevie Wonder - Isn't She Lovely)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-IV-V-iv (Bill Withers - Ain't No Sunshine)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (D'Angelo - Untitled (How Does It Feel))",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
	},
}

var FunkProgressions = ProgressionType{
	Name: FunkKind.String(),
	Progressions: []Progression{
		{
			Name: "I7-IV7 (Basic Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I9-IV9-V9 (Extended Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "Im7-IV7-bVII7 (Minor Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-IV7-V7-IV7 (Funk Turnaround)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I13-IV13-ii7-V7 (Jazz Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom13Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom13Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Im7-bIII7-bVI7-V7 (Chromatic Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-bVII7-IV7 (James Brown - I Got You (I Feel Good))",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Im7-IV7 (Stevie Wonder - Superstition)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-V7-IV7-I7 (Parliament - Give Up the Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Im7-bIII7-IV7-bII7 (Herbie Hancock - Chameleon)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I9-IV9-V9-IV9 (Earth, Wind & Fire - September)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "Im7-Vm7-IVm7 (Chic - Good Times)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "I7-IV7-ii7-V7 (Tower of Power - What Is Hip?)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Im9-bVII9-bIII9-bVI9 (Jamiroquai-style Funk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
	},
}

var PunkProgressions = ProgressionType{
	Name: PunkKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V (Basic Punk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (Pop Punk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV (Punk Rock)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-bVII-IV (Alternative Punk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-IV (Punk Power Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Melodic Punk)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-IV-IV (Ramones - Blitzkrieg Bop)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V (Green Day - Basket Case)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV (Blink-182 - All The Small Things)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-vi-IV (The Clash - Should I Stay or Should I Go)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV-I (Sex Pistols - Anarchy in the U.K.)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV-I-V (NOFX - Linoleum)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-IV (Bad Religion - 21st Century Digital Boy)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-bVI-V (Minor Punk Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var PopProgressions = ProgressionType{
	Name: PopKind.String(),
	Progressions: []Progression{
		// Theoretical Progressions
		{
			Name: "I-V-vi-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "ii-V-I-vi",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},

		// Real-world Pop Progressions
		{
			Name: "I-V-vi-IV (Four Chords - Axis of Awesome)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Don't Stop Believin' - Journey)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-III-IV (Every Breath You Take - The Police)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V (Stand By Me - Ben E. King)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-V (Despacito - Luis Fonsi & Daddy Yankee)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V (Achy Breaky Heart - Billy Ray Cyrus)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-V-IV-V (Zombie - The Cranberries)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-iii-vi-IV (Someone Like You - Adele)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV-I-V-iii-IV (Wherever You Will Go - The Calling)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Wonderwall - Oasis)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var RockProgressions = ProgressionType{
	Name: RockKind.String(),
	Progressions: []Progression{
		// Theoretical Rock Progressions
		{
			Name: "I-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-bVI-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-IV-IV",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-II-IV-I",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},

		// Real-world Rock Progressions
		{
			Name: "I-bVII-IV (Sweet Home Alabama - Lynyrd Skynyrd)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (When I Come Around - Green Day)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVI-III-bVII (Smoke on the Water - Deep Purple)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-IV (Louie Louie - The Kingsmen)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bIII-IV-IV (The Hardest Button to Button - The White Stripes)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-IV-V (Twist and Shout - The Beatles)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-vi-IV (Under the Bridge - Red Hot Chili Peppers)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "V-IV-I (Pinball Wizard - The Who)",
			Chords: []Chord{
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-bVI-bVII (Ain't Talkin' 'Bout Love - Van Halen)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Stairway to Heaven - Led Zeppelin)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Zombie - The Cranberries)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-iii-IV-I-V (Where is My Mind - Pixies)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var ProgressiveRockProgressions = ProgressionType{
	Name: ProgressiveRockKind.String(),
	Progressions: []Progression{
		// Theoretical Progressive Rock Progressions
		{
			Name: "I-V-vi-iii-IV-I-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-bVI-bVII-i-bIII-bII-bIII",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-vi-ii-iii-IV-V",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
			},
		},
		{
			Name: "I-bIII-IV-iv-bVI-V-I",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
			},
		},

		// Real-world Progressive Rock Progressions
		{
			Name: "i-bVII-bVI-bVII (Starless - King Crimson)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-ii-IV (Roundabout - Yes)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
			},
		},
		{
			Name: "i-bIII-bVII-IV-bVI (21st Century Schizoid Man - King Crimson)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV-I (Carry on Wayward Son - Kansas)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVI-bIII-bVII (Kashmir - Led Zeppelin)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-IV-bVII-IV (Tom Sawyer - Rush)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-V-vi (Firth of Fifth - Genesis)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-bVII-bVI-bVII-i-bIII-bII-bIII (Shine On You Crazy Diamond - Pink Floyd)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-ii-iii-IV-V-vi-vii° (Apocalypse in 9/8 - Genesis)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{DiminishedExtension}},
			},
		},
		{
			Name: "I-bIII-IV-bVI (Money - Pink Floyd)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var BluesProgressions = ProgressionType{
	Name: BluesKind.String(),
	Progressions: []Progression{
		{
			Name: "12-Bar Blues (Basic)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "12-Bar Blues (Quick Change)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "12-Bar Blues (Minor)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "8-Bar Blues",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "16-Bar Blues",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Stormy Monday Blues",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "St. Louis Blues",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Jazz Blues",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Boogie Woogie",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
	},
}

var JazzProgressions = ProgressionType{
	Name: JazzKind.String(),
	Progressions: []Progression{
		{
			Name: "ii-V-I (Basic)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "ii-V-I (Extended)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension, Add13Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension, Add9Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Basic)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Rhythm Changes (A Section)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "Autumn Leaves (First 8 bars)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Min7b5Extension}},
				{Base: 9, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "All The Things You Are (First 8 bars)",
			Chords: []Chord{
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 11, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 9, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "Giant Steps (First 8 bars)",
			Chords: []Chord{
				{Base: 9, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 11, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 11, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Blue Bossa",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7b5Extension}},
				{Base: 10, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Coltrane Changes",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 9, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 10, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "So What (Modal Jazz)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension, Add11Extension}},
			},
		},
		{
			Name: "Stella By Starlight (First 8 bars)",
			Chords: []Chord{
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 7, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 8, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
	},
}

var ClassicalProgressions = ProgressionType{
	Name: ClassicalKind.String(),
	Progressions: []Progression{
		{
			Name: "Perfect Authentic Cadence",
			Chords: []Chord{
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Plagal Cadence",
			Chords: []Chord{
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Deceptive Cadence",
			Chords: []Chord{
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "Half Cadence",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Pachelbel's Canon",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Circle of Fifths (Partial)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 8, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 11, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Romanesca",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "La Folia",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "Mozart's Sonata Facile (Opening)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "Beethoven's Für Elise (Opening)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "Bach's Prelude in C Major (Opening)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var FolkProgressions = ProgressionType{
	Name: FolkKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V (Basic Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-IV-I (Folk Rock)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V (50s Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Modern Folk)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (Pop Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-IV-I (Irish Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VII-VI-V (Andalusian Cadence)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-ii-iii-IV (Ascending Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-I-IV (Country Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-III-VII-iv (Minor Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-IV-I-V (Bluegrass Turn)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV-I (Modal Folk)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var CountryProgressions = ProgressionType{
	Name: CountryKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V (Basic Country)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-IV-I (Country Turnaround)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-I-V (Nashville Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Modern Country)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (Pop Country)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-IV (Country Walk-up)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-III-IV-iv (Country Blues)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-vi-ii-V (Country Jazz)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Extended Country)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-IV (Country Rock)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-ii-V-I (Country Ballad)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V-V7 (Country Shuffle)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
	},
}

var ReggaeProgressions = ProgressionType{
	Name: ReggaeKind.String(),
	Progressions: []Progression{
		{
			Name: "I-V (Basic Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-IV-V (Roots Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-VI (Minor Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V (Ska Influence)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-iv-V (Minor Roots)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV (Pop Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-bVI-V (Reggae Rock)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bIII-bVII-i (Dub Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-IV-I-V (Reggae Fusion)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVI-bVII (Roots Rocksteady)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-iii-IV-I (Lover's Rock)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-IV-I-IV (Extended Reggae)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var EDMProgressions = ProgressionType{
	Name: EDMKind.String(),
	Progressions: []Progression{
		{
			Name: "I-V-vi-IV (Progressive House)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-IV-I-V (Trance)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV (Melodic Dubstep)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VI-III-VII (Deep House)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-iii-vi-IV (Future Bass)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "i-bVI-bVII-V (Techno)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-V-vi-iii (Tropical House)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-bVII-IV-I (Big Room House)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "vi-V-IV-V (Electro House)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-iii-IV-vi (Chillstep)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Add9Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "i-bIII-bVII-IV (Drum and Bass)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-vi-IV-V-I-vi-IV-V (Extended EDM)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{Sus4Extension}},
			},
		},
	},
}

var RnBProgressions = ProgressionType{
	Name: RnBKind.String(),
	Progressions: []Progression{
		{
			Name: "ii-V-I-vi (Classic R&B)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "I-vi-IV-V (Doo-Wop Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-iv-VII-III (Neo-Soul)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "I-V-vi-IV (Contemporary R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "ii-V-I-IV (Jazz-Influenced R&B)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "vi-IV-I-V (90s R&B)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-iii-IV-V (Soul Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V (Minor R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Smooth R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Extended R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-IV-VII-III (Alternative R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "I-bIII-IV-iv (Blues-Inspired R&B)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
	},
}

var GospelProgressions = ProgressionType{
	Name: GospelKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V-IV (Basic Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Gospel Turnaround)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
			},
		},
		{
			Name: "I-III-IV-V (Traditional Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "vi-IV-V-iii-ii-I (Contemporary Gospel)",
			Chords: []Chord{
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
			},
		},
		{
			Name: "I-V-vi-IV (Modern Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min11Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
			},
		},
		{
			Name: "I-bIII-IV-iv (Gospel Blues)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "I-VI-ii-V (Jazz Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom13Extension}},
			},
		},
		{
			Name: "I-V-vi-iii-IV-I-IV-V (Extended Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-iii-IV-V (Soulful Gospel)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min11Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom13Extension}},
			},
		},
		{
			Name: "ii-V-I-vi (Gospel Jazz)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min9Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom13Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj9Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min11Extension}},
			},
		},
		{
			Name: "I-IV-V-III-vi-ii-V-I (Circle Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "I-V-vi-IV-I-V-iii-vi (Contemporary Worship)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
	},
}

var FlamencoProgressions = ProgressionType{
	Name: FlamencoKind.String(),
	Progressions: []Progression{
		{
			Name: "i-VII-VI-V (Andalusian Cadence - Soleá, Bulerías)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VII-iv-V (Variant Andalusian - Seguiriya)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-iv-V (Tango Flamenco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VII-i-VII (Fandangos de Huelva)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-v-VI-i (Phrygian Cadence - Soleá por Bulerías)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-V-I-V (Sevillanas)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-v-i-v (Tarantas)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-III-i-III (Farruca)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VI-V-i (Bulerías al Golpe)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i7-iv7-V7-i7 (Jazz Flamenco)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
			},
		},
		{
			Name: "i-iv-V7-i (Rumba Flamenca)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-VII-VI-V-i (Extended Andalusian - Alegrías)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-III-VI-i (Guajiras)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-iv-VII-VI (Peteneras)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VI-VII-i (Tanguillos)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-VII-VI-V-iv-III-II-I (Extended Phrygian - Granaína)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bII-bIII-i (Taranta Variant)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 2, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-VII-VI-v (Malagueña)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
	},
}

var MetalProgressions = ProgressionType{
	Name: MetalKind.String(),
	Progressions: []Progression{
		{
			Name: "I-VI-bVII-VI (Power Chord Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-bVII-bVI-V (Chromatic Descent)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 11, Extensions: []ExtensionKind{NoExtension}},
				{Base: 10, Extensions: []ExtensionKind{NoExtension}},
				{Base: 9, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bII-VI-iv (Phrygian Metal)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-V-VI-i (Harmonic Minor Metal)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-iii-IV-vi (Dorian Metal)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 2, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-bVI-I-bVI (Tritone-based Progression)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "I-VI-VIII-VI (Black Sabbath - Iron Man)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 7, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-IV-V-i (Metallica - Enter Sandman)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "i-bII-bVII-i (Slayer - Raining Blood)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 1, Extensions: []ExtensionKind{NoExtension}},
				{Base: 11, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I-VI-bVII-I (Judas Priest - Breaking the Law)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 0, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-vi-bII-I (Megadeth - Symphony of Destruction)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 8, Extensions: []ExtensionKind{NoExtension}},
				{Base: 7, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bVII-VI-IV (Dream Theater - Pull Me Under)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-VI-IV-V (System of a Down - Chop Suey!)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
				{Base: 3, Extensions: []ExtensionKind{NoExtension}},
				{Base: 4, Extensions: []ExtensionKind{NoExtension}},
			},
		},
		{
			Name: "i-bII-bVII-VI (Opeth - The Drapery Falls)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 8, Extensions: []ExtensionKind{NoExtension}},
				{Base: 6, Extensions: []ExtensionKind{NoExtension}},
				{Base: 5, Extensions: []ExtensionKind{NoExtension}},
			},
		},
	},
}

var BossaNovaProgressions = ProgressionType{
	Name: BossaNovaKind.String(),
	Progressions: []Progression{
		{
			Name: "ii-V-I (Basic Bossa Nova)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "i-IV7-iii7-bVI7-ii7-V7 (Minor Bossa Nova)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-vi-ii-V (Bossa Nova Turnaround)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-vi7-iimin7-V7 (Extended Bossa Nova)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-#Idim7-ii7-V7 (Chromatic Bossa Nova)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dim7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-bVIImaj7-bIIImaj7-bVImaj7 (Jobim - Corcovado)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Maj7Extension}},
			},
		},
		{
			Name: "Imaj7-#Idim7-ii7-V7 (Jobim - The Girl from Ipanema)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dim7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i7-IV7-iii7-VI7 (João Gilberto - Desafinado)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V7 (Luiz Bonfá - Manhã de Carnaval)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-vi7-iimin7-V7-iii7-VI7-ii7-V7 (Jobim - Wave)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-V7/III-IIImaj7-V7/VI-VImaj7-iv7-V7 (Jobim - How Insensitive)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 7, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-vi7-ii7-V7-iii7-VI7-ii7-V7 (João Donato - A Rã)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
	},
}

var SambaProgressions = ProgressionType{
	Name: SambaKind.String(),
	Progressions: []Progression{
		{
			Name: "I-ii-V7 (Basic Samba)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i7-iv7-V7 (Minor Samba)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-IV7-V7 (Partido Alto Samba)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-VI7-ii7-V7 (Samba Turnaround)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-III7-VI7-II7-V7 (Extended Samba)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-vi7-ii7-V7 (Jobim - Samba de Uma Nota Só)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i7-iv7-VII7-III7-VI7-ii7b5-V7 (Jorge Ben - Mas Que Nada)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 6, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7b5Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I6-vi7-ii7-V7 (Ary Barroso - Aquarela do Brasil)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Add13Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V7 (Cartola - As Rosas Não Falam)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "Imaj7-vi7-iimin7-V7-iii7-VI7-ii7-V7 (João Bosco - Linha de Passe)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 2, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-IV7-ii7-V7 (Zeca Pagodinho - Deixa a Vida Me Levar)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i7-IV7-V7 (Paulinho da Viola - Timoneiro)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
	},
}

var SalsaProgressions = ProgressionType{
	Name: SalsaKind.String(),
	Progressions: []Progression{
		{
			Name: "I-IV-V-IV (Basic Salsa)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-iv-V7 (Minor Salsa)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-V-vi-IV (Pop-influenced Salsa)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
			},
		},
		{
			Name: "ii7-V7-I7 (Jazz-influenced Salsa)",
			Chords: []Chord{
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I7-IV7-V7 (Son Montuno)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V (Andalusian Cadence in Salsa)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
			},
		},
		{
			Name: "I-vi-ii-V (Salsa Romantica)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Maj7Extension}},
				{Base: 5, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 1, Extensions: []ExtensionKind{Min7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-IV-V-i (Héctor Lavoe - El Cantante)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
		{
			Name: "I7-IV7-V7 (Celia Cruz - La Vida Es Un Carnaval)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 3, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "i-bVII-bVI-V7 (Marc Anthony - Valió la Pena)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 6, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
			},
		},
		{
			Name: "I-vi-IV-V (Rubén Blades - Pedro Navaja)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 5, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MajorExtension}},
				{Base: 4, Extensions: []ExtensionKind{MajorExtension}},
			},
		},
		{
			Name: "i-iv-V7-i (Willie Colón - Idilio)",
			Chords: []Chord{
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 3, Extensions: []ExtensionKind{MinorExtension}},
				{Base: 4, Extensions: []ExtensionKind{Dom7Extension}},
				{Base: 0, Extensions: []ExtensionKind{MinorExtension}},
			},
		},
	},
}

/*
Still want:

Salsa
Tango
Celtic
Bluegrass
Hip Hop
Electronic (different from EDM, focusing on genres like Ambient, IDM, Synthwave)
New Age
Indie Rock
Alternative Rock
Grunge
Ska
Swing
Bebop (as a distinct jazz subgenre)
Fusion (Jazz Fusion)
Baroque (as a distinct classical subgenre)
Romantic (as a distinct classical subgenre)
Impressionist (as a distinct classical subgenre)
Minimalist
Film Score
Video Game Music
J-Pop (Japanese Pop)
K-Pop (Korean Pop)
Bollywood
*/

var AllProgressionTypes = []ProgressionType{
	BluesProgressions,
	BossaNovaProgressions,
	ClassicalProgressions,
	CountryProgressions,
	DiscoProgressions,
	EDMProgressions,
	FlamencoProgressions,
	FolkProgressions,
	FunkProgressions,
	GospelProgressions,
	JazzProgressions,
	MetalProgressions,
	PopProgressions,
	ProgressiveRockProgressions,
	PunkProgressions,
	ReggaeProgressions,
	RnBProgressions,
	RockProgressions,
	SalsaProgressions,
	SambaProgressions,
	SoulProgressions,
}
