package app

import "regexp"

// Pluralize generates the plural form of a given word.
func Pluralize(word string) string {
	irregulars := map[string]string{
		"child":  "children",
		"man":    "men",
		"woman":  "women",
		"tooth":  "teeth",
		"foot":   "feet",
		"mouse":  "mice",
		"person": "people",
		"ox":     "oxen",
	}

	// Check irregular words
	if plural, found := irregulars[word]; found {
		return plural
	}

	// Words ending in s, sh, ch, x, or z -> add "es"
	if matched, _ := regexp.MatchString(`(s|sh|ch|x|z)$`, word); matched {
		return word + "es"
	}

	// Words ending in "y" (but not "ay", "ey", "oy", "uy") -> replace "y" with "ies"
	if matched, _ := regexp.MatchString(`[^aeiou]y$`, word); matched {
		return word[:len(word)-1] + "ies"
	}

	// Words ending in "f" or "fe" -> replace with "ves"
	if matched, _ := regexp.MatchString(`(f|fe)$`, word); matched {
		return regexp.MustCompile(`(f|fe)$`).ReplaceAllString(word, "ves")
	}

	// Default rule: just add "s"
	return word + "s"
}
