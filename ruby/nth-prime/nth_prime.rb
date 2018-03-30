# frozen_string_literal: true

module BookKeeping
  VERSION = 1
end

class Prime
  MIN_PRIME     = 2
  MIN_ODD_PRIME = 3

  class << self
    def nth(limit)
      raise ArgumentError, 'there is no zeroth prime' if limit.zero?

      primes = [MIN_PRIME]
      number = MIN_ODD_PRIME

      while primes.size < limit
        primes << number if prime?(number, primes)
        number += MIN_PRIME
      end

      primes.last
    end

    private

    def prime?(number, primes)
      primes.none? { |prime| (number % prime).zero? }
    end
  end
end
