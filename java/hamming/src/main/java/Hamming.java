public class Hamming {
    private int hammingDistance = 0;

    public Hamming(String leftStrand, String rightStrand) {
        validateStrands(leftStrand, rightStrand);
        calculateHammingDistance(leftStrand, rightStrand);
    }

    public int getHammingDistance() {
        return this.hammingDistance;
    }

    private void calculateHammingDistance(String leftStrand, String rightStrand) {
        for (int i = 0; i < leftStrand.length(); i++) {
            if (leftStrand.charAt(i) != rightStrand.charAt(i)) {
                this.hammingDistance += 1;
            }
        }
    }

    private void validateStrands(String leftStrand, String rightStrand) throws IllegalArgumentException {
        if (leftStrand.length() == rightStrand.length()) return;
        if (isEmpty(leftStrand)) throw new IllegalArgumentException("left strand must not be empty.");
        if (isEmpty(rightStrand)) throw new IllegalArgumentException("right strand must not be empty.");
        throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
    }

    private boolean isEmpty(String strand) {
        return strand.equals("");
    }
}
