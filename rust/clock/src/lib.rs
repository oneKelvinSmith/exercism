use std::fmt;

const HOURS_IN_A_DAY: i32 = 24;
const MINUTES_IN_AN_HOUR: i32 = 60;
const MINUTES_IN_A_DAY: i32 = MINUTES_IN_AN_HOUR * HOURS_IN_A_DAY;

fn divide(number: i32, divisor: i32) -> i32 {
    number / divisor
}

fn modulo(number: i32, modulus: i32) -> i32 {
    ((number % modulus) + modulus) % modulus
}

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let clock = Self { hours, minutes };

        clock.normalise()
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let clock = Self {
            hours: self.hours,
            minutes: self.minutes + minutes,
        };

        clock.normalise()
    }

    fn normalise(mut self) -> Self {
        let normalised_total_minutes = modulo(
            (self.hours * MINUTES_IN_AN_HOUR) + self.minutes,
            MINUTES_IN_A_DAY,
        );
        self.hours = divide(normalised_total_minutes, MINUTES_IN_AN_HOUR);
        self.minutes = modulo(normalised_total_minutes, MINUTES_IN_AN_HOUR);
        self
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
