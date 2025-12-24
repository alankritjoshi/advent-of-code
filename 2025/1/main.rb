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

  p LockPick.new(File.readlines(input_file_name, chomp: true)).solve
end

class LockPick
  def initialize(rotations, start = 50)
    @rotations = rotations
    @pointer = start
    @clicks = 0
  end

  def solve
    @rotations.each do |rotation|
      direction = rotation[0]
      distance = rotation[1..].to_i

      distance = case direction
                 when 'L' then distance * -1
                 when 'R' then distance
                 end

      new_distance = @pointer + distance

      # Positive incomplete rotation
      if new_distance.positive? && new_distance < 100
      # Positive complete rotations
      elsif new_distance >= 100
        @clicks += (new_distance / 100)
      # Stuck at zero
      elsif @pointer.zero? && new_distance.zero?
      # Started at zero so at least one negative rotation is needed
      elsif @pointer.zero? && new_distance.negative?
        @clicks += (-1 * new_distance) / 100
      # Negative complete rotations
      elsif @pointer.positive? && new_distance <= 0
        @clicks += ((-1 * new_distance) / 100) + 1
      end

      @pointer = new_distance % 100

      # @clicks += 1 if @pointer.zero?
    end

    @clicks
  end
end

main if $PROGRAM_NAME == __FILE__
