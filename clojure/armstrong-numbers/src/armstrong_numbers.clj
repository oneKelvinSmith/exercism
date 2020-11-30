(ns armstrong-numbers)

(defn- exp [base exponent] (apply * (repeat exponent base)))

(defn- digits [number]
  (let [base 10]
    (->> number
         (iterate #(quot % base))
         (take-while pos?)
         (map #(rem % base)))))

(defn- strong-arm [number]
  (let [ds (digits number)]
    (->> ds
         (map #(exp % (count ds)))
         (apply +))))

(defn armstrong? [number] (= number (strong-arm number)))
