package golang

type UnionFind struct{
    parents []int
}

func NewUnionFind(n int) UnionFind{
    parents := make([]int,n)
    
    for i:=0;i<n;i++{
        parents[i] = i
    }
    
    return UnionFind{
        parents : parents,
    }
}

func (uf *UnionFind) Find(i int) int{
    for i != uf.parents[i]{
        uf.parents[i] = uf.parents[uf.parents[i]]
        i = uf.parents[i]    
    }
    
    return uf.parents[i]
}

func (uf *UnionFind) Union(i int,j int) bool{
    parentP := uf.Find(i)
    parentQ := uf.Find(j)
    
    if parentP == parentQ{
        return false
    }
    
    if parentQ < parentP{
        uf.parents[parentP] = parentQ
        return true
    }
    
    uf.parents[parentQ] = parentP
    
    return true
}