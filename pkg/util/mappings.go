package util

type Coords struct {
	X int
	Z int
}

type BorderCoordinatesStruct struct {
	Min        Coords
	Max        Coords
	Rotation   float32
	MinRotated Coords
	MaxRotated Coords
}

// Weather maps the integer values to weather strings
var Weather = map[uint8]string{
	0: "clear",
	1: "light cloud",
	2: "overcast",
	3: "light rain",
	4: "heavy rain",
	5: "storm",
}

// SessionType maps the integer values to the actual names
var SessionType = map[uint8]string{
	0:  "unknown",
	1:  "P1",
	2:  "P2",
	3:  "P3",
	4:  "Short P",
	5:  "Q1",
	6:  "Q2",
	7:  "Q3",
	8:  "Short Q",
	9:  "OSQ",
	10: "R",
	11: "R2",
	12: "Time Trial",
}

// TrackName maps the integer values to an array of track name & country code
var TrackName = map[int8][2]string{
	-1: {"Unknown", ""},
	0:  {"Melbourne, Australia", "au"},
	1:  {"Le Castellet, France", "fr"},
	2:  {"Shanghai, China", "cn"},
	3:  {"Sakhir, Bahrain", "bh"},
	4:  {"Montmeló, Spain", "es"},
	5:  {"Monte Carlo, Monaco", "mc"},
	6:  {"Montréal, Canada", "ca"},
	7:  {"Silverstone, Great Britain", "gb"},
	8:  {"Hockenheim, Germany", "de"},
	9:  {"Mogyoród, Hungary", "hu"},
	10: {"Spa, Belgium", "be"},
	11: {"Monza, Italy", "it"},
	12: {"Singapore, Singapore", "sg"},
	13: {"Suzuka, Japan", "jp"},
	14: {"Abu Dhabi, UAE", "ae"},
	15: {"Austin, United States", "us"},
	16: {"São Paulo, Brazil", "br"},
	17: {"Spielberg, Austria", "at"},
	18: {"Sochi, Russia", "ru"},
	19: {"Mexico City, Mexico", "mx"},
	20: {"Baku, Azerbaijan", "az"},
	21: {"Sakhir Short, Bahrain", "bh"},
	22: {"Silverstone Short, Great Britain", "gb"},
	23: {"Austin Short, United States", "us"},
	24: {"Suzuka Short, Japan", "jp"},
	25: {"Hanoi, Vietnam", "vn"},
	26: {"Zandvoort, Netherlands", "nl"},
}

// Formula maps the integer values to the Formula classes
var Formula = map[uint8]string{
	0: "F1 Modern",
	1: "F1 Classic",
	2: "F2",
	3: "F1 Generic",
}

// SafetyCarStatus maps the integer values to the safety car phases
var SafetyCarStatus = map[uint8]string{
	0: "No Safety Car",
	1: "Full Safety Car",
	2: "Virtual Safety Car",
}

// FlagColor maps the integer values to the different flag colors
var FlagColor = map[int8]string{
	-1: "Unknown",
	0:  "None",
	1:  "Green",
	2:  "Blue",
	3:  "Yellow",
	4:  "Red",
}

