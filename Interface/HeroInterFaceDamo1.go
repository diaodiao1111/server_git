package Interface

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// 实现Interface接口

func (hs HeroSlice) Len() int {

	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {

	return hs[i].Age < hs[j].Age

}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}
