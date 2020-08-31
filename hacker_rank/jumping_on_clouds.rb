#!/bin/ruby

require 'json'
require 'stringio'

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c)
  s = 0
  i = 0

  while i < c.length
    if c[i + 2] && c[i + 2] != 1
      s += 1
      i += 1
    elsif c[i + 1] && c[i + 1] != 1
      s += 1
    end
    i += 1
  end

  return s
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

c = gets.rstrip.split(' ').map(&:to_i)

result = jumpingOnClouds c

fptr.write result
fptr.write "\n"

fptr.close()
