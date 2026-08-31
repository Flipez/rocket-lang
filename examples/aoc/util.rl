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
        a.append(true)
      elif (type == "STRING")
        a.append("")
      elif (type == "INTEGER")
        a.append(0)
      elif (type == "FLOAT")
        a.append(0.0)
      elif (type == "ARRAY")
        a.append([])
      elif (type == "HASH")
        a.append({})
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
