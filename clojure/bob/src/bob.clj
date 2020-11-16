(ns bob,
  (:require [clojure.string :as string]))

(defn response-for [s]
  (let [remark (string/trim s)
        question? #(string/ends-with? % "?")
        sensible? #(not (nil? (re-find #"[A-z]" %)))
        forceful? #(and (sensible? %) (= % (string/upper-case %)))
        exasperated? #(and (forceful? %) (question? %))
        silence? #(string/blank? %)]

    (condp apply [remark]
      silence? "Fine. Be that way!"
      exasperated? "Calm down, I know what I'm doing!"
      forceful? "Whoa, chill out!"
      question? "Sure."
      "Whatever."
      ))
  )
