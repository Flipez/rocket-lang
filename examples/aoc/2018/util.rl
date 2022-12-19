def Swap(a, i, j)
  buf = a[i]
  a[i] = a[j]
  a[j] = buf
end

def Make(type, size)
  a = []
  if (size > 0)
    foreach i in size
      if (type == "BOOLEAN")
        a.push(true)
      else if (type == "STRING")
        a.push("")
      else if (type == "INTEGER")
        a.push(0)
      else if (type == "FLOAT")
        a.push(0.0)
      else if (type == "ARRAY")
        a.push([])
      else if (type == "HASH")
        a.push({})
      end
    end
  end
  return a
end

def Contains(a, i)
  foreach b in a
    if (b == i)
      return true
    end
  end
  return false
end

def Format(a, b)
  foreach idx, item in b
    val = ""
    if (item.type() == "STRING")
      val = item
    else
      val = item.to_s()
    end
    a = a.replace("{"+idx.to_s()+"}", val)
  end
  return a
end
