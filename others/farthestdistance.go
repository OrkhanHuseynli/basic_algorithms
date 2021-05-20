package others

//func findSuitableBlock(blocks []map[string]bool) map[string]bool {
//	var selected  map[string]bool
//	if  len(blocks) <2 {
//		return blocks[0]
//	}
//	prev := blocks[0]
//	current := blocks[1]
//	next := blocks[2]
//	for i:=1; i< len(blocks)-1; i++ {
//		if isSuitable(i, blocks) {
//			selected = current
//		}
//		prev = current
//		current = blocks[i]
//		next = blocks[i+1]
//	}
//}