# you can write to stdout for debugging purposes, e.g.
# puts "this is a debug message"

def solution(a)
  # write your code in Ruby 2.2
  return a[0] if a.length == 1

  for i in 0..a.length
    unpaired = a[i + 2] if a[i+2] && a[i] != a[i+2] && a[i+2] % 2 == 0
  end

  return unpaired
end

puts solution([])