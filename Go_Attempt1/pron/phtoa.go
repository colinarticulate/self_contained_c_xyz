package pron

import (
	"fmt"
	"strings"
)

type phonToAlphas struct {
	phons  []phoneme
	alphas string
}

func (p phonToAlphas) equal(q phonToAlphas) bool {
	if p.alphas == q.alphas {
		if len(p.phons) != len(q.phons) {
			return false
		}
		for i := 0; i < len(p.phons); i++ {
			if p.phons[i] != q.phons[i] {
				return false
			}
		}
		return true
	}
	return false
}

type phToAError struct {
	phon phoneme
	word string
}

func (p phToAError) Error() string {
	return fmt.Sprintf("Phoneme translation failure, translating phoneme %s in word, %s", string(p.phon), p.word)
}

func stringAt(s string, start, length int, ts ...string) bool {
	if start < 0 || start+length > len(s)+1 {
		return false
	}
	subS := s[start:]
	for _, t := range ts {
		if len(t) != length {
			continue
		}
		if strings.HasPrefix(subS, t) {
			return true
		}
	}
	return false
}

/*
func getStringAt(s string, start, length int, ts ...string) (string, bool) {
  if start < 0 || start + length > len(s) + 1 {
    return "", false
  }
  subS := s[start:]
  for _, t := range ts {
    if len(t) != length {
      continue
    }
    if strings.HasPrefix(subS, t) {
      return t, true
    }
  }
  return "", false
}
*/

// Returns the first string in ts that appears as a prefix of s[start:] - so
// the order in which strings are passed to this function is important
func getStringAt(s string, start, length int, ts ...string) (string, bool) {
	if start < 0 {
		return "", false
	}
	subS := s[start:]
	for _, t := range ts {
		if strings.HasPrefix(subS, t) {
			return t, true
		}
	}
	return "", false
}

func nextPhIsConsonant(ps []phoneme, start int) bool {
	if start < 0 || start+1 >= len(ps) {
		return false
	}
	return isConsonant(ps[start+1])
}

// Returns the first array of phonemes in qs that appears as a prefix of
// ps[start:] - so the order in which the arrays in qs are passed to this
// function is important
func phsAt(ps []phoneme, start int, qs ...[]phoneme) ([]phoneme, bool) {
	hasPrefix := func(p, q []phoneme) bool {
		if len(q) > len(p) {
			// q is too long to be a prefix
			return false
		}
		for i, q1 := range q {
			if p[i] != q1 {
				// Check to see if q is a 'prefix' of p, phoneme by phoneme
				return false
			}
		}
		return true
	}
	if start < 0 || start >= len(ps) {
		return []phoneme{}, false
	}
	subPh := ps[start:]
	for _, q := range qs {
		if hasPrefix(subPh, q) {
			return q, true
		}
	}
	return []phoneme{}, false
}

func isConsonant(ph phoneme) bool {
	consonants := []phoneme{
		b, bl, ch, d, dh, dz, f, fl, g, gr, hh, jh, k, kl, kr, ks, kw, l, m, n, ng, p, pl, pr, r, s, sh, st, t, th, thr, tr, ts, v, w, z, zh,
	}
	for _, c := range consonants {
		if ph == c {
			return true
		}
	}
	return false
}

func isVowel(ph phoneme) bool {
	vowels := []phoneme{
		aa, ae, ah, ao, aw, ay, eh, er, ey, ih, iy, ow, oy, uw, uh, y,
	}
	for _, v := range vowels {
		if ph == v {
			return true
		}
	}
	return false
}

func trailingSilentAlphas(abc string, currAbc int, phons []phoneme, currPh int, currMap phonToAlphas) string {
	unmappedAlphas := ""
	if currAbc+len(currMap.alphas) < len(abc) {
		unmappedAlphas = abc[currAbc+len(currMap.alphas):]
	}
	if (currPh+len(currMap.phons) == len(phons)) &&
		(len(unmappedAlphas) != 0) {
		return unmappedAlphas
	}
	return ""
}

func isSilentE(abc string, currAbc int, phons []phoneme, currPh int, currMap phonToAlphas) bool {
	// Words with 'hidden' silent e
	hiddenSilentEs := map[string]int{
		"changeover":  5,
		"changeovers": 5,
		"closeup":     4,
		"closeups":    4,
		"fadeout":     3,
		"graveyard":   4,
		"graveyards":  4,
		"hereafter":   3,
		"hereof":      3,
		"hereunder":   3,
		"hereupon":    3,
		"hideout":     3,
		"hideouts":    3,
		"homeowner":   3,
		"homeowners":  3,
		"lineup":      3,
		"lineups":     3,
		"makeover":    3,
		"moreover":    3,
		"pineapple":   3,
		"pineapples":  3,
		"shakeups":    4,
		"takeover":    3,
		"takeovers":   3,
		"timeout":     3,
		"vineyard":    3,
		"vineyards":   3,
	}
	nextPhIndex := currPh + len(currMap.phons)
	if nextPhIndex > len(phons)-1 {
		// There is no next phoneme so check to see if the next character is an e
		if stringAt(abc, currAbc+len(currMap.alphas), 1, "e") {
			return true
		}
		return false
	}
	nextPh := phons[nextPhIndex]

	if e, ok := hiddenSilentEs[abc]; ok {
		if e == currAbc+len(currMap.alphas) {
			return true
		}
	}

	if !isConsonant(nextPh) {
		return false
	}

	// If we get this far then we have two adjacent consonant phonemes.
	// Check to see if the next letter is an e
	//

	// BUT...
	// This is pretty ugly but having two adjacent phoneme consonants and a next letter
	// e isn't enough. Take G AH V N for instance (and there are plenty of other words
	// like it)... So we test for that explictly here
	if stringAt(abc, currAbc+len(currMap.alphas), 3, "ern") && nextPh == n {
		return false
	}

	if stringAt(abc, currAbc+len(currMap.alphas), 1, "e") {
		return true
	}
	return false
}

func (m phonToAlphas) mapB(phons []phoneme, currPh int, alphas string, currAbc int) phonToAlphas {
	new := m
	currAbc += len(new.alphas)
	currPh += len(new.phons)

	// Treat february as a special case
	if stringAt(alphas, currAbc, 6, "bruary") {
		new = phonToAlphas{
			append(m.phons, []phoneme{
				b,
			}...),
			m.alphas + "br",
		}
		return new
	}
	if stringAt(alphas, currAbc, 2, "bb") {
		if _, ok := phsAt(phons, currPh, []phoneme{b, b}); !ok {
			// As in stuBBorn,...
			new = phonToAlphas{
				append(m.phons, []phoneme{
					b,
				}...),
				m.alphas + "bb",
			}
			return new
		}
	}
	if stringAt(alphas, currAbc, 2, "pb") {
		// There's a silent p here
		// As in cuPBoard,...
		new = phonToAlphas{
			append(m.phons, []phoneme{
				b,
			}...),
			m.alphas + "pb",
		}
		return new
	}
	if stringAt(alphas, currAbc, 1, "b") {
		new = phonToAlphas{
			append(m.phons, []phoneme{
				b,
			}...),
			m.alphas + "b",
		}
		return new
	}
	return new
}

func (m phonToAlphas) mapL(phons []phoneme, currPh int, alphas string, currAbc int) phonToAlphas {
	new := m
	currAbc += len(new.alphas)
	currPh += len(new.phons)

	if s, ok := getStringAt(alphas, currAbc, 0, "lel"); ok {
		if _, ok := phsAt(phons, currPh, []phoneme{l, l}, []phoneme{l, ey}); !ok {
			// As in candLELight,...
			// But not as in candLelight, ukeLele,...
			new = phonToAlphas{
				[]phoneme{
					l,
				},
				s,
			}
			return new
		}
	}
	if s, ok := getStringAt(alphas, currAbc, 0, "hl", "ll"); ok {
		if _, ok := phsAt(phons, currPh, []phoneme{l, l}); !ok {
			// As in daHLia, yeLLow,...
			// But not as in goaLLess, we...
			new = phonToAlphas{
				[]phoneme{
					l,
				},
				s,
			}
			return new
		}
	}
	// Check for the silent h in delhi
	if s, ok := getStringAt(alphas, currAbc, 0, "lh"); ok {
		if _, ok := phsAt(phons, currPh, []phoneme{l, hh}); !ok {
			// The h isn't sounded
			// As in deLHi,...
			new = phonToAlphas{
				[]phoneme{
					l,
				},
				s,
			}
			return new
		}
	}
	if stringAt(alphas, currAbc, 1, "l") {
		new = phonToAlphas{
			[]phoneme{
				l,
			},
			"l",
		}
		return new
	}
	return new
}