// Teams maps the integer values to the team names
var Teams = map[uint8][2]string{
	0:   {"Mercedes-AMG F1 Team", "00d2be"},
	1:   {"Scuderia Ferrari", "dc0000"},
	2:   {"Red Bull Racing", "0600ef"},
	3:   {"Williams Racing", "0082fa"},
	4:   {"Racing Point F1 Team", "F596C8"},
	5:   {"Renault F1 Team", "FFF500"},
	6:   {"Scuderia AlphaTauri Honda", "ffffff"},
	7:   {"Haas F1 Team", "787878"},
	8:   {"McLaren F1 Team", "FF8700"},
	9:   {"Alfa Romeo Racing", "960000"},
	10:  {"1988 McLaren MP4/4", "7d7d7d"},
	11:  {"1991 McLaren MP4/6", "7d7d7d"},
	12:  {"1992 Williams FW14B", "7d7d7d"},
	13:  {"1995 Ferrari 412 T2", "7d7d7d"},
	14:  {"1996 Williams FW18", "7d7d7d"},
	15:  {"1998 McLaren MP4-13", "7d7d7d"},
	16:  {"2002 Ferrari F2002", "d9d9d9"},
	17:  {"2004 Ferrari F2004", "d9d9d9"},
	18:  {"2006 Renault R26", "d9d9d9"},
	19:  {"2007 Ferrari F2007", "d9d9d9"},
	20:  {"2008 McLaren MP4-23", "d9d9d9"},
	21:  {"2010 Red Bull Racing RB6", "d9d9d9"},
	22:  {"1976 Ferrari 312 T2", "7d7d7d"},
	23:  {"ART Grand Prix", "640005"}, // 2018 Teams
	24:  {"Campos Vexatec Racing", "f0dc05"},
	25:  {"Carlin", "0a505a"},
	26:  {"Charouz Racing System", "d7300a"},
	27:  {"DAMS", "00e1eb"},
	28:  {"Russian Time", "05329b"},
	29:  {"MP Motorsport", "f06e28"},
	30:  {"Pertamina Prema", "e1504b"},
	31:  {"1990 McLaren MP4/5B", "7d7d7d"},
	32:  {"Trident", "2d5ad7"},
	33:  {"BWT Arden", "faa0be"},
	34:  {"1976 McLaren M23-D", "7d7d7d"},
	35:  {"1972 Team Lotus 72D", "7d7d7d"},
	36:  {"1979 Ferrari 312 T4", "7d7d7d"},
	37:  {"1982 McLaren MP4/1B", "7d7d7d"},
	38:  {"2003 Williams FW25", "d9d9d9"},
	39:  {"2009 Brawn GP BGP-001", "d9d9d9"},
	40:  {"1978 Team Lotus 79", "7d7d7d"},
	41:  {"F1 Generic Car", "7d7d7d"},
	42:  {"ART Grand Prix", "640005"}, // 2019 Teams
	43:  {"Campos Racing", "fa6400"},
	44:  {"Carlin", "0a505a"},
	45:  {"Sauber Junior Charouz", "9e002a"},
	46:  {"DAMS", "00e1eb"},
	47:  {"UNI-Virtuosi Racing", "ffd600"},
	48:  {"MP Motorsport", "fe6557"},
	49:  {"Prema Racing", "fe2805"},
	50:  {"Trident", "2d5ad7"},
	51:  {"BWT Arden", "faa0be"},
	53:  {"1994 Benetton B194", "7d7d7d"},
	54:  {"1995 Benetton B195", "7d7d7d"},
	55:  {"2000 Ferrari F1-2000", "d9d9d9"},
	56:  {"1991 Jordan 191", "7d7d7d"},
	63:  {"1990 Ferrari F1-90", "7d7d7d"},
	64:  {"2010 McLaren MP4-25", "d9d9d9"},
	65:  {"2010 Ferrari F10", "d9d9d9"},
	70:  {"ART Grand Prix", "b0acae"}, // 2020 Teams
	71:  {"Campos Racing", "e66600"},
	72:  {"Carlin", "2840e5"},
	73:  {"Charouz Racing System", "0e0e6b"},
	74:  {"DAMS", "00d2f7"},
	75:  {"UNI-Virtuosi Racing", "f3ea33"},
	76:  {"MP Motorsport", "fd3613"},
	77:  {"Prema Racing", "f4000c"},
	78:  {"Trident", "0f0e86"},
	79:  {"BWT HWA Racelab", "f4b1d3"},
	80:  {"Hitech Grand Prix", "fdfcfd"},
	255: {"F1 2020 My Team", "7d7d7d"},
}

