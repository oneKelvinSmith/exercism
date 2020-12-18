(ns rna-transcription
  (:require [clojure.string :as string]))

(defn to-rna [dna]
  (let [validate #(if (= (count %) (count dna))
                    %
                    (throw (AssertionError. (str "Invalid DNA: " dna))))
        dna->rna {\C \G
                  \G \C
                  \A \U
                  \T \A}]
    (validate (string/join (map dna->rna dna)))))
