#!/usr/bin/env ruby
# frozen_string_literal: true

require "optparse"

def main
  options = {}

  parser = OptionParser.new do |opts|
    opts.banner = "Usage: aoc_runner.rb [options]"

    opts.on("-i", "--input FILE", "Input file name") do |file|
      options[:input] = file
    end
  end

  parser.parse!

  unless options[:input]
    warn "Missing required option: --input FILE"
    warn parser
    exit 1
  end

  input_file_name = options[:input]

  File.open(input_file_name, "r") do |f|
    f.each_line do |line|
      puts line
    end
  end
end

if $PROGRAM_NAME == __FILE__
  main
end
