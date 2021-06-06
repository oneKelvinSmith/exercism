(ns run-length-encoding
  (:import (java.lang Integer))
  (:require [clojure.string :as string]))

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (let [groups (partition-by identity plain-text)
        encoder (fn [encoded group]
                  (let [character (first group)
                        frequency (count group)
                        code (if (= 1 frequency) "" frequency)]
                    (str encoded code character)))]
    (reduce encoder "" groups)))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (let [groups (re-seq #"(\d+)?[\w\s]" cipher-text)
        decoder (fn [group]
                  (let [frequency ((fnil #(Integer. %) 1) (last group))
                        character (last (first group))]
                    (repeat frequency character)))]
    (string/join (flatten (map decoder groups)))))
