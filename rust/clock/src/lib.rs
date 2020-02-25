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
        let normalised_total_minutes = modulo((hours * MINUTES_IN_AN_HOUR) + minutes, MINUTES_IN_A_DAY);
        let normalised_hours = divide(normalised_total_minutes, MINUTES_IN_AN_HOUR);
        let normalised_minutes = modulo(normalised_total_minutes, MINUTES_IN_AN_HOUR);

        Self {
            hours: normalised_hours,
            minutes: normalised_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
