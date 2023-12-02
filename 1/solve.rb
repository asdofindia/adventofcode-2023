require 'set'

class SubstringSearch
  def initialize(words)
    @words = words
    @substrings = generate_substrings(words)
  end

  def search(search_string)
    matches = []
    current_slices = []

    search_string.each_char do |char|
      new_slices = current_slices.map { |slice| slice + char }
      new_slices << char

      current_slices = new_slices.select { |slice| @substrings.include?(slice) }

      current_slices.each do |slice|
        matches << slice if @words.include?(slice)
      end
    end

    matches
  end

  private

  def generate_substrings(words)
    substrings = Set.new

    words.each do |word|
      (0..word.length - 1).each do |i|
        substrings << word[0..i]
      end
    end

    substrings.to_a
  end
end

replacements = {}

# Part 1
# search_words = []

# Part 2
search_words = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

(0..9).each do |i|
  search_words << i.to_s
  replacements[search_words[i]] = i.to_s
  replacements[i.to_s] = i.to_s
end
substring_search = SubstringSearch.new(search_words)

sum = 0

File.readlines("input", chomp: true).each do |line|
    matches = substring_search.search(line)
    calibration_value = replacements[matches.first] + replacements[matches.last]
    puts calibration_value + "----" + line
    sum += calibration_value.to_i
end

puts sum

