(ns armstrong-numbers)

(defn- digits
  ([number]
   (digits number ()))
  ([n acc]
   (if (zero? n)
     acc
     (digits (quot n 10) (conj acc (rem n 10))))))

(defn- exp [base exponent] (reduce * (repeat exponent base)))

(defn- strongarm [number]
  (let [ds (digits number)]
    (reduce + (map #(exp % (count ds)) ds))))

(defn armstrong? [number] (= number (strongarm number)))
