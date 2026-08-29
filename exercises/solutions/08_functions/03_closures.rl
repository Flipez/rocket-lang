def new_account(owner, balance)
  return {
    "deposit":  def(n) balance = balance + n return balance end,
    "describe": def() return owner + ": " + balance.to_s() end
  }
end

a = new_account("robert", 100)
b = new_account("someone", 0)

puts(a.deposit(50))
puts(a.describe())
puts(b.describe())
