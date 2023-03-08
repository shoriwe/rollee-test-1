package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	accumulator := initialValue
	for _, value := range l.Values {
		accumulator = f(accumulator, value)
	}
	return map[ID]int{l.ID: accumulator}
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	result := map[ID]int{}
	for list := range ch {
		if _, found := result[list.ID]; !found {
			result[list.ID] = initialValue
		}
		for _, value := range list.Values {
			result[list.ID] = f(result[list.ID], value)
		}
	}
	return result
}

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[ID]int {
	var (
		result      = map[ID]int{}
		resultMutex sync.Mutex // Why mutex? We want to write to result as soon as possible, not when workers complete
		wg          sync.WaitGroup
	)
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(channel chan List) {
			defer wg.Done()
			localResult := FoldChan(initialValue, f, channel) // This line should always be computed in parallel
			resultMutex.Lock()
			defer resultMutex.Unlock()
			for id, value := range localResult {
				if current, found := result[id]; found {
					result[id] = f(current, value)
				} else {
					result[id] = value
				}
			}
		}(ch)
	}
	wg.Wait()
	return result
}
