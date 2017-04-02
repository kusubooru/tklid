package tklid

type words []string

type byName words

func (w byName) Len() int           { return len(w) }
func (w byName) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w byName) Less(i, j int) bool { return w[i] < w[j] }

var (
	allWords = []words{
		part1,
		part2,
		part3,
		part4,
		part5,
		part6,
		part7,
		part8,
		part9,
		part10,
		part11,
	}
	part1 = words{
		"ferocious",
		"naive",
		"shy",
		"lost",
		"sleepy",
		"asleep",
		"nosy",
		"cocky",
		"bratty",
		"peppy",
		"snooty",
		"fearless",
		"laughing",
		"chuckling",
		"giggling",
		"giddy",
		"silly",
		"grumpy",
		"brave",
		"muscular",
		"strong",
		"buffed",
		"chubby",
		"plump",
		"tall",
		"short",
		"cute",
		"enchanted",
		"hypnotized",
		"bewitched",
		"spellbound",
		"captured",
		"enthralled",
		"friendly",
		"unfriendly",
		"enemy",
		"ally",
	}
	part2 = words{
		"blue",
		"brown",
		"hazel",
		"green",
		"black",
		"grey",
		"amber",
		"violet",
		"red",
		"golden",
		"silver",
		"fuchsia",
		"crystal",
		"turquoise",
		"aqua",
	}
	part3 = words{
		"eyed",
	}
	part4 = words{
		"brown",
		"dark",
		"black",
		"pink",
		"silver",
		"blue",
		"green",
		"red",
		"blond",
		"purple",
		"fuchsia",
		"grey",
		"auburn",
		"orange",
		"yellow",
		"white",
		"caramel",
		"amber",
		"copper",
	}
	part5 = words{
		"haired",
	}
	part6 = words{
		"human",
		"elf",
		"gnome",
		"hobbit",
		"monster",
		"cat",
		"neko",
		"wolf",
		"bunny",
		"sheep",
		"lion",
		"tiger",
		"panther",
		"dog",
		"duck",
		"phoenix",
		"swan",
		"raccoon",
		"tanuki",
		"fox",
		"kitsune",
		"oni",
		"alien",
		"android",
		"cyborg",
		"moogle",
		"pixie",
		"fairy",
		"angel",
		"devil",
		"smurf",
		"frog",
		"ninja",
		"knight",
		"hunter",
		"monk",
		"mage",
		"wizard",
		"lizard",
		"pigeon",
		"lama",
		"alpaca",
	}
	part7 = words{
		"boy",
		"girl",
		"lady",
	}
	part8 = words{
		"with",
		"having",
		"possessing",
		"boasting",
		"featuring",
	}
	part9 = words{
		"very",
		"pretty",
		"exquisitely",
		"hilariously",
		"ridiculously",
		"incredibly",
		"unbelievably",
		"deliciously",
		"dangerously",
		"surprisingly",
		"unexpectedly",
		"comically",
		"magically",
		"tragically",
		"unbearably",
		"infamously",
		"famously",
		"lovingly",
		"enjoyably",
		"hysterically",
		"humorously",
		"amusingly",
		"laughably",
		"pricelessly",
		"strangely",
		"oddly",
		"funnily",
		"prominently",
	}
	part10 = words{
		"ticklish",
		"sensitive",
		"delicate",
		"tender",
		"dainty",
		"touchy",
	}
	part11 = words{
		"neck",
		"nose",
		"ears",
		"chin",
		"chest",
		"armpits",
		"pits",
		"underarms",
		"arms",
		"sides",
		"ribs",
		"tummy",
		"belly",
		"abs",
		"face",
		"shoulders",
		"navel",
		"bellybutton",
		"hips",
		"thighs",
		"calves",
		"feet",
		"soles",
		"toes",
		"insteps",
		"heels",
		"tootsies",
		"legs",
		"back",
		"body",
		"skin",
		"knees",
		"waist",
		"waistline",
	}
)