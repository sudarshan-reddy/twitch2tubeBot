package main

func main() {
	filename := setFileName("91a54c825a398c5c1840adaec6324136.mp4")
	title := setTitle("Arteezy Dodge", "Arteezy reaction")
	desc := setdesc("Test Desc", "Arteezy insane reaction")
	category := setCategory()
	keywords := setKeywords("Dota2, Arteezy, Reddit")
	privacy := setPrivacy()

	upload2Tube(filename, title, desc, category, keywords, privacy)
}