// Drivers maps the integer values to an array of first & last names
var Drivers = map[uint8][2]string{
	0:   {"Carlos", "Sainz"},
	1:   {"Daniil", "Kvyat"},
	2:   {"Daniel", "Ricciardo"},
	6:   {"Kimi", "Räikkönen"},
	7:   {"Lewis", "Hamilton"},
	9:   {"Max", "Verstappen"},
	10:  {"Nico", "Hülkenberg"},
	11:  {"Kevin", "Magnussen"},
	12:  {"Romain", "Grosjean"},
	13:  {"Sebastian", "Vettel"},
	14:  {"Sergio", "Perez"},
	15:  {"Valtteri", "Bottas"},
	17:  {"Esteban", "Ocon"},
	19:  {"Lance", "Stroll"},
	20:  {"Arron", "Barnes"},
	21:  {"Martin", "Giles"},
	22:  {"Alex", "Murray"},
	23:  {"Lucas", "Roth"},
	24:  {"Igor", "Correia"},
	25:  {"Sophie", "Levasseur"},
	26:  {"Jonas", "Schiffer"},
	27:  {"Alain", "Forest"},
	28:  {"Jay", "Letourneau"},
	29:  {"Esto", "Saari"},
	30:  {"Yasar", "Atiyeh"},
	31:  {"Callisto", "Calabresi"},
	32:  {"Naota", "Izum"},
	33:  {"Howard", "Clarke"},
	34:  {"Wilheim", "Kaufmann"},
	35:  {"Marie", "Laursen"},
	36:  {"Flavio", "Nieves"},
	37:  {"Peter", "Belousov"},
	38:  {"Klimek", "Michalski"},
	39:  {"Santiago", "Moreno"},
	40:  {"Benjamin", "Coppens"},
	41:  {"Noah", "Visser"},
	42:  {"Gert", "Waldmüller"},
	43:  {"Julian", "Quesada"},
	44:  {"Daniel", "Jones"},
	45:  {"Artem", "Markelov"},
	46:  {"Tadasuke", "Makino"},
	47:  {"Sean", "Gelael"},
	48:  {"Nyck", "De Vries"},
	49:  {"Jack", "Aitken"},
	50:  {"George", "Russell"},
	51:  {"Maximilian", "Günther"},
	52:  {"Nirei", "Fukuzumi"},
	53:  {"Luca", "Ghiotto"},
	54:  {"Lando", "Norris"},
	55:  {"Sérgio", "Sette Câmara"},
	56:  {"Louis", "Delétraz"},
	57:  {"Antonio", "Fuoco"},
	58:  {"Charles", "Leclerc"},
	59:  {"Pierre", "Gasly"},
	62:  {"Alexander", "Albon"},
	63:  {"Nicholas", "Latifi"},
	64:  {"Dorian", "Boccolacci"},
	65:  {"Niko", "Kari"},
	66:  {"Roberto", "Merhi"},
	67:  {"Arjun", "Maini"},
	68:  {"Alessio", "Lorandi"},
	69:  {"Ruben", "Meijer"},
	70:  {"Rashid", "Nair"},
	71:  {"Jack", "Tremblay"},
	74:  {"Antonio", "Giovinazzi"},
	75:  {"Robert", "Kubica"},
	78:  {"Nobuharu", "Matsushita"},
	79:  {"Nikita", "Mazepin"},
	80:  {"Guanyu", "Zhou"},
	81:  {"Mick", "Schumacher"},
	82:  {"Callum", "Ilott"},
	83:  {"Juan Manuel", "Correa"},
	84:  {"Jordan", "King"},
	85:  {"Mahaveer", "Raghunathan"},
	86:  {"Tatiana", "Calderón"},
	87:  {"Anthoine", "Hubert"},
	88:  {"Guiliano", "Alesi"},
	89:  {"Ralph", "Boschung"},
	91:  {"Dan", "Ticktum"},
	92:  {"Marcus", "Armstrong"},
	93:  {"Christian", "Lundgaard"},
	94:  {"Yuki", "Tsunoda"},
	95:  {"Jehan", "Daruvala"},
	96:  {"Guilherme", "Samaia"},
	97:  {"Pedro", "Piquet"},
	98:  {"Felipe", "Drugovich"},
	99:  {"Robert", "Shwartzman"},
	100: {"Roy", "Nissany"},
	101: {"Marino", "Satō"},
}

