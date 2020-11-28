(ns armstrong-numbers)

(defn- digits [number]
  (let [base 10 initial-ds '()]
    (->> [initial-ds number]
         (iterate (fn [[ds n]] [(conj ds (rem n base)) (quot n base)]))
         (drop-while #((comp not zero?) (second %)))
         (first)
         (first))))

(defn- exp [base exponent] (apply * (repeat exponent base)))

(defn- strong-arm [number]
  (let [ds (digits number)]
    (->> ds
         (map #(exp % (count ds)))
         (apply +))))

(defn armstrong? [number] (= number (strong-arm number)))
