(ns run-length-encoding
  (:import (java.lang Integer)))

(defn run-length-encode
  "encodes a string with run-length-encoding"
  [plain-text]
  (let [character-groups (partition-by identity plain-text)
        group-encoder (fn [encoded group]
                        (let [character (first group)
                              frequency (count group)
                              code (if (= 1 frequency) "" frequency)]
                          (str encoded code character)))]
    (reduce group-encoder "" character-groups)))

(defn run-length-decode
  "decodes a run-length-encoded string"
  [cipher-text]
  (let [matches (re-seq #"(\d+)?[\w\s]" cipher-text)]
    (clojure.string/join (map (fn [group]
                                (if (nil? (last group))
                                  (first group)
                                  (clojure.string/join (repeat (Integer. (last group)) (last (first group)))))) matches))))
