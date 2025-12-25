#!/usr/bin/env ruby
# frozen_string_literal: true

require 'optparse'

def main
  options = {}

  parser = OptionParser.new do |opts|
    opts.banner = 'Usage: aoc_runner.rb [options]'

    opts.on('-i', '--input FILE', 'Input file name') do |file|
      options[:input] = file
    end
  end

  parser.parse!

  unless options[:input]
    warn 'Missing required option: --input FILE'
    warn parser
    exit 1
  end

  input_file_name = options[:input]

  p ProductIDValidator.new(File.readlines(input_file_name, chomp: true)).solve
end

class ProductIDValidator
  def initialize(product_ids)
    @product_ids = product_ids.first.split(',').map do |id|
      id.split('-').map(&:to_i)
    end
    @invalid_sum = 0
  end

  def num_repeats?(num)
    num.to_s.match?(/^(.+)\1$/)
  end

  def solve
    invalid_groups = @product_ids.map do |first, second|
      (first..second).filter(&method(:num_repeats?))
    end
    invalid_groups.reject(&:empty?).flatten.sum
  end
end

main if $PROGRAM_NAME == __FILE__
