package main

import (
	"C"
)
import (
	"errors"
	"fmt"

	"github.com/colinarticulate/pron"
)

//export MockPron
func MockPron(audiofile string, word string, outputfolder string, dictfile string, phdictfile string, featparams string, hmm string) string {
	// results, err := doPronounce("", audiofile, word, "", "", "")
	fmt.Println("Hello from Go!")
	fmt.Println(audiofile)
	fmt.Println(word)
	fmt.Println(outputfolder)
	fmt.Println(dictfile)
	fmt.Println(phdictfile)
	fmt.Println(featparams)
	fmt.Println(hmm)

	response := `{\\
    \"word\": \"climbed\",\\
    \"results\": [\\
        {\\
            \"letters\": \"cl\",\\
            \"phonemes\": \"kl\",\\
            \"verdict\": \"good\"\\
        },\\
        {\\
            \"letters\": \"i\",\\
            \"phonemes\": \"ɑɪ\",\\
            \"verdict\": \"good\"\\
        },\\
        {\\
            \"letters\": \"mb\",\\
            \"phonemes\": \"m\",\\
            \"verdict\": \"good\"\\
        },\\
        {\\
            \"letters\": \"ed\",\\
            \"phonemes\": \"d\",\\
            \"verdict\": \"good\"\\
        }\\
    ],\\
    \"percent_move\": 100,\\
    \"err\": null\\
}`
	return response //string(pron.ToJSON(results, err))
}

//export Pronounce
func Pronounce(outfolder string, audiofile string, word string, dictfile string, phdictfile string, featparams string, hmm string) (ret []pron.LettersVerdict, err error) {
	defer func() {
		if r := recover(); r != nil {
			ret = []pron.LettersVerdict{}
			err = errors.New("pron.Pronounce panicked!")
		}
	}()
	ret, err = pron.Pronounce(
		outfolder,
		audiofile,
		word,
		dictfile,
		phdictfile,
		featparams,
		hmm,
	)
	return ret, err
}

func main() {}
