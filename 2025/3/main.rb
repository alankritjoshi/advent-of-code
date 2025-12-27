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

  p Solver.new(File.readlines(input_file_name, chomp: true)).solve
end

class Solver
  def initialize(banks)
    @banks = banks
  end

  def bank_processor(bank)
    bank_max = 0

    bank.chars.each_with_index do |chf, idx|
      first_digit = chf.to_i

      bank.chars[idx + 1..].each do |chs|
        second_digit = chs.to_i
        bank_max = [bank_max, "#{first_digit}#{second_digit}".to_i].max
      end
    end

    bank_max
  end

  def solve
    @banks.map(&method(:bank_processor)).sum
  end
end

main if $PROGRAM_NAME == __FILE__
