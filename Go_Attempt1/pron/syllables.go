package pron

var extendedVowels = []phoneme{
	aa, ae, ah, ao, axl, axm, axn, axr, ax, aw, ay,
	eh, ehr, ehl, er, ey,
	ih, ihl, ing, iy,
	oh, ow, oy,
	uh, uw, uwl, uwm, uwn, yuw,
}

// Looking for triple consonant endings of phonetic spellings. These may indicate
// a missing schwa in the spelling. Where a schwa is missing we need to add a
// syllable to the count
// Typically a triple with L, M or N as a middle phoneme has a missing schwa with
// the following exceptions
var notMissingSchwa = [][3]phoneme{
	{
		l, m, d, // As in fiLMED,...
	},
	{
		l, m, m, // As in fiLMMakers,...
	},
	{
		l, m, z, // As in eLMS,...
	},
	{
		l, n, z, // As in kiLNS,...
	},
}

// And for triples not containing L, M or N as a middle phoneme the following
// triples have a missing schwa
var missingSchwa = [][3]phoneme{
	{
		ng, th, n, // As in streNGTHEN,...
	},
	{
		s, t, v, // As in muST'VE
	},
	{
		n, jh, l, // As in aNGEL,...
	},
}

func isExtendedVowel(ph phoneme) bool {
	for _, v := range extendedVowels {
		if v == ph {
			return true
		}
	}
	return false
}

// This is very heuristic!
// It's based on the dictionary defined at the time of writing (7 Jan 2022) and may
// change
//
// This basically checks for triples of consonants contained in a pronunciation
// with a middle consonant of l, m or n. Typically this indicates a missing schwa
// except for triples included in the array notMissingSchwa.
// There are exceptions with a middle consonant other than l, m or n which nevertheless
// indicate a missing schwa. These are included in the array missingSchwa.
//
// We also check for a missing schwa at the end of a word. Here we check for a final
// consonant of l, m, or n preceded by another consonant. A final consonant of
// m or n preceded by l (as in fiLM or kiLN) does not indicate a missing schwa
func isMissingSchwa(phonemes []phoneme) bool {
	// We need at least 3 phonemes to be missing a schwa
	if len(phonemes) < 3 {
		return false
	}
	for i := 0; i < len(phonemes)-2; i++ {
		foundVowel := false
		three := [3]phoneme{phonemes[i], phonemes[i+1], phonemes[i+2]}
		for _, ph := range three {
			if isExtendedVowel(ph) {
				foundVowel = true
				break
			}
		}
		if foundVowel {
			continue
		}
		if three[1] == l || three[1] == m || three[1] == n {
			notMissing := false
			for _, triple := range notMissingSchwa {
				if triple == three {
					notMissing = true
					continue
				}
			}
			if !notMissing {
				return true
			}
		} else {
			for _, triple := range missingSchwa {
				if triple == three {
					return true
				}
			}
		}
	}
	// Now check for a missing schwa in the last two phonemes
	lastPh := phonemes[len(phonemes)-1]
	nextToLastPh := phonemes[len(phonemes)-2]
	if isVowel(lastPh) || isVowel(nextToLastPh) {
		return false
	} else {
		if lastPh == l || lastPh == m || lastPh == n {
			// As in fiLM, kiLN,...
			return nextToLastPh != l
		} else {
			return false
		}
	}
}

// Counting syllables boils down to counting vowels where a vowel is defined as
// one of the phonemes in the vowels array above.
// However, some pronunciations contain consonant clusters that are pretty hard
// to say without introducing a schwa (which in turn adds a syllable to the word)
// so we add one to the count when a missing schwa is detected.
// NOTE: this implementation assumes there will only ever by up to one missing schwa
// in word (which I think is true with the current dictionary).
func syllables(phonemes []phoneme) int {
	isExtendedVowel := func(ph phoneme) bool {
		for _, v := range extendedVowels {
			if v == ph {
				return true
			}
		}
		return false
	}
	count := 0
	for _, phoneme := range phonemes {
		if isExtendedVowel(phoneme) {
			count++
		}
	}
	if isMissingSchwa(phonemes) {
		count++
	}
	return count
}
