package twofer

// ShareWith returns a string containing the name of a person to share with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
