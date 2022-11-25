class Car
  def initialize(b)
    this.brand = b
  end

  def build()
    puts(this.brand)
    puts('New car from %s'.format(this.brand))
  end
end


fab = Car.new("audi")

fab.build()