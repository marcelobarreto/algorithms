# https://www.hackerrank.com/challenges/counting-valleys/
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the countingValleys function below.
def countingValleys(n, s)
  sea_level = 0
  valleys = 0
  i = 0

  while i < s.length do
    s[i] == 'D' ? sea_level -= 1 : sea_level += 1
    valleys += 1 if sea_level == 0 && s[i] == 'U'
    i += 1
  end

  return valleys
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

s = gets.to_s.rstrip

result = countingValleys n, s

fptr.write result
fptr.write "\n"

fptr.close()
