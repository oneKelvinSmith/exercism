trait DivisibleBy {
    fn is_divisible_by(self, num: i32) -> bool;
    fn is_not_divisible_by(self, num: i32) -> bool;
}

impl DivisibleBy for i32 {
    fn is_divisible_by(self, num: i32) -> bool {
        self % num == 0
    }
    fn is_not_divisible_by(self, num: i32) -> bool {
        !self.is_divisible_by(num)
    }
}

pub fn is_leap_year(year: i32) -> bool {
    year.is_divisible_by(4) &&
        year.is_not_divisible_by(100) ||
        year.is_divisible_by(400)
}
