(ns reverse-string
  (:require [clojure.string :as str]))

(defn reverse-string [s]
  (loop [[head & rest :as characters] (vec s) reversed []]
    (if (empty? characters)
      (str/join reversed)
      (recur rest (cons head reversed)))
    )
  )
