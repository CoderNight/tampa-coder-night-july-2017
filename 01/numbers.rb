# find elements in a that aren't in b
def find_missing(a,b)
  compare freq(a), freq(b)
end

def compare(list_a, list_b)
  puts list_b.select { |el, count| list_a[el] < count }
    .keys
    .map(&:to_i)
    .sort
    .join " "
end

def freq(list)
  list.reduce(Hash.new(0)) do |memo, el|
    memo[el] += 1
    memo
  end
end


lines = []
ARGF.each_slice(2) do |_, line|
  lines << line.split
end

find_missing lines[0], lines[1]
