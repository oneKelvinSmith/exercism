(ns armstrong-numbers)

(defn- digits [number]
  (let [base 10]
    (loop [n number acc ()]
      (if (zero? n)
        acc
        (recur (quot n base) (conj acc (rem n base)))))))

(defn- exp [base exponent] (reduce * (repeat exponent base)))

(defn- strong-arm [number]
  (let [ds (digits number)]
    (reduce + (map #(exp % (count ds)) ds))))

(defn armstrong? [number] (= number (strong-arm number)))