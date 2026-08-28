import "../fixtures/module" as Inner

export def SumViaInner(a, b)
    return Inner.Sum(a, b)
end
