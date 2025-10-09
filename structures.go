package task

type Move struct {
	x int
	y int
}

type Phone struct {
	n, m  int
	field [][]int
}

func (p *Phone) setNum(val, i, j int) {
	p.field[i][j] = val
}

func (p *Phone) getNum(i, j int) int {
	return p.field[i][j]
}

func (p *Phone) getSize() (int, int) {
	return len(p.field), len(p.field[0])
}

func (p *Phone) checkFiled(i, j int) bool {
	if i < p.n && i >= 0 && j < p.m && j >= 0 {
		return true
	}
	return false
}

func (p *Phone) create(N, M int) {
	p.n, p.m = N, M
	p.field = make([][]int, N)
	for idx := range N {
		p.field[idx] = make([]int, M)
	}
}

type Visit struct {
	n, m  int
	field [][]bool
}

func (v *Visit) set(i, j int) {
	v.field[i][j] = true
}

func (v *Visit) get(i, j int) bool {
	return v.field[i][j]
}

func (p *Visit) create(N, M int) {
	p.n, p.m = N, M
	p.field = make([][]bool, N)
	for idx := range N {
		p.field[idx] = make([]bool, M)
	}
}

func (p *Visit) copy() Visit {
	pCopy := Visit{}
	pCopy.create(p.n, p.m)
	for idx := range p.n {
		for jdx := range p.m {
			pCopy.field[idx][jdx] = p.field[idx][jdx]
		}
	}
	return pCopy
}
