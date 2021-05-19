(ns rna-transcription
  (:require [clojure.string :as string]))

(def dna->rna {\C \G \G \C \A \U \T \A})
(defn to-rna [dna]
  (let [rna (string/join (map dna->rna dna))]
    (assert (= (count rna) (count dna)) (str "Invalid DNA: " dna))
    rna))

