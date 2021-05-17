(ns rna-transcription
  (:require [clojure.string :as string]))

(defn to-rna [dna]
  (let [rna (#(string/join (map {\C \G \G \C \A \U \T \A} %)) dna)]
    (assert (= (count rna) (count dna)) (str "Invalid DNA: " dna))
    rna))

