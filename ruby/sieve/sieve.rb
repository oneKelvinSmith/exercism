# frozen_string_literal: true

module BookKeeping
  VERSION = 1
end

class Sieve
  SMALLEST_PRIME = 2

  def initialize(limit = 0)
    @limit  = limit
    @sieve  = (SMALLEST_PRIME..limit).to_a
    @primes = []
  end

  def primes
    return [] if limit < SMALLEST_PRIME

    until sieve.empty?
      prime = sieve.shift
      sieve.reject! { |number| (number % prime).zero? }
      @primes << prime
    end

    @primes
  end

  private

  attr_reader :limit, :sieve
end
