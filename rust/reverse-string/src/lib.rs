pub fn reverse(input: &str) -> String {
    let mut characters: Vec<char> = input.chars().collect();
    characters.reverse();
    characters.iter().collect()
}
