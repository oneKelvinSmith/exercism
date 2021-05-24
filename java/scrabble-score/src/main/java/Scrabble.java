class Scrabble {

    private final int score;

    Scrabble(String word) {
        this.score = word.codePoints().map(this::characterScore).sum();
    }

    public int getScore() {
        return score;
    }

    private int characterScore(int c) {
        return switch (Character.toLowerCase((char) c)) {
            case 'q', 'z' -> 10;
            case 'j', 'x' -> 8;
            case 'k' -> 5;
            case 'f', 'h', 'v', 'w', 'y' -> 4;
            case 'b', 'c', 'm', 'p' -> 3;
            case 'd', 'g' -> 2;
            default -> 1;
        };
    }
}
