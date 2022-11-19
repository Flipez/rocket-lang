class Car
  def initialize(b)
    this.brand = b
  end

  def build()
    puts(brand)
    puts('New car from %s'.format(brand))
  end
end


fab = Car.new("audi")

fab.build()