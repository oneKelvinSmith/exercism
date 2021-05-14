(ns rna-transcription
  (:require [clojure.string :as string]))

(defn to-rna [dna]
  (let [lookup-rna #({\C \G \G \C \A \U \T \A} %)
        dna->rna #(string/join (map lookup-rna %))
        rna (dna->rna dna)]
    (assert (= (count rna) (count dna)) (str "Invalid DNA: " dna))
    rna))