// Nationality maps the integer values to the driver nationalities
var Nationality = map[uint8][2]string{
	1:  {"American", "us"},
	2:  {"Argentinean", "ar"},
	3:  {"Australian", "au"},
	4:  {"Austrian", "at"},
	5:  {"Azerbaijani", "az"},
	6:  {"Bahraini", "bh"},
	7:  {"Belgian", "be"},
	8:  {"Bolivian", "bo"},
	9:  {"Brazilian", "br"},
	10: {"British", "gb"},
	11: {"Bulgarian", "bg"},
	12: {"Cameroonian", "cm"},
	13: {"Canadian", "ca"},
	14: {"Chilean", "cl"},
	15: {"Chinese", "cn"},
	16: {"Colombian", "co"},
	17: {"Costa Rican", "cr"},
	18: {"Croatian", "hr"},
	19: {"Cypriot", "cy"},
	20: {"Czech", "cz"},
	21: {"Danish", "dk"},
	22: {"Dutch", "nl"},
	23: {"Ecuadorian", "ec"},
	24: {"English", "gb-eng"},
	25: {"Emirian", "ae"},
	26: {"Estonian", "ee"},
	27: {"Finnish", "fi"},
	28: {"French", "fr"},
	29: {"German", "de"},
	30: {"Ghanaian", "gh"},
	31: {"Greek", "gr"},
	32: {"Guatemalan", "gt"},
	33: {"Honduran", "hn"},
	34: {"Hong Konger", "hk"},
	35: {"Hungarian", "hu"},
	36: {"Icelander", "is"},
	37: {"Indian", "in"},
	38: {"Indonesian", "id"},
	39: {"Irish", "ie"},
	40: {"Israeli", "il"},
	41: {"Italian", "it"},
	42: {"Jamaican", "jm"},
	43: {"Japanese", "jp"},
	44: {"Jordanian", "jo"},
	45: {"Kuwaiti", "kw"},
	46: {"Latvian", "lv"},
	47: {"Lebanese", "lb"},
	48: {"Lithuanian", "lt"},
	49: {"Luxembourger", "lu"},
	50: {"Malaysian", "my"},
	51: {"Maltese", "mt"},
	52: {"Mexican", "mx"},
	53: {"Monegasque", "mc"},
	54: {"New Zealander", "nz"},
	55: {"Nicaraguan", "ni"},
	56: {"North Korean", "kp"},
	57: {"Northern Irish", "gb-nir"},
	58: {"Norwegian", "no"},
	59: {"Omani", "om"},
	60: {"Pakistani", "pk"},
	61: {"Panamanian", "pa"},
	62: {"Paraguayan", "py"},
	63: {"Peruvian", "pe"},
	64: {"Polish", "pl"},
	65: {"Portuguese", "pt"},
	66: {"Qatari", "qa"},
	67: {"Romanian", "ro"},
	68: {"Russian", "ru"},
	69: {"Salvadoran", "sv"},
	70: {"Saudi", "sa"},
	71: {"Scottish", "gb-sct"},
	72: {"Serbian", "rs"},
	73: {"Singaporean", "sg"},
	74: {"Slovakian", "sk"},
	75: {"Slovenian", "si"},
	76: {"South Korean", "kr"},
	77: {"South African", "za"},
	78: {"Spanish", "es"},
	79: {"Swedish", "se"},
	80: {"Swiss", "ch"},
	81: {"Thai", "th"},
	82: {"Turkish", "tr"},
	83: {"Uruguayan", "uy"},
	84: {"Ukrainian", "ua"},
	85: {"Venezuelan", "ve"},
	86: {"Welsh", "gb-wls"},
	87: {"Barbadian", "bb"},
	88: {"Vietnamese", "vn"},
}

// SurfaceTypes maps the integer values to the different surface types
var SurfaceTypes = map[uint8]string{
	0:  "Tarmac",
	1:  "Rumble strip",
	2:  "Concrete",
	3:  "Rock",
	4:  "Gravel",
	5:  "Mud",
	6:  "Sand",
	7:  "Grass",
	8:  "Water",
	9:  "Cobblestone",
	10: "Metal",
	11: "Ridged",
}

// FuelMix maps the integer values to the available fuel mixes
var FuelMix = map[uint8]string{
	0: "lean",
	1: "standard",
	2: "rich",
	3: "max",
}

// ERSMode maps the integer values to the ERS modes
var ERSMode = map[uint8]string{
	0: "none",
	1: "medium",
	2: "overtake",
	3: "hotlap",
}

// ActualTyreCompound maps the integer values to technical tyre names
var ActualTyreCompound = map[uint8]string{
	7:  "inter",
	8:  "wet",
	9:  "dry",
	10: "wet",
	11: "supersoft",
	12: "soft",
	13: "medium",
	14: "hard",
	15: "wet",
	16: "C5",
	17: "C4",
	18: "C3",
	19: "C2",
	20: "C1",
}

// VisualTyreCompound maps the integer values to the standard tyre names
var VisualTyreCompound = map[uint8]string{
	7:  "inter",
	8:  "wet",
	9:  "dry",
	10: "wet_old",
	15: "wet",
	16: "soft",
	17: "medium",
	18: "hard",
	19: "supersoft",
	20: "soft",
	21: "medium",
	22: "hard",
	23: "supersoft",
	24: "soft",
	25: "medium",
	26: "hard",
}

