(ns rna-transcription
  (:require [clojure.string :as string]))

(defn to-rna [dna]
  (let [mapping {\C \G
                 \G \C
                 \A \U
                 \T \A}
        dna->rna #(if-let [[_ rna] (find mapping %)]
                    rna
                    (throw (AssertionError. (str "Invalid DNA: " dna))))]
    (string/join (map dna->rna dna))))
