lines = File.readlines('1.in').map{_1.split.map(&:to_i)}
a, b = lines.transpose.map(&:sort)
p a.zip(b).map{(_2 - _1).abs}.sum
p a.map{b.count(_1)*_1}.sum
