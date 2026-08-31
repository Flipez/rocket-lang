export def Swap(a, i, j)
  buf = a[i]
  a[i] = a[j]
  a[j] = buf
end

export def Make(type, size)
  a = []
  if (size > 0)
    foreach i in size
      if (type == "BOOLEAN")
        a.push(true)
      elif (type == "STRING")
        a.push("")
      elif (type == "INTEGER")
        a.push(0)
      elif (type == "FLOAT")
        a.push(0.0)
      elif (type == "ARRAY")
        a.push([])
      elif (type == "HASH")
        a.push({})
      end
    end
  end
  return a
end

export def Contains(a, i)
  foreach b in a
    if (b == i)
      return true
    end
  end
  return false
end

export def Format(a, b)
  foreach idx, item in b
    val = ""
    if (item.type() == "STRING")
      val = item
    else
      val = item.to_string()
    end
    a = a.replace("{"+idx.to_string()+"}", val)
  end
  return a
end
