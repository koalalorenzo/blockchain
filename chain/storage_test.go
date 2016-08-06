package chain

import "testing"

func TestBlockChainSorting(t *testing.T) {
	firstBlockHex := "313436393536333933393330393231323238322c353537373030363739313934373737393431302c67656e657369732c333533343337333533363335333233303334363133373335333636333332333033333332333333363332333033333332333333323333363133333331333333323333363133333331333333393332333033343333333433353335333333353334333233303333333233333330333333313333333632633335333433363338333633393337333333323330333633393337333333323330333733343336333833363335333233303337333533363635333733333336333933363337333636353336333533363334333233303336333733363335333636353336333533373333333633393337333333323330333633343336333133373334333633313263"
	secondBlockHex := "313436393536333939393331343136373032372c383637343636353232333038323135333535312c646439326639643132336564393635313533343531346637323463373062323964376366343136643536653661313233383134613862353230316561303531372c333533343337333533363335333233303334363133373335333636333332333033333332333333363332333033333332333333323333363133333331333333333333363133333331333333393332333033343333333433353335333333353334333233303333333233333330333333313333333632633335333433363338333633393337333333323330333633393337333333323330333733343336333833363335333233303337333333363335333633333336363633363635333633343332333033363332333636333336363633363333333636323263"

	var blocks []Block

	blocks = append(blocks, BlockFromHex(secondBlockHex))
	blocks = append(blocks, BlockFromHex(firstBlockHex))

	blocks = SortChain(blocks)

	if blocks[0].PreviousHash != "genesis" {
		t.Error("Sorting process for genesis block failed")
	}

	if blocks[1].PreviousHash == "genesis" {
		t.Error("Sorting process failed")
	}
}
