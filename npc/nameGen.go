package npc

type MarkovChain struct {
	attempts         int
	n_grams          map[string]string
	vowels           []string
	accepted_bigrams []string
}

// TODO(wholesomeow): Use RogueBasin link to create more advanced Markov Chain
// LINK: http://www.roguebasin.com/index.php?title=Names_from_a_high_order_Markov_Process_and_a_simplified_Katz_back-off_scheme
func buildNGram(mc MarkovChain) (map[string]string, error) {
	//Get data from some place here, if no data then error
	data := 55

	for i := 0; i < data; i++ {
		if i == data-1 {
			break
		} else if data == 25 {
			continue
		}
	}
	return mc.n_grams, nil
}