// DriverStatus maps the integer values to the current driver status
var DriverStatus = map[uint8]string{
	0: "in garage",
	1: "flying lap",
	2: "in lap",
	3: "out lap",
	4: "on track",
}

// ResultStatus maps the integer values to the current result status
var ResultStatus = map[uint8]string{
	0: "invalid",
	1: "inactive",
	2: "active",
	3: "finished",
	4: "disqualified",
	5: "not classified",
	6: "retired",
}

// BorderCoordinates stores the outer coordinates for each track
// and the degrees of rotation to show it better
var BorderCoordinates = map[int8]BorderCoordinatesStruct{
	-1: {
		Min:        Coords{X: 0, Z: 0},
		Max:        Coords{X: 0, Z: 0},
		Rotation:   0,
		MinRotated: Coords{X: 0, Z: 0},
		MaxRotated: Coords{X: 0, Z: 0},
	},
	0: {
		Min:        Coords{X: -780, Z: -920},
		Max:        Coords{X: 780, Z: 920},
		Rotation:   -44,
		MinRotated: Coords{X: -940, Z: -610},
		MaxRotated: Coords{X: 1100, Z: 460},
	},
	1: {
		Min:        Coords{X: -1120, Z: -700},
		Max:        Coords{X: 1090, Z: 610},
		Rotation:   -4,
		MinRotated: Coords{X: -1160, Z: -630},
		MaxRotated: Coords{X: 1120, Z: 550},
	},
	2: {
		Min:        Coords{X: -660, Z: -580},
		Max:        Coords{X: 660, Z: 580},
		Rotation:   123.5,
		MinRotated: Coords{X: -830, Z: -680},
		MaxRotated: Coords{X: 590, Z: 280},
	},
	3: {
		Min:        Coords{X: -460, Z: -650},
		Max:        Coords{X: 450, Z: 640},
		Rotation:   -92.5,
		MinRotated: Coords{X: -630, Z: -430},
		MaxRotated: Coords{X: 650, Z: 440},
	},
	4: {
		Min:        Coords{X: -580, Z: -640},
		Max:        Coords{X: 490, Z: 610},
		Rotation:   57,
		MinRotated: Coords{X: -760, Z: -400}, // +/-100px on z-axis
		MaxRotated: Coords{X: 700, Z: 310},   // to achieve a 2:1 ratio
	},
	5: {
		Min:        Coords{X: -400, Z: -500},
		Max:        Coords{X: 440, Z: 550},
		Rotation:   53,
		MinRotated: Coords{X: -600, Z: -360}, // +/-50px on z-axis
		MaxRotated: Coords{X: 640, Z: 250},   // to achieve a 2:1 ratio
	},
	6: {
		Min:        Coords{X: -240, Z: -490},
		Max:        Coords{X: 480, Z: 1520},
		Rotation:   110,
		MinRotated: Coords{X: -1570, Z: -640}, // +/-160px on z-axis
		MaxRotated: Coords{X: 460, Z: 380},    // to achieve a 2:1 ratio
	},
	7: {
		Min:        Coords{X: -660, Z: -800},
		Max:        Coords{X: 450, Z: 1030},
		Rotation:   90,
		MinRotated: Coords{X: -1030, Z: -660},
		MaxRotated: Coords{X: 800, Z: 450},
	},
	8: {
		Min:        Coords{X: 0, Z: 0},
		Max:        Coords{X: 0, Z: 0},
		Rotation:   0,
		MinRotated: Coords{X: 0, Z: 0},
		MaxRotated: Coords{X: 0, Z: 0},
	},
	9: {
		Min:        Coords{X: -650, Z: -660},
		Max:        Coords{X: 510, Z: 640},
		Rotation:   51,
		MinRotated: Coords{X: -500, Z: -490},
		MaxRotated: Coords{X: 580, Z: 530},
	},
	10: {
		Min:        Coords{X: -770, Z: -1110},
		Max:        Coords{X: 600, Z: 1020},
		Rotation:   -95,
		MinRotated: Coords{X: -1060, Z: -640},
		MaxRotated: Coords{X: 1000, Z: 710},
	},
	11: {
		Min:        Coords{X: -680, Z: -1130},
		Max:        Coords{X: 680, Z: 1140},
		Rotation:   -95,
		MinRotated: Coords{X: -1170, Z: -610},
		MaxRotated: Coords{X: 1170, Z: 620},
	},
	12: {
		Min:        Coords{X: -760, Z: -470},
		Max:        Coords{X: 780, Z: 530},
		Rotation:   0,
		MinRotated: Coords{X: -760, Z: -470},
		MaxRotated: Coords{X: 780, Z: 530},
	},
	13: {
		Min:        Coords{X: -1040, Z: -550},
		Max:        Coords{X: 1050, Z: 540},
		Rotation:   0,
		MinRotated: Coords{X: -1040, Z: -550},
		MaxRotated: Coords{X: 1050, Z: 540},
	},
	14: {
		Min:        Coords{X: -780, Z: -360},
		Max:        Coords{X: 880, Z: 710},
		Rotation:   0,
		MinRotated: Coords{X: -780, Z: -360},
		MaxRotated: Coords{X: 880, Z: 710},
	},
	15: {
		Min:        Coords{X: -870, Z: -90},
		Max:        Coords{X: 1040, Z: 1080},
		Rotation:   0,
		MinRotated: Coords{X: -870, Z: -90},
		MaxRotated: Coords{X: 1040, Z: 1080},
	},
	16: {
		Min:        Coords{X: -610, Z: -390},
		Max:        Coords{X: 150, Z: 750},
		Rotation:   -90,
		MinRotated: Coords{X: -390, Z: -150},
		MaxRotated: Coords{X: 740, Z: 610},
	},
	17: {
		Min:        Coords{X: -580, Z: -530},
		Max:        Coords{X: 780, Z: 350},
		Rotation:   0,
		MinRotated: Coords{X: -580, Z: -530},
		MaxRotated: Coords{X: 780, Z: 350},
	},
	18: {
		Min:        Coords{X: -920, Z: -620},
		Max:        Coords{X: 1000, Z: 670},
		Rotation:   0,
		MinRotated: Coords{X: -920, Z: -620},
		MaxRotated: Coords{X: 1000, Z: 670},
	},
	19: {
		Min:        Coords{X: -1070, Z: -1070},
		Max:        Coords{X: 560, Z: 110},
		Rotation:   0,
		MinRotated: Coords{X: -1070, Z: -1070},
		MaxRotated: Coords{X: 560, Z: 110},
	},
	20: {
		Min:        Coords{X: -1230, Z: -940},
		Max:        Coords{X: 940, Z: 640},
		Rotation:   21.5,
		MinRotated: Coords{X: -1290, Z: -810}, // +/-200px on z-axis
		MaxRotated: Coords{X: 1100, Z: 480},   // to achieve a 2:1 ratio
	},
	21: {
		Min:        Coords{X: -460, Z: -650},
		Max:        Coords{X: 360, Z: 640},
		Rotation:   -92.5,
		MinRotated: Coords{X: -630, Z: -370},
		MaxRotated: Coords{X: 650, Z: 440},
	},
	22: {
		Min:        Coords{X: -660, Z: -80},
		Max:        Coords{X: 370, Z: 1030},
		Rotation:   63,
		MinRotated: Coords{X: -1000, Z: -350},
		MaxRotated: Coords{X: 130, Z: 430},
	},
	23: {
		Min:        Coords{X: -870, Z: 130},
		Max:        Coords{X: 230, Z: 1080},
		Rotation:   0,
		MinRotated: Coords{X: -870, Z: 130},
		MaxRotated: Coords{X: 230, Z: 1080},
	},
	24: {
		Min:        Coords{X: 260, Z: -320},
		Max:        Coords{X: 1040, Z: 540},
		Rotation:   -25,
		MinRotated: Coords{X: 140, Z: -470},
		MaxRotated: Coords{X: 1130, Z: 110},
	},
	25: {
		Min:        Coords{X: -690, Z: -790},
		Max:        Coords{X: 790, Z: 830},
		Rotation:   112,
		MinRotated: Coords{X: -1040, Z: -780},
		MaxRotated: Coords{X: 690, Z: 470},
	},
	26: {
		Min:        Coords{X: -510, Z: -460},
		Max:        Coords{X: 550, Z: 490},
		Rotation:   15.5,
		MinRotated: Coords{X: -580, Z: -490},
		MaxRotated: Coords{X: 530, Z: 390},
	},
}
