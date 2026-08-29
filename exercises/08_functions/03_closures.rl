// A function remembers where it was written, so it can see the variables that
// were in scope there. That is enough to build an object: a hash of functions
// closing over a constructor's locals, with private state.
//
// A value stored in a hash under a name can be called with a dot, so
// account.deposit(50) works.
//
// Make this print:
//
//   150
//   robert: 150
//   someone: 0

def new_account(owner, balance)
  return {
    "deposit":  def(n) balance = balance + n return balance end,
    "describe": def() return owner + ": " + balance.to_s() end
  }
end

a = new_account("robert", 100)
b = new_account("someone", 0)

// TODO: deposit 50 into a and print what it answers
// TODO: print a.describe() and then b.describe()
