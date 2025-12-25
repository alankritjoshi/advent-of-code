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
  def initialize(lines)
    @lines = lines
  end

  def solve
    @lines
  end
end

main if $PROGRAM_NAME == __FILE__