func mapPhToA(phons []phoneme, alphas string) ([]phonToAlphas, error) {
	fail := func(p phoneme, a string) ([]phonToAlphas, error) {
		ret := []phonToAlphas{}
		err := phToAError{
			p,
			a,
		}
		return ret, err
	}
	ret := []phonToAlphas{}
	// current := 0
	currPh := 0
	currAbc := 0
	punctuationSkipped := ""
	for currPh < len(phons) {
		// for _, phon := range phons {
		var new phonToAlphas
		phon := phons[currPh]
		// Check we haven't run out of letters and return if we have
		if currAbc > len(alphas)-1 {
			return fail(phons[currPh], alphas)
		}
		if (alphas[currAbc] == '\'' && phon != ih) || alphas[currAbc] == '-' {
			// ' does not typically get expressed as a phoneme so move the current
			// character on. If the current phoneme is ih though then it does, for
			// instance in james's which has phonemes jh ey m z ih z
			// - has no effect on pronunciation
			//
			punctuationSkipped = string(alphas[currAbc])
			currAbc++
			continue
		}
		switch phon {
		case aa:
			if s, ok := getStringAt(alphas, currAbc, 0, "aar", "arr", "ear", "har", "er", "aa", "ar", "or"); ok {
				// As in AARdvark, stARRed, hEARt, philHARmonic, clERk, aardvARk, tomORrow,...
				// Now check whether we have a US r phoneme following
				if _, ok := phsAt(phons, currPh, []phoneme{aa, r}); ok {
					s = s[:len(s)-1]
				}
				new = phonToAlphas{
					[]phoneme{
						aa,
					},
					s,
				}
				break
			}

			// Note that aa can sound like oh in cot and ah as in barn
			if s, ok := getStringAt(alphas, currAbc, 2, "au", "ha", "ho", "ou"); ok {
				// As in cAUght, gymkHAna, HOnest, cOUgh,...
				new = phonToAlphas{
					[]phoneme{
						aa,
					},
					s,
				}
				break
			}
			// Some words have an -ah ending with the 'h' silent
			if stringAt(alphas, currAbc, 2, "ah") {
				if _, ok := phsAt(phons, currPh, []phoneme{aa, hh}); !ok {
					// The 'h' isn't sounded
					// As in shAH,...
					new = phonToAlphas{
						[]phoneme{
							aa,
						},
						"ah",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "o"); ok {
				// As in , gEnre, ,...
				new = phonToAlphas{
					[]phoneme{
						aa,
					},
					s,
				}
				break
			}
		case ae:
			if s, ok := getStringAt(alphas, currAbc, 2, "ai", "au", "ei"); ok {
				// As in plAIts, drAUght, revEIlle,,...
				new = phonToAlphas{
					[]phoneme{
						ae,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "i"); ok {
				// As in bAt, rEbate, merIngue,..
				new = phonToAlphas{
					[]phoneme{
						ae,
					},
					s,
				}
				break
			}
		case ah:
			if alphas == "hiccough" {
				new = phonToAlphas{
					[]phoneme{
						ah, p,
					},
					"ough",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "wer") {
				// As in ansWER,...
				s := "wer"
				// But only if the 'r' isn't sounded
				// As in ansWErable,...
				if _, ok := phsAt(phons, currPh, []phoneme{ah, r}); ok {
					s = "we"
				}
				new = phonToAlphas{
					[]phoneme{
						ah,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "our", "ure"); ok {
				// As in flavOUR, futURE...
				// But not if the r is sounded, for instance in armOURy, usUREr...
				if _, ok := phsAt(phons, currPh, []phoneme{ah, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ah,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ur") {
				// As in dUration,...
				// If the r is sounded, we should only grab the 'u'
				if _, ok := phsAt(phons, currPh, []phoneme{ah, r}); ok {
					new = phonToAlphas{
						[]phoneme{
							ah,
						},
						"u",
					}
					break
				}
			}
			// Handling this separately as this is a mix of a silent e followed by a
			// vowel so I may handle this more generally at some point
			if stringAt(alphas, currAbc, 2, "ea") {
				// As in likEAble,...
				new = phonToAlphas{
					[]phoneme{
						ah,
					},
					"ea",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "au", "ei", "ia", "ie", "io", "oo", "ou", "re", "ua", "ui", "ur", "yr"); ok {
				// As in becAUse, forEIgn, catERpillar, RussIA, conscIEnce, percussIOn, blOOd, rOUgh, theatRE, usUAlly, biscUIt, treasURe, zephYR,...
				new = phonToAlphas{
					[]phoneme{
						ah,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ar", "er", "or"); ok {
				// As in wizARd, pERcussion tractOR...
				// Need to be careful here. I don't want to consume the r if there's
				// an r phoneme in the phonetic spelling, for instance as in
				// d(d)o(ao)c(k)u(y uh)m(m)e(eh)n(n)t(t)a(ah)r(r)y(iy)
				if _, ok := phsAt(phons, currPh, []phoneme{ah, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ah,
						},
						s,
					}
					break
				}
			}
			// Some words have an -ah ending with the 'h' silent
			if stringAt(alphas, currAbc, 2, "ah") {
				if _, ok := phsAt(phons, currPh, []phoneme{ah, hh}); !ok {
					// The 'h' isn't sounded
					// As in purdAH,...
					new = phonToAlphas{
						[]phoneme{
							ah,
						},
						"ah",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "i", "o", "u", "y"); ok {
				// As in ..., propYlene,...
				new = phonToAlphas{
					[]phoneme{
						ah,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "r") {
				// As in houR,...
				new = phonToAlphas{
					[]phoneme{
						ah,
					},
					"r",
				}
				break
			}
		case ao:
			if stringAt(alphas, currAbc, 3, "hon") {
				// Words like HOnest start with a silent 'h'
				// As in HOnour,... and they don't necessarily start with hon-, for
				// instance disHOnourable,...
				new = phonToAlphas{
					[]phoneme{
						ao,
					},
					"ho",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ort") {
				// As in the borrowed French word, rappORT,...
				if _, ok := phsAt(phons, currPh, []phoneme{ao, ch}, []phoneme{ao, dh}, []phoneme{ao, sh}, []phoneme{ao, t}, []phoneme{ao, tr}, []phoneme{ao, th}, []phoneme{ao, thr}); !ok {
					// But not as in fORTunate, nORTHern, abORTion, repORT, pORTrait, nORTH,...
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						"ort",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "orp"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ao, p}, []phoneme{ao, f}); !ok {
					// The p is not sounded so swallow it here
					// As in the French cORPs,...
					// But not as in cORpulent, amORphous,...
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						s,
					}
					break
				}
			}
			// Borrowed French words ending -eur
			if stringAt(alphas, currAbc, 3, "eur") {
				if _, ok := phsAt(phons, currPh, []phoneme{ao, r}); !ok {
					// As in sabotEUR,...
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						"eur",
					}
					break
				}
			}
			// The test for a phonetic 'r' following the 'ao' doesn't work for
			// words like storeroom and forerunner so pulling these out as a special
			// case
			if (strings.HasPrefix(alphas, "storeroom") || strings.HasPrefix(alphas, "forerunner")) && stringAt(alphas, currAbc, 3, "ore") {
				new = phonToAlphas{
					[]phoneme{
						ao,
					},
					// The phonetic 'r' that follows 'ao' is for the lexical 'r' in room
					"ore",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "aor", "aur", "oar", "oor", "orr", "our", "uor", "wor", "or"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ao, r}); !ok {
					// As in extrAORdinary, dinosAUR, bOARd, flOOR, abhORRed, yOURs, flUOResce, sWORd, fORtnight...
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						s,
					}
					break
				} else {
					// As in AUric, hOAry, mOOrish, pOUring, stOry,...
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						// The 'r' is sounded so don't grab it here
						s[:len(s)-1],
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 4, "awer") {
				// if _, ok := phsAt(phons, currPh, []phoneme{ao, ax}, []phoneme{ao, axr}); !ok {
				// 	// The 'er' is not sounded
				// 	// As in drAWER,...
				// 	new = phonToAlphas{
				// 		[]phoneme{
				// 			ao,
				// 		},
				// 		"awer",
				// 	}
				// 	break
				// }
				if p, ok := phsAt(phons, currPh, []phoneme{ao, r, ax}, []phoneme{ao, ax}, []phoneme{ao, r, axr}, []phoneme{ao, axr}); ok {
					// As in gnAWers, gnAWer,...
					new = phonToAlphas{
						p[:len(p)-1],
						"aw",
					}
					break
				}
				// As in drAWER,...
				new = phonToAlphas{
					[]phoneme{
						ao,
					},
					"awer",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "au", "aw", "oa", "ou"); ok {
				// As in applAUd, clAWs, brOAd, thOUght,...

				// But check for linking r. I think this'll only be found for
				// aw, as in drAWing...
				p := []phoneme{
					ao,
				}
				if _, ok := phsAt(phons, currPh, []phoneme{ao, r}); ok {
					// But check the string! If it's 'awr' as in outlAWRy then
					//  don't swallow the r here.
					if !stringAt(alphas, currAbc, 3, "awr") {
						p = []phoneme{
							ao, r,
						}
					}
				}

				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "arr", "ar"); ok {
				// As in wARRed, wAR,...
				if _, ok := phsAt(phons, currPh, []phoneme{ao, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ao,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "o", "a"); ok {
				// As in fOr, wAter...
				new = phonToAlphas{
					[]phoneme{
						ao,
					},
					s,
				}
				break
			}
		case aw:
			if stringAt(alphas, currAbc, 4, "ough") {
				// As in bOUGH,...
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w}, []phoneme{aw}); ok {
					// As in plOUGH,...
					new = phonToAlphas{
						p,
						"ough",
					}
					break
				}

				new = phonToAlphas{
					[]phoneme{
						aw,
					},
					"ough",
				}
				break
			}
			if stringAt(alphas, currAbc, 4, "hour") {
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w, ax}, []phoneme{aw, w, axr}, []phoneme{aw, ax}, []phoneme{aw, axr}); ok {
					// As in HOURglass, HOUR,...
					new = phonToAlphas{
						p,
						"hour",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w, axl}, []phoneme{aw, axl}); ok {
					// As in HOURly,...
					new = phonToAlphas{
						p[:len(p)-1],
						"hour",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "our"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w, axn}, []phoneme{aw, axn}); ok {
					// As in sOURNess,...
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w, ax, r}, []phoneme{aw, ax, r}); ok {
					// As in sOURest,...
					// Leave the 'r' as it's sounded
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}

			}
			// if stringAt(alphas, currAbc, 3, "our") {
			// 	new = phonToAlphas{
			// 		[]phoneme{
			// 			aw,
			// 		},
			// 		"ou",
			// 	}
			// 	break
			// }
			if stringAt(alphas, currAbc, 2, "ow") {
				if p, ok := phsAt(phons, currPh, []phoneme{aw, w}); ok {
					// The lexical w is sounded as in bOWing,...
					new = phonToAlphas{
						p,
						"ow",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ao", "au", "ho", "ou", "ow"); ok {
				// As in mAO, sAUerkrAUt, HOur, sOUnd, gOWn,...

				// But check for linking w. Examples include flOUr, mAOist,
				// sAUerkraut,...
				p := []phoneme{
					aw,
				}
				if _, ok := phsAt(phons, currPh, []phoneme{aw, w}); ok {
					p = []phoneme{
						aw, w,
					}
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case ay:
			if stringAt(alphas, currAbc, 4, "eigh") {
				// As in hEIGHt,...
				new = phonToAlphas{
					[]phoneme{
						ay,
					},
					"eigh",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ai") {
				// As in nAIve,...
				// So check for this as two separate phonemes
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, iy}, []phoneme{ay, iy}); ok {
					// As in hawAII,...
					new = phonToAlphas{
						p[:len(p)-1],
						"a",
					}
					break
				}
			}
			if str, ok := getStringAt(alphas, currAbc, 0, "ais", "is"); ok {
				// As in AISle, ISland,...

				// Check for linking y.
				if _, ok := phsAt(phons, currPh, []phoneme{ay, y}); ok {
					new = phonToAlphas{
						[]phoneme{
							ay, y,
						},
						str,
					}
					break
				}
				// But NOT as in ISolate, chrIST, demonISe...
				if _, ok := phsAt(phons, currPh, []phoneme{ay, s}, []phoneme{ay, st}, []phoneme{ay, z}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ay,
						},
						str,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "aye", "igh", "ai", "ay"); ok {
				p, _ := phsAt(phons, currPh, []phoneme{ay, y}, []phoneme{ay})
				// As in AYE, hIGH, thAI, paraguAY...
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			// Catching -ihi- words.
			if stringAt(alphas, currAbc, 3, "ihi") {
				if _, ok := phsAt(phons, currPh, []phoneme{ay, y, hh}, []phoneme{ay, hh}); !ok {
					// No need to check for ok here. We know we have at least got
					// the phoneme ay
					p, _ := phsAt(phons, currPh, []phoneme{ay, y}, []phoneme{ay})
					// The h is not sounded so swallow it here
					// As in nIHilism,...
					new = phonToAlphas{
						p,
						"ih",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "ire") {
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, axr}, []phoneme{ay, axr}); ok {
					// This sounds like ire for words ending -ire
					// As in wIRE,...
					new = phonToAlphas{
						p,
						"ire",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, ax}, []phoneme{ay, ax}); ok {
					// This sounds like ire for words containing -ire-
					// As in fIREplace,...

					// But we need to check that the r isn't sounded before
					// swallowing the lexical r here.
					s := "ire"
					if _, ok := phsAt(phons, currPh, []phoneme{ay, y, ax, r}, []phoneme{ay, ax, r}); ok {
						// The r is sounded!
						s = "i"
					}
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, axl}, []phoneme{ay, y, axm}, []phoneme{ay, y, axn}, []phoneme{ay, axl}, []phoneme{ay, axm}, []phoneme{ay, axn}); ok {
					// As in entIRELy, fIREMan, dIRENess,...
					// Don't swallow the ax* here. Let it map to the lexical *
					new = phonToAlphas{
						p[:len(p)-1],
						"ire",
					}
					break
				}

			}
			if stringAt(alphas, currAbc, 2, "ir") {
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, ax, r}, []phoneme{ay, ax, r}, []phoneme{ay, r}); ok {
					// The r is sounded so don't swallow the r here
					// As in IRonies, Ironic,...
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
				} else {
					// The r isn't sounded so swallow it now
					// As in IRon,...
					p, _ := phsAt(phons, currPh, []phoneme{ay, y, ax}, []phoneme{ay, ax}, []phoneme{ay, y, axn}, []phoneme{ay, axn}, []phoneme{ay})
					new = phonToAlphas{
						p[:len(p)-1],
						"ir",
					}
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ie") {
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, ah}, []phoneme{ay, ah}, []phoneme{ay, y, ax}, []phoneme{ay, ax}, []phoneme{ay, y, axl}, []phoneme{ay, axl}, []phoneme{ay, y, eh}, []phoneme{ay, eh}, []phoneme{ay, y, ih}, []phoneme{ay, ih}, []phoneme{ay, y, iy}, []phoneme{ay, iy}); ok {
					// Looks like the 'e' is sounded so don't swallow it here
					// As in dIEt, clIEnt, quIEscent, socIEtal, quIEtus...
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ia") {
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y, ae}, []phoneme{ay, ae}, []phoneme{ay, y, ax}, []phoneme{ay, ax}, []phoneme{ay, y, axl}, []phoneme{ay, axl}, []phoneme{ay, y, axm}, []phoneme{ay, axm}, []phoneme{ay, y, axn}, []phoneme{ay, axn}, []phoneme{ay, y, axr}, []phoneme{ay, axr}, []phoneme{ay, y, ey}, []phoneme{ay, ey}); ok {
					// Looks like the 'a' is sounded so don't swallow it here
					// As in trIAngle, allIAnce, denIAL, dIAMond, gIANtess, lIAr, strIAtion,...
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "io") {

			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eye", "ae", "ei", "ie", "uy"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{ay, y}, []phoneme{ay}); ok {
					// As in EYEing, mAEstro, mEIosis, frIEd, bUY,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			// if str, ok := getStringAt(alphas, currAbc, 2, "ei", "ie"); ok {
			// 	// As in dIAl, EIther, frIEd,...
			// 	new = phonToAlphas{
			// 		[]phoneme{
			// 			ay,
			// 		},
			// 		str,
			// 	}
			// 	break
			// }
			// if s, ok := getStringAt(alphas, currAbc, 0, "ae"); ok {
			// 	// As in mAEstro,...
			// 	new = phonToAlphas{
			// 		[]phoneme{
			// 			ay,
			// 		},
			// 		s,
			// 	}
			// 	break
			// }
			if s, ok := getStringAt(alphas, currAbc, 1, "i", "u", "y"); ok {
				// As in tIme, flUtist, wrYly,...
				p := []phoneme{
					ay,
				}
				// Check for linking y. AS in beguIle, , cYanide,...
				if _, ok := phsAt(phons, currPh, []phoneme{ay, y}); ok {
					p = []phoneme{
						ay, y,
					}
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case b:
			new = new.mapB(
				phons, currPh, alphas, currAbc,
			)
			break
		case bl:
			// Catch deleted schwas. There are several patterns...
			if s, ok := getStringAt(alphas, currAbc, 0, "bell", "boll", "bal", "bel", "bol"); ok {
				// As in laBELLed, gamBOLLing, suBALtern, laBEL, gamBOL,...
				new = phonToAlphas{
					[]phoneme{
						bl,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "bbl", "bl"); ok {
				new = phonToAlphas{
					[]phoneme{
						bl,
					},
					s,
				}
				break
			}
		case ch:
			if stringAt(alphas, currAbc, 3, "tch") {
				new = phonToAlphas{
					// As in stiTCH,...
					[]phoneme{
						ch,
					},
					"tch",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "cc", "ch", "cz", "tt", "c"); ok {
				// As in cappuCCino, CHeese, CZech, aTTune, Cello,...
				new = phonToAlphas{
					[]phoneme{
						ch,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "t") {
				// As in signature, overture, etc
				new = phonToAlphas{
					[]phoneme{
						ch,
					},
					"t",
				}
				break
			}
		case d:
			if s, ok := getStringAt(alphas, currAbc, 2, "dd", "ld"); ok {
				// As in riDDle, and the silent l in wouLD,...
				p := []phoneme{d}
				if _, ok := phsAt(phons, currPh, []phoneme{d, d}); ok {
					// As in miDDay,... There's a double phonetic 'd' though, so don't grab
					//both lexical 'd's here
					s = "d"
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ed") {
				// As in wanderED,...
				// Although there is a kind of silent e here it doesn't get picked up
				// by isSilentE in this case because the preceding phoneme is er and so
				// not a consonant
				new = phonToAlphas{
					[]phoneme{
						d,
					},
					"ed",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "d") {
				new = phonToAlphas{
					[]phoneme{
						d,
					},
					"d",
				}
				break
			}
		case dh:
			if stringAt(alphas, currAbc, 2, "th") {
				new = phonToAlphas{
					[]phoneme{
						dh,
					},
					"th",
				}
				break
			}
		case eh:
			if stringAt(alphas, currAbc, 1, "x") {
				// As in Xmas,...
				if p, ok := phsAt(phons, currPh, []phoneme{eh, k, s}); ok {
					new = phonToAlphas{
						p,
						"x",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{eh, ks}); ok {
					new = phonToAlphas{
						p,
						"x",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "hei") {
				// As in HEIr,...
				new = phonToAlphas{
					[]phoneme{
						eh,
					},
					"hei",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "are") {
				// As in awARE,...
				// But not if the r is sounded as in aRena,...
				if _, ok := phsAt(phons, currPh, []phoneme{eh, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							eh,
						},
						"are",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "air", "are", "ear"); ok {
				// As in chAIR, fARE, EARthenware,...
				if _, ok := phsAt(phons, currPh, []phoneme{eh, r}, []phoneme{eh, er}); !ok {
					new = phonToAlphas{
						[]phoneme{
							eh,
						},
						s,
					}
					break
				} // else we handle the case with no r phoneme below
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ae", "ai", "ay", "ea", "ei", "eo", "ie"); ok {
				// As in AErobic, sAId, sAYs, endEAvour, thEIr, lEOpard, frIEnd,...
				new = phonToAlphas{
					[]phoneme{
						eh,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ar") {
				if _, ok := phsAt(phons, currPh, []phoneme{eh, r}); !ok {
					// As in scARce,...
					new = phonToAlphas{
						[]phoneme{
							eh,
						},
						"ar",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eh"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{eh, hh}); !ok {
					// The h isn't sounded so grab it here
					// As in tEHran,...
					new = phonToAlphas{
						[]phoneme{
							eh,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "i", "u"); ok {
				// As in contrAry, rEd, squIrrel, bUried,...
				new = phonToAlphas{
					[]phoneme{
						eh,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "r") {
				// As in houR,...
				new = phonToAlphas{
					[]phoneme{
						eh,
					},
					"r",
				}
				break
			}
		case ehl:
			if s, ok := getStringAt(alphas, currAbc, 1, "eal", "elh", "ell", "el"); ok {
				// As in wEALth, dELHi, wELLington, hELicopter,...
				new = phonToAlphas{
					[]phoneme{
						ehl,
					},
					s,
				}
				break
			}
		case er:
			// Pulling this out as a special case for now because I can't think of
			// any other words this occurs in.
			if stringAt(alphas, currAbc-1, 4, "iron") {
				// As in iROn,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"ro",
				}
				break
			}
			// Another special case. olo as in colonel sounds as er.
			if stringAt(alphas, currAbc, 3, "olo") {
				// As in cOLOnel,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"olo",
				}
				break
			}
			// Lots of herb- words with a silent 'h' phonetic variant
			if strings.HasPrefix(alphas, "herb") && currAbc == 0 {
				// I'm being deliberately specific here and only looking for words
				// that start herb-
				// As in HERbal,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"her",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ure") {
				// As in treasURE, futURE,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"ure",
				}
				break
			}
			// Some er sounds can have a sounded r when combined with for instance -ing.
			if s, ok := getStringAt(alphas, currAbc, 3, "err", "eur", "irr", "urr", "er"); ok {
				// As in refERRed, entreprenEURial, stIRRed, pURRed, transfERable,...
				if _, ok := phsAt(phons, currPh, []phoneme{er, r}); ok {
					// As in refErring, etc...
					// In this case don't swallow the r now, save it for the r phoneme
					s = s[:1]
				}
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "arr", "ear", "err", "eur", "irr", "our", "urr", "wer"); ok {
				// As in ARRay, hEARd, transfERRed, sabotEUR, stIRRed, yOURself, pURRed, ansWERed,...
				// What the hell! yourself is y er s eh l f but your is y ao r - why
				// the diffference?
				// That has to be a US English thing...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ah", "ar", "er", "eu", "ia", "ir", "or", "re", "ur", "yr"); ok {
				// As in hookAH,..., massEUse,..., theatRE, fUR, zephYR,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "a", "i") {
				// As in umbrellA, anImal...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"i",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "r") {
				// As in houR,...
				new = phonToAlphas{
					[]phoneme{
						er,
					},
					"r",
				}
				break
			}
		case ey:
			if stringAt(alphas, currAbc, 4, "eigh") {
				// As in wEIGH...
				// But check for a linking y
				p := []phoneme{
					ey,
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{ey, y}, []phoneme{ey}); ok {
					p = p1
				}
				new = phonToAlphas{
					p,
					"eigh",
				}
				break
			}
			// Capture borrowed French word -er endings
			if stringAt(alphas, currAbc, 2, "er") {
				// As in dossiER,...
				if _, ok := phsAt(phons, currPh, []phoneme{ey, r}); !ok {
					// But not as in paYRoll,...
					new = phonToAlphas{
						[]phoneme{
							ey,
						},
						"er",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ai") {
				// As in pAId
				s := "ai"
				p := []phoneme{
					ey,
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{ey, y, ax}, []phoneme{ey, ax}, []phoneme{ey, y, axl}, []phoneme{ey, axl}); ok {
					// As in AIl, AIl,...
					p = p1[:len(p1)-1]
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{ey, y, ih}, []phoneme{ey, ih}); ok {
					// As in algebrAist,...
					// The ai is split across two phonetic vowels
					s = "a"
					p = p1[:len(p1)-1]
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ao") {
				// As in gAOl,...
				s := "ao"
				p, _ := phsAt(phons, currPh, []phoneme{ey, y}, []phoneme{ey})
				if _, ok := phsAt(phons, currPh, []phoneme{ey, y, ao}, []phoneme{ey, ao}, []phoneme{ey, y, ax}, []phoneme{ey, ax}, []phoneme{ey, y, oh}, []phoneme{ey, oh}); ok {
					// As in Aorta, bAobab, chAotic,...
					s = "a"
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ea") {
				if p, ok := phsAt(phons, currPh, []phoneme{ey, y, aa}, []phoneme{ey, aa}, []phoneme{ey, ax}, []phoneme{ey, y, axr}, []phoneme{ey, axr}); ok {
					// As in rEAl (the unit of currency), eritrEA,...
					// Only swallow the e, the a will be picked up later
					new = phonToAlphas{
						p,
						"e",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ei") {
				// As in cunEIform,...
				if p, ok := phsAt(phons, currPh, []phoneme{ey, y, ih}, []phoneme{ey, ih}); ok {
					// The ei is split across two phonetic vowels so only grab
					// the first one
					new = phonToAlphas{
						p[:len(p)-1],
						"e",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ay", "ey"); ok {
				// If there is a linking y we should map it to the lexical y
				// later
				p := []phoneme{
					ey,
				}
				if _, ok := phsAt(phons, currPh, []phoneme{ey, y}); ok {
					// There's is linking y so map it now
					s = s[:len(s)-1]
				}
				new = phonToAlphas{
					p,
					s,
				}
				break

			}
			if stringAt(alphas, currAbc, 2, "ei") {
				// As in EIght,...
				// But check for a linking y
				p := []phoneme{
					ey,
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{ey, y}, []phoneme{ey}); ok {
					p = p1
				}
				new = phonToAlphas{
					p,
					"ei",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "au", "ea", "ee", "ei"); ok {
				// As in gAUge, wAY, grEAt, , soirEE, EIght,...
				new = phonToAlphas{
					[]phoneme{
						ey,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "et") {
				if _, ok := phsAt(phons, currPh, []phoneme{ey, t}); !ok {
					// Swallow the t here, as in cabarET,...
					// But also check for a linking y
					p := []phoneme{
						ey,
					}
					if p1, ok := phsAt(phons, currPh, []phoneme{ey, y}); ok {
						// As in ricochETing,...
						p = p1
					}
					new = phonToAlphas{
						p,
						"et",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "a") {
				p, _ := phsAt(phons, currPh, []phoneme{ey, y}, []phoneme{ey})
				// As in Ale,...
				new = phonToAlphas{
					p,
					"a",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "e") {
				// As in borrowed French words like touchE,...
				p, _ := phsAt(phons, currPh, []phoneme{ey, y}, []phoneme{ey})
				new = phonToAlphas{
					p,
					"e",
				}
				break
			}
		case f:
			if stringAt(alphas, currAbc, 2, "gh") {
				// As in couGH,...
				new = phonToAlphas{
					[]phoneme{
						f,
					},
					"gh",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "pph", "ph"); ok {
				// As in saPPhire, PHosPHorus,...
				new = phonToAlphas{
					[]phoneme{
						f,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ff") {
				new = phonToAlphas{
					[]phoneme{
						f,
					},
					"ff",
				}
				break
			}
			// Catch a possible silent l
			if stringAt(alphas, currAbc, 2, "lf") {
				// As in caLf,...
				new = phonToAlphas{
					[]phoneme{
						f,
					},
					"lf",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ft") {
				// A 't' following 'f' can be silent. Trying to catch this here because
				// it's harder to catch on the next phoneme
				if _, ok := phsAt(phons, currPh, []phoneme{f, t}, []phoneme{f, th}); !ok {
					// As in soFTen, but not as in liFT, fiFTh ...
					new = phonToAlphas{
						[]phoneme{
							f,
						},
						"ft",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "f") {
				new = phonToAlphas{
					[]phoneme{
						f,
					},
					"f",
				}
				break
			}
		case fl:
			if s, ok := getStringAt(alphas, currAbc, 0, "full", "fel", "ffl", "ghl", "phl", "fl"); ok {
				// As in awFULLy, liFELess, aFFLuent, rouGHLy, pamPHLet, FLorin,...
				new = phonToAlphas{
					[]phoneme{
						fl,
					},
					s,
				}
				break
			}
		case g:
			// A very special case to start with
			if alphas == "blackguard" && stringAt(alphas, currAbc, 2, "ck") {
				// There's a silent ck, and also a silent u after the g
				new = phonToAlphas{
					[]phoneme{
						g,
					},
					"ckgu",
				}
				break
			}
			// Trying to trap the silent h that can follow ex-
			if stringAt(alphas, currAbc, 2, "xh") {
				// As in eXHaust,...
				_, ok_gz := phsAt(phons, currPh, []phoneme{g, z})
				_, ok_gzh := phsAt(phons, currPh, []phoneme{g, z, hh})
				if ok_gz && !ok_gzh {
					// Okay, the h is silent so include it now
					new = phonToAlphas{
						[]phoneme{
							g, z,
						},
						"xh",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "x") {
				// As in eXactly but the g phoneme must be followed by a z...
				if _, ok := phsAt(phons, currPh, []phoneme{g, z}, []phoneme{g, zh}); ok {
					new = phonToAlphas{
						[]phoneme{
							g, z,
						},
						"x",
					}
				}
				break
			}
			// Handling silent u which typically follows a g
			if stringAt(alphas, currAbc, 3, "gue") {
				// As in dialoGUE, and many other words...
				if currPh == len(phons)-1 {
					new = phonToAlphas{
						[]phoneme{
							g,
						},
						"gue",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{g, ah}, []phoneme{g, ax}, []phoneme{g, er}, []phoneme{g, eh}, []phoneme{g, y}, []phoneme{g, ih}, []phoneme{g, yuw}); !ok {
					// So NOT as in beleaGUEred, vaGUEst, beleaGUEred, GUEss, arGUE, vaGUEst, arGUE,...
					// If we get here then the ue is not sounded
					new = phonToAlphas{
						[]phoneme{
							g,
						},
						"gue",
					}
					break
				} else if _, ok := phsAt(phons, currPh, []phoneme{g, y}, []phoneme{g, yuw}); !ok {
					// So NOT as in arGUE, arGUE,...
					// In other cases the u is silent and the vowel phoneme can be
					// processed next time round the loop - I think anyway...
					new = phonToAlphas{
						[]phoneme{
							g,
						},
						"gu",
					}
					break
				}
			}
			// There's also a silent u in other non-gue words
			if stringAt(alphas, currAbc, 2, "gu") {
				if _, ok := phsAt(phons, currPh, []phoneme{g, aa}, []phoneme{g, ae}, []phoneme{g, eh}, []phoneme{g, ay}); ok {
					// As in GUard, GUarantee, GUarantee, GUide,...
					new = phonToAlphas{
						[]phoneme{
							g,
						},
						"gu",
					}
					break
				}
			}
			// A silent h can also follow a g
			if stringAt(alphas, currAbc, 2, "gh") {
				// As in GHost,...
				// But check the h is silent first
				if _, ok := phsAt(phons, currPh, []phoneme{g, hh}); !ok {
					new = phonToAlphas{
						[]phoneme{
							g,
						},
						"gh",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "gg") {
				new = phonToAlphas{
					[]phoneme{
						g,
					},
					"gg",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "g") {
				new = phonToAlphas{
					[]phoneme{
						g,
					},
					"g",
				}
				break
			}
		case gr:
			if s, ok := getStringAt(alphas, currAbc, 0, "gar", "ggr", "gr"); ok {
				// As in marGARet, aGGRegate, GRipe,...
				// TDOD: Don't like adding margaret here as it's a proper name and the only
				// example of this mapping I can find in the dictionary. Should margaret be
				// removed from the dictionary?
				new = phonToAlphas{
					[]phoneme{
						gr,
					},
					s,
				}
				break
			}
		case hh:
			if stringAt(alphas, currAbc, 2, "wh") {
				// As in WHo, WHole,...
				new = phonToAlphas{
					[]phoneme{
						hh,
					},
					"wh",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "h") {
				new = phonToAlphas{
					[]phoneme{
						hh,
					},
					"h",
				}
				break
			}
		case ih:
			// First catch a silent h in some pronunciations of forehead
			if stringAt(alphas, currAbc, 3, "hea") {
				// As in foreHEAd,...
				new = phonToAlphas{
					[]phoneme{
						ih,
					},
					"hea",
				}
				break
			}
			// Catch borrowed French words ending -ier
			if stringAt(alphas, currAbc, 3, "ier") {
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ey}, []phoneme{ih, ey}, []phoneme{ih, y, ehr}, []phoneme{ih, ehr}); ok {
					// As in atelIer, concIerge,...
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "ior") {
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ao}, []phoneme{ih, ao}); ok {
					// As in fIord,...
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			// Some more French stuff, this time -ez- in rendezvous
			if stringAt(alphas, currAbc, 2, "ez") {
				// Check the lexical 'z' isn't sounded though
				if _, ok := phsAt(phons, currPh, []phoneme{ih, z}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ih,
						},
						"ez",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "aeo", "oeo"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ax}, []phoneme{ih, ax}, []phoneme{ih, y, axn}, []phoneme{ih, axn}, []phoneme{ih, y, oh}, []phoneme{ih, oh}); ok {
					// As in [palAEontology, archAEology], homOEopathy,...
					// The 'o' is sounded so don't swallow the 'o' here
					new = phonToAlphas{
						// []phoneme{
						// 	ih,
						// },
						p[:len(p)-1],
						s[:2],
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "io") {
				// As in biblIographic, perIodic, cardIoid,...
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ax}, []phoneme{ih, ax}, []phoneme{ih, y, oh}, []phoneme{ih, oh}, []phoneme{ih, y, oy}, []phoneme{ih, oy}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ei") {
				// As in nuclEi,...
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ay}, []phoneme{ih, ay}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						"e",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ii") {
				// As in radIi,...
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, ay}, []phoneme{ih, ay}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						"i",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ea", "ia", "ie"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, aa}, []phoneme{ih, aa}, []phoneme{ih, y, ax}, []phoneme{ih, ax}, []phoneme{ih, y, axr}, []phoneme{ih, axr}, []phoneme{ih, y, ey}, []phoneme{ih, ey}, []phoneme{ih, y, ae}, []phoneme{ih, ae}, []phoneme{ih, y, eh}, []phoneme{ih, eh}, []phoneme{ih, y, iy}, []phoneme{ih, iy}); ok {
					// As in cavIAr, folIAte, latvIA, nausEAte, asIAtic, fIEsta, medIEval...
					// Don't swallow the second letter here
					new = phonToAlphas{
						p[:len(p)-1],
						s[:1],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eo", "io", "yo"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{ih, y, oh}, []phoneme{ih, oh}, []phoneme{ih, y, ow}, []phoneme{ih, ow}); ok {
					// As in [gEology, embrYology], cheerIo,...
					// The 'o' is sounded so don't swallow the 'o' here (as we do below in thEOry)
					new = phonToAlphas{
						p[:len(p)-1],
						s[:1],
					}
					break
				}
			}
			if _, ok := getStringAt(alphas, currAbc, 2, "ey"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ih, y}); ok {
					// The y is sounded so don't swallow it here
					// as in bEyond
					new = phonToAlphas{
						[]phoneme{
							eh,
						},
						"e",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ae", "ai", "ay", "ea", "ee", "ei", "eo", "ey", "ia", "ie", "ui"); ok {
				// As in archAEology, portrAIt, mondAY, EAr (x! TODO: This example is wrong!), bEEn, wherEIn, thEOry, convERsation, donkEYs, carrIAge, sIEve, bUIlding,...
				new = phonToAlphas{
					[]phoneme{
						ih,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "er"); ok {
				// As in ERupt,...
				if _, ok := phsAt(phons, currPh, []phoneme{ih, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ih,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "hi"); ok {
				// It looks like we have a silent h
				// As in exHIbit,...
				new = phonToAlphas{
					[]phoneme{
						ih,
					},
					s,
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "i", "o", "u", "y"); ok {
				// As in encourAgIng, Erupt, bIn ,wOmen, bUsily, abYss,...
				p, _ := phsAt(phons, currPh, []phoneme{ih, y}, []phoneme{ih})
				new = phonToAlphas{
					// []phoneme{
					// 	ih,
					// },
					p,
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "'s", "s"); ok {
				// As in James'S, bridgeS,... and many other plural nouns
				if _, ok := phsAt(phons, currPh, []phoneme{ih, z}); ok {
					new = phonToAlphas{
						[]phoneme{
							ih, z,
						},
						s,
					}
					break
				}
			}
		case ihl:
			if s, ok := getStringAt(alphas, currAbc, 0, "uill", "ell", "eyl", "ill", "uil", "yll", "el", "il", "yl", "l"); ok {
				// As in in gUILLemot, ELLipse, monEYLender, guerILLa, gUILd, chlorophYLL, bELittle, untIL, pterodactYL, dieLectric*,...
				// *This is a mapping error elsewhere d AY is swallowing die,
				// leaving the lexical l to map to the phonetic ihl,
				new = phonToAlphas{
					[]phoneme{
						ihl,
					},
					s,
				}
				break
			}
		case ing:
			p := []phoneme{
				ing,
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eng", "ing"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ing, g}); ok {
					// As in ENGland, lINGer,...
					// We should save the g for later
					new = phonToAlphas{
						p,
						s[:len(s)-1],
					}
				} else {
					// As in waxING, and just about any -ing word...
					new = phonToAlphas{
						p,
						s,
					}
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eing", "uing"); ok {
				// As in agEING, catalogUING,...
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "inc", "ink", "inq", "inx", "ync", "ynx"); ok {
				// As in zINc, twINking, delINquent, sphINx, sYNchronous, pharYNx,...
				new = phonToAlphas{
					[]phoneme{
						ing,
					},
					s[:len(s)-1],
				}
				break
			}
		case iy:
			if stringAt(alphas, currAbc-1, 4, "peop") {
				// Treating pEOple as a special case here. There are other occurrences
				// of the letters eo which break otherwise, like stero
				new = phonToAlphas{
					[]phoneme{
						iy,
					},
					"eo",
				}
				break
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "aeo"); ok {
				// As in palAEontology...
				// Check for a linking y though
				p := []phoneme{
					iy,
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy}); ok {
					p = p1
				}
				new = phonToAlphas{
					p,
					"ae",
				}
				break
			}
			if strings.HasPrefix(alphas, "here") {
				// I'm being quite specific here. I don't want to break words like etHEREal.
				// As in words like HEREafter,...
				if p, ok := phsAt(phons, currPh, []phoneme{iy, ax, r}); ok {
					new = phonToAlphas{
						p,
						"ere",
					}
					break
				}
			}
			// Catch -eying words early, else they get caught up in the patterns that follow
			if stringAt(alphas, currAbc, 5, "eying") {
				// As in curtsEYing,...
				if p, ok := phsAt(phons, currPh, []phoneme{iy, y, ing}, []phoneme{iy, ing}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						"ey",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ear", "eer"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{iy, y, axm}, []phoneme{iy, axm}, []phoneme{iy, y, axn}, []phoneme{iy, axn}); ok {
					// As in EARMark, shEERNess,...
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{iy, y, ax, r}, []phoneme{iy, ax, r}); ok {
					// As in EArring,...
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{iy, y, ax}, []phoneme{iy, ax}); ok {
					// As in EAR, bEER...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "eir"); ok {
				p := []phoneme{}
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y, ax, r}, []phoneme{iy, ax, r}); ok {
					// The r is sounded so leave it for later
					s = "ei"
					p = p1[:len(p1)-1]
				}
				// As in wEIR, wEIRd,...
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y, axr}, []phoneme{iy, axr}, []phoneme{iy, y, ax}, []phoneme{iy, ax}); ok {
					p = p1
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ae", "ee"); ok {
				// As in AEgean, agrEE...
				// But what about pAEan, or agrEEing which might have a linking y
				p := []phoneme{
					iy,
				}
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy}); ok {
					p = p1
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ay") {
				// As in quAYside,...
				new = phonToAlphas{
					[]phoneme{
						iy,
					},
					"ay",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ea"); ok {
				// As in mEAt,...
				p := []phoneme{
					iy,
				}
				// Check to see if this is two vowels though, as in rEarm, rEact, aegEan, linEaments, aegEan, azalEa, crEate, linEage,...
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y, aa}, []phoneme{iy, aa}, []phoneme{iy, y, ae}, []phoneme{iy, ae}, []phoneme{iy, y, ax}, []phoneme{iy, ax}, []phoneme{iy, y, axm}, []phoneme{iy, axm}, []phoneme{iy, y, axn}, []phoneme{iy, axn}, []phoneme{iy, y, axr}, []phoneme{iy, axr}, []phoneme{iy, y, ey}, []phoneme{iy, ey}, []phoneme{iy, y, ih}, []phoneme{iy, ih}); ok {
					p = p1[:len(p1)-1]
					s = "e"
				}
				// As in annEAl,...
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y, axl}, []phoneme{iy, axl}); ok {
					p = p1[:len(p1)-1]
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "eh") {
				// As in rEHabilitate, vEHicle,...
				if _, ok := phsAt(phons, currPh, []phoneme{iy, hh}); !ok {
					// But not as in vEHicular,... - where the phonetic 'h' is sounded
					p, _ := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy})
					new = phonToAlphas{
						// []phoneme{
						// 	iy,
						// },
						p,
						"eh",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ei"); ok {
				// As in EIther, or plebEIan,...
				p, _ := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy})
				// p := []phoneme{
				// 	iy,
				// }
				// Is this two vowels?
				// As in homogenEity, thEism, bEing,...
				if p1, ok := phsAt(phons, currPh, []phoneme{iy, y, ax}, []phoneme{iy, ax}, []phoneme{iy, y, ih}, []phoneme{iy, ih}, []phoneme{iy, y, ing}, []phoneme{iy, ing}); ok {
					p = p1[:len(p1)-1]
					s = "e"
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ey") {
				// As in jockEYing, monEY,...
				p, _ := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy})
				new = phonToAlphas{
					// []phoneme{
					// 	iy,
					// },
					p,
					"ey",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ie") {
				// As in shrIEked,...
				// All other examples of 'ie' are two adjacent vowels so just
				// check for another vowel or a linking y, as dirtIest and in
				// pretty much any word ending -iest, experIential, filthIest, alIen,...
				if _, ok := phsAt(phons, currPh, []phoneme{iy, ax}, []phoneme{iy, eh}, []phoneme{iy, ih}, []phoneme{iy, y}); !ok {
					// There's no linking phoneme so map all of 'ie' now
					new = phonToAlphas{
						[]phoneme{
							iy,
						},
						"ie",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "oe") {
				// As in homOEopathy, phOEnix,...
				p, _ := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy})
				new = phonToAlphas{
					p,
					"oe",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "e"); ok {
				// There seems to be a pattern that 'e' followed by a sounded 'r' can
				// be sounded as iy ah
				// As in bactEria,...
				if p, ok := phsAt(phons, currPh, []phoneme{iy, y, ax, r}, []phoneme{iy, ax, r}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ae", "oe"); ok {
				// As in pAEan, diarrhOEa,...
				// if _, ok := phsAt(phons, currPh, []phoneme{iy, ax, r}); ok {
				new = phonToAlphas{
					[]phoneme{
						iy, ax,
					},
					s,
				}
				break
				// }
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "e", "i", "y"); ok {
				// As in shE, catchIest, shylY,...
				p, _ := phsAt(phons, currPh, []phoneme{iy, y}, []phoneme{iy})
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case jh:
			if s, ok := getStringAt(alphas, currAbc, 2, "ch", "dg", "di", "dj", "gg"); ok {
				// As in sandwiCH, heDGing, solDIer, aDJust, suGGest,...
				new = phonToAlphas{
					[]phoneme{
						jh,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "g", "j"); ok {
				// As in ginGer,... We don't want to swallow up the e in er
				new = phonToAlphas{
					[]phoneme{
						jh,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "de", "du") {
				// As in granDEur, eDUcate... But leave the u to be picked up by the following
				// phoneme
				new = phonToAlphas{
					[]phoneme{
						jh,
					},
					"d",
				}
				break
			}
		case k:
			if alphas == "ok" {
				new = phonToAlphas{
					[]phoneme{
						k, ey,
					},
					"k",
				}
				break
			}
			// The k phoneme can optionally appear in -ngth words according to Google
			// and they do appear in the CMU dictionary
			if stringAt(alphas, currAbc-2, 4, "ngth") {
				if _, ok := phsAt(phons, currPh-1, []phoneme{ng, k, th}); ok {
					// As in lengTH,...
					new = phonToAlphas{
						[]phoneme{
							k, th,
						},
						"th",
					}
					break
				}
			}
			// Detecting words ending que...
			if stringAt(alphas, currAbc, 3, "que") && currPh == len(phons)-1 {
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					"que",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "kh"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{k, hh}); !ok {
					// It looks like the h is silent so swallow it now.
					new = phonToAlphas{
						[]phoneme{
							k,
						},
						s,
					}
					break
				}
			}
			// Catch a possible silent l
			if stringAt(alphas, currAbc, 2, "lk") {
				// As in waLK, foLK,...
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					"lk",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ct") {
				// But make sure the lexical 't' isn't sounded - and 't' isn't always sound as t!
				if _, ok := phsAt(phons, currPh, []phoneme{k, ch}, []phoneme{k, sh}, []phoneme{k, t}, []phoneme{k, tr}); !ok {
					// As in aCTuary, traCTion, striCTly, buTTRess,...
					new = phonToAlphas{
						[]phoneme{
							k,
						},
						"ct",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "cqu", "qu"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{k, w}); ok {
					// As aCQUire, QUoth and all sorts of words containing qu...
					new = phonToAlphas{
						[]phoneme{
							k, w,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "cqu"); ok {
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					s,
				}
				break
			}
			// Check for q on it's own. These will be in borrowed words or foreign placenames
			if stringAt(alphas, currAbc, 1, "q") {
				// As in Qatar,...
				// new = phonToAlphas{
				// 	[]phoneme{
				// 		k,
				// 	},
				// 	"q",
				// }
				// break
			}
			if stringAt(alphas, currAbc, 4, "xion") {
				if _, ok := phsAt(phons, currPh, []phoneme{k, sh, n}); ok {
					new = phonToAlphas{
						[]phoneme{
							k, sh, n,
						},
						"xion",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ckc", "kc"); ok {
				// As in blaCKCurrant,...
				if _, ok := phsAt(phons, currPh, []phoneme{k, k}, []phoneme{k, ch}, []phoneme{k, kl}, []phoneme{k, kr}); !ok {
					// But not as in blaCKcurrant, baCKchat, saCKcloth, coCKcrow...
					new = phonToAlphas{
						[]phoneme{
							k,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "cch", "ch", "ck", "qu"); ok {
				// As in saCCHarine, CHoir (or loCH approximately), chiCKen, QUay,...
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "kk"); ok {
				p := []phoneme{k}
				if _, ok := phsAt(phons, currPh, []phoneme{k, k}); ok {
					// As in treKKing,...
					// There's a double phonetic 'k' though, so don't grab both
					// lexical 'k's here
					s = "k"
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			// Trapping words like exchequer, else the 'c' gets swallowed by
			// processing of 'xc' further on
			if stringAt(alphas, currAbc, 3, "xch") {
				if _, ok := phsAt(phons, currPh, []phoneme{k, s, ch}); ok {
					// As in eXchange
					new = phonToAlphas{
						[]phoneme{
							k, s,
						},
						"x",
					}
					break
				}
			}
			if str, ok := getStringAt(alphas, currAbc, 2, "cc", "xc"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{k, s, ch}, []phoneme{k, s, k}, []phoneme{k, s, kl}, []phoneme{k, s, kr}); ok {
					// As in eXCHange, eXCoriate, eXClude,, eXCRete...
					// Save the (second) k (or kl, kr) for later...
					new = phonToAlphas{
						[]phoneme{
							k, s,
						},
						str[:1],
					}
					break
				}
				// We need to be careful here and check the phonemes for k, s
				if _, ok := phsAt(phons, currPh, []phoneme{k, s}); ok {
					// As in suCCess, eXCited,...
					new = phonToAlphas{
						[]phoneme{
							k, s,
						},
						str,
					}
					break
				}
				// As in suCCour,...
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					str,
				}
				break
			}
			if str, ok := getStringAt(alphas, currAbc, 2, "xs"); ok {
				// As in coXSwain,...
				new = phonToAlphas{
					[]phoneme{
						k, s,
					},
					str,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "c", "k", "q"); ok {
				// Note the 'q' here
				// As in ...,Qatar,...
				new = phonToAlphas{
					[]phoneme{
						k,
					},
					s,
				}
				break
			}
			// Trying to trap the silent h that can follow ex-
			if stringAt(alphas, currAbc, 2, "xh") {
				// As in eXHibition,...
				_, ok_gz := phsAt(phons, currPh, []phoneme{k, s})
				_, ok_gzh := phsAt(phons, currPh, []phoneme{k, s, hh})
				if ok_gz && !ok_gzh {
					// Okay, the h is silent so include it now
					new = phonToAlphas{
						[]phoneme{
							k, s,
						},
						"xh",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "xts"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{k, sts}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "xed", "xt"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{k, st}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			//The 't' following x is sometimes not pronounced so catch it here
			if _, ok := getStringAt(alphas, currAbc, 0, "xtb"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{k, s, b}); ok {
					// The 't' is silent so swallow it now
					// As in teXTbook,... (I think this and its plural are the only examples)
					new = phonToAlphas{
						[]phoneme{
							k, s,
						},
						"xt",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "x") {
				if p, ok := phsAt(phons, currPh, []phoneme{k, s}, []phoneme{k, sh}); ok {
					// As in eXit, anXious,...
					new = phonToAlphas{
						p,
						"x",
					}
				}
			}
		case kl:
			// if s, ok := getStringAt(alphas, currAbc, 0, "call", "ckcl", "quel", "ccl", "chl", "ckl", "col", "cul", "kel", "khl", "ctl", "cl", "kl"); ok {
			if s, ok := getStringAt(alphas, currAbc, 0, "call", "ckcl", "quel", "ccl", "chl", "ckl", "col", "cul", "kel", "ctl", "cl", "kl"); ok {
				// As in dramatiCALLy, saCKCLoth, uniQUELy, aCCLaim, CHLorine, tiCKLer, choCOLate, faCULty, liKELy,
				// striCTLy, CLock, KLingon,...
				new = phonToAlphas{
					[]phoneme{
						kl,
					},
					s,
				}
				break
			}
		case kr:
			if s, ok := getStringAt(alphas, currAbc, 0, "ckcr", "cker", "ccr", "chr", "ckr", "cr", "kr"); ok {
				// As in coCKCRow, ..., aCCRete, laCHRymose, coCKRoach, maCKERel, CRisis, sauerKRaut,...
				new = phonToAlphas{
					[]phoneme{
						kr,
					},
					s,
				}
				break
			}
		case kt:
			// This is a bit of a hack. We have a pronunciation of actualities,
			// A KT CH Y UW AE L IH T IY Z and we need to find a home for the CH
			if s, ok := getStringAt(alphas, currAbc, 0, "ct"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{kt, ch}); ok {
					new = phonToAlphas{
						[]phoneme{
							kt, ch,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ched", "cked", "kked", "lked", "qued", "ced", "ckt", "ked", "lkt", "ct", "kt"); ok {
				// As in aCHED, blaCKED, treKKed, bauLKED, piQUED, arCED, baCKTracked*, tusKED, foLKTale*, interaCT, desKTop*,...
				// *These maybe shouldn't be a kt
				new = phonToAlphas{
					[]phoneme{
						kt,
					},
					s,
				}
				break
			}
		case l:
			new = new.mapL(
				phons, currPh, alphas, currAbc,
			)
			if !new.equal(phonToAlphas{}) {
				break
			}
			// Check for borrowed Italian words like seraglio
			if s, ok := getStringAt(alphas, currAbc, 0, "gl"); ok {
				// As in intaGLio,...
				new = phonToAlphas{
					[]phoneme{
						l,
					},
					s,
				}
				break
			}
			// Check for deleted schwa in -ally, -ully word endings
			if s, ok := getStringAt(alphas, currAbc, 0, "all", "ial", "ill", "oll", "ull"); ok {
				// As in principALLy, specIAL, pencILLed, gambOLLing, wonderfULLy,...
				new = phonToAlphas{
					[]phoneme{
						l,
					},
					s,
				}
				break
			}
			// Check for a deleted schwa, for instance in a(ae)n(n)i(ih)m(m)al(l).
			// This should be dealt with more generically (ie.e not just for the l
			// phoneme but for now this will catch some of the more common examples
			if s, ok := getStringAt(alphas, currAbc, 2, "al", "el", "il", "ol", "ul"); ok {
				// As in animAL, squirrEL, councIL, cathOLic, facULty,...
				new = phonToAlphas{
					[]phoneme{
						l,
					},
					s,
				}
				break
			}
		case m:
			// A couple of abbreviations whcih can't easily be mapped phoneme
			// by phoneme
			if alphas == "mr" || alphas == "mrs" {
				new = phonToAlphas{
					phons,
					alphas,
				}
				break
			}
			if strings.HasPrefix(alphas, "mc") && currAbc == 0 {
				// As in MCgregor
				if p, ok := phsAt(phons, currPh, []phoneme{m, ax}); ok {
					if _, ok = phsAt(phons, currPh, []phoneme{m, ax, k}); !ok {
						// The c isn't being sounded so swallow it here.
						new = phonToAlphas{
							p,
							"mc",
						}
					} else {
						new = phonToAlphas{
							p,
							"m",
						}
					}
					break
				}

			}
			if _, ok := getStringAt(alphas, currAbc, 3, "med", "mes"); ok {
				// Don't want to swallow these characters unless they're at the end of
				// a word. Think of MEDical, tiMEShift,...
				if currAbc+3 == len(alphas) {
					// As in tiMED, liMES,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"me",
					}
					break
				}
			}
			// Treating dumbbell as special case. The first b should really be
			// attached to the m phoneme
			if alphas == "dumbbell" {
				new = phonToAlphas{
					[]phoneme{
						m,
					},
					"mb",
				}
				break
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "mem"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{m, m}); ok {
					// As in hoMEmade,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"me",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{m, ax, m}, []phoneme{m, ih, m}, []phoneme{m, eh, m}); !ok {
					// Not as in MEMber, MEMento, imMEMorial,...
					// But as in hoMEMade,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mem",
					}
					break
				}
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "mbm"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{m, m}); ok {
					// As in entoMBment,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mb",
					}
					break
				} else {
					// As in entoMBMent,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mbm",
					}
					break
				}
			}
			// Catch a possible mbl, before we test for mb
			if _, ok := getStringAt(alphas, currAbc, 0, "mboll", "mbl"); ok {
				// As in gamBOLLed, tuMBLer,...
				if _, ok := phsAt(phons, currPh, []phoneme{m, bl}); ok {
					// As in raMBLer,... but leave the bl for later
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"m",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "mb") {
				// As in cliMber,...
				if _, ok := phsAt(phons, currPh, []phoneme{m, b}); !ok {
					// Not as in claMBer,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mb",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "mn") {
				// As in hyMN,...
				if _, ok := phsAt(phons, currPh, []phoneme{m, n}); !ok {
					// Not as in hyMNal,...
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mn",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "mm") {
				// As in diMMer,...
				if _, ok := phsAt(phons, currPh, []phoneme{m, m}); !ok {
					// As in rooMMate,...
					// There's a double phonetic 'm' though, so don't grab both
					// lexical 'm's here
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						"mm",
					}
					break
				}
			}
			// Catch a possible silent g, or l
			if s, ok := getStringAt(alphas, currAbc, 2, "gm", "lm"); ok {
				// As in diaphraGM, caLM,...
				new = phonToAlphas{
					[]phoneme{
						m,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "mp"); ok {
				// As in redeMPtion,...
				// Check for a silent p though. A lexical p can also be sounded as f (as in eMPhatic)!
				if _, ok := phsAt(phons, currPh, []phoneme{m, f}, []phoneme{m, fl}, []phoneme{m, p}, []phoneme{m, pl}, []phoneme{m, pr}); !ok {
					// The lexical 'p' is silent so grab it now.
					new = phonToAlphas{
						[]phoneme{
							m,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "m") {
				new = phonToAlphas{
					[]phoneme{
						m,
					},
					"m",
				}
				break
			}
			// In the word drachm, the ch is silent
			if stringAt(alphas, currAbc, 3, "chm") {
				new = phonToAlphas{
					[]phoneme{
						m,
					},
					"chm",
				}
			}
		case n:
			if s, ok := getStringAt(alphas, currAbc, 2, "ln", "mn"); ok {
				// As in lincoLN, mnemonic, ...
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "nn") {
				// As in naNNy,...
				s := "nn"
				if _, ok := phsAt(phons, currPh, []phoneme{n, n}); ok {
					// But not as in uNnatural,...
					// There's a double-n phoneme so save the other one for the other
					// lexical 'n'
					s = "n"
				}
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "nkn"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{n, n}, []phoneme{n, k}); ok {
					// As in uNknown, baNknote,...
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						"n",
					}
				} else {
					// Asin uNKNown,...
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						s,
					}
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "nen"); ok {
				// This could be represented by one phoneme as in oNENess but
				// we need to be careful
				if _, ok := phsAt(phons, currPh, []phoneme{n, n}); ok {
					// As in oNeness,...
					// But not as in oNENess,...
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						"n",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{n, ax, n}, []phoneme{n, eh, n}, []phoneme{n, ih, n}); ok {
					// Not as in oppoNENt, uNENding, lINEN,...
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						"n",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{n, axn}); !ok {
					// Not as in oppoNENt,...
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						s,
					}
					break
				}
			}
			// Catch the borrowed 'ñ' (n-yah)
			if s, ok := getStringAt(alphas, currAbc, 2, "gn", "nh"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{n, y}); ok {
					// As in siGNor, piraNHa,...
					new = phonToAlphas{
						[]phoneme{
							n, y,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "gnn"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{n, n}); !ok {
					// As in foreiGNNess
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						s,
					}
					break
				}
			}

			if s, ok := getStringAt(alphas, currAbc, 2, "dn", "gn", "kn", "mp", "pn"); ok {
				// As in weDNesday, siGN, KNee, coMPtroller, PNeumatic,...
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "nd") {
				if _, ok := phsAt(phons, currPh, []phoneme{n, jh}); ok {
					// As in graNdeur,...
					// Leave the d to be processed later
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						"n",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{n, d}, []phoneme{n, dz}); !ok {
					// Looks like the d is not sounded as in laNDs...
					// So swallow it now
					new = phonToAlphas{
						[]phoneme{
							n,
						},
						"nd",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 4, "wain") {
				// As in coxswain,...
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					"wain",
				}
				break
			}
			// Check for deleted schwa
			if s, ok := getStringAt(alphas, currAbc, 0, "ain", "an", "ern", "ian", "ion", "ten"); ok {
				// As in certAIN, tartAN, govERNment, alsatIAN, fashIONable, sofTEN...
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "en", "in", "on"); ok {
				// As in christEN, basIN, arsON,...
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "n") {
				new = phonToAlphas{
					[]phoneme{
						n,
					},
					"n",
				}
				break
			}
		case ng:
			if stringAt(alphas, currAbc, 4, "ngue") {
				// As in toNGUE,...
				new = phonToAlphas{
					[]phoneme{
						ng,
					},
					"ngue",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "nc") {
				// As in the place name, Altrincham, which Wikipedia confirms is pronounced with a
				// phonetic 'ng'
				if _, ok := phsAt(phons, currPh, []phoneme{ng, k}, []phoneme{ng, kr}, []phoneme{ng, ks}, []phoneme{ng, kt}); !ok {
					// But not as in acupuNCture, paNCReas...
					new = phonToAlphas{
						[]phoneme{
							ng,
						},
						"nc",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "nd") {
				// As in haNDkerchief,...
				new = phonToAlphas{
					[]phoneme{
						ng,
					},
					"nd",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ng") {
				// If next phoneme is gr don't swallow the g here
				if _, ok := phsAt(phons, currPh, []phoneme{ng, gr}); ok {
					// As in coNgregate,...
					new = phonToAlphas{
						[]phoneme{
							ng,
						},
						"n",
					}
					break
				}
				// If next phoneme is g then both phonemes map to the letters ng
				// As in aNGuish,...
				if _, ok := phsAt(phons, currPh, []phoneme{ng, g}); ok {
					new = phonToAlphas{
						[]phoneme{
							ng,
						},
						"n",
					}
				} else {
					// As in wiNG,...
					new = phonToAlphas{
						[]phoneme{
							ng,
						},
						"ng",
					}
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "n") {
				// As in think...
				new = phonToAlphas{
					[]phoneme{
						ng,
					},
					"n",
				}
				break
			}
		case ow:
			if s, ok := getStringAt(alphas, currAbc, 0, "ough", "aoh"); ok {
				// As in pharAOH, furlOUGH,... at least with a US accent
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					s,
				}
				break
			}
			if stringAt(alphas, 0, 6, "brooch") {
				// I think this is about the only word in the English language with oo
				// represented by the phoneme ow. In other oo words the oo is sounded
				// as two separate vowels, as in cooperate, microorganism,...
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					"oo",
				}
				break
			}
			// Borrowed French words ending -eau, -eaus, -eaux. Do we really
			// pluralise borrowed French words with an s?
			if s, ok := getStringAt(alphas, currAbc, 0, "eaus", "eaux", "eau"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ow, z}); ok {
					// Don't swallow the s, or x as it's sounded
					s = "eau"
				}
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "au") {
				// As in chAUvanist,...
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					"au",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "oa") {
				// As in cOAt,...
				if p, ok := phsAt(phons, currPh, []phoneme{ow, w, aa}, []phoneme{ow, aa}, []phoneme{ow, w, ae}, []phoneme{ow, ae}, []phoneme{ow, ao}, []phoneme{ow, w, ao}, []phoneme{ow, w, ax}, []phoneme{ow, ax}, []phoneme{ow, w, axl}, []phoneme{ow, axl}, []phoneme{ow, w, axn}, []phoneme{ow, axn}, []phoneme{ow, w, axr}, []phoneme{ow, axr}, []phoneme{ow, w, ey}, []phoneme{ow, ey}, []phoneme{ow, w, ih}, []phoneme{ow, ih}); ok {
					// As in kOala, radiOactive, prOactive, cOordinate, jerbOas, cOalesce, psychOanalysis, bOa, crOatia, inchOate,...
					new = phonToAlphas{
						p[:len(p)-1],
						"o",
					}
				} else {
					new = phonToAlphas{
						[]phoneme{
							ow,
						},
						"oa",
					}
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "oe") {
				if p, ok := phsAt(phons, currPh, []phoneme{ow, w, eh}, []phoneme{ow, eh}, []phoneme{ow, w, er}, []phoneme{ow, er}, []phoneme{ow, w, iy}, []phoneme{ow, iy}); ok {
					// As in whosOever, cOerce, micrOelectronics,...
					new = phonToAlphas{
						p[:len(p)-1],
						"o",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ou") {
				// As in sOUl,...
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					"ou",
				}
				break

			}
			if _, ok := getStringAt(alphas, currAbc, 0, "oww"); ok {
				// As in glOWWorm,...
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					"ow",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ew", "ow"); ok {
				// As in micrOwave,...
				if _, ok := phsAt(phons, currPh, []phoneme{ow, w}); ok {
					// The lexical 'w' is sounded so leave it for the phonetic 'w'
					new = phonToAlphas{
						[]phoneme{
							ow,
						},
						s[:len(s)-1],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eo", "ew", "ow"); ok {
				// As in yEOman, bOW,...
				new = phonToAlphas{
					[]phoneme{
						ow,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "oh") {
				// As in OHm,...
				if _, ok := phsAt(phons, currPh, []phoneme{ow, hh}); !ok {
					// But not as in bOHemia,...
					// Check for linking w, as in prOHibition,...
					p, _ := phsAt(phons, currPh, []phoneme{ow, w}, []phoneme{ow})
					new = phonToAlphas{
						p,
						"oh",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ot") {
				// As in depot,...
				if _, ok := phsAt(phons, currPh, []phoneme{ow, t}, []phoneme{ow, dh}, []phoneme{ow, sh}, []phoneme{ow, th}, []phoneme{ow, tr}); !ok {
					// But not as in rOTe, bOTH, pOTion, clOTHE, synchrOTRon,...
					// So it looks like the t is silent here
					new = phonToAlphas{
						[]phoneme{
							ow,
						},
						"ot",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "o"); ok {
				// As in Over,...
				p, _ := phsAt(phons, currPh, []phoneme{ow, w}, []phoneme{ow})
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case oy:
			p := []phoneme{
				oy,
			}
			if _, ok := phsAt(phons, currPh, []phoneme{oy, y}); ok {
				p = []phoneme{
					oy, y,
				}
			}
			if stringAt(alphas, currAbc, 3, "uoy") {
				new = phonToAlphas{
					p,
					"uoy",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "oi", "oy"); ok {
				// As in chOIce, tOY...
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case p:
			if stringAt(alphas, currAbc, 2, "pp") {
				new = phonToAlphas{
					[]phoneme{
						p,
					},
					"pp",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "bp") {
				// As in subpoena,...
				new = phonToAlphas{
					[]phoneme{
						p,
					},
					"bp",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "pt") {
				if _, ok := phsAt(phons, currPh, []phoneme{p, ch}, []phoneme{p, sh}, []phoneme{p, t}, []phoneme{p, tr}, []phoneme{p, th}, []phoneme{p, thr}); !ok {
					// It's not caPTure, descriPTIon, comPTRoller, uPTurn, upTHrust,...
					// So it looks like there's a silent t in words like bankruptcy
					new = phonToAlphas{
						[]phoneme{
							p,
						},
						"pt",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "p") {
				new = phonToAlphas{
					[]phoneme{
						p,
					},
					"p",
				}
				break
			}
		case pl:
			if s, ok := getStringAt(alphas, currAbc, 0, "pall", "p-l", "pel", "ppl", "pul", "pl"); ok {
				// As in princiPALLy, hooP-La, hoPELess, suPPLy, sePULchre, couPLet,...
				new = phonToAlphas{
					[]phoneme{
						pl,
					},
					s,
				}
				break
			}
		case pr:
			if s, ok := getStringAt(alphas, currAbc, 0, "par", "per", "pir", "por", "ppr", "pr"); ok {
				// as in comPARably, temPERature, asPIRin, contemPORary, aPPRaise, PRove,...
				new = phonToAlphas{
					[]phoneme{
						pr,
					},
					s,
				}
				break
			}
		case r:

			if s, ok := getStringAt(alphas, currAbc, 2, "aur", "our", "rrh", "ar", "ir", "or", "rh", "rr", "ur", "wr"); ok {
				// As in restAURant, labOURer, ciRRHosis, solitARy, aspIRin, conservatORy, RHythm, eRRor, natURal, WRite,...
				// But NOT as in speaRHead,...
				if _, ok := phsAt(phons, currPh, []phoneme{r, hh}); !ok {
					new = phonToAlphas{
						[]phoneme{
							r,
						},
						s,
					}
					break
				}
			}
			// Is there a better way of handling these 'special' cases?
			//
			// Handle forehead(s) as a special case
			if strings.HasPrefix(alphas, "forehead") {
				new = phonToAlphas{
					[]phoneme{
						r,
					},
					"re",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "r") {
				new = phonToAlphas{
					[]phoneme{
						r,
					},
					"r",
				}
				break
			}
		case s:
			// Treating this a special case. There are spellings of conversation in
			// the dictionary which drop the er phoneme
			if stringAt(alphas, currAbc, 3, "ers") {
				new = phonToAlphas{
					[]phoneme{
						s,
					},
					"ers",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ces") {
				// We have to be careful here. Thre are lots of options
				if _, ok := phsAt(phons, currPh, []phoneme{s, ax}, []phoneme{s, eh}, []phoneme{s, ih}, []phoneme{s, iy}, []phoneme{s, sh}, []phoneme{s, s}); !ok {
					// So, not as in neCEssary, anCEstry, spiCEs, faeCEs, spaCESHip, iCESkate,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"ces",
					}
					break
				}
			}
			// Focussing on postscript(s) which can be rendered P OH S K R IH P T (S)
			if stringAt(alphas, currAbc, 5, "stscr") {
				if _, ok := phsAt(phons, currPh, []phoneme{s, k}, []phoneme{s, kr}); ok {
					// As in poSTScript,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"sts",
					}
					break
				}
			}
			// breaststroke is a right mare to parse! The -ststr- can be represented phonetically as
			// S T S T R, S S T R, S T S tR, S tS T R, S S tR, S T R, S tS tR, S tR.
			if _, ok := getStringAt(alphas, currAbc, 0, "sts"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{s, t, r}); ok {
					// As in breaSTStroke,...
					// By looking at three phonemes we know tha the t being sounded
					// is the second t.
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"sts",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{s, tr}); ok {
					// As in breaSTStroke,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"sts",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{s, st}); ok {
					// As in breaSTSTroke,...
					// The first lexical t isn't sounded so swallow it here
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"st",
					}
					break
				}
			}
			// if str, ok := getStringAt(alphas, currAbc, 0, "sts"); ok {
			// 	if _, ok := phsAt(phons, currPh, []phoneme{s, t}, []phoneme{s, s}); !ok {
			// 		// The t doesn't appear to be sounded,
			// 		new = phonToAlphas{
			// 			[]phoneme{
			// 				s,
			// 			},
			// 			str,
			// 		}

			// 	}
			// }
			if stringAt(alphas, currAbc, 3, "sth") {
				if _, ok := phsAt(phons, currPh, []phoneme{s, th}, []phoneme{s, t}); !ok {
					// There is a silent th so swallow it now
					// As in iSTHmus,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"sth",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "st") {
				// Look out for a silent 't'. It typically follows an 's'.
				// if _, ok := phsAt(phons, currPh, []phoneme{s, ch}, []phoneme{s, t}, []phoneme{s, th}, []phoneme{s, tr}, []phoneme{s, ts}); !ok {
				if _, ok := phsAt(phons, currPh, []phoneme{s, ax}, []phoneme{s, axl}, []phoneme{s, axn}, []phoneme{s, k}, []phoneme{s, l}, []phoneme{s, m}, []phoneme{s, n}, []phoneme{s, p}, []phoneme{s, s}); ok {
					// As in caSTle, caSTle, muSTn't, waiSTcoat, neSTle, adjuSTment, cheSTnut, poSTpone, breaSTstroke,...
					// There is a silent t
					// As in thiSTle,... but not as in reStless,..
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"st",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 4, "sten") {
				// As in faSTEN,... The t here is typically silent so include it
				// here
				// But be careful, we don't want to swallow the t in sTencil...
				if _, ok := phsAt(phons, currPh, []phoneme{s, t}); !ok {
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"st",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "tz") {
				// As in walTZ,...
				new = phonToAlphas{
					[]phoneme{
						s,
					},
					"tz",
				}
				break
			}
			if stringAt(alphas, currAbc-1, 3, "nds") {
				if _, ok := phsAt(phons, currPh-1, []phoneme{n, s}); ok {
					// We've found a silent (or suppressed) 'd' as in wiNDSwept, wiNDSor...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"ds",
					}
					break
				}
			}
			if stringAt(alphas, currAbc-1, 3, "cts") {
				if _, ok := phsAt(phons, currPh-1, []phoneme{k, s}); ok {
					// We've found a silent (or suppressed) 't' as in reflecTs,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"ts",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ic") {
				// As in the posh pronunciation of medICine,...
				new = phonToAlphas{
					[]phoneme{
						s,
					},
					"ic",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "sc") {
				// As in SCene,...
				if _, ok := phsAt(phons, currPh, []phoneme{s, ch}, []phoneme{s, k}, []phoneme{s, kr}, []phoneme{s, kl}, []phoneme{s, ks}); !ok {
					// But not as in miSChief, SChool, SCrum, diSClose, molluSCs,...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"sc",
					}
				} else {
					// As in SCHism, aSCRibe, diSCLaimer...
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"s",
					}
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ps") {
				// As in PSychic,...
				new = phonToAlphas{
					[]phoneme{
						s,
					},
					"ps",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ss") {
				// As in paSS,...
				if _, ok := phsAt(phons, currPh, []phoneme{s, s}, []phoneme{s, sh}, []phoneme{s, st}); !ok {
					// But not as in miSSpell, miSSHapen, miSSTatement,... and many other mis- words (in which we leave the second
					// lexical 's' for the second phonetic 's')
					new = phonToAlphas{
						[]phoneme{
							s,
						},
						"ss",
					}
					break
				}
			}
			if st, ok := getStringAt(alphas, currAbc, 1, "c", "s", "z"); ok {
				// As in truCe, whiSt, glitZy,...
				new = phonToAlphas{
					[]phoneme{
						s,
					},
					st,
				}
				break
			}
		case sh:
			if s, ok := getStringAt(alphas, currAbc, 4, "cian", "sion", "tion"); ok {
				// As in electriCIAN, penSION, naTION,...
				if _, ok := phsAt(phons, currPh, []phoneme{sh, n}); ok {
					new = phonToAlphas{
						[]phoneme{
							sh, n,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "cesh", "sesh"); ok {
				// As in apprentiCESHip, horSESHoe,...
				new = phonToAlphas{
					[]phoneme{
						sh,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 5, "ssion"); ok {
				// As in percuSSION,...
				if _, ok := phsAt(phons, currPh, []phoneme{sh, n}); ok {
					new = phonToAlphas{
						[]phoneme{
							sh, n,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "sch", "ssi"); ok {
				// As in SCHnapps, discuSSIon,...
				new = phonToAlphas{
					[]phoneme{
						sh,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ci", "ti"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{sh, ih}, []phoneme{sh, iy}, []phoneme{sh, y}); ok {
					// Only swallow the c or t in this case because the following i is
					// sounded separately.
					// As in appreCiate, negoTiable, assoCiative,...
					new = phonToAlphas{
						[]phoneme{
							sh,
						},
						s[:1],
					}
					break
				}
			}
			// Catch threshold, where the h is pronounced...
			// But be careful. There are words like fishhook where the h is pronounced
			// and in these words we do want to swallow the lexical sh
			_, shh := getStringAt(alphas, currAbc, 0, "shh")
			if _, ok := getStringAt(alphas, currAbc, 0, "sh"); ok && !shh {
				if _, ok := phsAt(phons, currPh, []phoneme{sh, hh}); ok {
					new = phonToAlphas{
						[]phoneme{
							sh,
						},
						"s",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ch", "ci", "sc", "sh", "ss", "ti"); ok {
				// As in CHagrin, vivaCIous, conSCience, SHed, seSSion, raTIon,...
				new = phonToAlphas{
					[]phoneme{
						sh,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "c", "s", "x"); ok {
				// As in oCeanography, Sugar, anXious...
				new = phonToAlphas{
					[]phoneme{
						sh,
					},
					s,
				}
				break
			}
		case st:
			if str, ok := getStringAt(alphas, currAbc, 0, "stst"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{st, s}, []phoneme{st, st}); !ok {
					// We're good to swallow all of this now
					// As in breaSTSTroke,...
					new = phonToAlphas{
						[]phoneme{
							st,
						},
						str,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "sth"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{st, hh}); !ok {
					// The lexical h looks like it's silent so swallow it here.
					// TODO: Actually it would be nicer to swallow it later but fo rnow do it here.
					// As in poSTHumous,...
					new = phonToAlphas{
						[]phoneme{
							st,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "sced", "ssed", "tzed", "ced", "cet", "sed", "set", "sst", "st"); ok {
				// As in colaeSCED, obseSSED, walTZED, notiCED, peaCETime, abaSED, mouSETrap, croSSTalk, theoriST,...
				new = phonToAlphas{
					[]phoneme{
						st,
					},
					s,
				}
				break
			}
		case sts:
			if s, ok := getStringAt(alphas, currAbc, 0, "stes", "sts"); ok {
				// As in taSTES, geologiSTS,...
				new = phonToAlphas{
					[]phoneme{
						sts,
					},
					s,
				}
				break
			}
		case t:
			if stringAt(alphas, currAbc, 2, "ts") {
				if _, ok := phsAt(phons, currPh, []phoneme{t, s}, []phoneme{t, sh}, []phoneme{t, st}); !ok {
					// A special case, the lexical 's' isn't sounded, as in TSetse,...
					new = phonToAlphas{
						[]phoneme{
							t,
						},
						"ts",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ed") {
				// As in askED,...
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					"ed",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ort") {
				// As in comfORTable,...
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					"ort",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "th") {
				if _, ok := phsAt(phons, currPh, []phoneme{t, th}); ok {
					// A special case, as in eighTH,...
					// The phonetic t is not represented in the lexical spelling. I'm not
					// aware of any other examples.
					new = phonToAlphas{
						[]phoneme{
							t, th,
						},
						"th",
					}
					break
				}
				// Some words sound 'th' as t like THyme
				if _, ok := phsAt(phons, currPh, []phoneme{t, hh}); !ok {
					// As in discoTHeque, but not as in poTHole,...
					new = phonToAlphas{
						[]phoneme{
							t,
						},
						"th",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "ght") {
				// As in wriGHT,...
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					"ght",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "cht") {
				// As in yachT...
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					"cht",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "tt") {
				if _, ok := phsAt(phons, currPh, []phoneme{t, t}, []phoneme{t, tr}, []phoneme{t, thr}); ok {
					// As in posTTraumatic, posTTRaumatic, cuTTHRoat,...
					// Only swallow one t if the next phoneme is also a t (or is a tr)
					new = phonToAlphas{
						[]phoneme{
							t,
						},
						"t",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{t, th}); ok {
					// Leave the second t for the 'th' phoneme, as in cuTthroat,...
					new = phonToAlphas{
						[]phoneme{
							t,
						},
						"t",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "bt", "ct", "pt", "tt"); ok {
				// As in suBTle, indiCT, emPTy, boTTle,...
				// But only
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					s,
				}
				break
			}
			if st, ok := getStringAt(alphas, currAbc, 0, "zz", "z"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{t, s}); ok {
					// As in schmalZ, piZZa...
					new = phonToAlphas{
						p,
						st,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "d", "t"); ok {
				// As in learneD, lookeD,...
				new = phonToAlphas{
					[]phoneme{
						t,
					},
					s,
				}
				break
			}
		case th:
			if s, ok := getStringAt(alphas, currAbc, 0, "dth", "fth", "tth", "th"); ok {
				// As in thousanDTH, fiFTH, maTTHew, wealTH,...
				new = phonToAlphas{
					[]phoneme{
						th,
					},
					s,
				}
				break
			}
		case thr:
			if s, ok := getStringAt(alphas, currAbc, 0, "thr"); ok {
				// As in forTHRight,...
				new = phonToAlphas{
					[]phoneme{
						thr,
					},
					s,
				}
				break
			}
		case tr:
			if s, ok := getStringAt(alphas, currAbc, 0, "taur", "tar", "ter", "tor", "ttr", "ptr", "tr"); ok {
				// As in resTAURant, planeTARy, cemeTERy, hisTORy, aTTRibute, temPTRess, sTRike,...
				new = phonToAlphas{
					[]phoneme{
						tr,
					},
					s,
				}
				break
			}
		case ts:
			new = phonToAlphas{
				[]phoneme{
					ts,
				},
				"ts",
			}
		case uh:
			if s, ok := getStringAt(alphas, currAbc, 3, "our"); ok {
				// Catch words ending in -our (the axr phoneme is only used at
				//  the end of a word)
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axr}, []phoneme{uh, axr}); ok {
					// As in velOUR,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				// As in gOURmand,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axm}, []phoneme{uh, axm}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "oor"); ok {
				// As in boor,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axr}, []phoneme{uh, axr}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				// As in mOORland,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axl}, []phoneme{uh, axl}); ok {
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
			}
			// Catch an uh phoneme transitioning to r
			if s, ok := getStringAt(alphas, currAbc, 3, "eur", "ewer", "oor", "our"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, ax, r}, []phoneme{uh, ax, r}); ok {
					// As in plEUral, brEWEry, mOOrish, tOUring,...
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}
				// And just in case there's no ax
				if p, ok := phsAt(phons, currPh, []phoneme{uh, r}); ok {
					// As in plEUral, brEWEry, mOOrish, tOUring,...
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axn}, []phoneme{uh, axn}); ok {
					// As in pOORness, tOURniquet,...
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
				// As in brEWER, mOOR, tOUR,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, ax}, []phoneme{uh, ax}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "ure"); ok {
				// Check for a sounded phonetic 'r'
				if p, ok := phsAt(phons, currPh, []phoneme{uh, ax, r}); ok {
					// As in assUredly,... - so don't swallow the 'r' here
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
				// Check for 'axm'
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, axm}, []phoneme{uh, axm}); ok {
					// As in allUREMent,... -
					new = phonToAlphas{
						// Leave the axm for later
						p[:len(p)-1],
						s,
					}
					break
				}
				// As in allURE,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, w, ax}, []phoneme{uh, ax}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ar") {
				// As in onwARd,...
				new = phonToAlphas{
					[]phoneme{
						uh,
					},
					"ar",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "oul") {
				// As in wOULd,...
				new = phonToAlphas{
					[]phoneme{
						uh,
					},
					"oul",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ou") {
				p, _ := phsAt(phons, currPh, []phoneme{uh, w}, []phoneme{uh})
				// As in bedOUin, shOUld,...
				new = phonToAlphas{
					p,
					"ou",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "oo") {
				// As in wOOl,...
				new = phonToAlphas{
					[]phoneme{
						uh,
					},
					"oo",
				}
				break
			}
			// Transitioning from u to r is often rendered phonetically as uh ax r
			// so trying to capture that here
			if stringAt(alphas, currAbc, 1, "u") {
				// As in allUring,...
				if p, ok := phsAt(phons, currPh, []phoneme{uh, ax, r}); ok {
					new = phonToAlphas{
						// If the r is sounded process it separately
						p[:len(p)-1],
						"u",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "o", "u"); ok {
				// As in wOlf, fUll,...
				p, _ := phsAt(phons, currPh, []phoneme{uh, w}, []phoneme{uh})
				if _, ok := getStringAt(alphas, currAbc, 0, "ow", "uw"); ok {
					// It looks like the phonetic w is represented lexically
					//  so don't swallow the phonetic w here
					p = []phoneme{
						uh,
					}
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case uw:
			if stringAt(alphas, currAbc, 4, "ough") {
				// As in thrOUGH,...
				p, _ := phsAt(phons, currPh, []phoneme{uw, w}, []phoneme{uw})
				new = phonToAlphas{
					// []phoneme{
					// 	uw,
					// },
					p,
					"ough",
				}
				break
			}
			// Some borrowed French words to deal with first
			if stringAt(alphas, currAbc, 3, "hou") {
				// As in silHOUette,...
				// But check for linking w
				p, _ := phsAt(phons, currPh, []phoneme{uw, w}, []phoneme{uw})
				new = phonToAlphas{
					// []phoneme{
					// 	uw,
					// },
					p,
					"hou",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oul"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ax}, []phoneme{uw, w, axl}, []phoneme{uw, ax}, []phoneme{uw, axl}); ok {
					// As in cagOUle,...
					// Leave the ax for the lexical l
					new = phonToAlphas{
						// []phoneme{
						// 	uw,
						// },
						p[:len(p)-1],
						s[:2],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oup"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{uw, p}); !ok {
					// The p is not sounded so swallow it here
					// As in the French cOUP,...
					new = phonToAlphas{
						[]phoneme{
							uw,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "oeu") {
				// As in manOEUvre,...
				new = phonToAlphas{
					[]phoneme{
						uw,
					},
					"oeu",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ieu") {
				// As in lIEU,...
				new = phonToAlphas{
					[]phoneme{
						uw,
					},
					"ieu",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ous") {
				// As in rendezvOUS,...
				if _, ok := phsAt(phons, currPh, []phoneme{uw, s}, []phoneme{uw, st}, []phoneme{uw, z}); !ok {
					// The lexical 's' isn't sounded, so swallow it now
					new = phonToAlphas{
						[]phoneme{
							uw,
						},
						"ous",
					}
					break
				}
			}
			if _, ok := getStringAt(alphas, currAbc, 4, "uest", "uism"); ok {
				p, _ := phsAt(phons, currPh, []phoneme{uw, w}, []phoneme{uw})
				// As in trUest, trUism,...
				new = phonToAlphas{
					p,
					"u",
				}
				// Don't let this drop through to the next if because "ui" as in jUIce
				// will swallow the "i"
				break
			}
			if stringAt(alphas, currAbc, 2, "oo") {
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, oh}, []phoneme{uw, oh}, []phoneme{uw, w, ax}, []phoneme{uw, ax}, []phoneme{uw, w, axl}, []phoneme{uw, axl}); ok {
					// As in zOOlogy,...
					// Only swallow the first 'o'. The second 'o' belongs to the phonetic oh
					new = phonToAlphas{
						p[:len(p)-1],
						"o",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, axr}, []phoneme{uw, axr}, []phoneme{uw, w, iy}, []phoneme{uw, iy}); ok {
					// As in wOOer, gOOey,...
					new = phonToAlphas{
						p[:len(p)-1],
						"oo",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oe", "oo"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ih}, []phoneme{uw, ih}, []phoneme{uw, w, ing}, []phoneme{uw, ing}); ok {
					// As in canOEing, mOOing,...
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ax}, []phoneme{uw, ax}, []phoneme{uw, w, axr}, []phoneme{uw, axr}); ok {
					// As in evildOers, evildOer,...
					new = phonToAlphas{
						p[:len(p)-1],
						"o",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ue") {
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ax}, []phoneme{uw, ax}, []phoneme{uw, w, axl}, []phoneme{uw, axl}, []phoneme{uw, w, axn}, []phoneme{uw, axn}, []phoneme{uw, w, axr}, []phoneme{uw, axr}, []phoneme{uw, w, ih}, []phoneme{uw, ih}); ok {
					// As in grUelling, grUelling, constitUency, trUer, sUEt,...
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ing}, []phoneme{uw, ing}, []phoneme{uw, w, iy}, []phoneme{uw, iy}); ok {
					// As in glUEing, glUEy,...
					new = phonToAlphas{
						// []phoneme{
						// 	uw,
						// },
						p[:len(p)-1],
						"ue",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ew") {
				if _, ok := phsAt(phons, currPh, []phoneme{uw, w}); ok {
					// As in sEwerage,...,...
					new = phonToAlphas{
						[]phoneme{
							uw,
						},
						"e",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ew", "eu", "oe", "oo", "ou", "ue", "wo"); ok {
				// As in nEW, slEUth, shOE, fOOl, sOUp, blUE, tWOsome...
				if _, ok := phsAt(phons, currPh, []phoneme{uw, w, eh}, []phoneme{uw, eh}); !ok {
					// But not as in whOEver,...
					p, _ := phsAt(phons, currPh, []phoneme{uw, w}, []phoneme{uw})
					new = phonToAlphas{
						// []phoneme{
						// 	uw,
						// },
						p,
						s,
					}
					break
				}
			}
			// Handle fluid separately else it screws up other uw words
			if stringAt(alphas, currAbc, 2, "ui") {
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ing}, []phoneme{uw, ing}); ok {
					// Leave the lexical 'i' to be processed as part of the ing later
					// As in constrUing, and many others...
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
				// As in jUIce,...
				// But NOT as in flUId,...
				if p, ok := phsAt(phons, currPh, []phoneme{uw, w, ih}, []phoneme{uw, ih}, []phoneme{uw, w, ah}, []phoneme{uw, ah}); !ok {
					new = phonToAlphas{
						[]phoneme{
							uw,
						},
						"ui",
					}
					break
				} else {
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "o", "u"); ok {
				// As in whO, tUne...
				p, _ := phsAt(phons, currPh, []phoneme{uw, w}, []phoneme{uw})
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case uwm:
			if s, ok := getStringAt(alphas, currAbc, 0, "ombm", "oomm"); ok {
				// We have pronunciation variants with uwm and uwm, m phonemes
				p := []phoneme{
					uwm,
				}
				if _, ok := phsAt(phons, currPh, []phoneme{uwm, m}); ok {
					// Leave the phonetic m for procesing later
					// As in entOMBment, rOOMmate...
					new = phonToAlphas{
						p,
						s[:len(s)-1],
					}
				} else {
					// As in entOMBMent, rOOMMate...
					new = phonToAlphas{
						p,
						s,
					}
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ulme", "ewm", "eum", "oom", "omb", "oum", "om", "um"); ok {
				// As in levenshULME, crEWMan, rhEUMatism, rOOM, khartOUM, wOMB, whOMsoever, costUMe,...
				new = phonToAlphas{
					[]phoneme{
						uwm,
					},
					s,
				}
				break
			}
		case uwn:
			if s, ok := getStringAt(alphas, currAbc, 0, "uen", "ewn", "oon", "oun", "on", "un"); ok {
				// As in blUENess, strEWN, pontOON, wOUNded, cantON*, fortUNe,...
				// *This pronunciation looks suspicious. TODO: Check to see if
				// this should be removed from the dictionary
				new = phonToAlphas{
					[]phoneme{
						uwn,
					},
					s,
				}
				break
			}
		case v:
			if s, ok := getStringAt(alphas, currAbc, 2, "ph", "vv"); ok {
				// As in nePHew, saVVy,...
				new = phonToAlphas{
					[]phoneme{
						v,
					},
					s,
				}
				break
			}
			// Catch a possible silent l
			if stringAt(alphas, currAbc, 2, "lv") {
				// As in haLVe,...
				new = phonToAlphas{
					[]phoneme{
						v,
					},
					"lv",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "f", "v"); ok {
				new = phonToAlphas{
					[]phoneme{
						v,
					},
					s,
				}
				break
			}
		case w:
			if stringAt(alphas, currAbc, 2, "wh") {
				new = phonToAlphas{
					[]phoneme{
						w,
					},
					"wh",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "w", "u"); ok {
				// As in Wall, langUid...
				new = phonToAlphas{
					[]phoneme{
						w,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oirm"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{w, ay, y, axm}, []phoneme{w, ay, axm}); ok {
					// As in chOIRMaster,...
					// BUt don't swallow the lexical m and phonetic axm just yet
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oir"); ok {
				// Check for the borrowed French oir
				if p, ok := phsAt(phons, currPh, []phoneme{w, aa}); ok {
					if _, ok := phsAt(phons, currPh, []phoneme{w, aa, r}); ok {
						// We have an r phoneme, it's the last phoneme in the word
						// As in sOIRee,...
						if len(phons) == currPh+3 {
							p = append(p, r)
						} else {
							// It isn't the last phoneme so don't swallow the lexical r here
							// As in sOIree,...
							s = s[:len(s)-1]
						}
					}
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{w, ay, y, ax}, []phoneme{w, ay, ax}, []phoneme{w, ay, y, axr}, []phoneme{w, ay, axr}); ok {
					// This is probably choir, check to see if there's an r to swallow
					// As in chOIRs, choir,...
					if _, ok := phsAt(phons, currPh, []phoneme{w, ay, y, ax, r}, []phoneme{w, ay, ax, r}); ok {
						p = append(p, r)
					}
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oi"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{w, aa}); ok {
					// As in cOIffure,...
					new = phonToAlphas{
						[]phoneme{
							w, aa,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 1, "o") {
				// As in vOyeur, One,...
				if p, ok := phsAt(phons, currPh, []phoneme{w, aa}, []phoneme{w, ah}, []phoneme{w, oh}); ok {
					new = phonToAlphas{
						p,
						"o",
					}
				}
				break
			}
		case y:
			if stringAt(alphas, currAbc, 3, "aeo") {
				// As in palAEOntology,...
				new = phonToAlphas{
					[]phoneme{
						y,
					},
					"aeo",
				}
				break
			}
			// An early check for the word ewe.
			if alphas == "ewe" {
				// Note we can't check for the substring ewe else we'll break handling of words
				// like newest
				new = phonToAlphas{
					[]phoneme{
						y, uw,
					},
					"ewe",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "ean") {
				// As in cetacEAN,...
				new = phonToAlphas{
					[]phoneme{
						y,
					},
					"e",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eue"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uw, w}, []phoneme{y, uw}); ok {
					// As in quEUEing,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ule"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uwl, ax}, []phoneme{y, uwl, ih}); ok {
					// The e is probably sounded so don't swallow it here
					// As in amULet, amULet, ...
					new = phonToAlphas{
						p[:len(p)-1],
						s[:len(s)-1],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 6, "uel", "ule", "eul", "ul"); ok {
				// As in valUELess, capsULE, EULogy, reticULar...
				if p, ok := phsAt(phons, currPh, []phoneme{y, uwl}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eum", "um"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uwm}); ok {
					// As in pnEUMonia, hUMan,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eun", "ewn", "ugn", "un"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uwn}); ok {
					// As in EUNuch, hEWN, impUGN, mUNicipal,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			// The many sounds of 'ure'...
			if _, ok := getStringAt(alphas, currAbc, 3, "eur"); ok {
				// As in EURope,...
				// Catch a possible sounded phonetic 'r'
				if _, ok := phsAt(phons, currPh, []phoneme{y, ax, r}, []phoneme{y, uh, r}); ok {
					new = phonToAlphas{
						[]phoneme{
							y, ax,
						},
						"eu",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "eur", "ure", "ur"); ok {
				// As in EURope, tenUREd, cURious,
				// Catch a possible sounded phonetic 'r'
				if _, ok := phsAt(phons, currPh, []phoneme{y, ax, r}, []phoneme{y, uh, r}); ok {
					// As in failURE,...
					new = phonToAlphas{
						[]phoneme{
							y, ax,
						},
						"u",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{y, uh, ax, r}, []phoneme{y, uw, ax, r}, []phoneme{y, uh, ah, r}, []phoneme{y, uw, ah, r}); ok {
					new = phonToAlphas{
						// Leave the r to be processed separately
						p[:len(p)-1],
						"u",
					}
					break
				}
				// Okay, we've covered the cases with a phonetic 'r'
				if p, ok := phsAt(phons, currPh, []phoneme{y, ax}, []phoneme{y, axr}, []phoneme{y, er}); ok {
					// As in failURE, failURE,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{y, uh, axm}, []phoneme{y, uw, axm}); ok {
					// As in procUREment,...
					new = phonToAlphas{
						p[:len(p)-1],
						s,
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{y, uh, w, axr}, []phoneme{y, uh, axr}, []phoneme{y, uh, w, ax}, []phoneme{y, uh, ax}, []phoneme{y, uh, w, er}, []phoneme{y, uh, er}, []phoneme{y, uw, w, axr}, []phoneme{y, uw, axr}, []phoneme{y, uw, w, ax}, []phoneme{y, uw, ax}, []phoneme{y, uw, w, er}, []phoneme{y, uw, er}); ok {
					// As in cURE, liquEUR, pURE,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ut"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{y, uw, sh}, []phoneme{y, uw, ch}, []phoneme{y, uw, tr}, []phoneme{y, uw, t}); !ok {
					// Not as in restitUtion, fUture, nUtrition, tUtor,...
					if _, ok := phsAt(phons, currPh, []phoneme{y, uw}); ok {
						// As in debUT,...
						new = phonToAlphas{
							[]phoneme{
								y, uw,
							},
							s,
						}
						break
					}
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "eau", "eu"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uw, w}, []phoneme{y, uw}); ok {
					// As in bEAUty, tEUtonic,...
					new = phonToAlphas{
						p,
						s,
					}
				}
				break
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "eo"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{y, ax}); ok {
					// The 'o' appears to be represented phonetically so just take
					// the 'e' now and leave the 'o' for future processing
					// As in metEorological,...
					new = phonToAlphas{
						[]phoneme{
							y,
						},
						"e",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ui"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uw, w, ih}, []phoneme{y, uw, ih}, []phoneme{y, uw, w, ing}, []phoneme{y, uw, ing}); ok {
					// The i is sounded so leave it for later
					// As in continUING,...
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
				if _, ok := phsAt(phons, currPh, []phoneme{y, uw}); ok {
					if _, ok := phsAt(phons, currPh, []phoneme{y, uw, w, ih}, []phoneme{y, uw, ih}); !ok {
						// The lexical 'i' isn't sounded separately so swallow it all
						new = phonToAlphas{
							[]phoneme{
								y, uw,
							},
							s,
						}
						break
					} else {

						// Check for linking y
						p, _ := phsAt(phons, currPh, []phoneme{y, uw, w}, []phoneme{y, uw})
						// As in acUity,...
						new = phonToAlphas{
							p,
							"u",
						}
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ew"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{y, uw, w}, []phoneme{y, uh, w}); ok {
					// As in stEWard, dont swallow the w yet
					new = phonToAlphas{
						p[:len(p)-1],
						"e",
					}
					break
				} else {
					if p, ok := phsAt(phons, currPh, []phoneme{y, uw}, []phoneme{y, uh}); ok {
						new = phonToAlphas{
							p,
							s,
						}
						break
					}
				}
				if _, ok := phsAt(phons, currPh, []phoneme{y, uwl}); ok {
					// As in nEWLy,...
					new = phonToAlphas{
						// Leave the l for uwL
						[]phoneme{
							y,
						},
						s,
					}
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "i", "y"); ok {
				// As in millIon, onIon, You,...
				new = phonToAlphas{
					[]phoneme{
						y,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "j") {
				// As in the German Junker,...
				new = phonToAlphas{
					[]phoneme{
						y,
					},
					"j",
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "u") {
				if p, ok := phsAt(phons, currPh, []phoneme{y, ah}, []phoneme{y, uh, w}, []phoneme{y, uh}, []phoneme{y, uw, w}, []phoneme{y, uw}); ok {
					// As in articUlate, tUreen, fUture,... At least according to the CMU dictionary
					new = phonToAlphas{
						p,
						"u",
					}
				}
				break
			}
		case yuw:
			if _, ok := getStringAt(alphas, currAbc, 0, "ure"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{yuw, ax}, []phoneme{yuw, axm}, []phoneme{yuw, axr}); ok {
					new = phonToAlphas{
						[]phoneme{
							yuw,
						},
						"u",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eur"); ok {
				// As in liquEUR,...
				if p, ok := phsAt(phons, currPh, []phoneme{yuw, w, axr}, []phoneme{yuw, axr}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}

			}
			if _, ok := getStringAt(alphas, currAbc, 0, "ue"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{yuw, w, ax}, []phoneme{yuw, ax}, []phoneme{yuw, w, axl}, []phoneme{yuw, axl}, []phoneme{yuw, w, eh}, []phoneme{yuw, eh}, []phoneme{yuw, w, ih}, []phoneme{yuw, ih}); ok {
					// Looks like the lexical e is sounded
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "ui"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{yuw, w, ih}, []phoneme{yuw, ih}, []phoneme{yuw, w, ing}, []phoneme{yuw, ing}); ok {
					// Looks like the lexical i is sounded
					new = phonToAlphas{
						p[:len(p)-1],
						"u",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ut"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{yuw, sh}, []phoneme{yuw, ch}, []phoneme{yuw, tr}, []phoneme{yuw, t}); !ok {
					// Not as in restitUtion, fUture, nUtrition, tUtor,...
					// As in debUT,...
					new = phonToAlphas{
						[]phoneme{
							yuw,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "eau", "ieu", "iew", "eue", "yew", "you", "eu", "ew", "hu", "iu", "ue", "ui", "yu", "u"); ok {
				// As in bEAUtiful, adIEU, vIEW, quEUE, YEW, YOUth, EUcalyptus, anEW, postHUmous, jIUjitsu, argUE, sUIt, YUle, Use,...
				p, _ := phsAt(phons, currPh, []phoneme{yuw, w}, []phoneme{yuw})
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
		case z:
			if s, ok := getStringAt(alphas, currAbc, 0, "ds"); ok {
				// As in bonDS,...
				new = phonToAlphas{
					[]phoneme{
						z,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "spb") {
				// As in raSPBerry,...
				// Check for a silent 'p'
				if _, ok := phsAt(phons, currPh, []phoneme{z, b}); ok {
					// Looks like the 'p' is silent
					new = phonToAlphas{
						[]phoneme{
							z,
						},
						"sp",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "cz"); ok {
				// As in CZar,...
				new = phonToAlphas{
					[]phoneme{
						z,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "se", "ze"); ok {
				_, ok := phsAt(phons, currPh, []phoneme{z, d})
				if ok || currPh == len(phons)-1 {
					// As in raiSE, raiSEd, raZE, raZEd,...
					new = phonToAlphas{
						[]phoneme{
							z,
						},
						s,
					}
					break
				}
			}
			if str, ok := getStringAt(alphas, currAbc, 2, "ss", "zz"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{z, s}, []phoneme{z, st}); ok {
					// As in newSStand,...
					new = phonToAlphas{
						[]phoneme{
							z,
						},
						// Leave the second lexical 's for the phonetic 's'
						"s",
					}
					break
				}
				// As in sciSSors, fiZZed,...
				new = phonToAlphas{
					[]phoneme{
						z,
					},
					str,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "s", "x", "z"); ok {
				// As in many (but not all) plural word, for instance dogS, Xylem, zoo,...
				new = phonToAlphas{
					[]phoneme{
						z,
					},
					s,
				}
				break
			}
		case zh:
			if s, ok := getStringAt(alphas, currAbc, 4, "sion", "tion"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{zh, n}); ok {
					// As in fuSION, equaTION,...
					new = phonToAlphas{
						[]phoneme{
							zh, n,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "j", "s", "t", "z", "g"); ok {
				// As in beiJing, uSual, equaTion, aZure, beiGe,...
				new = phonToAlphas{
					[]phoneme{
						zh,
					},
					s,
				}
				break
			}
		case ax:
			if stringAt(alphas, currAbc, 2, "ve") {
				// As in should'VE,...
				new = phonToAlphas{
					[]phoneme{
						ax, v,
					},
					"ve",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "wer") {
				// As in ansWER,...
				s := "wer"
				// But only if the 'r' isn't sounded
				// As in ansWErable,...
				if _, ok := phsAt(phons, currPh, []phoneme{ax, r}); ok {
					s = "we"
				}
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "wai") {
				// As in coxsWAIn,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					"wai",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "anc") {
				if _, ok := phsAt(phons, currPh, []phoneme{ax, n, k}, []phoneme{ax, n, s}, []phoneme{ax, n, st}); !ok {
					// As in blANCmange,..
					// But not as in melANCholy, vagrANCy, balanCED, and many other words...
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						"anc",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "har", "her"); ok {
				// As in philHARmonic, shepHERd,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ha") {
				// As in FulHAm and many other place names
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					"ha",
				}
				break
			}
			if stringAt(alphas, currAbc, 3, "lel") {
				// As in candLELight,...
				if _, ok := phsAt(phons, currPh, []phoneme{ax, l, l}); !ok {
					// The second lexical l isn't represented phonetically so swallow it here
					new = phonToAlphas{
						[]phoneme{
							ax, l,
						},
						"lel",
					}
					break
				}
			}
			if p, ok := phsAt(phons, currPh, []phoneme{ax, l}, []phoneme{ax, m}, []phoneme{ax, n}); ok {
				// Trying to spot an inserted schwa as in bottLe, theisM, wasN't,...
				if stringAt(alphas, currAbc, 2, "le") {
					// Silent e?
					new = phonToAlphas{
						p,
						"le",
					}
					break
				}
				if stringAt(alphas, currAbc, 2, "ll") {
					// As in genteeLLy,...
					// Check for a phonetic l l
					s := "ll"
					if _, ok := phsAt(phons, currPh, []phoneme{ax, l, l}); ok {
						s = "l"
					}
					new = phonToAlphas{
						[]phoneme{
							ax, l,
						},
						s,
					}
					break
				}
				if s, ok := getStringAt(alphas, currAbc, 0, "l", "m", "n"); ok {
					// As in  maiL, theisM, isn't,...
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "eau", "eou", "iou"); ok {
				// As in burEAUcrat, outragEOUs, suspicIOUs,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "eur", "oar", "our", "ure"); ok {
				// As in amatEUR, cupbOARd, flavOUR, futURE...
				// But not if the r is sounded, for instance in armOURy, usUREr...
				if _, ok := phsAt(phons, currPh, []phoneme{ax, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ai") {
				if p, ok := phsAt(phons, currPh, []phoneme{ax, r, ih}, []phoneme{ax, ih}); ok {
					// As in contrAindication,...
					new = phonToAlphas{
						p[:len(p)-1],
						"a",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ao") {
				if p, ok := phsAt(phons, currPh, []phoneme{ax, r, ao}, []phoneme{ax, ao}); ok {
					// As in extrAordinary,...
					new = phonToAlphas{
						p[:len(p)-1],
						"a",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "ur") {
				// As in aubURn,...
				// If the r is sounded, we should only grab the 'u'
				if _, ok := phsAt(phons, currPh, []phoneme{ax, r}); ok {
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						"u",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 4, "ough") {
				// As in thorOUGH,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					"ough",
				}
				break
			}
			// Handling this separately as this is a mix of a silent e followed by a
			// vowel so I may handle this more generally at some point
			if stringAt(alphas, currAbc, 2, "ea") {
				// As in likEAble,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					"ea",
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "yr") {
				if _, ok := phsAt(phons, currPh, []phoneme{ax, r}); ok {
					// The lexical 'r' is sounded so leave it for later
					// As in labYRinth,...
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						"y",
					}
				} else {
					// As in zephYR,...
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						"yr",
					}
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ei", "ia", "ie", "io", "iu", "oi", "ou", "ua", "ui"); ok {
				// As in forEIgn, russIA, conscIEnce, percussIOn, tortOIse, belgIUm, luxuriOUs, usUAlly, biscUIt,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
			// Treat 're' separately.
			if stringAt(alphas, currAbc, 2, "re") {
				// As in theatREs,...
				s := "re"
				p := []phoneme{ax}
				if p1, ok := phsAt(phons, currPh, []phoneme{ax, r}); ok {
					// It's hard to distinguish severest and acreage, both have
					//  ax r ih but need to be mapped differently
					p = p1
					// As in acREage,...
					if stringAt(alphas, currAbc, 4, "rest") {
						// But as in seveRest, and many others...
						// We only want to swallow the lexical r here
						s = "r"
					}
				}
				new = phonToAlphas{
					p,
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 4, "erwr") {
				// As in undERWRite,...
				// Don't swallow the r phoneme here! It belong with the
				// lexical wr
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					"er",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ar", "er", "ir", "or", "ur"); ok {
				// As in wizARd, pERcussion weIRd, tractOR, pURveyor,...
				// Need to be careful here. I don't want to consume the r if there's
				// an r phoneme in the phonetic spelling, for instance as in
				// d(d)o(ao)c(k)u(y uh)m(m)e(eh)n(n)t(t)a(ax)r(r)y(iy)
				if _, ok := phsAt(phons, currPh, []phoneme{ax, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						s,
					}
					break
				}
			}
			// Some words have an -ah ending with the 'h' silent
			if stringAt(alphas, currAbc, 2, "ah") {
				if _, ok := phsAt(phons, currPh, []phoneme{ax, hh}); !ok {
					// The 'h' isn't sounded
					// As in purdAH,...
					new = phonToAlphas{
						[]phoneme{
							ax,
						},
						"ah",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e", "i", "o", "u", "y"); ok {
				// As in - pretty much anything you can think of and in particular, as in pYjamas,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "rr", "r"); ok {
				if p, ok := phsAt(phons, currPh, []phoneme{ax, r}); ok {
					new = phonToAlphas{
						p,
						s,
					}
					break
				}
				// As in houR,...
				new = phonToAlphas{
					[]phoneme{
						ax,
					},
					s,
				}
				break
			}
		case axr:
			if s, ok := getStringAt(alphas, currAbc, 0, "ough", "eur", "our", "ure", "wer", "ar", "er", "ha", "ia", "ir", "or", "re", "ur", "yr", "a", "e", "o", "r"); ok {
				// As in thorOUGH, amatEUR, favOUR, sutURE, ansWER, liAR, harriER, piranHA, nostalgIA, fIR, tailOR, metRE, lemUR, zephYR, troikA, timbrE, ontO*, cobbleR,...
				// *onto is a bit suspect for a standalone word pronunciation. In the case of cobbler the test for silent e might
				// swallow the lexical e so that all we're left with is the lexical r
				new = phonToAlphas{
					[]phoneme{
						axr,
					},
					s,
				}
				break
			}
		case oh:
			if s, ok := getStringAt(alphas, currAbc, 3, "eau", "au", "ho", "oh", "ou"); ok {
				// As in burEAUcracy, sAUsage, HOnest, jOHn, cOUgh,...
				new = phonToAlphas{
					[]phoneme{
						oh,
					},
					s,
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 1, "ow", "a", "e", "o"); ok {
				// As in knOWledge, whAt, Encore, cOttage,...
				new = phonToAlphas{
					[]phoneme{
						oh,
					},
					s,
				}
				break
			}
		case ehr:
			if stringAt(alphas, currAbc, 4, "heir") {
				// As in HEIR, but
				if _, ok := phsAt(phons, currPh, []phoneme{ehr, r}); !ok {
					// As in HEIR,...
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						"heir",
					}
				} else {
					// Don't swallow the r up, as in HEIress,...
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						"hei",
					}
				}
			}
			if strings.HasPrefix(alphas, "there") || strings.HasPrefix(alphas, "where") {
				// I'm being quite specific here. I don't want to break words like etHEREal.
				// As in words like THEREafter, WHEREof...
				if p, ok := phsAt(phons, currPh, []phoneme{ehr, r, eh}); ok {
					// We don't want to swallow the second lexical e in words like whERever,...
					new = phonToAlphas{
						p[:len(p)-1],
						"er",
					}
					break
				}
				if p, ok := phsAt(phons, currPh, []phoneme{ehr, r}); ok {
					new = phonToAlphas{
						p,
						"ere",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 4, "ayor"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ehr, r}); ok {
					// As in mAYOral,...
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						s[:len(s)-1],
					}
				} else {
					// As in mAYOR,...
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						s,
					}
				}
				break
			}
			if _, ok := getStringAt(alphas, currAbc, 0, "ar"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ehr, r}); ok {
					// The r is sounded so don't grab it here
					// As in phAraoh,...
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						"a",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "aire", "ayer", "air", "are", "ear", "eir", "ere", "ar", "er"); ok {
				// As in millionAIRE, prAYER, eclAIR, bewARE, wEAR, thEIR, whERE, scARce, concERto,...
				// But only if the r isn't sounded
				if _, ok := phsAt(phons, currPh, []phoneme{ehr, r}); !ok {
					new = phonToAlphas{
						[]phoneme{
							ehr,
						},
						s,
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ae", "ai", "ea"); ok {
				// As in AErial, AIring, wEAr (for alternative W EHR R phonetic spelling)...
				new = phonToAlphas{
					[]phoneme{
						ehr,
					},
					s,
				}
				break
			}

			if s, ok := getStringAt(alphas, currAbc, 1, "a", "e"); ok {
				// As in aquArium, whEreas,...
				new = phonToAlphas{
					[]phoneme{
						ehr,
					},
					s,
				}
				break
			}
		case axl:
			if stringAt(alphas, currAbc, 3, "lel") {
				if _, ok := phsAt(phons, currPh, []phoneme{axl, l}); !ok {
					// As in candLELight,...
					new = phonToAlphas{
						[]phoneme{
							axl,
						},
						"lel",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "le") {
				if _, ok := phsAt(phons, currPh, []phoneme{axl, ax}, []phoneme{axl, eh}, []phoneme{axl, ih}); ok {
					// The lexical 'e' is sounded, so leave it for later
					// As in tirelEss, coaLEsce, wirelEss,...
					new = phonToAlphas{
						[]phoneme{
							axl,
						},
						"l",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "rl") {
				// As in douRLy,... and many others in which the r isn't sounded
				new = phonToAlphas{
					[]phoneme{
						axl,
					},
					"rl",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ou'll", "oughl", "iall", "ourl", "uall", "urel", "wale", "ael", "all", "arl", "aul", "ell", "erl", "ial", "ill", "oll", "orl", "rel", "ull", "al", "el", "il", "le", "ol", "ul", "yl"); ok {
				// As in yOU'LL, thorOUGHLy, partIALLy, odOURLess, usUALLy, leisURELy, gunWALE, michAEL, nationALLy, regulARLy, epAULettes, modELLing, eastERLy, specIAL, councILLor, pOLLute, fORLorn, seveRELy, awfULLy, typicAL, caramEL, civIL, articLE, viOLate, awfUL, sibYL,...
				new = phonToAlphas{
					[]phoneme{
						axl,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 2, "ll") {
				// As in genteeLLy,...
				// Check for a phonetic l l
				s := "ll"
				if _, ok := phsAt(phons, currPh, []phoneme{axl, l}); ok {
					s = "l"
				}
				new = phonToAlphas{
					[]phoneme{
						axl,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "l") {
				// As in maiL,...
				new = phonToAlphas{
					[]phoneme{
						axl,
					},
					"l",
				}
				break
			}
		case axm:
			// Catch a possible double phonetic m first
			if stringAt(alphas, currAbc, 3, "omm") {
				if _, ok := phsAt(phons, currPh, []phoneme{axm, m}); ok {
					// As in bottOMmost,...
					new = phonToAlphas{
						[]phoneme{
							axm,
						},
						// Leave the second m until later
						"om",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 3, "emn", "umn"); ok {
				// As in solEMNly, colUMN,...
				if _, ok := phsAt(phons, currPh, []phoneme{axm, n}); !ok {
					// But not as in solEMNity, colUMNist,...
					new = phonToAlphas{
						[]phoneme{
							axm,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 3, "ham") {
				// As in many place names like wrexHAM, ...
				new = phonToAlphas{
					[]phoneme{
						axm,
					},
					"ham",
				}
				break
			}
			if stringAt(alphas, currAbc, 4, "harm") {
				// As in philHARMonic,...
				new = phonToAlphas{
					[]phoneme{
						axm,
					},
					"harm",
				}
				break
			}
			if s, ok := getStringAt(alphas, currAbc, 2, "ancm", "urem", "amm", "arm", "erm", "iam", "irm", "ium", "olm", "omm", "orm", "rem", "umm", "urm", "am", "em", "im", "om", "um"); ok {
				// As in bLANCMange, measUREMent, grAMMatical, philhARMonic, vERMillion, parlIAMent, affIRMation, belgIUM, malcOLM, cOMMercial, infORMation, procuREMent, consUMMation, sURMise, fAMiliar, acadEMy, anIMal (not sure this should
				// be a schwa though), incOMe, vacuUM...
				new = phonToAlphas{
					[]phoneme{
						axm,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "m") {
				// As in prisM,...
				new = phonToAlphas{
					[]phoneme{
						axm,
					},
					"m",
				}
				break
			}
		case axn:
			if s, ok := getStringAt(alphas, currAbc, 0, "eignn", "enn", "ann", "onn", "ornn"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{axn, n}); ok {
					// As in forEIGNness, unevENness, humaNness, commONness, stubbORNness,...
					// The second lexical 'n' is sounded so dont swallow it here
					new = phonToAlphas{
						[]phoneme{
							axn,
						},
						s[:len(s)-1],
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "oughn", "eignn", "eign", "erin", "ionn", "ornn", "ourn", "wain", "ain", "ann", "ean", "enn", "ien", "eon", "ern", "ian", "ign", "ion", "oen", "oln", "omp", "onn",
				"orn", "ren", "urn", "an", "en", "in", "on", "un", "n"); ok {
				// As in thorOUGHNess, forEIGNNess, forEIGN, vetERINary, legIONNaire, stubbORNNess, sojOURN, coxsWAIN, certAIN, ANNexe, pagEANt, unevENNess, conscIENce, burgEON, hibERNate, clinicIAN, ensIGN, equatION, rOENtgen, lincOLN, cOMPtroller, cONNection,
				// holbORN, meagRENess, aubURN, laymAN, conferENce, origINal, sextON, volUNteer, shouldN't,...
				new = phonToAlphas{
					[]phoneme{
						axn,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "n") {
				// This is a syllabic consonant, as in ?
				new = phonToAlphas{
					[]phoneme{
						axn,
					},
					"n",
				}
			}
		case ks:
			if s, ok := getStringAt(alphas, currAbc, 2, "xe"); ok {
				// Catch things like aXE, but not eXercise...
				if nextPhIsConsonant(phons, currPh) {
					new = phonToAlphas{
						[]phoneme{
							ks,
						},
						s,
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "xc") {
				if _, ok := phsAt(phons, currPh, []phoneme{ks, ch}, []phoneme{ks, k}, []phoneme{ks, kl}, []phoneme{ks, kr}); !ok {
					// As in eXCHange, eXCise but not as in eXCoriate, eXCLude, eXCRete, ...
					new = phonToAlphas{
						[]phoneme{
							ks,
						},
						"xc",
					}
					break
				}
			}
			if stringAt(alphas, currAbc, 2, "xs") {
				if _, ok := phsAt(phons, currPh, []phoneme{ks, s}); !ok {
					// The phonetic 's' isn't sounded so grab it now
					// As in coXSwain,...
					new = phonToAlphas{
						[]phoneme{
							ks,
						},
						"xs",
					}
					break
				}
			}
			// Trying to trap the silent h that can follow ex-
			if stringAt(alphas, currAbc, 2, "xh") {
				// As in eXHibition,...
				if _, ok := phsAt(phons, currPh, []phoneme{ks, hh}); !ok {
					// Okay, the h is silent so include it now
					new = phonToAlphas{
						[]phoneme{
							ks,
						},
						"xh",
					}
					break
				}
			}
			//The 't' following x is sometimes not pronounced so catch it here
			if _, ok := getStringAt(alphas, currAbc, 0, "xtb"); ok {
				if _, ok := phsAt(phons, currPh, []phoneme{ks, b}); ok {
					// The 't' is silent so swallow it now
					// As in teXTbook,... (I think this and its plural are the only examples)
					new = phonToAlphas{
						[]phoneme{
							ks,
						},
						"xt",
					}
					break
				}
			}
			if s, ok := getStringAt(alphas, currAbc, 0, "ches", "ques", "chs", "cks", "cts", "kes", "khs", "lks", "cc", "cs", "cz", "ks"); ok {
				// As in aCHES, antiQUES, daCHShund, triCKSter, refleCTS, spoKESperson, sheiKHS, waLKS, aCCede, froliCSome, eczema, pranKSter,...
				new = phonToAlphas{
					[]phoneme{
						ks,
					},
					s,
				}
				break
			}
			// As in eXit,...
			new = phonToAlphas{
				[]phoneme{
					ks,
				},
				"x",
			}
		case kw:
			if s, ok := getStringAt(alphas, currAbc, 2, "cqu", "qu"); ok {
				// As in aCQUaint, QUiet,...
				new = phonToAlphas{
					[]phoneme{
						kw,
					},
					s,
				}
				break
			}
		case dz:
			if s, ok := getStringAt(alphas, currAbc, 2, "des", "d's", "ds"); ok {
				// As in blonDES, world's, trenDS,...
				new = phonToAlphas{
					[]phoneme{
						dz,
					},
					s,
				}
				break
			}
		case uwl:
			if s, ok := getStringAt(alphas, currAbc, 3, "o'll", "oel", "ool", "oul", "ou'll", "uel", "ule", "ul"); ok {
				// As in whO'LL, shOELace, schOOL, ampOULe, yOU'LL, clUELess, rULE, unrULy...
				new = phonToAlphas{
					[]phoneme{
						uwl,
					},
					s,
				}
				break
			}
			if stringAt(alphas, currAbc, 1, "l") {
				new = phonToAlphas{
					[]phoneme{
						uwl,
					},
					"l",
				}
				break
			}
		default:
			break
		}
		if new.equal(phonToAlphas{}) {
			return fail(phon, alphas)
		}
		if isSilentE(alphas, currAbc, phons, currPh, new) {
			new.alphas += "e"
		}
		if punctuationSkipped != "" {
			new.alphas = punctuationSkipped + new.alphas
			currAbc -= len(punctuationSkipped)
			punctuationSkipped = ""
		}
		// Catch any trailing characters not yet mapped to phonemes.
		// TODO: This is a bit crude but will do for now.
		new.alphas += trailingSilentAlphas(alphas, currAbc, phons, currPh, new)
		currPh += len(new.phons)
		currAbc += len(new.alphas)
		ret = append(ret, new)
	}
	if currAbc != len(alphas) || currPh != len(phons) {
		// This is also a failure. Since we loop on phonemes it's most likely
		// that currAbc != len(alphas)
		return fail(phons[len(phons)-1], alphas)
	}
	return ret, nil
}
