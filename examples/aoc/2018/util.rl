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
        a.yoink(true)
      else if (type == "STRING")
        a.yoink("")
      else if (type == "INTEGER")
        a.yoink(0)
      else if (type == "FLOAT")
        a.yoink(0.0)
      else if (type == "ARRAY")
        a.yoink([])
      else if (type == "HASH")
        a.yoink({})
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
      val = item.plz_s()
    end
    a = a.replace("{"+idx.plz_s()+"}", val)
  end
  return a
end
