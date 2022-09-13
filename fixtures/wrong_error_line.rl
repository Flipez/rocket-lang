def isMatrix(m)
  if (m.type() != "ARRAY")
    //raise(1, "unexpected type, got=" + m.type() + " want=ARRAY")
    return false
  end

  itemTypes = []
  foreach idx, item in m
    itemTypes = itemTypes + item.type()
  end
  if (itemTypes.uniq().size() > 1)
    //raise(1, "unconsistent item types, got=" + itemTypes.to_json() + " want=ARRAY")
    return false
  end

  itemSizes = []
  foreach idx, item in m
    itemSizes = itemSizes + item.size()
  end
  if (itemSizes.uniq().size() > 1 )
    //raise(1, "unconsistent item sizes: " + itemSizes.to_json())
    return false
  end

  subItemTypes = []
  foreach rIdx, row in m
    foreach cIdx, col in row 
      subItemTypes = subItemTypes + col.type()
    end
  end
  if ((subItemTypes.uniq().size() > 1) || (subItemTypes.uniq().first().type() != "INTEGER") || (subItemTypes.uniq().first().type() != "FLOAT"))
    //raise(1, "unconsistent subitem types, got=" + subItemTypes.to_json() + " want=INTEGER|FLOAT")
    return false
  end

  return true
end

def dim(m)
  return [m.size(), m.first().size()]
end

def mul(a, b)
  // Check whether both arguments are matrices
  if (!isMatrix(a))
    raise(1, "first argument needs to be a valid matrix")
  end
  if (!isMatrix(b))
    raise(1, "first argument needs to be a valid matrix")
  end

  aDim, bDim = dim(a), dim(b)

  puts(aDim, bDim)
end

mul([[1,2],[3,4]], [[4,3],[2,1]])
