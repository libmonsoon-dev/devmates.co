package solution

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func empty(str string) bool {
	return len(str) == 0
}

func include(items []string, item string) bool {
	for i := range items {
		if item == items[i] {
			return true
		}
	}

	return false
}

func Solution(beginWord, endWord string, wordList []string) int {
	if empty(endWord) || empty(beginWord) || !include(wordList, endWord) {
		fmt.Println(empty(endWord), empty(beginWord), !include(wordList, endWord))
		return 0
	}

	combos := make(map[string][]string)

	for i := 0; i < len(wordList); i++ {
		word := wordList[i]

		for j := 0; j < len(beginWord); j++ {
			key := word[:j] + "*" + word[j+1:]

			combos[key] = append(combos[key], word)
		}
	}

	queue := newQueue()

	queue.Push(queueItem{
		beginWord,
		1,
	})

	//visitedKey := beginWord
	visited := make(map[string]struct{})

	for !queue.q.Empty() {
		item, ok := queue.Pop()
		if !ok {
			panic(fmt.Errorf("ok=false, item=%v ", item))
		}

		for i := 0; i < len(beginWord); i++ {

			combosKey := item.word[:i] + "*" + item.word[i+1:]

			possibleWords := combos[combosKey]

			for j := 0; j < len(possibleWords); j++ {
				word := possibleWords[j]

				if word == endWord {
					return item.level + 1
				}

				if _, ok := visited[word]; !ok {
					visited[word] = struct{}{}
					queue.Push(queueItem{word, item.level + 1})

				}
			}
			delete(combos, combosKey)
		}
	}

	return 0
}

type queueItem struct {
	word  string
	level int
}

type queue struct {
	q *deque.Deque
}

func (q *queue) Push(item queueItem) {
	q.q.PushRight(item)
}

func (q *queue) Pop() (item queueItem, ok bool) {
	if ok = !q.q.Empty(); !ok {
		return
	}

	item, ok = q.q.PopLeft().(queueItem)
	return
}

func newQueue() *queue {
	return &queue{deque.New()}
}
